package models

import "encoding/xml"

type AcclRequest struct {
	XMLName xml.Name   `xml:"ws:acclRequest"`
	Text    string     `xml:",chardata"`
	Version string     `xml:"version,attr"`
	Header  AcclHeader `xml:"v3:header"`
	Body    AcclBody   `xml:"ws:body"`
}

type AcclHeader struct {
	Text      string          `xml:",chardata"`
	Channel   string          `xml:"v1:channel"`
	Reference string          `xml:"v1:reference"`
	Date      string          `xml:"v1:date"`
	Language  string          `xml:"v1:language"`
	Token     *AcclTypedField `xml:"v1:token,omitempty"`
	Timeout   *string         `xml:"v1:timeout,omitempty"`
}

type AcclBody struct {
	Text        string         `xml:",chardata"`
	Operation   string         `xml:"accl:operation"`
	Description string         `xml:"accl:description"`
	Account     AcclTypedField `xml:"accl:account"`
	XData       *AcclXData     `xml:"xdat:xData,omitempty"`
}

type AcclTypedField struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type AcclXData struct {
	Text string   `xml:",chardata"`
	Trn  []string `xml:"xdat:trn,omitempty"`
}
