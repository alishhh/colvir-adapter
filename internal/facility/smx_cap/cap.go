package smx_cap

import (
	"context"
	"fmt"

	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/mapper"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/storage"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/internal/storage/soap"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/logger"
)

type CAP struct {
	conf   config.IConfigStorage
	log    *logger.ContextLogManager
	mm     *mapper.MapperManager
	caller storage.IStorage
}

func BuildCAPFacility(conf config.IConfigStorage, log *logger.ContextLogManager) (*CAP, error) {
	const fn = "facility.cap.buildCAPFacility"

	section, err := conf.GetSection("app.soap")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	soapCaller, err := soap.BuildSoapFromConfig(section, log)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	mm := mapper.NewMapperManager(conf, log)
	mm.Register(
		&mapper.FinaMapper{},
		&mapper.AcclMapper{},
		&mapper.InfoMapper{},
		&mapper.AdmnMapper{},
		&mapper.StatMapper{},
	)

	return &CAP{
		conf:   conf,
		log:    log,
		mm:     mm,
		caller: soapCaller,
	}, nil
}

func (c *CAP) SendRequest(tracerCtx context.Context, command string, payload map[string]interface{}) ([]byte, error) {
	const fn = "facility.cap.SendRequest"

	reqBytes, err := c.mm.Map(command, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	respBytes := []byte{}
	// ctx, cncl := context.WithTimeout(tracerCtx, caller.TimeOut)
	// defer cncl()

	// reqStart := time.Now()

	err = c.caller.Call(tracerCtx, &reqBytes, &respBytes)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	c.log.DebugContext(tracerCtx, "request and response body", map[string]interface{}{"request_body": string(reqBytes), "response_body": string(respBytes)})
	// c.observeSmixDuration(action, operation, time.Since(reqStart).Seconds())

	return respBytes, nil
}
