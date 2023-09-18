package soap

func XML() []byte {
	return []byte(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" {{.namespaces}} xmlns:v1="http://bus.colvir.com/common/v1">
    <soapenv:Header xmlns:wsa="http://www.w3.org/2005/08/addressing">
        <wsa:Action>{{.action}}</wsa:Action>
        <wsa:MessageID>{{.messageID}}</wsa:MessageID>
		<wsa:To>http://www.w3.org/2005/08/addressing/anonymous</wsa:To>
		<wsa:RelatesTo>http://www.w3.org/2005/08/addressing/unspecified</wsa:RelatesTo>
    </soapenv:Header>
    <soapenv:Body>
       {{.body}}   
	</soapenv:Body>
</soapenv:Envelope>`)
}
