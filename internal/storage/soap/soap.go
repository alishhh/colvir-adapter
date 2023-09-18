package soap

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"

	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/storage"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/logger"
)

type Soap struct {
	Host    string
	Url     string
	TimeOut time.Duration
	log     *logger.ContextLogManager
	Client  *http.Client
}

func NewSoap(host string, url string, timeOut time.Duration, log *logger.ContextLogManager) (*Soap, error) {
	const fn = "storage.soap.NewSoap"
	if host == "" || url == "" || timeOut == 0 {
		return nil, fmt.Errorf("%s: %w", fn, storage.EmptyArguments)
	}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}}
	return &Soap{Host: host, Url: url, TimeOut: timeOut, log: log, Client: &http.Client{Transport: tr, Timeout: timeOut}}, nil
}

func BuildSoapFromConfig(section config.ConfigSection, log *logger.ContextLogManager) (*Soap, error) {
	const fn = "storage.soap.BuildFromConfig"
	url, err := section.GetString("url")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	if url[0] != '/' {
		url = "/" + url
	}
	host, err := section.GetString("host")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	timeout, err := section.GetInt("timeout")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	return NewSoap(host, url, time.Duration(timeout)*time.Second, log)
}

func (s *Soap) Call(ctx context.Context, req *[]byte, resp *[]byte) error {
	const fn = "storage.soap.Call"
	var (
		err      error
		request  *http.Request
		response *http.Response
		data     io.Reader
	)

	clientTrace := &httptrace.ClientTrace{}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)
	data = bytes.NewReader(*req)
	request, err = http.NewRequestWithContext(traceCtx, http.MethodPost, s.Host+s.Url, data)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	if request != nil {
		request.Header.Set("Content-Type", "text/xml;charset=UTF-8")
		request.Header.Set("User-Agent", "go-http")
	}
	if response, err = s.Client.Do(request); err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	} else {
		defer response.Body.Close()
		if response.StatusCode != 200 {
			*resp, _ = io.ReadAll(response.Body)
			return fmt.Errorf("%s: <soap> #response status=%d body=%s", fn, response.StatusCode, string(*resp))
		} else {
			if *resp, err = io.ReadAll(response.Body); err != nil {
				return fmt.Errorf("%s: <soap> #read status=%d body=%w", fn, response.StatusCode, err)
			}
		}
	}

	return nil
}
