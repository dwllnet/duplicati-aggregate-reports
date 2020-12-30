package clientstatus

import (
	"../../lib/flight"
	"../../model/duplicati"
	"../../model/duplicati/storage"
	"encoding/json"
	"github.com/blue-jay/core/router"
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

var (
	uri = "/clientstatus"
)

// Load the routes.
func Load() {
	//TODO:ACL using keys
	router.Get(uri, Index)
	router.Post("/clientstatus/:server", RecordClientStatus)
}

// Index displays the login page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	c.View.New("clientstatus/index").Render(w, r)
}

func RecordClientStatus(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	duplicatiClient := c.Param("server")

	// TODO: Move this into a helper
	if r.Header.Get("Content-Type") != "" {
		reportBody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if reportBody != "application/json" {
			msg := "Content-Type json header missing."
			c.FlashNotice(msg)
			return
		} else {
			//Enforce a maximum read
			r.Body = http.MaxBytesReader(w, r.Body, 1048576)
			reportJSON := json.NewDecoder(r.Body)
			reportJSON.DisallowUnknownFields()
			var report storage.JsonReport

			err := reportJSON.Decode(&report)
			if err != nil {
				c.FlashWarning("JSON Decode Error")
			}

			//log.Println(report)

			summaryReport := duplicati.CreateSummary(duplicatiClient, report)
			_, err = duplicati.Create(c.DB, summaryReport)

			if err != nil {
				c.FlashErrorGeneric(err)
				//log.Println(err)
			} else {
				c.FlashSuccess("Data stored: ")
				return
			}
		}
	} else {
		msg := "Conten-Type is empty."
		c.FlashWarning(msg)
	}

	v := c.View.New("clientstatus/index")
	v.Vars["Server"] = duplicatiClient
	v.Render(w, r)
}
