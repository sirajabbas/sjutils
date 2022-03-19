package utilsgo

import "testing"

func TestHMACSha256(t *testing.T) {
	t.Log("HMAC sha 256 test")
	op := HMACSha256("sirajabbas", "mysecret")
	if len(op) == 0 || op != "3f834c5e773660671d576d3990d797f04c5cba41d02e340a6bcf43ae2bab909c" {
		t.Fatal("HMAC sha 256  failed")
	}
	t.Log("HMAC sha 256 complete")
}
