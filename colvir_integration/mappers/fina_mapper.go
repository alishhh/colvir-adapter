package mappers

import (
	"bytes"
	"encoding/xml"
	"math"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/soap"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/helpers/typeconv"
)

type FinaMapper struct{}

func (fm *FinaMapper) Map(src map[string]interface{}) ([]byte, error) {
	channel := makeStringValue(src["channel"])
	operation := strings.Replace(makeStringValue(src["operation"]), ".00", "", -1)
	description := makeStringValue(src["description"])
	account := makeStringValue(src["account"])
	accType := makeStringValue(src["account_type"])
	reference := makeStringValue(src["reference"])
	timeout := makeInt(src["timeout"])
	var timeoutS *string
	if timeout > 0 {
		s := (strconv.FormatInt(timeout, 10))
		timeoutS = &s
	}
	account2 := makeStringValue(src["account2"])
	accType2 := makeStringValue(src["account2_type"])

	amount := typeconv.MakeString(src["amount"])
	amountCur := makeStringValue(src["currency"])

	fee := typeconv.MakeFloat64(src["fee"])
	feeCur := typeconv.MakeString(src["fee_currency"])
	if feeCur == "" {
		feeCur = amountCur
	}

	accountLinked := makeStringValue(src["linked_account"])
	accTypeLinked := makeStringValue(src["linked_account_type"])
	retObj := &models.FinaRequest{
		Version: getVersion(operation),
		Header: models.FinaHeader{
			Channel:   channel,
			Language:  getLang(operation),
			Reference: reference,
			Date:      time.Now().Format("2006-01-02T15:04:05"),
			Timeout:   timeoutS,
		},
		Body: models.FinaBody{
			Operation:   operation,
			Description: description,
			Amount: models.FinaAmount{
				Text:     amount,
				Currency: amountCur,
			},
		},
	}

	if !math.IsNaN(fee) && feeCur != "" {
		retObj.Body.Fee = &models.FinaAmount{
			Text:     strconv.FormatFloat(fee, 'f', 2, 64),
			Currency: feeCur,
		}
	}

	if account != "" && accType != "" {
		retObj.Body.Account = &models.FinaTypedField{
			Text: account,
			Type: accType,
		}
	}

	if account2 != "" && accType2 != "" {
		retObj.Body.Account2 = &models.FinaTypedField{
			Text: account2,
			Type: accType2,
		}
	}
	if accountLinked != "" && accTypeLinked != "" {
		retObj.Body.LinkedAccount = &models.FinaTypedField{
			Text: accountLinked,
			Type: accTypeLinked,
		}
	}
	if src["xData"] != nil {
		xdata, ok := src["xData"].(map[string]interface{})
		if ok {
			retObj.Body.XData = fulfillXData(xdata)
		}
	}
	bd, _ := xml.Marshal(retObj)
	tmpl, err := template.New("soap").Parse(string(soap.XML()))
	if err != nil {
		return nil, err
	}
	var ret bytes.Buffer
	err = tmpl.Execute(
		&ret, map[string]string{
			"action":     `http://bus.colvir.com/service/cap/v3/FINA/FINA` + operation,
			"body":       string(bd),
			"messageID":  uuid.NewString(),
			"namespaces": `xmlns:fina="http://bus.colvir.com/service/cap/v3/FINA" xmlns:ws="http://bus.colvir.com/service/cap/v3/ws" xmlns:v3="http://bus.colvir.com/service/cap/v3" xmlns:xdat="http://bus.colvir.com/service/cap/v3/FINA/xdata"`,
		})
	if err != nil {
		return nil, err
	}
	return ret.Bytes(), nil
}

func (fm *FinaMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "fina")
}

func getVersion(operation string) string {
	ver := "3.0.3"
	if operation == "2512" || operation == "2517" {
		ver = "3.0.2"
	}
	return ver
}

func getLang(operation string) string {
	lang := "RU"
	if operation == "2512" || operation == "2517" {
		lang = "EN"
	}
	return lang
}

