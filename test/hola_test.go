package test

import(
	"testing"
)

func TestHolaMundo(t *testing.T)  {
	str := "hola mundo"

	if str != "hola mundo"{
		t.Error("No es posible saludor a tus amigos",nil)
	}
}