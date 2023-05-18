package test

import (
	"goapi/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)
type arg struct {
	key   string
	Value string
}
func TestDeleteKey(t *testing.T){
	testCases:=[]struct{
		name	string
		arg	arg
		key	string
		want	map[string]string
	}{
		{
			name: "deleting key-value from map",
			arg: arg{"talal","56"},
			key: "talal",
			want: map[string]string{},
		},
	}
	for _,tc:=range testCases{
		t.Run(tc.name,func(t *testing.T) {
			output:=controller.DeleteKeyAPI(tc.arg.key)
			assert.Equal(t,tc.want,output)
		})
	}
}