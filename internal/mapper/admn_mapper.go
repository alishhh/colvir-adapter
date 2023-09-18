package mapper

import (
	"bytes"
	"encoding/xml"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/soap"
)

type AdmnMapper struct{}

func (am *AdmnMapper) Map(src map[string]interface{}) ([]byte, error) {
	channel := makeStringValue(src["channel"])
	reference := makeStringValue(src["reference"])
	operation := strings.Replace(makeStringValue(src["operation"]), ".00", "", -1)
	description := makeStringValue(src["description"])

	retObj := &models.AdmnRequest{
		Version: "3.0",
		Header: models.AdmnHeader{
			Channel:   channel,
			Language:  "ru",
			Reference: reference,
			Date:      time.Now().Format("2006-01-02T15:04:05"),
		},
		Body: models.AdmnBody{
			Operation:   operation,
			Description: description,
		},
	}

	if src["account"] != nil {
		acc, ok := src["account"].(map[string]interface{})
		if ok {
			account := makeStringValue(acc["value"])
			accType := makeStringValue(acc["type"])
			retObj.Body.Account = &models.AdmnTypedField{
				Text: account,
				Type: accType,
			}
		}
	}

	if src["amount"] != nil {
		srcB := src["amount"].(map[string]interface{})
		amount := makeStringValue(srcB["value"])
		currency := makeStringValue(srcB["currency"])
		retObj.Body.Amount = &models.AdmnAmount{
			Text:     amount,
			Currency: currency,
		}
	}

	if src["account2"] != nil {
		srcC := src["account2"].(map[string]interface{})
		account2 := makeStringValue(srcC["value"])
		accType2 := makeStringValue(srcC["type"])
		retObj.Body.Account2 = &models.AdmnTypedField{
			Text: account2,
			Type: accType2,
		}
	}

	if src["xData"] != nil {
		xdata, ok := src["xData"].(map[string]interface{})
		if ok {
			retObj.Body.XData = fulfillXDataAdmn(xdata)
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
			"action":     `http://bus.colvir.com/service/cap/v3/ADMN/ADMN` + operation,
			"body":       string(bd),
			"messageID":  uuid.NewString(),
			"namespaces": `xmlns:admn="http://bus.colvir.com/service/cap/v3/ADMN" xmlns:ws="http://bus.colvir.com/service/cap/v3/ws" xmlns:v3="http://bus.colvir.com/service/cap/v3" xmlns:xdat="http://bus.colvir.com/service/cap/v3/ADMN/xdata"`,
		})
	if err != nil {
		return nil, err
	}

	return ret.Bytes(), nil
}

func (am *AdmnMapper) IsAppliable(command string) bool {
	return strings.HasPrefix(strings.ToLower(command), "admn")
}

func fulfillXDataAdmn(src map[string]interface{}) *models.AdmnXData {
	ret := &models.AdmnXData{}
	if src["trn"] != nil {
		if trn, ok := src["trn"].(map[string]interface{}); ok {
			ret.Trn = &models.AdmnTrn{}
			if trn["intra_account_payment"] != nil {
				if iap, ok := trn["intra_account_payment"].(map[string]interface{}); ok {
					ret.Trn.IntraAccountPayment = fulfillIntraPaymentAdmn(iap)
				}
			}

			if trn["card"] != nil {
				if card, ok := trn["card"].(map[string]interface{}); ok {
					ret.Trn.Card = fulfillCardAdmn(card)
				}
			}

			if trn["demand"] != nil {
				if demand, ok := trn["demand"].(map[string]interface{}); ok {
					ret.Trn.Demand = fulfillDemandAdmn(demand)
				}
			}

			if trn["cardinfo"] != nil {
				if card, ok := trn["cardinfo"].(map[string]interface{}); ok {
					ret.Trn.CardInfo = fulfillCardInfoAdmn(card)
				}
			}

			if trn["cardsrc"] != nil {
				if card, ok := trn["cardsrc"].(map[string]interface{}); ok {
					ret.Trn.CardSrc = fulfillCardSrcAdmn(card)
				}
			}

			if trn["inform"] != nil {
				if in, ok := trn["inform"].(map[string]interface{}); ok {
					ret.Trn.IntraAccountPayment = fulfillIntraPaymentAdmn(in)
				}
			}

			if trn["wu_transfer"] != nil {
				if wt, ok := trn["wu_transfer"].(map[string]interface{}); ok {
					ret.Trn.WuTransfer = fulfillWuTransferAdmn(wt)
				}
			}

			if trn["extfee"] != nil {
				if extfee, ok := trn["extfee"].(map[string]interface{}); ok {
					ret.Trn.ExtFee = fulfillExtFeeAdmn(extfee)
				}
			}

			if trn["national_bank_payment"] != nil {
				if nbp, ok := trn["national_bank_payment"].(map[string]interface{}); ok {
					ret.Trn.NationalBankPayment = fulfillNBPAdmn(nbp)
				}
			}

			if trn["installment"] != nil {
				if ff, ok := trn["installment"].(map[string]interface{}); ok {
					ret.Trn.Installment = makeStringValuePtr(ff)
				}
			}

		}
	}
	return ret
}

