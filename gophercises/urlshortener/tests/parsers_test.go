package tests

import(
	"testing"
)

func TestOk(t *testing.T) {
	return
}

func TestNotOk(t *testing.T) {
	t.Errorf("not ok")
	return
}