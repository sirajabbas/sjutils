package piutils

import "testing"

func TestPrintFailure(t *testing.T) {
	PrintFailure("printing failure")
}

func TestPrintSuccess(t *testing.T) {
	PrintSuccess("printing success")
}