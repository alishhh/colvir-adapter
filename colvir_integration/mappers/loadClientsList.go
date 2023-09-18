package mappers

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/clients"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/schema"
)

type LoadClientsListMapper struct{}

func (am *LoadClientsListMapper) Map(src map[string]interface{}) ([]byte, error) {
	var dq *clients.ClientsListFilter
	var clq *schema.CustomQuery
	var cq *clients.ClientsListFilter

	//var dcFilter []schema.ICustomOperator
	//var cmFilter []schema.ICustomOperator

	if clqI, ok := src["clientsListQuery"]; ok && clqI != nil {
		clqM, ok := clqI.(map[string]interface{})
		if ok {
			var clFilter []schema.ICustomOperator
			eqs, _ := am.getEquals(clqM)
			likes, _ := am.getLike(clqM)
			ins, _ := am.getIn(clqM)
			clFilter = []schema.ICustomOperator{}
			if eqs != nil {
				clFilter = append(clFilter, eqs)
			}
			if likes != nil {
				clFilter = append(clFilter, likes)
			}
			if ins != nil {
				clFilter = append(clFilter, ins)
			}
			if len(clFilter) > 0 {
				clq = &schema.CustomQuery{
					Operators: clFilter,
				}
			}
		}
	}

	if dqI, ok := src["documentsQuery"]; ok && dqI != nil {
		dqM, ok := dqI.(map[string]interface{})
		if ok {
			var dFilter []schema.ICustomOperator
			eqs, _ := am.getEquals(dqM)
			likes, _ := am.getLike(dqM)
			ins, _ := am.getIn(dqM)
			dFilter = []schema.ICustomOperator{}
			if eqs != nil {
				dFilter = append(dFilter, eqs)
			}
			if likes != nil {
				dFilter = append(dFilter, likes)
			}
			if ins != nil {
				dFilter = append(dFilter, ins)
			}
			if len(dFilter) > 0 {
				dq = &clients.ClientsListFilter{
					Filter: &schema.CustomQuery{
						Operators: dFilter,
					},
				}

			}
		}
	}

	if cqI, ok := src["contactsQuery"]; ok && cqI != nil {
		cqM, ok := cqI.(map[string]interface{})
		if ok {
			eqs, _ := am.getEquals(cqM)
			likes, _ := am.getLike(cqM)
			ins, _ := am.getIn(cqM)
			cqFilter := []schema.ICustomOperator{}
			if eqs != nil {
				cqFilter = append(cqFilter, eqs)
			}
			if likes != nil {
				cqFilter = append(cqFilter, likes)
			}
			if ins != nil {
				cqFilter = append(cqFilter, ins)
			}
			if len(cqFilter) > 0 {
				cq = &clients.ClientsListFilter{
					Filter: &schema.CustomQuery{
						Operators: cqFilter,
					},
				}

			}
		}
	}
	hdetails := 0
	if src["hideDetails"] != nil {
		hdetails = makeIntValue(src["hideDetails"])
	}
	if dq == nil {
		dq = &clients.ClientsListFilter{}
	}
	if cq == nil {
		cq = &clients.ClientsListFilter{}
	}

	ret := models.SoapEnvClients{
		SoapenvAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		V1Attr:      "http://bus.colvir.com/service/clients/v1",
		V11Attr:     "http://bus.colvir.com/common/support/v1",
		V12Attr:     "http://bus.colvir.com/common/query/v1",
		Header: models.SoapHeader{
			NamespaceAttr: "http://www.w3.org/2005/08/addressing",
			Action:        "http://bus.colvir.com/service/clients/loadClientsList",
			To:            "http://www.w3.org/2005/08/addressing/anonymous",
			RelatesTo:     "http://www.w3.org/2005/08/addressing/unspecified",
			MessageID:     uuid.NewString(),
		},
		Body: clients.LoadClientsBody{
			Clients: clients.LoadClientsListElem{
				HideDetails: &hdetails,
				ClientRole:  makeStringValue(src["clientsRole"]),
				AbstractRequest: &schema.AbstractRequest{
					Head: &schema.RequestHeader{
						RequestId: "0",
						Params: &schema.SessionParams{
							InterfaceVersion: "1.0",
							ClientType:       "GWS",
							Language:         "ru",
							OperationalDate:  time.Now().Format("2006-01-02T15:04:05"),
						},
					},
				},
				ClientsListQuery: clq,
				DocumentsQuery:   *dq,
				ContactsQuery:    *cq,
			},
		},
	}
	return xml.MarshalIndent(ret, "", " ")
}

func (lcl *LoadClientsListMapper) IsAppliable(command string) bool {
	return strings.Contains(strings.ToLower(command), "loadclientslist")
}

func (lcl *LoadClientsListMapper) getEquals(querysrc map[string]interface{}) ([]*schema.CustomEquals, error) {
	eq, ok := querysrc["equals"].([]interface{})
	if ok {
		ret := []*schema.CustomEquals{}
		for _, e := range eq {
			eMap, ok := e.(map[string]interface{})
			if ok {
				if _, ok = eMap["attr"]; !ok {
					continue
				}
				if _, ok = eMap["value"]; !ok {
					continue
				}
				attr := makeStringValue(eMap["attr"])
				val := makeStringValue(eMap["value"])
				eq := schema.CustomEquals{
					Value: schema.Value{
						Value:     val,
						AttrField: attr,
					},
				}
				ret = append(ret, &eq)
			}
		}
		return ret, nil
	}
	return nil, nil

}

func (lcl *LoadClientsListMapper) getLike(querysrc map[string]interface{}) ([]*schema.CustomLike, error) {
	eq, ok := querysrc["like"].([]interface{})
	if ok {
		ret := []*schema.CustomLike{}
		for _, e := range eq {
			eMap, ok := e.(map[string]interface{})
			if ok {
				if _, ok = eMap["attr"]; !ok {
					continue
				}
				if _, ok = eMap["value"]; !ok {
					continue
				}
				attr := makeStringValue(eMap["attr"])
				val := makeStringValue(eMap["value"])
				eq := schema.CustomLike{
					Value: schema.Value{
						Value:     val,
						AttrField: attr,
					},
				}
				ret = append(ret, &eq)
			}
		}
		return ret, nil
	}
	return nil, nil
}

func (lcl *LoadClientsListMapper) getIn(querysrc map[string]interface{}) ([]*schema.CustomIn, error) {
	eq, ok := querysrc["in"].([]interface{})
	if ok {
		ret := []*schema.CustomIn{}
		for _, e := range eq {
			eMap, ok := e.(map[string]interface{})
			if ok {
				if _, ok = eMap["attr"]; !ok {
					continue
				}
				if _, ok = eMap["value"]; !ok {
					continue
				}
				attr := makeStringValue(eMap["attr"])
				valArr, ok := eMap["value"].([]interface{})
				vals := []string{}
				if ok {
					for _, v := range valArr {
						vals = append(vals, makeStringValue(v))
					}
				} else {
					vals = append(vals, makeStringValue(eMap["value"]))
				}

				eq := schema.CustomIn{
					ValueArray: schema.ValueArray{
						AttrField: attr,
						Value:     vals,
					},
				}
				ret = append(ret, &eq)
			}
		}
		return ret, nil
	}
	return nil, nil
}
