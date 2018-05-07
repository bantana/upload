// curl -F "data=@./testfile/test1.txt" http://localhost:3000/upload
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// "io"

	// "log"
	// "net/http"
	// "mime"
	// "mime/multipart"
	"github.com/julienschmidt/httprouter"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
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

	// log.Println("something run tips here!")
	log.Printf("tips: staring server at port: %s use local dir: %s,", myserver.Port, myserver.Dir)
	log.Printf("tips: access path: %s", myserver.URL)

	mux.ServeFiles("/file/*filepath", http.Dir(myserver.Dir))
	mux.POST("/upload", uploadHandle)

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(mux)
	n.Run(myserver.Port)
	// log.Fatal(http.ListenAndServe(myserver.Port, mux))
}

func uploadHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("data")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "upload files Header: %v!\n", header)
	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatal(err)
	}
}
