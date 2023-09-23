package encrypt

import (
	"testing"
)

func TestEncryptMobile(t *testing.T) {
	mobile := "13800138000"
	encryptedMobile, err := EncMobile(mobile)
	if err != nil {
		t.Fatal(err)
	}
	decryptedMobile, err := DecMobile(encryptedMobile)
	if err != nil {
		t.Fatal(err)
	}
	if mobile != decryptedMobile {
		t.Fatalf("expected %s, but got %s", mobile, decryptedMobile)
	}
}
