package sql

import "testing"

func TestInit(t *testing.T) {

	SetEngine()
}

func TestS2t(t *testing.T) {
	S2T("3")
}
