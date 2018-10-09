package print_test

import (
	"testing"

	"tmt/print"
)

func TestCenter(t *testing.T) {
	checkCentered(t, "", ``)

	checkCentered(t, "CP123456", `

     CP123456`)

	checkCentered(t, " abc\ndef", `

       abc
       def`)

	checkCentered(t, "1\n123\n12345", `
        1
       123
      12345`)
}

func checkCentered(t *testing.T, have, want string) {
	got := print.CenterLabelText(have)
	if got != want {
		t.Errorf("got\n---\n%s\n---\nbut want\n---\n%s\n---", got, want)
	}
}
