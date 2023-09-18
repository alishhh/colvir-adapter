package mapper

import (
	"bytes"
	"encoding/xml"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/soap"
)

type AcclMapper struct{}

func (am *AcclMapper) Map(src map[string]interface{}) ([]byte, error) {
	channel := makeStringValue(src["channel"])
	operation := strings.Replace(makeStringValue(src["operation"]), ".00", "", -1)
	description := makeStringValue(src["description"])
	account := ""
	account_type := ""
	if accm, ok := src["account"].(map[string]interface{}); ok {
		account = makeStringValue(accm["value"])
		account_type = makeStringValue(accm["type"])
	}
	reference := makeStringValue(src["reference"])
	timeout := makeStringValue(src["timeout"])
	var timeoutS *string
	if timeout == "" {
		timeout = "30"
		timeoutS = &timeout
	}
	retObj := &models.AcclRequest{
		Version: "3.0",
		Header: models.AcclHeader{
			Channel:   channel,
			Language:  "RU",
			Reference: reference,
			Date:      time.Now().Format("2006-01-02T15:04:05"),
			Timeout:   timeoutS,
		},
		Body: models.AcclBody{
			Operation:   operation,
			Description: description,
			Account: models.AcclTypedField{
				Text: account,
				Type: account_type,
			},
		},
	}
	bd, _ := xml.Marshal(retObj)
	tmpl, err := template.New("soap").Parse(string(soap.XML()))
	if err != nil {
		return nil, err
	}
	var ret bytes.Buffer
	err = tmpl.Execute(
		&ret, map[string]string{
			"action":     `http://bus.colvir.com/service/cap/v3/ACCL/ACCL` + operation,
			"body":       string(bd),
			"messageID":  uuid.NewString(),
			"namespaces": `xmlns:accl="http://bus.colvir.com/service/cap/v3/ACCL" xmlns:ws="http://bus.colvir.com/service/cap/v3/ws" xmlns:v3="http://bus.colvir.com/service/cap/v3" xmlns:xdat="http://bus.colvir.com/service/cap/v3/ACCL/xdata"`,
		})
	if err != nil {
		return nil, err
	}
	return ret.Bytes(), nil
}

func (am *AcclMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "accl")
}
