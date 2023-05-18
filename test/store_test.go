package test

import (
	"goapi/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreKey(t *testing.T) {
	type args struct {
		key   string
		Value string
	}
	testCases := []struct {
		name string

		args args

		want map[string]string
		
	}{
		{
			name: "store map data by key",

			args: args{"talal", "tala"},
			
			want: map[string]string{"talal": "tala"},

		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			output := controller.StoringDataAPI(tc.args.key, tc.args.Value)
			assert.Equal(t, tc.want, output)
		})
	}
}
