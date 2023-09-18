package adapter

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/beevik/etree"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/facility"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/facility/smx_cap"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/adapter"
	aerror "gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/adapter/errors"
	aerror_i "gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/adapter/errors/implementation"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config"
	dbp_event "gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/externals/dbp-event"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/helpers/x2m"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/logger"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/tracer"
	"go.opentelemetry.io/otel/codes"
)

type ColvirAdapter struct {
	conf               config.IConfigStorage
	log                *logger.ContextLogManager
	facility           facility.IFacility
	rStart             *regexp.Regexp
	rEnd               *regexp.Regexp
	rNamespaces        *regexp.Regexp
	metric             *prometheus.CounterVec
	durationMetric     *prometheus.HistogramVec
	smixDurationMetric *prometheus.HistogramVec
	processedCounter   map[string][]prometheus.Counter
	durationHists      map[string]prometheus.Observer
	smixDurationHists  map[string]prometheus.Observer
}

type ServiceBusError struct {
	XMLName   xml.Name `xml:"ExceptionReport"`
	ErrorCode string
	ErrorText string
}

var (
	defaultError ServiceBusError = ServiceBusError{ErrorCode: "BIP3711", ErrorText: "Internal service error occured"}
)

func BuildColvirAdapter(conf config.IConfigStorage, log *logger.ContextLogManager) (adapter.IAdapter, error) {
	appname, err := conf.GetString("framework.appname")
	if err != nil {
		return nil, err
	}
	ft, err := smx_cap.BuildCAPFacility(conf, log)
	if err != nil {
		return nil, err
	}
	r, err := regexp.Compile(`\<[^/][a-z,A-Z,0-9,\-]*(:+?)`)
	if err != nil {
		return nil, err
	}
	r2, err := regexp.Compile(`\<[/][a-z,A-Z,0-9,\-]*(:+?)`)
	if err != nil {
		return nil, err
	}
	r3, err := regexp.Compile(`xmlns:[^=]+=\"[^\"]+\"`)
	if err != nil {
		return nil, err
	}
	clv := &ColvirAdapter{
		conf:        conf,
		log:         log,
		facility:    ft,
		rStart:      r,
		rEnd:        r2,
		rNamespaces: r3,
	}

	//Processed count
	clv.metric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "dbp_adapter_processed_commands",
		Help: "Counter for gathering all commands processed by adapter",
		ConstLabels: map[string]string{
			"app_name": appname,
		},
	}, []string{"command", "operation", "result"})
	err = prometheus.Register(clv.metric)
	if err != nil {
		log.Error("", fmt.Errorf("failed to register metric with error: %s", err.Error()), nil)
	}

	clv.processedCounter = map[string][]prometheus.Counter{}

	buckets := []float64{}
	buckets = append(buckets, prometheus.DefBuckets...)
	buckets = append(buckets, []float64{20, 30, 40, 60}...)

	//Total duration metric
	clv.durationMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "dbp_adapter_processed_duration_seconds",
		Help: "Histogram for gathering request duration info from adapter",
		ConstLabels: map[string]string{
			"app_name": appname,
		},
		Buckets: buckets,
	}, []string{"command", "operation"})
	clv.durationHists = map[string]prometheus.Observer{}
	prometheus.Register(clv.durationMetric)
	if err != nil {
		log.Error("", fmt.Errorf("failed to register processed duration metric with error: %s", err.Error()), nil)
	}

	//ServiceMix metric
	clv.smixDurationMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "dbp_adapter_smix_duration_seconds",
		Help: "Histogram for gathering request duration info from adapter",
		ConstLabels: map[string]string{
			"app_name": appname,
		},
		Buckets: buckets,
	}, []string{"command", "operation"})
	clv.smixDurationHists = map[string]prometheus.Observer{}
	prometheus.Register(clv.smixDurationMetric)
	if err != nil {
		log.Error("", fmt.Errorf("failed to register smix duration metric with error: %s", err.Error()), nil)
	}

	return clv, nil
}