func fulfillXData(src map[string]interface{}) *models.FinaXData {
	ret := &models.FinaXData{}
	if src["trn"] != nil {
		if trn, ok := src["trn"].(map[string]interface{}); ok {
			ret.Trn = &models.FinaTrn{}
			if trn["card"] != nil {
				if card, ok := trn["card"].(map[string]interface{}); ok {
					ret.Trn.Card = fulfillCard(card)
				}
			}
			if trn["cardinfo"] != nil {
				if card, ok := trn["cardinfo"].(map[string]interface{}); ok {
					ret.Trn.CardInfo = fulfillCardInfo(card)
				}
			}
			if trn["cardsrc"] != nil {
				if card, ok := trn["cardsrc"].(map[string]interface{}); ok {
					ret.Trn.CardSrc = fulfillCardSrc(card)
				}
			}

			if trn["extfee"] != nil {
				if extfee, ok := trn["extfee"].(map[string]interface{}); ok {
					ret.Trn.ExtFee = fulfillExtFee(extfee)
				}
			}

			if trn["demand"] != nil {
				if demand, ok := trn["demand"].(map[string]interface{}); ok {
					ret.Trn.Demand = fulfillDemand(demand)
				}
			}

			if trn["intra_account_payment"] != nil {
				if iap, ok := trn["intra_account_payment"].(map[string]interface{}); ok {
					ret.Trn.IntraAccountPayment = fulfillIntraPayment(iap)
				}
			}

			if trn["national_bank_payment"] != nil {
				if nbp, ok := trn["national_bank_payment"].(map[string]interface{}); ok {
					ret.Trn.NationalBankPayment = fulfillNBP(nbp)
				}
			}

			if trn["fast_transfer"] != nil {
				if ff, ok := trn["fast_transfer"].(map[string]interface{}); ok {
					ret.Trn.FastTransfer = fulfillFastTransfer(ff)
				}
			}

			if trn["payment"] != nil {
				if p, ok := trn["payment"].(map[string]interface{}); ok {
					ret.Trn.Payment = fulfillPayment(p)
				}
			}
		}
	}
	return ret
}

func fulfillCard(card map[string]interface{}) *models.FinaCard {
	ret := &models.FinaCard{}
	ret.A = makeStringValue(card["a"])
	ret.C = makeStringValue(card["c"])
	ret.Panm = makeStringValue(card["panm"])
	ret.Dt = makeStringValue(card["dt"])
	ret.Rrn = makeStringValue(card["rrn"])
	ret.Mcc = makeStringValue(card["mcc"])
	ret.Loc = makeStringValue(card["loc"])
	ret.TermType = makeStringValue(card["termtype"])
	ret.Psys = makeStringValue(card["psys"])

	ret.Ba = makeStringValuePtr(card["ba"])
	ret.Bc = makeStringValuePtr(card["bc"])
	ret.Pan = makeStringValuePtr(card["pan"])
	ret.Sttldt = makeStringValuePtr(card["sttldt"])
	ret.Auth = makeStringValuePtr(card["auth"])
	ret.Iss = makeStringValuePtr(card["iss"])
	ret.Acq = makeStringValuePtr(card["acq"])
	ret.Acqc = makeStringValuePtr(card["acqc"])
	ret.Term = makeStringValuePtr(card["term"])
	ret.Caid = makeStringValuePtr(card["caid"])
	ret.Country = makeStringValuePtr(card["country"])
	ret.City = makeStringValuePtr(card["city"])
	ret.Adr = makeStringValuePtr(card["adr"])
	ret.Posmode = makeStringValuePtr(card["posmode"])
	ret.Eci = makeStringValuePtr(card["eci"])
	ret.Cond = makeStringValuePtr(card["cond"])
	ret.McName = makeStringValuePtr(card["mcname"])
	return ret
}

func fulfillCardInfo(info map[string]interface{}) *models.FinaCardInfo {
	ret := &models.FinaCardInfo{}

	ret.Bin = makeStringValuePtr(info["bin"])
	ret.Country = makeStringValuePtr(info["country"])
	ret.Rrn = makeStringValuePtr(info["rrn"])
	ret.Srn = makeStringValuePtr(info["srn"])

	return ret
}

func fulfillCardSrc(src map[string]interface{}) *models.FinaCardSrc {
	ret := &models.FinaCardSrc{}

	ret.Bin = makeStringValuePtr(src["bin"])
	ret.Country = makeStringValuePtr(src["country"])
	ret.Rrn = makeStringValuePtr(src["rrn"])
	ret.Srn = makeStringValuePtr(src["srn"])

	return ret
}

