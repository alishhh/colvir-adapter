package models

import "encoding/xml"

//ISoapBody - пустой интерфейс, нужен для смыслового обозначения поля Body в SoapEnvelope
//В поля, обозначенные этим интерфейсом нельзя передавать map, interface{}, или иные неявные на этапе компиляции структуры
//Подобное присваивание приведёт к ошибкам при сериализации запроса
type ISoapBody interface{}

type SoapEnvClients struct {
	XMLName     xml.Name   `xml:"soapenv:Envelope"`
	SoapenvAttr string     `xml:"xmlns:soapenv,attr,omitempty"`
	V1Attr      string     `xml:"xmlns:v1,attr,omitempty"`
	V11Attr     string     `xml:"xmlns:v11,attr,omitempty"`
	V12Attr     string     `xml:"xmlns:v12,attr,omitempty"`
	Header      SoapHeader `xml:"soapenv:Header,omitempty"`
	Body        ISoapBody  `xml:"soapenv:Body,omitempty"`
}

type SoapHeader struct {
	NamespaceAttr string `xml:"xmlns:wsa,attr"`
	Action        string `xml:"wsa:Action"`
	MessageID     string `xml:"wsa:MessageID"`
	To            string `xml:"wsa:To"`
	RelatesTo     string `xml:"wsa:RelatesTo"`
}