func fulfillCardAdmn(card map[string]interface{}) *models.AdmnCard {
	ret := &models.AdmnCard{}
	ret.A = makeStringValue(card["a"])
	ret.C = makeStringValue(card["c"])
	ret.Fa = makeStringValue(card["fa"])
	ret.Aa = makeStringValue(card["aa"])
	ret.Ac = makeStringValue(card["ac"])
	ret.Ba = makeStringValue(card["ba"])
	ret.Bc = makeStringValue(card["bc"])
	ret.Sa = makeStringValue(card["sa"])
	ret.Pan = makeStringValue(card["pan"])
	ret.Panm = makeStringValue(card["panm"])
	ret.Dt = makeStringValue(card["dt"])
	ret.Sttldt = makeStringValue(card["sttldt"])
	ret.Rrn = makeStringValue(card["rrn"])
	ret.Auth = makeStringValue(card["auth"])
	ret.Iss = makeStringValue(card["iss"])
	ret.Acq = makeStringValue(card["acq"])
	ret.Mcc = makeStringValue(card["mcc"])
	ret.Acqc = makeStringValue(card["acqc"])
	ret.Term = makeStringValue(card["term"])
	ret.Caid = makeStringValue(card["caid"])
	ret.Loc = makeStringValue(card["loc"])
	ret.Country = makeStringValue(card["country"])
	ret.Adr = makeStringValue(card["adr"])
	ret.Termtype = makeStringValue(card["termtype"])
	ret.Psys = makeStringValue(card["psys"])
	ret.Stan = makeStringValue(card["stan"])
	ret.Ft = makeStringValue(card["ft"])
	ret.Posmode = makeStringValue(card["posmode"])
	ret.Phonenumber = makeStringValue(card["phone_number"])
	return ret
}

func fulfillCardInfoAdmn(info map[string]interface{}) *models.AdmnCardInfo {
	ret := &models.AdmnCardInfo{}
	ret.Bin = makeStringValuePtr(info["bin"])
	ret.Country = makeStringValuePtr(info["country"])
	ret.Rrn = makeStringValuePtr(info["rrn"])
	ret.Srn = makeStringValuePtr(info["srn"])
	return ret
}

func fulfillCardSrcAdmn(src map[string]interface{}) *models.AdmnCardSrc {
	ret := &models.AdmnCardSrc{}
	ret.Bin = makeStringValuePtr(src["bin"])
	ret.Country = makeStringValuePtr(src["country"])
	return ret
}

func fulfillIntraPaymentAdmn(src map[string]interface{}) *models.AdmnIntraAccountPayment {
	ret := &models.AdmnIntraAccountPayment{}
	ret.PaymentDetails = makeStringValuePtr(src["payment_details"])
	return ret
}

