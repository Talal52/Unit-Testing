package test

import (
	"goapi/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)
type args struct {
	key   string
	Value string
}
func TestUpdateKey(t *testing.T){
	testCases:=[]struct{
		name	string
		agrs	args
		key		string
		want	map[string]string
	}{
		{
			name: "update the value",
			agrs: args{"talal","56"},
			key: "10",
			want: map[string]string{"talal":"10"},
		},
	}
	for _,tc:=range testCases{
		t.Run(tc.name,func(t *testing.T){
			output:=controller.UpdatingKeyAPI(tc.agrs.key,tc.agrs.Value,tc.key)
			assert.Equal(t,tc.want,output)
		})
	}
}