func (ca *ColvirAdapter) ProcessRequest(ctx context.Context, event dbp_event.IEventManager, resource string) (interface{}, aerror.IError) {
	ca.log.InfoContext(ctx, "started processing request", map[string]any{"command": event.GetCommand().Dso, "event_id": event.GetEventMeta().Event.Id})
	command := event.GetEventMeta().Command.Dso
	trCtx, span := tracer.GetTracer().Start(context.Background(), "processing")
	defer span.End()

	start := time.Now()
	action, operation, err := ca.getKey(command)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to get action/operation")
		ca.log.ErrorContext(ctx, fmt.Errorf("failed to extract action/operation"), map[string]any{"command": command, "operation": operation, "action": action})
		return ca.GetErrorPayload("BIP3711", "Internal server error from adapter"), nil
	}

	ca.checkInitCounter(action, operation)
	ca.checkInitDurationHist(action, operation)
	ca.checkInitSMixHist(action, operation)

	payload, ok := event.GetEvent().Payload.(map[string]interface{})
	if !ok {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to get payload")
		ca.incFailedCounter(action, operation, time.Since(start).Seconds())
		ca.log.FatalContext(ctx, fmt.Errorf("payload is not an object/map[string]interface"), nil)
		return ca.GetErrorPayload("BIP3711", "Input payload unavailable"), nil
	}

	resp, err := ca.facility.SendRequest(trCtx, command, payload)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to process soap request")
		ca.incFailedCounter(action, operation, time.Since(start).Seconds())
		ca.log.ErrorContext(ctx, fmt.Errorf("failed to send request to service-mix: %s", err.Error()), nil)
		return ca.GetErrorPayload("BIP3711", "Remote service unavailable: "+err.Error()), nil
	}

	resp = ca.rStart.ReplaceAll(resp, []byte("<"))
	resp = ca.rEnd.ReplaceAll(resp, []byte("</"))
	resp = ca.rNamespaces.ReplaceAll(resp, []byte{})
	cleanedBytes, err := ca.cutEnvelope(resp)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to parse response")
		ca.incFailedCounter(action, operation, time.Since(start).Seconds())
		ca.log.ErrorContext(ctx, fmt.Errorf("failed to parse service-mix response: %s", err.Error()), nil)
		return ca.GetErrorPayload("BIP3711", "Remote service unavailable: "+err.Error()), nil
	}
	if strings.ToLower(command) == "loadcolvirreportdata" {
		respS := string(cleanedBytes)
		fInd := strings.Index(respS, "<reportData>")
		if fInd != -1 {
			lInd := strings.Index(respS, "</reportData>")
			if lInd != -1 {
				toDecode := respS[fInd+len("<reportData>") : lInd]
				decoded, err := base64.StdEncoding.DecodeString(toDecode)
				if err == nil {
					cleanedBytes = decoded
				}
			}
		}
	}

	cleanedBytes = ca.replaceKaz(cleanedBytes)
	ret, err := x2m.ToMap(cleanedBytes)
	if err != nil {
		return nil, &aerror_i.AdapterError{ErrCode: "DecodeFailed", ErrMessage: "failed to decode xml response from external system:" + err.Error(), Result: "unknown"}
	}
	ca.incSuccessCounter(action, operation, time.Since(start).Seconds())
	ca.log.InfoContext(ctx, "successfully processed request", nil)
	return ret, nil
}

func (ca *ColvirAdapter) ProcessRequestAsync(ctx context.Context, event dbp_event.IEventManager, resource string) (interface{}, aerror.IError) {
	return ca.ProcessRequest(ctx, event, resource)
}

func (ca *ColvirAdapter) cutEnvelope(data []byte) ([]byte, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromBytes(data)
	if err != nil {
		return nil, errors.New("failed to read xml body: " + err.Error())
	}
	el := doc.SelectElement("Envelope")
	if el == nil {
		return doc.WriteToBytes()
	}
	elBody := el.SelectElement("Body")
	if elBody == nil {
		if els := el.ChildElements(); len(els) > 0 {
			doc2 := etree.NewDocument()
			doc2.SetRoot(els[0])
			return doc2.WriteToBytes()
		} else {
			return nil, errors.New("no childs in the Envelope/Body")
		}
	}
	doc2 := etree.NewDocument()
	elBody = elBody.Copy()
	if els := elBody.ChildElements(); len(els) > 0 {
		doc2.SetRoot(els[0])
		return doc2.WriteToBytes()
	}
	return nil, errors.New("no childs in the body")
}

func (ca *ColvirAdapter) replaceKaz(data []byte) []byte {
	str := []rune(string(data))
	dict := map[rune]rune{
		'ә': 'ј',
		'ң': 'ѕ',
		'ғ': 'є',
		'ү': 'ї',
		'ұ': 'ў',
		'қ': 'ќ',
		'ө': 'ґ',
		'һ': 'ћ',
		'Ә': 'Ј',
		'Ң': 'Ѕ',
		'Ғ': 'Є',
		'Ү': 'Ї',
		'Ұ': 'Ў',
		'Қ': 'Ќ',
		'Ө': 'Ґ',
		'Һ': 'Ћ',
	}
	for i := 0; i < len(str); i++ {
		if v, ok := dict[str[i]]; ok {
			str[i] = v
		}
	}
	return []byte(string(str))
}

