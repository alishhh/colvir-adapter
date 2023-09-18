package models

import "encoding/xml"

type AdmnRequest struct {
	XMLName xml.Name   `xml:"ws:admnRequest"`
	Text    string     `xml:",chardata"`
	Version string     `xml:"version,attr"`
	Header  AdmnHeader `xml:"v3:header"`
	Body    AdmnBody   `xml:"ws:body"`
}

type AdmnHeader struct {
	Text      string `xml:",chardata"`
	Channel   string `xml:"v1:channel"`
	Reference string `xml:"v1:reference"`
	Date      string `xml:"v1:date"`
	Language  string `xml:"v1:language"`
}

type AdmnBody struct {
	Text        string          `xml:",chardata"`
	Operation   string          `xml:"admn:operation"`
	Description string          `xml:"admn:description"`
	Amount      *AdmnAmount     `xml:"admn:amount"`
	Account     *AdmnTypedField `xml:"admn:account,omitempty"`
	Account2    *AdmnTypedField `xml:"admn:account2,omitempty"`
	XData       *AdmnXData      `xml:"xdat:xData,omitempty"`
}

type AdmnAmount struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type AdmnTypedField struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type AdmnXData struct {
	Trn          *AdmnTrn          `xml:"xdat:trn,omitempty"`
	FastTransfer *AdmnFastTransfer `xml:"xdat:fast_transfer,omitempty"`
}

type AdmnTrn struct {
	IntraAccountPayment *AdmnIntraAccountPayment `xml:"xdat:intra_account_payment,omitempty"`
	Card                *AdmnCard                `xml:"xdat:card,omitempty"`
	Demand              *AdmnDemand              `xml:"xdat:demand,omitempty"`
	CardInfo            *AdmnCardInfo            `xml:"xdat:cardinfo,omitempty"`
	CardSrc             *AdmnCardSrc             `xml:"xdat:cardsrc,omitempty"`
	Inform              *AdmnInform              `xml:"xdat:inform,omitempty"`
	WuTransfer          *AdmnWuTransfer          `xml:"xdat:wu_transfer,omitempty"`
	ExtFee              *AdmnExtFee              `xml:"xdat:extfee,omitempty"`
	NationalBankPayment *AdmnNationalBankPayment `xml:"xdat:national_bank_payment,omitempty"`
	Installment         *string                  `xml:"xdat:installment,omitempty"`
}

type AdmnIntraAccountPayment struct {
	PaymentDetails *string `xml:"xdat:paymentDetails,omitempty"`
}

type AdmnCard struct {
	A           string `xml:"xdat:a,omitempty"`
	C           string `xml:"xdat:c,omitempty"`
	Fa          string `xml:"xdat:fa,omitempty"`
	Aa          string `xml:"xdat:aa,omitempty"`
	Ac          string `xml:"xdat:ac,omitempty"`
	Ba          string `xml:"xdat:ba,omitempty"`
	Bc          string `xml:"xdat:bc,omitempty"`
	Sa          string `xml:"xdat:sa,omitempty"`
	Pan         string `xml:"xdat:pan,omitempty"`
	Panm        string `xml:"xdat:panm,omitempty"`
	Dt          string `xml:"xdat:dt,omitempty"`
	Sttldt      string `xml:"xdat:sttldt,omitempty"`
	Rrn         string `xml:"xdat:rrn,omitempty"`
	Auth        string `xml:"xdat:auth,omitempty"`
	Iss         string `xml:"xdat:iss,omitempty"`
	Acq         string `xml:"xdat:acq,omitempty"`
	Mcc         string `xml:"xdat:mcc,omitempty"`
	Acqc        string `xml:"xdat:acqc,omitempty"`
	Term        string `xml:"xdat:term,omitempty"`
	Caid        string `xml:"xdat:caid,omitempty"`
	Loc         string `xml:"xdat:loc,omitempty"`
	Country     string `xml:"xdat:country,omitempty"`
	City        string `xml:"xdat:city,omitempty"`
	Adr         string `xml:"xdat:adr,omitempty"`
	Termtype    string `xml:"xdat:termtype,omitempty"`
	Psys        string `xml:"xdat:psys,omitempty"`
	Stan        string `xml:"xdat:stan,omitempty"`
	Ft          string `xml:"xdat:ft,omitempty"`
	Posmode     string `xml:"xdat:posmode,omitempty"`
	Phonenumber string `xml:"xdat:phone_number,omitempty"`
}

