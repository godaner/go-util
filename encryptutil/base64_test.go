package encryptutil

import "testing"

func TestEncryptBase64(t *testing.T) {
	tests := []struct {
		in   string
	}{
		{
			"",
		},
		{
			"a123",
		},
		{
			"aassssssssss46sa4d6a4d64a56d4a564d56a4d6a456da46d4a6d46a4d56a46d46a4d4ad46a54d894a89wqewqeq",
		},
	}

	for _,tt :=range tests  {
		var s string
		if s=EncodeBase64(tt.in);s==""{
			t.Error("EncodeBase64 is error !")
		}else{
			t.Log("base64 is : ",s," , len is : ",len(s))
		}
		if s=DecodeBase64(s);tt.in!=s{
			t.Error("DecodeBase64 is error !")
		}else{
			t.Log("origin str is : ",s," , len is : ",len(s))
		}

	}
}