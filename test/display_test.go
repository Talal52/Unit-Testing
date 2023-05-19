package test

import (
	"goapi/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayKeys(t *testing.T){
	testCases:=[]struct{
		name	string
		want	map[string]string

	}{
		{
			name: "display the map",
			want: map[string]string{"talal":"56"},
		},
	}
	for _,tc:=range testCases{
		t.Run(tc.name,func (t *testing.T)  {
			output:=controller.DisplayingKeyAPI()
			assert.Equal(t,tc.want,output)
		})
	}
}