type AdmnFastTransfer struct {
	Docid        string             `xml:"xdat:doc_id"`
	Date         string             `xml:"xdat:date"`
	State        string             `xml:"xdat:state"`
	Feeaccount   string             `xml:"xdat:fee_account"`
	Order        string             `xml:"xdat:order"`
	Type         string             `xml:"xdat:type"`
	Authorstatus string             `xml:"xdat:author_status"`
	Beneficiary  AdmnCapBeneficiary `xml:"xdat:beneficiary"`
}

type AdmnCapBeneficiary struct {
	Bin        string `xml:"xdat:bin"`
	Department string `xml:"xdat:department"`
	Rnn        string `xml:"xdat:rnn"`
	Name       string `xml:"xdat:name"`
	Resident   bool   `xml:"xdat:resident"`
	Kpp        string `xml:"xdat:kpp"`
	Purpose    string `xml:"xdat:purpose"`
}

type AdmnDemand struct {
	CustomerInfo *AdmnCustomerInfo `xml:"xdat:customerInfo"`
	AccountInfo  *AdmnAccountInfo  `xml:"xdat:accountInfo"`
	Status       *AdmnStatus       `xml:"xdat:status"`
	Deposit      *AdmnDeposit      `xml:"xdat:deposit"`
}

type AdmnCustomerInfo struct {
	Email  string `xml:"xdat:email"`
	Mobile string `xml:"xdat:mobile"`
	Locale string `xml:"xdat:locale"`
}

type AdmnAccountInfo struct {
	NickName string `xml:"xdat:nickName"`
}

type AdmnStatus struct {
	Code   int     `xml:"xdat:code"`
	State  *string `xml:"xdat:state"`
	Reason string  `xml:"xdat:reason"`
	DateTo string  `xml:"xdat:date_to"`
}

type AdmnDeposit struct {
	Account string `xml:"xdat:account"`
}

type AdmnCardInfo struct {
	Bin     *string `xml:"xdat:bin"`
	Country *string `xml:"xdat:country"`
	Rrn     *string `xml:"xdat:rrn"`
	Srn     *string `xml:"xdat:srn"`
}

type AdmnCardSrc struct {
	Bin     *string `xml:"xdat:bin"`
	Country *string `xml:"xdat:country"`
}

type AdmnInform struct {
	MsgCode    string `xml:"xdat:msgcode"`
	MsgChannel string `xml:"xdat:msgchannel"`
	MsgProfile string `xml:"xdat:msgprofile"`
}

type AdmnExtFee struct {
	Code string `xml:"xdat:code"`
}

type AdmnNationalBankPayment struct {
	ProcessingMethod string           `xml:"xdat:processingMethod"`
	Beneficiary      *AdmnBeneficiary `xml:"xdat:beneficiary"`
	BeneficiaryBank  string           `xml:"xdat:beneficiaryBank"`
	PurposeCode      string           `xml:"xdat:purposeCode"`
	PaymentDetails   string           `xml:"xdat:paymentDetails"`
	DocumentNumber   string           `xml:"xdat:documentNumber"`
	Date             string           `xml:"xdat:date"`
}

type AdmnBeneficiary struct {
	AccountIban             string `xml:"xdat:accountIban"`
	Name                    string `xml:"xdat:name"`
	LegalIdentificationCode string `xml:"xdat:legalIdentificationCode"`
	PartyCode               string `xml:"xdat:partyCode"`
}

