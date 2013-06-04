package controllers

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"net/http"
	"time"

	"gonuts"
)

func ahHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	d := make(ContentData)

	d["Message"] = "Hello from _ah."
	ServeJSON(w, http.StatusOK, d)
	return
}

func ahCronSearchHandler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	d := make(ContentData)

	var nut gonuts.Nut
	for i := datastore.NewQuery("Nut").Run(c); ; {
		_, err := i.Next(&nut)
		if err == datastore.Done {
			break
		}
		if err == nil {
			err = gonuts.AddToSearchIndex(c, &nut)
		}
		if err != nil {
			gonuts.LogError(c, err)
			ServeJSONError(w, http.StatusInternalServerError, err, d)
			return
		}
	}

	m := fmt.Sprintf("Search index updated in %d seconds.", time.Since(start)/time.Second)
	c.Infof("%s", m)
	d["Message"] = m
	ServeJSON(w, http.StatusOK, d)
	return
}
