package authorize

import (
	"testing"
)

var settings = map[string]string{
	"ApiLoginId":     "",
	"TransactionKey": "",
	"Key":            "",
}

func TestAuthenticationTest(t *testing.T) {
	auth, err := New(
		settings["ApiLoginId"],
		settings["TransactionKey"],
		settings["Key"],
		true,
	)
	if err != nil {
		t.Errorf("Error instantiate authorize library: `%s`", err)
		return
	}

	if err = auth.AuthenticationTest(); err != nil {
		t.Errorf("Authentication test faild: `%s`", err)
	}

	t.Logf("Authentication test success")
}
