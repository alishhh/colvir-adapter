package soap

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/logger"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/tracer"
	//"go.opentelemetry.io/otel"
	//"go.opentelemetry.io/otel/attribute"
	//"go.opentelemetry.io/otel/trace"
)

type Soap struct {
	Host    string
	Url     string
	TimeOut time.Duration
	log     *logger.ContextLogManager
}

func NewSoap(host string, url string, timeOut time.Duration, log *logger.ContextLogManager) (*Soap, error) {
	if host == "" || url == "" || timeOut == 0 || log == nil {
		return nil, errors.New("arguments should not be nil or empty")
	}
	return &Soap{Host: host, Url: url, TimeOut: timeOut, log: log}, nil
}

func (s *Soap) Call(ctx context.Context, req *[]byte, resp *[]byte) error {
	var (
		err      error
		request  *http.Request
		response *http.Response
		data     io.Reader
	)
	trCtx, span := tracer.GetTracer().Start(context.Background(), "call_service_mix")
	defer span.End()
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	client := &http.Client{Transport: tr, Timeout: s.TimeOut}

	data = bytes.NewReader(*req)

	request, err = http.NewRequest(http.MethodPost, s.Host+s.Url, data)
	if err != nil {
		return err
	}
	if request != nil {
		tracer.SetHTTPTracingContext(trCtx, request)
		request.Header.Set("Content-Type", "text/xml;charset=UTF-8")
		request.Header.Set("User-Agent", "go-http")
	}

	if response, err = client.Do(request); err != nil {
		return fmt.Errorf("<soap> #call error=%s", err.Error())
	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			*resp, _ = ioutil.ReadAll(response.Body)
			return fmt.Errorf("<soap> #response status=%d body=%s", response.StatusCode, string(*resp))
		} else {
			if *resp, err = ioutil.ReadAll(response.Body); err != nil {
				return fmt.Errorf("<soap> #read status=%d body=%s", response.StatusCode, err.Error())
			}
		}
	}

	return nil
}

func BuildFromConfig(cfg config.ConfigSection, log *logger.ContextLogManager) (*Soap, error) {
	url, err := cfg.GetString("url")
	if url == "" {
		return nil, err
	}
	if url[0] != '/' {
		url = "/" + url
	}
	host, err := cfg.GetString("host")
	if host == "" {
		return nil, err
	}
	defTimeout := 60 * time.Second
	timeout, err := cfg.GetDuration("timeout")
	if timeout == 0 {
		timeout = defTimeout
	}
	return NewSoap(host, url, time.Duration(timeout)*time.Second, log)
}

/*
func setTrAttr(span trace.Span, key, value string) {
	span.SetAttributes(attribute.KeyValue{Key: attribute.Key(key), Value: attribute.StringValue(value)})
}
*/
