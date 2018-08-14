package encryptutil

import (
	"testing"
)

func TestEncryptMd5(t *testing.T) {
	tests := []struct {
		in   string
	}{
		{
			"a",
		},
		{
			"aassssssssss46sa4d6a4d64a56d4a564d56a4d6a456da46d4a6d46a4d56a46d46a4d4ad46a54d894a89wqewqeq",
		},
	}

	for _,tt :=range tests  {
		if s:=EncryptMd5(tt.in);s==""{
			t.Error("str is empty !")
		}else{
			t.Log("md5 is : ",s," , len is : ",len(s))
		}
	}
}