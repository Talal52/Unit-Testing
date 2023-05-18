package test

import (
	"goapi/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKey(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		data map[string]string
		want string
	}{
		{
			name: "get map data by key",
			key:  "talal",
			data: map[string]string{"talal": ""},
			want: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//output := controller.GettingKeyAPI(tc.key)
			// assert.Equal(t, tc.want, tc.data[tc.key])
			output := controller.GettingKeyAPI(tc.key)
			assert.Equal(t, tc.want, output)
		})
	}
}
