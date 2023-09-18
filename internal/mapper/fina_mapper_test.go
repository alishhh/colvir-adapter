package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/mappers"
	event_i "gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/externals/dbp-event/implementation"
)

func Test(t *testing.T) {
	data := `{"type":"SYNC","convertKaz":1,"body":{"FINA2590":{"channel":"ATM","reference":"5B6EAFDB33443C4567B65040B955067E","dt":"2022-08-04T10:00:46","operation":2590,"description":"\u041A\u043E\u043D\u0432\u0435\u0440\u0442\u0430\u0446\u0438\u044F \u0447\u0435\u0440\u0435\u0437 \u0431\u0430\u043D\u043A\u043E\u043C\u0430\u0442 KZT-USD","amount":"30000.00","fee":0,"fee_currency":"KZT","currency":"KZT","account":"00000000000000000000","account_type":"IBN","xData":{"trn":{"card":{"a":"50.0","c":"USD","dt":"2022-08-04T10:00:46","rnn":"","term":"00004778","country":"KZ","loc":"ALM","termtype":"ATM","psys":"LOCAL","pan":"000107500551"}}}}}}`
	old, _ := event_i.OldInputJson([]byte(data))
	pld := old.GetEvent().Payload
	mp := mappers.FinaMapper{}
	_, err := mp.Map(pld.(map[string]interface{}))
	assert.NotNil(t, err)

}
