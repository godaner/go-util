package encryptutil

import (
	"testing"
)

func TestEncryptDes(t *testing.T) {
	tests := []struct {
		text   string
		key   []byte
	}{
		{
			"a",
			[]byte("aasd897s7"),
		},
		{
			"a",
			[]byte("aasd8977"),
		},
		{
			"a",
			[]byte("a"),
		},
	}

	for _,tt :=range tests  {
		var s string
		if s,_=EncryptDes(tt.text,tt.key);s==""{
			t.Error("EncryptDes is error !")
		}else{
			t.Log("text is : ",s," , len is : ",len(s))
		}
		if s,_=DecryptDes(s,tt.key);tt.text!=s{
			t.Error("DecryptDes is error !")
		}else{
			t.Log("origin str is : ",s," , len is : ",len(s))
		}

	}
}