type AdmnWuTransfer struct {
	MTSSessionID   string `xml:"xdat:MTSSessionID"`
	Infl           string `xml:"xdat:Infl"`
	Name1Cr        string `xml:"xdat:Name1_Cr"`
	Name2Cr        string `xml:"xdat:Name2_Cr"`
	Name3Cr        string `xml:"xdat:Name3_Cr"`
	Name4Cr        string `xml:"xdat:Name4_Cr"`
	Question       string `xml:"xdat:Question"`
	Answer         string `xml:"xdat:Answer"`
	Message        string `xml:"xdat:Message"`
	MTSClientType  string `xml:"xdat:MTSClientType"`
	D2BFields      string `xml:"xdat:D2B_Fields"`
	Rnncr          string `xml:"xdat:rnn_cr"`
	ResidFl        string `xml:"xdat:ResidFl"`
	AccountIban    string `xml:"xdat:accountIban"`
	PurposePayee   string `xml:"xdat:purposePayee"`
	PartyCode      string `xml:"xdat:partyCode"`
	PurposeCode    string `xml:"xdat:purposeCode"`
	PaymentDetails string `xml:"xdat:paymentDetails"`
}

type AdmnGoldenCrown struct {
	UIN                     string `xml:"xdat:UIN"`
	Oid                     string `xml:"xdat:oid"`
	LastOperationDate       string `xml:"xdat:lastOperationDate"`
	Status                  string `xml:"xdat:status"`
	ToCountryISO            string `xml:"xdat:toCountryISO"`
	SendLocationID          string `xml:"xdat:sendLocationID"`
	FromCountryISO          string `xml:"xdat:fromCountryISO"`
	SenderlastName          string `xml:"xdat:sender_lastName"`
	SenderfirstName         string `xml:"xdat:sender_firstName"`
	SendermiddleName        string `xml:"xdat:sender_middleName"`
	SenderdocserialNumber   string `xml:"xdat:sender_doc_serialNumber"`
	Senderdocissuer         string `xml:"xdat:sender_doc_issuer"`
	SenderdocissueDate      string `xml:"xdat:sender_doc_issueDate"`
	SenderdocissuerCode     string `xml:"xdat:sender_doc_issuerCode"`
	SenderdocissueCity      string `xml:"xdat:sender_doc_issueCity"`
	SenderdocexpireDate     string `xml:"xdat:sender_doc_expireDate"`
	Senderdoctype           string `xml:"xdat:sender_doc_type"`
	Senderdocnumber         string `xml:"xdat:sender_doc_number"`
	Senderphone             string `xml:"xdat:sender_phone"`
	SendercountryISO        string `xml:"xdat:sender_countryISO"`
	SenderbirthDate         string `xml:"xdat:sender_birthDate"`
	SenderbirthPlace        string `xml:"xdat:sender_birthPlace"`
	SenderregCountry        string `xml:"xdat:sender_regCountry"`
	SenderregCity           string `xml:"xdat:sender_regCity"`
	SenderregAddress        string `xml:"xdat:sender_regAddress"`
	Sendertin               string `xml:"xdat:sender_tin"`
	Fundsamount             string `xml:"xdat:funds_amount"`
	Fundscurrency           string `xml:"xdat:funds_currency"`
	Commissionamount        string `xml:"xdat:commission_amount"`
	Commissioncurrency      string `xml:"xdat:commission_currency"`
	SenderAgentCommamount   string `xml:"xdat:senderAgentComm_amount"`
	SenderAgentCommcurrency string `xml:"xdat:senderAgentComm_currency"`
	MaskedCardPAN           string `xml:"xdat:maskedCardPAN"`
	ExchangeRate            string `xml:"xdat:exchangeRate"`
	PayFundsamount          string `xml:"xdat:payFunds_amount"`
	PayFundscurrency        string `xml:"xdat:payFunds_currency"`
	SendFundsamount         string `xml:"xdat:sendFunds_amount"`
	SendFundscurrency       string `xml:"xdat:sendFunds_currency"`
	FeeAgentamount          string `xml:"xdat:feeAgent_amount"`
	FeeAgentcurrency        string `xml:"xdat:feeAgent_currency"`
	ReceiverlastName        string `xml:"xdat:receiver_lastName"`
	ReceiverfirstName       string `xml:"xdat:receiver_firstName"`
	ReceivermiddleName      string `xml:"xdat:receiver_middleName"`
	Receiverphone           string `xml:"xdat:receiver_phone"`
}
