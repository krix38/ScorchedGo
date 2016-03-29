package page

import (
	"net/http"

	"github.com/krix38/ScorchedGo/web/viewResolver"
)

func Main(w http.ResponseWriter, r *http.Request) {
	viewResolver.Render(w, "main", nil)
}
