package mappers_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/mappers"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		desc string
		data []byte
	}{
		{
			desc: "non-empty list with clientCode as filter",
			data: []byte(`{"clientsListQuery":{"equals":[{"attr":"clientCode","value":"07296932"},{"attr":"clientCode","value":"123"}]}}`),
		},
		//{
		//	desc: "empty list",
		//	data: []byte(`{"clientsListQuery":{"equals":[{}]}}`),
		//},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mp := mappers.LoadClientsListMapper{}
			src := map[string]interface{}{}
			json.Unmarshal(tC.data, &src)
			resp, err := mp.Map(src)
			assert.Nil(t, err, "deserialization error should be nil")
			assert.Nil(t, resp)
		})
	}
}
