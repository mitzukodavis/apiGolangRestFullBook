package test

import(
	"testing"
	"github.com/mitzukodavis/apirestgolang/models"
)

func TestConnection(t *testing.T)  {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("no es posible realizar ", nil)
	}
}
