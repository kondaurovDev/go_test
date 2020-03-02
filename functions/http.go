package functions

import (
	"fmt"
	"net/http"
)

func StartHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		name := q.Get("name")
		fmt.Fprintf(w, "Hello "+name)
	})
	http.ListenAndServe(":8099", nil)
}
