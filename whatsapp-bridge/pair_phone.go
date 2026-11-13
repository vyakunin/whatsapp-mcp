package main

import (
	"os"
	"path/filepath"
	"strings"
)

// pairPhoneEnv names the phone number to link by pairing code instead of a QR scan. The account
// owner types the code into WhatsApp → Linked devices → Link with phone number, so no second
// screen is needed. Unset, the bridge pairs by QR as before.
const pairPhoneEnv = "WHATSAPP_PAIR_PHONE"

// pairCodePath holds the current pairing code while one is pending, for whoever relays it to the
// account owner; it is removed once pairing succeeds.
const pairCodePath = "store/pair_code.txt"

func pairPhoneFromEnv() string {
	return strings.TrimSpace(os.Getenv(pairPhoneEnv))
}

func writePairCode(path, code string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(code+"\n"), 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
