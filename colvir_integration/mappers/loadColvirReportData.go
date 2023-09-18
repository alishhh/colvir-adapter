package mappers

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/schema"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/statements"
)

type LCRDMapper struct{}

func (am *LCRDMapper) Map(src map[string]interface{}) ([]byte, error) {
	var prm strings.Builder
	prms := ""
	mpI := src["reportParams"]
	if mpI != nil {
		if mp, ok := mpI.(map[string]interface{}); ok {
			for k, v := range mp {
				prm.WriteString(k + "=>" + makeStringValue(v) + ",")
			}
		}
	}
	prms = prm.String()
	if prms != "" {
		prms = prms[:len(prms)-1]
	}
	ret := models.SoapEnvClients{
		SoapenvAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		V1Attr:      "http://bus.colvir.com/service/statement/v1",
		V11Attr:     "http://bus.colvir.com/common/support/v1",
		Header: models.SoapHeader{
			NamespaceAttr: "http://www.w3.org/2005/08/addressing",
			Action:        "http://bus.colvir.com/service/statement/loadColvirReportData",
			To:            "http://www.w3.org/2005/08/addressing/anonymous",
			RelatesTo:     "http://www.w3.org/2005/08/addressing/unspecified",
			MessageID:     uuid.NewString(),
		},
		Body: statements.LoadColvirReportDataBody{
			Elem: statements.LoadColvirReportDataElem{
				AbstractRequest: &schema.AbstractRequest{
					Head: &schema.RequestHeader{
						RequestId: "0",
						Params: &schema.SessionParams{
							InterfaceVersion: "1.0",
							ClientType:       "GWS",
							Language:         "ru",
							OperationalDate:  time.Now().Format("2006-01-02T15:04:05"),
							ClientTimeout:    0,
						},
					},
				},
				ReportCode:   makeStringValue(src["reportCode"]),
				ReportParams: prms,
			},
		},
	}
	return xml.Marshal(ret)
}

func (lcl *LCRDMapper) IsAppliable(command string) bool {
	return strings.Contains(strings.ToLower(command), "loadcolvirreportdata")
}
