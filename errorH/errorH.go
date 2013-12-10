package errorH

import (
	"net/http"
	"html/template"
)

func Handle(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				w.WriteHeader(http.StatusInternalServerError)
				t, _ := template.ParseFiles("error.html")
    			t.Execute(w, err)
			}
		}()
		fn(w, r)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}