package mappers

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

type StatMapper struct{}

func (sm *StatMapper) Map(src map[string]interface{}) ([]byte, error) {
	channel := makeStringValue(src["channel"])
	referenceHeader := makeStringValue(src["reference"])

	retObj := &models.StatRequest{
		Version: "3.0",
		Header: models.StatHeader{
			Channel:   channel,
			Language:  "ru",
			Reference: referenceHeader,
			Timeout:   "0",
			Date:      time.Now().Format("2006-01-02T15:04:05"),
		},
	}

	transaction := makeStringValuePtr(src["transaction"])
	if src["transaction"] != nil {
		retObj.Body.Transaction = transaction
	}
	extReference := makeStringValuePtr(src["extReference"])
	if src["extReference"] != nil {
		retObj.Body.Reference = extReference
	}

	bd, _ := xml.Marshal(retObj)
	tmpl, err := template.New("soap").Parse(string(soap.XML()))
	if err != nil {
		return nil, err
	}
	var ret bytes.Buffer
	err = tmpl.Execute(
		&ret, map[string]string{
			"action":     `http://bus.colvir.com/service/cap/v3/STAT/STAT`,
			"body":       string(bd),
			"messageID":  uuid.NewString(),
			"namespaces": `xmlns:fina="http://bus.colvir.com/service/cap/v3/STAT" xmlns:ws="http://bus.colvir.com/service/cap/v3/ws" xmlns:v3="http://bus.colvir.com/service/cap/v3" xmlns:xdat="http://bus.colvir.com/service/cap/v3/STAT/xdata"`,
		})
	if err != nil {
		return nil, err
	}

	return ret.Bytes(), nil
}

func (sm *StatMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "stat")
}
