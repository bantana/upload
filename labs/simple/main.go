// curl -F "data=@./testfile/test1.txt" http://localhost:3000/upload
package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/meatballhat/negroni-logrus"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

//  fileServer define Upload file director and Port
type fileServer struct {
	Dir  string
	Port string
	URL  string
}

var (
	log = logrus.New()
)

func main() {
	logfile := flag.String("logfile", "", "logrus log file path")
	level := flag.String("level", "", "log level can be debug, error, default is info")
	flag.Parse()

	// log := logrus.New()
	log.Infof("log level is %s", *level)
	switch *level {
	case "debug":
		log.Level = logrus.DebugLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "info":
		log.Level = logrus.InfoLevel
	default:
		log.Fatal("loglevel must debug , error , info")
	}

	if *logfile != "" {
		file, err := os.OpenFile(*logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Out = file
	} else {
		log.Out = os.Stdout
	}
	mux := httprouter.New()

	myserver := &fileServer{
		Dir:  "./public",
		Port: ":3001",
		URL:  "/file",
	}

	// log.Println("something run tips here!")
	log.Infof("tips: staring server at port: %s use local dir: %s,", myserver.Port, myserver.Dir)
	log.Infof("tips: access path: %s", myserver.URL)

	mux.ServeFiles("/file/*filepath", http.Dir(myserver.Dir))
	mux.POST("/upload", uploadHandle)

	n := negroni.New()
	// n.Use(negronilogrus.NewMiddleware())
	// n.Use(negronilogrus.NewCustomMiddleware(loglevel, &logrus.JSONFormatter{}, "web"))
	n.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	n.UseHandler(mux)
	n.Run(myserver.Port)
	// log.Fatal(http.ListenAndServe(myserver.Port, mux))
}

func uploadHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("data")
	if err != nil {
		log.Error(err)
	}
	log.Debugf("upload files Header: %v!\n", header)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fid, err := uuid.NewV4()
	if err != nil {
		log.Error("fid uuid : %s", err)
	}
	log.Debugf("fid: %s", fid)
	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatal(err)
	}
}
