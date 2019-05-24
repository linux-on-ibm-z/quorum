package vault

import (
	"testing"
)

func getMinimumValidConfig() HashicorpWalletConfig {
	return HashicorpWalletConfig{
		Client: HashicorpClientConfig{
			Url: "someurl",
		},
		Secrets: []HashicorpSecret{
			{
				Name:         "name",
				SecretEngine: "eng",
				Version:      1,
				AccountID:    "acct",
				KeyID:        "key",
			},
			{
				Name:         "othername",
				SecretEngine: "othereng",
				Version:      1,
				AccountID:    "otheracct",
				KeyID:        "otherkey",
			},
		},
	}
}

func TestValidateValidReturnsNil(t *testing.T) {
	w := getMinimumValidConfig()

	err := w.Validate(false)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestValidateNoVaultUrlReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Client.Url = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateNoSecretNameReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].Name = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateNoSecretEngineReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].SecretEngine = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateZeroVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].Version = 0

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateNegativeVersionReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].Version = -1

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateNoAccountIdReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].AccountID = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateNoKeyIdReturnsError(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].KeyID = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}
}

func TestValidateMultipleErrorsAreCombined(t *testing.T) {
	w := getMinimumValidConfig()
	w.Secrets[0].Name = ""
	w.Secrets[1].Name = ""

	err := w.Validate(false)

	if err == nil {
		t.Error("Wanted error")
	}

	want := "\nInvalid vault secret config, vault=someurl: Name must be provided\nInvalid vault secret config, vault=someurl: Name must be provided"

	got := err.Error()

	if got != want {
		t.Errorf("Incorrect error\nwant: %v\ngot : %v", want, got)
	}
}