func fulfillIntraPayment(src map[string]interface{}) *models.FinaIntraAccountPayment {
	ret := &models.FinaIntraAccountPayment{}
	if src["beneficiary"] != nil {
		if benf, ok := src["beneficiary"].(map[string]interface{}); ok {
			ret.Beneficiary = &models.FinaIPABeneficiary{}
			ret.Beneficiary.Account = makeStringValue(benf["account"])
			ret.Beneficiary.LegalIdentificationCode = makeStringValuePtr("legal_identification_code")
		}

	}

	ret.PaymentDetails = makeStringValuePtr(src["payment_details"])

	return ret
}

func fulfillDemand(src map[string]interface{}) *models.FinaDemand {
	var ret *models.FinaDemand
	if dep, ok := src["deposit"].(map[string]interface{}); ok {
		ret = &models.FinaDemand{}
		ret.Deposit = models.FinaDemandDeposit{
			Account: makeStringValue(dep["account"]),
		}
	}
	return ret
}

func fulfillNBP(src map[string]interface{}) *models.FinaNationalBankPayment {
	ret := &models.FinaNationalBankPayment{}
	ret.ProcessingMethod = makeStringValue(src["processingMethod"])
	ret.PurposeCode = makeStringValue(src["purposeCode"])
	ret.PaymentDetails = makeStringValue(src["paymentDetails"])
	ret.BeneficiaryBank = makeStringValue(src["beneficiaryBank"])
	ret.DocumentNumber = makeStringValuePtr(src["documentNumber"])
	ret.Date = makeStringValuePtr(src["purposeCode"])
	if src["beneficiary"] != nil {
		if benf, ok := src["beneficiary"].(map[string]interface{}); ok {
			ret.Beneficiary = models.FinaNBPBeneficiary{
				AccountIban:             makeStringValue(benf["accountIban"]),
				LegalIdentificationCode: makeStringValue(benf["legalIdentificationCode"]),
				PartyCode:               makeStringValue(benf["partyCode"]),
			}
		}
	}

	return ret
}

func fulfillExtFee(src map[string]interface{}) *models.FinaExtFee {
	ret := &models.FinaExtFee{
		Code: makeStringValue(src["code"]),
	}
	return ret
}

func fulfillFastTransfer(src map[string]interface{}) *models.FinaFastTransfer {
	ret := &models.FinaFastTransfer{
		DocId:        makeStringValue(src["doc_id"]),
		Date:         makeStringValue(src["date"]),
		State:        makeStringValue(src["state"]),
		FeeAccount:   makeStringValue(src["fee_account"]),
		Order:        makeStringValue(src["order"]),
		AuthorStatus: makeStringValue(src["author_status"]),
		Type:         makeStringValue(src["type"]),
	}
	if src["beneficiary"] != nil {
		if bff, ok := src["beneficiary"].(map[string]interface{}); ok {
			ret.Beneficiary = models.FinaFFBeneficiary{
				Bin:        makeStringValue(bff["bin"]),
				Department: makeStringValue(bff["department"]),
				Rnn:        makeStringValue(bff["rnn"]),
				Name:       makeStringValue(bff["name"]),
				Resident:   makeStringValue(bff["resident"]),
				Kpp:        makeStringValue(bff["kpp"]),
				Purpose:    makeStringValue(bff["purpose"]),
			}
		}
	}

	return ret
}

func fulfillPayment(src map[string]interface{}) *models.FinaPayment {
	ret := &models.FinaPayment{
		TerminalType:      makeStringValue(src["terminalType"]),
		Knp:               makeStringValue(src["knp"]),
		AgentAccount:      makeStringValue(src["agentAccount"]),
		PaymentIdentifier: makeStringValue(src["paymentIdentifier"]),
		Kno:               makeStringValue(src["kno"]),
		Kbe:               makeStringValue(src["kbe"]),
	}
	if src["provider"] != nil {
		if provider, ok := src["provider"].(map[string]interface{}); ok {
			ret.Provider = models.FinaPaymentProvider{
				BIC:                makeStringValue(provider["bic"]),
				Name:               makeStringValue(provider["name"]),
				BIN:                makeStringValue(provider["bin"]),
				ContractType:       makeStringValue(provider["contractType"]),
				ProviderAccount:    makeStringValue(provider["providerAccount"]),
				ProviderAccountTrn: makeStringValue(provider["providerAccountTrn"]),
				ProviderNetting:    makeStringValue(provider["providerNetting"]),
				ProviderFee:        makeStringValue(provider["providerFee"]),
				FeeType:            makeStringValue(provider["feeType"]),
			}
		}
	}

	return ret
}
