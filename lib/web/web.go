package web

import (
	"io"
	"os"
	"log"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	)

/*
Ideas for future:
1. Config static dirs file system path
2. Config url prefixes, since this main responsibility is configuration of mux
3. Copy/Close db sessions for mongo
*/

type JSONApiHandler func(bson.M, *http.Request)

func (h JSONApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var response bson.M = bson.M{}
	w.Header().Set("Content-type", "application/json")
	log.Println(req.URL.Path)
	req.ParseMultipartForm(100 << 20)
	rv, err1 := json.Marshal(req.Form)
	if err1 == nil {
		log.Println(string(rv))
	}
	h(response, req)
	j, err2 := json.Marshal(response)
	if err2 != nil { log.Fatal(err2) }
	w.Write(j)
}

type ApiBuilder struct {
	mux *http.ServeMux
}

func NewApiBuilder() *ApiBuilder {
	apiBuilder := &ApiBuilder{
		mux: http.NewServeMux(),
	}
	return apiBuilder
}

func (b *ApiBuilder) HandleJson(pattern string, h JSONApiHandler) {
	b.mux.Handle(pattern, h)
}

func (b *ApiBuilder) AddStaticDir(dir string) {
	b.mux.HandleFunc(dir, staticGetter)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	r.URL.Path = "index.html"
	staticGetter(w, r)
}

func staticGetter(w http.ResponseWriter, r *http.Request) {
	var resource string = r.URL.Path
	var staticRoot string = "./www/"
	var fullResource string = staticRoot + resource;
	f, err := os.Open(fullResource)
	if err != nil {
		log.Println("Static file not found:", resource)
		http.NotFound(w, r)
		return
	}
	io.Copy(w, f)		
}

func (b *ApiBuilder) Build() *http.ServeMux {
	b.mux.HandleFunc("/", indexHandler)
	return b.mux;
}