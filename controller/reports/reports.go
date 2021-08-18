package reports

import (
	"net/http"

	"../../lib/flight"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/reports", Index)
}

// Index displays the reporting page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("reports/index")
	//if c.Sess.Values["id"] != nil {
	//	v.Vars["first_name"] = c.Sess.Values["first_name"]
	//}
	v.Render(w, r)
}
