package models

import (
	"encoding/xml"
)

type FinaRequest struct {
	XMLName xml.Name   `xml:"ws:finaRequest"`
	Text    string     `xml:",chardata"`
	Version string     `xml:"version,attr"`
	Header  FinaHeader `xml:"v3:header"`
	Body    FinaBody   `xml:"ws:body"`
}

type FinaHeader struct {
	Text      string          `xml:",chardata"`
	Channel   string          `xml:"v1:channel"`
	Reference string          `xml:"v1:reference"`
	Date      string          `xml:"v1:date"`
	Language  string          `xml:"v1:language"`
	Token     *FinaTypedField `xml:"v1:token,omitempty"`
	Timeout   *string         `xml:"v1:timeout,omitempty"`
}

type FinaBody struct {
	Text          string          `xml:",chardata"`
	Operation     string          `xml:"fina:operation"`
	Description   string          `xml:"fina:description"`
	Amount        FinaAmount      `xml:"fina:amount"`
	Fee           *FinaAmount     `xml:"fina:fee,omitempty"`
	Account       *FinaTypedField `xml:"fina:account,omitempty"`
	LinkedAccount *FinaTypedField `xml:"fina:linkedAccount,omitempty"`
	Account2      *FinaTypedField `xml:"fina:account2,omitempty"`
	XData         *FinaXData      `xml:"xdat:xData,omitempty"`
}

type FinaAmount struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type FinaTypedField struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type FinaXData struct {
	Trn *FinaTrn `xml:"xdat:trn,omitempty"`
}

type FinaTrn struct {
	Card                 *FinaCard                 `xml:"xdat:card,omitempty"`
	ExtFee               *FinaExtFee               `xml:"xdat:extfee,omitempty"`
	Demand               *FinaDemand               `xml:"xdat:demand,omitempty"`
	CardInfo             *FinaCardInfo             `xml:"xdat:card_info,omitempty"`
	CardSrc              *FinaCardSrc              `xml:"xdat:card_src,omitempty"`
	IntraAccountPayment  *FinaIntraAccountPayment  `xml:"xdat:intra_account_payment,omitempty"`
	NationalBankPayment  *FinaNationalBankPayment  `xml:"xdat:national_bank_payment,omitempty"`
	FastTransfer         *FinaFastTransfer         `xml:"xdat:fast_transfer,omitempty"`
	InternationalPayment *FinaInternationalPayment `xml:"xdat:international_payment,omitempty"` //unused
	Inform               *FinaInform               `xml:"xdat:inform,omitempty"`                //unused
	Payment              *FinaPayment              `xml:"xdat:Payment,omitempty"`
}

type FinaCard struct {
	A        string  `xml:"xdat:a,omitempty"`
	C        string  `xml:"xdat:c,omitempty"`
	Ba       *string `xml:"xdat:ba,omitempty"`
	Bc       *string `xml:"xdat:bc,omitempty"`
	Pan      *string `xml:"xdat:pan,omitempty"`
	Panm     string  `xml:"xdat:panm,omitempty"`
	Dt       string  `xml:"xdat:dt,omitempty"`
	Sttldt   *string `xml:"xdat:sttldt,omitempty"`
	Rrn      string  `xml:"xdat:rrn,omitempty"`
	Auth     *string `xml:"xdat:auth,omitempty"`
	Iss      *string `xml:"xdat:iss,omitempty"`
	Mcc      string  `xml:"xdat:mcc,omitempty"`
	Acq      *string `xml:"xdat:acq,omitempty"`
	Acqc     *string `xml:"xdat:acqc,omitempty"`
	Term     *string `xml:"xdat:term,omitempty"`
	Caid     *string `xml:"xdat:caid,omitempty"`
	Loc      string  `xml:"xdat:loc,omitempty"`
	Country  *string `xml:"xdat:country,omitempty"`
	City     *string `xml:"xdat:city,omitempty"`
	Adr      *string `xml:"xdat:adr,omitempty"`
	TermType string  `xml:"xdat:termtype,omitempty"`
	Psys     string  `xml:"xdat:psys,omitempty"`
	Posmode  *string `xml:"xdat:posmode,omitempty"`
	Eci      *string `xml:"xdat:eci,omitempty"`
	Cond     *string `xml:"xdat:cond,omitempty"`
	McName   *string `xml:"xdat:mcname,omitempty"`
}

type FinaExtFee struct {
	Code string `xml:"xdat:code"`
}

