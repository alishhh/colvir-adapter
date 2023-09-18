package mappers

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/soap"
)

type InfoMapper struct{}

func (fm *InfoMapper) Map(src map[string]interface{}) ([]byte, error) {
	channel := makeStringValue(src["channel"])
	operation := strings.Replace(makeStringValue(src["operation"]), ".00", "", -1)
	description := makeStringValue(src["description"])
	var account string
	var accType string
	if srcA, ok := src["account"].(map[string]interface{}); ok {
		account = makeStringValue(srcA["value"])
		accType = makeStringValue(srcA["type"])
	}
	reference := makeStringValue(src["reference"])
	timeout := makeInt(src["timeout"])
	var timeoutS *string
	if timeout > 0 {
		s := (strconv.FormatInt(timeout, 10))
		timeoutS = &s
	}
	retObj := &models.InfoRequest{
		Version: getVersion(operation),
		Header: models.InfoHeader{
			Channel:   channel,
			Language:  getLang(operation),
			Reference: reference,
			Date:      time.Now().Format("2006-01-02T15:04:05"),
			Timeout:   timeoutS,
		},
		Body: models.InfoBody{
			Operation:   operation,
			Description: description,
		},
	}
	if account != "" && accType != "" {
		retObj.Body.Account = &models.FinaTypedField{
			Text: account,
			Type: accType,
		}
	}
	bd, _ := xml.Marshal(retObj)
	tmpl, err := template.New("soap").Parse(string(soap.XML()))
	if err != nil {
		return nil, err
	}
	var ret bytes.Buffer
	err = tmpl.Execute(
		&ret, map[string]string{
			"action":     `http://bus.colvir.com/service/cap/v3/INFO/INFO` + operation,
			"body":       string(bd),
			"messageID":  uuid.NewString(),
			"namespaces": `xmlns:info="http://bus.colvir.com/service/cap/v3/INFO" xmlns:ws="http://bus.colvir.com/service/cap/v3/ws" xmlns:v3="http://bus.colvir.com/service/cap/v3"`,
		})
	if err != nil {
		return nil, err
	}
	return ret.Bytes(), nil
}

func (fm *InfoMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "info2520")
}
