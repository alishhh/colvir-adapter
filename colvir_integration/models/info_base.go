package models

import (
	"encoding/xml"
)

type InfoRequest struct {
	XMLName xml.Name   `xml:"ws:infoRequest"`
	Text    string     `xml:",chardata"`
	Version string     `xml:"version,attr"`
	Header  InfoHeader `xml:"v3:header"`
	Body    InfoBody   `xml:"ws:body"`
}

type InfoHeader struct {
	Text      string          `xml:",chardata"`
	Channel   string          `xml:"v1:channel"`
	Reference string          `xml:"v1:reference"`
	Date      string          `xml:"v1:date"`
	Language  string          `xml:"v1:language"`
	Token     *FinaTypedField `xml:"v1:token,omitempty"`
	Timeout   *string         `xml:"v1:timeout,omitempty"`
}

type InfoBody struct {
	Text        string          `xml:",chardata"`
	Operation   string          `xml:"info:operation"`
	Description string          `xml:"info:description"`
	Account     *FinaTypedField `xml:"info:account,omitempty"`
	XData       *InfoXData      `xml:"xdat:xData,omitempty"`
}
type InfoXData struct {
	Trn *map[string]interface{} `xml:"xdat:trn,omitempty"`
}
