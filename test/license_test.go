package test

import "testing"

func TestShowLicense(t *testing.T) {
	license, err := client.ShowLicense(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", license)
}