func (ca *ColvirAdapter) GetErrorPayload(code, text string) []byte {
	errMsg := ServiceBusError{
		ErrorCode: code,
		ErrorText: text,
	}
	ret, _ := xml.Marshal(errMsg)
	return ret
}

func (ca *ColvirAdapter) getKey(dso string) (action string, operation string, err error) {
	isOperationString := false
	var comBuilder strings.Builder
	var opBuilder strings.Builder
	for _, symbol := range dso {
		if unicode.IsDigit(symbol) {
			isOperationString = true
		}
		if isOperationString {
			opBuilder.WriteRune(symbol)
		} else {
			comBuilder.WriteRune(symbol)
		}
	}
	action = comBuilder.String()
	operation = opBuilder.String()
	if action == "" {
		err = errors.New("no operation found")
	}
	return
}

func (ca *ColvirAdapter) checkInitCounter(command, operation string) {
	if ca.processedCounter == nil {
		ca.processedCounter = map[string][]prometheus.Counter{}
	}
	key := command + operation
	if _, ok := ca.processedCounter[key]; ok {
		return
	}
	suc, err := ca.metric.GetMetricWithLabelValues(command, operation, "success")
	if err != nil {
		ca.log.Error("", fmt.Errorf("failed to build success counter for command %s with error: %s", key, err.Error()), nil)
	}
	f, err := ca.metric.GetMetricWithLabelValues(command, operation, "fail")
	if err != nil {
		ca.log.Error("", fmt.Errorf("failed to build failed counter for command %s with error: %s", key, err.Error()), nil)
	}
	ca.processedCounter[key] = []prometheus.Counter{suc, f}
}

func (ca *ColvirAdapter) checkInitDurationHist(command, operation string) {
	if ca.durationHists == nil {
		ca.durationHists = map[string]prometheus.Observer{}
	}
	key := command + operation
	if _, ok := ca.durationHists[key]; ok {
		return
	}
	h, err := ca.durationMetric.GetMetricWithLabelValues(command, operation)
	if err != nil {
		ca.log.Error("", fmt.Errorf("failed to build duration hist for command %s with error: %s", key, err.Error()), nil)
	}
	ca.durationHists[key] = h
}

func (ca *ColvirAdapter) checkInitSMixHist(command, operation string) {
	if ca.smixDurationHists == nil {
		ca.smixDurationHists = map[string]prometheus.Observer{}
	}
	key := command + operation
	if _, ok := ca.smixDurationHists[key]; ok {
		ca.log.Debug("", fmt.Errorf("smix observer already presented").Error(), nil)
		return
	}
	h, err := ca.smixDurationMetric.GetMetricWithLabelValues(command, operation)
	if err != nil {
		ca.log.Error("", fmt.Errorf("failed to build smix duration hist for command %s with error: %s", key, err.Error()), nil)
	}
	ca.smixDurationHists[key] = h
}

func (ca *ColvirAdapter) incSuccessCounter(command, operation string, duration float64) {
	key := command + operation
	if h, ok := ca.durationHists[key]; ok && h != nil {
		h.Observe(duration)
	}
	if cs, ok := ca.processedCounter[key]; ok {
		if cs[0] == nil {
			return
		}
		cs[0].Inc()
	}
}

func (ca *ColvirAdapter) incFailedCounter(command, operation string, duration float64) {
	key := command + operation
	if h, ok := ca.durationHists[key]; ok && h != nil {
		h.Observe(duration)
	}
	if cs, ok := ca.processedCounter[key]; ok {
		if cs[1] == nil {
			return
		}
		cs[1].Inc()
	}
}

func (ca *ColvirAdapter) observeSmixDuration(command, operation string, duration float64) {
	key := command + operation
	if h, ok := ca.smixDurationHists[key]; ok && h != nil {
		h.Observe(duration)
	} else {
		ca.log.Error("", fmt.Errorf("no smix observer presented"), nil)
	}
}

/*
func setTrAttr(span trace.Span, key, value string) {
	span.SetAttributes(attribute.KeyValue{Key: attribute.Key(key), Value: attribute.StringValue(value)})
}
*/
