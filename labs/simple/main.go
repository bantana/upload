package main

import (
	"log"
	"net/http"
	// "io"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "mime"
	// "mime/multipart"
	"github.com/julienschmidt/httprouter"
)

//  fileServer define Upload file director and Port
type fileServer struct {
	Dir  string
	Port string
	URL  string
}

func main() {
	// mux := http.NewServeMux()
	mux := httprouter.New()

	myserver := &fileServer{
		Dir:  "./public",
		Port: ":3001",
		URL:  "/file",
	}

	// fmt.Println("starting server")
	log.Println("something run tips here!")
	log.Printf("tips: staring server at port: %s use local dir: %s,", myserver.Port, myserver.Dir)
	log.Printf("tips: access path: %s", myserver.URL)

	mux.ServeFiles("/file/*filepath", http.Dir(myserver.Dir))

	log.Fatal(http.ListenAndServe(myserver.Port, mux))
}
