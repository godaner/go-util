package encryptutil

import (
	"testing"
)

func TestEncryptAes(t *testing.T) {
	tests := []struct {
		text   string
		key   []byte
	}{
		{
			"a",
			[]byte("aasd8977"),
		},
	}

	for _,tt :=range tests  {

		var s string
		if s,_=EncryptAes(tt.text,tt.key);s==""{
			t.Error("EncryptAes is error !")
		}else{
			t.Log("text is : ",s," , len is : ",len(s))
		}
		if s,_=DecryptAes(s,tt.key);tt.text!=s{
			t.Error("DecryptAes is error !")
		}else{
			t.Log("origin str is : ",s," , len is : ",len(s))
		}
	}
}