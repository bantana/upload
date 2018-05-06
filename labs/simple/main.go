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
)

//  fileServer define Upload file director and Port
type fileServer struct {
	Dir  string
	Port string
	URL  string
}

func main() {
	mux := http.NewServeMux()
	myserver := &fileServer{"./public", ":3001", "/file"}
	// fmt.Println("starting server")
	log.Println("something run tips here!")
	log.Printf("tips: staring server at port: %s use local dir: %s,", myserver.Port, myserver.Dir)
	log.Printf("tips: access path: %s", myserver.URL)

	// mux.Handle("/", welcomeHandler())
	mux.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(myserver.Dir))))

	log.Fatal(http.ListenAndServe(myserver.Port, mux))
}

// func welcomeHandler() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Welome"))
// 	})
// }