func fulfillDemandAdmn(src map[string]interface{}) *models.AdmnDemand {
	ret := &models.AdmnDemand{}
	if cust, ok := src["customerInfo"].(map[string]interface{}); ok {
		ret.CustomerInfo = &models.AdmnCustomerInfo{
			Email:  makeStringValue(cust["email"]),
			Mobile: makeStringValue(cust["mobile"]),
			Locale: makeStringValue(cust["locale"]),
		}
	}
	if acc, ok := src["accountInfo"].(map[string]interface{}); ok {
		ret.AccountInfo = &models.AdmnAccountInfo{
			NickName: makeStringValue(acc["nickName"]),
		}
	}
	if stat, ok := src["status"].(map[string]interface{}); ok {
		ret.Status = &models.AdmnStatus{
			Code:   makeIntValue(stat["code"]),
			State:  makeStringValuePtr(stat["state"]),
			Reason: makeStringValue(stat["reason"]),
			DateTo: makeStringValue(stat["date_to"]),
		}
	}
	if dep, ok := src["deposit"].(map[string]interface{}); ok {
		ret.Deposit = &models.AdmnDeposit{
			Account: makeStringValue(dep["account"]),
		}
	}
	return ret
}

func fulfillNBPAdmn(src map[string]interface{}) *models.AdmnNationalBankPayment {
	ret := &models.AdmnNationalBankPayment{}
	ret.ProcessingMethod = makeStringValue(src["processingMethod"])
	ret.PurposeCode = makeStringValue(src["purposeCode"])
	ret.PaymentDetails = makeStringValue(src["paymentDetails"])
	ret.BeneficiaryBank = makeStringValue(src["beneficiaryBank"])
	ret.DocumentNumber = makeStringValue(src["documentNumber"])
	ret.Date = makeStringValue(src["purposeCode"])
	ret.BeneficiaryBank = makeStringValue(src["beneficiaryBank"])
	if src["beneficiary"] != nil {
		if benf, ok := src["beneficiary"].(map[string]interface{}); ok {
			ret.Beneficiary = &models.AdmnBeneficiary{
				AccountIban:             makeStringValue(benf["accountIban"]),
				Name:                    makeStringValue(benf["name"]),
				LegalIdentificationCode: makeStringValue(benf["legalIdentificationCode"]),
				PartyCode:               makeStringValue(benf["partyCode"]),
			}
		}
	}

	return ret
}

func fulfillExtFeeAdmn(src map[string]interface{}) *models.AdmnExtFee {
	ret := &models.AdmnExtFee{
		Code: makeStringValue(src["code"]),
	}
	return ret
}

func fulfillWuTransferAdmn(src map[string]interface{}) *models.AdmnWuTransfer {
	ret := &models.AdmnWuTransfer{}
	ret.MTSSessionID = makeStringValue(src["MTSSessionID"])
	ret.Infl = makeStringValue(src["Infl"])
	ret.Name1Cr = makeStringValue(src["Name1_Cr"])
	ret.Name2Cr = makeStringValue(src["Name2_Cr"])
	ret.Name3Cr = makeStringValue(src["Name3_Cr"])
	ret.Name4Cr = makeStringValue(src["Name4_Cr"])
	ret.Question = makeStringValue(src["Question"])
	ret.Answer = makeStringValue(src["Answer"])
	ret.Message = makeStringValue(src["Message"])
	ret.MTSClientType = makeStringValue(src["MTSClientType"])
	ret.D2BFields = makeStringValue(src["D2B_Fields"])
	ret.Rnncr = makeStringValue(src["rnn_cr"])
	ret.ResidFl = makeStringValue(src["ResidFl"])
	ret.AccountIban = makeStringValue(src["accountIban"])
	ret.PurposePayee = makeStringValue(src["purposePayee"])
	ret.PartyCode = makeStringValue(src["partyCode"])
	ret.PurposeCode = makeStringValue(src["purposeCode"])
	ret.PaymentDetails = makeStringValue(src["paymentDetails"])
	return ret
}