type FinaDemand struct {
	Deposit FinaDemandDeposit `xml:"xdat:deposit"`
}

type FinaDemandDeposit struct {
	Account string `xml:"xdat:account"`
}

type FinaCardInfo struct {
	Bin     *string `xml:"xdat:bin,omitempty"`
	Country *string `xml:"xdat:country,omitempty"`
	Rrn     *string `xml:"xdat:rrn,omitempty"`
	Srn     *string `xml:"xdat:srn,omitempty"`
}

type FinaCardSrc struct {
	Bin     *string `xml:"xdat:bin,omitempty"`
	Country *string `xml:"xdat:country,omitempty"`
	Rrn     *string `xml:"xdat:rrn,omitempty"`
	Srn     *string `xml:"xdat:srn,omitempty"`
}

type FinaIntraAccountPayment struct {
	Beneficiary    *FinaIPABeneficiary `xml:"xdat:beneficiary,omitempty"`
	PaymentDetails *string             `xml:"xdat:paymentDetails,omitempty"`
}

type FinaNationalBankPayment struct {
	ProcessingMethod string             `xml:"xdat:processingMethod"`
	Beneficiary      FinaNBPBeneficiary `xml:"xdat:beneficiary"`
	BeneficiaryBank  string             `xml:"xdat:beneficiaryBank"`
	PurposeCode      string             `xml:"xdat:purposeCode"`
	PaymentDetails   string             `xml:"xdat:paymentDetails"`
	DocumentNumber   *string            `xml:"xdat:documentNumber,omitempty"`
	Date             *string            `xml:"xdat:date,omitempty"`
}

type FinaIPABeneficiary struct {
	Account                 string  `xml:"xdat:account"`
	LegalIdentificationCode *string `xml:"xdat:legalIdentificationCode,omitempty"`
}

type FinaNBPBeneficiary struct {
	AccountIban             string `xml:"xdat:accountIban"`
	Name                    string `xml:"xdat:name"`
	LegalIdentificationCode string `xml:"xdat:legalIdentificationCode"`
	PartyCode               string `xml:"xdat:partyCode"`
}

// объявленные, но не используемые структуры
type FinaInternationalPayment struct {
}

type FinaFastTransfer struct {
	Beneficiary  FinaFFBeneficiary `xml:"xdat:beneficiary"`
	DocId        string            `xml:"xdat:doc_id"`
	Date         string            `xml:"xdat:date"`
	State        string            `xml:"xdat:state"`
	FeeAccount   string            `xml:"xdat:fee_account"`
	Order        string            `xml:"xdat:order"`
	AuthorStatus string            `xml:"xdat:author_status"`
	Type         string            `xml:"xdat:type"`
}

type FinaFFBeneficiary struct {
	Bin        string `xml:"xdat:bin"`
	Department string `xml:"xdat:department"`
	Rnn        string `xml:"xdat:rnn"`
	Name       string `xml:"xdat:name"`
	Resident   string `xml:"xdat:resident"`
	Kpp        string `xml:"xdat:kpp"`
	Purpose    string `xml:"xdat:purpose"`
}

type FinaInform struct {
}

type FinaPayment struct {
	TerminalType      string              `xml:"xdat:TerminalType,omitempty"`
	Knp               string              `xml:"xdat:knp,omitempty"`
	Provider          FinaPaymentProvider `xml:"xdat:Provider,omitempty"`
	AgentAccount      string              `xml:"xdat:AgentAccount,omitempty"`
	PaymentIdentifier string              `xml:"xdat:paymentIdentifier,omitempty"`
	Kno               string              `xml:"xdat:kno,omitempty"`
	Kbe               string              `xml:"xdat:kbe,omitempty"`
}

type FinaPaymentProvider struct {
	BIC                string `xml:"xdat:BIC,omitempty"`
	Name               string `xml:"xdat:name,omitempty"`
	BIN                string `xml:"xdat:BIN,omitempty"`
	ContractType       string `xml:"xdat:ContractType,omitempty"`
	ProviderAccount    string `xml:"xdat:ProviderAccount,omitempty"`
	ProviderAccountTrn string `xml:"xdat:ProviderAccountTrn,omitempty"`
	ProviderNetting    string `xml:"xdat:ProviderNetting,omitempty"`
	ProviderFee        string `xml:"xdat:ProviderFee,omitempty"`
	FeeType            string `xml:"xdat:FeeType,omitempty"`
}
