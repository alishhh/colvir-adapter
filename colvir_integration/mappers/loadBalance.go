package mappers

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/clients"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/schema"
)

type LoadBalanceListMapper struct{}

func (am *LoadBalanceListMapper) Map(src map[string]interface{}) ([]byte, error) {

	hdrI := src["headerParams"]
	var sr, mr int
	if hdrI != nil {

		hdr, ok := hdrI.(map[string]interface{})
		if !ok {
			return nil, errors.New("loadBalanceList is not an object")
		}
		sr = makeIntValue(hdr["startRow"])
		mr = makeIntValue(hdr["maxRows"])
	}

	ret := models.SoapEnvClients{
		SoapenvAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		V1Attr:      "http://bus.colvir.com/service/clients/v1",
		V11Attr:     "http://bus.colvir.com/common/support/v1",
		V12Attr:     "http://bus.colvir.com/common/query/v1",
		Header: models.SoapHeader{
			NamespaceAttr: "http://www.w3.org/2005/08/addressing",
			Action:        "http://bus.colvir.com/service/clients/loadBalanceList",
			To:            "http://www.w3.org/2005/08/addressing/anonymous",
			RelatesTo:     "http://www.w3.org/2005/08/addressing/unspecified",
			MessageID:     uuid.NewString(),
		},
		Body: clients.LoadBalanceBody{
			Clients: clients.LoadBalanceListElem{
				AccountCode: []string{makeStringValue(src["account"])},
				AbstractRequest: &schema.AbstractRequest{
					Head: &schema.RequestHeader{
						RequestId: "0",
						Params: &schema.SessionParams{
							InterfaceVersion: "1.0",
							ClientType:       "GWS",
							Language:         "ru",
							OperationalDate:  time.Now().Format("2006-01-02T15:04:05"),
						},
						StartRow: sr,
						MaxRows:  mr,
					},
				},
			},
		},
	}
	return xml.MarshalIndent(ret, "", " ")
}

func (lcl *LoadBalanceListMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "loadbalancelist")
}
