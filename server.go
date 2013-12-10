package goshort

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "errorH"
    "html/template"
    "crypto/sha1"
    "appengine"
    "appengine/datastore"
)

// this is the type of entity stored in the datastore
type Url struct {
	OrigUrl string // these fields are uppercase to give other functions access to them
	Id string
}

// Our collection of templates
var templates = template.Must(template.ParseFiles(
	"index.html",
	"error.html",
))

// sets up http handlers with urls
func init() {
    r := mux.NewRouter()
    r.HandleFunc("/", errorH.Handle(rootH))
    r.HandleFunc("/{urlId}", errorH.Handle(urlH))
    http.Handle("/", r)
}

// this function shows a form to enter a url to be shortened
// it generates a url id, which should be copied and pasted after the goshort address
// example when running goshort on port 8080: localhost:8080/urlId
func rootH(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
		templates.ExecuteTemplate(w, "index.html", nil)
	} else {
		shortUrl := r.FormValue("shortUrl")
		id := genUrlId([]byte(shortUrl))
		c := appengine.NewContext(r)
		key := datastore.NewKey(c, "Url", id, 0, nil)
        url := new(Url)
        err := datastore.Get(c, key, url)
		if err == datastore.ErrNoSuchEntity {
        	newUrl := &Url{shortUrl, id}
       		_, err := datastore.Put(c, key, newUrl)
        	errorH.Check(err)
    	}
        templates.ExecuteTemplate(w, "index.html", id)
	}
}

// this function receives a url id, and trys to get the url with that id from the datastore
// if the url is stored, it redirects to it
// if not, it throws an error
func urlH(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlId := vars["urlId"]
	
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Url", urlId, 0, nil)
    queryUrl := new(Url)
    err := datastore.Get(c, key, queryUrl)
    if err == datastore.ErrNoSuchEntity {
    	templates.ExecuteTemplate(w, "error.html", "this id does not correspond with a url")
    } else {
        http.Redirect(w, r, queryUrl.OrigUrl, http.StatusFound)
    }
}

// this function hashes a url and prints the first 8 characters, to generate a unique id
func genUrlId(inUrl []byte) string {
	sha := sha1.New()
	sha.Write(inUrl)
	return fmt.Sprintf("%x", string(sha.Sum(nil))[0:4])
}