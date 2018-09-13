package handlers

import (
	"net/http"
	"github.com/mitzukodavis/apirestgolang/utils"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	utils.RenderTemplate(w, "application/index", nil)
}