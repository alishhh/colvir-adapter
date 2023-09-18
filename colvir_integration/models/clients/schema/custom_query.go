package schema

import "encoding/xml"

type CustomQuery struct {
	//XMLName   xml.Name `xml:"v1:clientsListQuery"`
	Operators []ICustomOperator
}

type CustomFilter struct {
	XMLName   xml.Name `xml:"v1:filter"`
	Operators []ICustomOperator
}

type CustomAnd struct {
	XMLName xml.Name `xml:"v12:and"`
	CustomOperator
}

type CustomOr struct {
	XMLName        xml.Name `xml:"v12:or"`
	CustomOperator `xml:",omitempty"`
}

type CustomEquals struct {
	XMLName xml.Name `xml:"v12:equals"`
	Value
}

type CustomLike struct {
	XMLName xml.Name `xml:"v12:equals"`
	Value
}

type CustomIn struct {
	XMLName xml.Name `xml:"v12:in"`
	ValueArray
}

type CustomNotEquals struct {
	XMLName xml.Name `xml:"v12:notEquals"`
	Value
}

type Value struct {
	AttrField string            `xml:"attr,attr"`
	Value     string            `xml:"v12:value"`
	Operators []ICustomOperator `xml:",omitempty"`
}

type ValueArray struct {
	AttrField string            `xml:"attr,attr"`
	Value     []string          `xml:"v12:value"`
	Operators []ICustomOperator `xml:",omitempty"`
}

type CustomOperator struct {
	Equals    []*Equals    `xml:"equals,omitempty"`
	NotEquals []*NotEquals `xml:"notEquals,omitempty"`
	In        []*In        `xml:"in,omitempty"`
	Like      []*Like      `xml:"like,omitempty"`
	Operators []ICustomOperator
}

type ICustomOperator interface{}
