package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kevydotvinu/curl-paste/paste"
)

var configFile = flag.String("c", "./curl-paste.conf", "Configuration file for curl-paste")

var pConf = Config{
	server:    "localhost",
	address:   "0.0.0.0",
	port:      "9990",
	directory: "/",
	size:      4294967295,
	logfile:   "/curl-paste.log",
}

func min(a, b int) int {

	if a > b {
		return b
	} else {
		return a
	}

}

func handleGetPaste(w http.ResponseWriter, r *http.Request) {

	var pasteName, origName string

	origName = filepath.Clean(r.URL.Path)
	pasteName = pConf.directory + "/" + origName

	origIP := r.RemoteAddr

	log.Printf("Received GET from %s for  '%s'\n", origIP, origName)

	if (origName == "/") || (origName == "/index.html") {
		http.ServeFile(w, r, pConf.directory+"/index.html")
	} else {

		_, _, content, err := paste.Retrieve(pasteName)

		if err == nil {
			fmt.Fprintf(w, "%s", content)
			return
		} else {
			fmt.Fprintf(w, "%s\n", err)
			return
		}
	}
}

func handlePutPaste(w http.ResponseWriter, r *http.Request) {

	err1 := r.ParseForm()
	err2 := r.ParseMultipartForm(int64(2 * pConf.size))

	if err1 != nil && err2 != nil {
		http.ServeFile(w, r, pConf.directory+"/index.html")
	} else {
		reqBody := r.PostForm

		origIP := r.RemoteAddr

		log.Printf("Received new POST from %s\n", origIP)

		title := reqBody.Get("title")
		date := time.Now().String()
		content := reqBody.Get("paste")

		content = content[0:min(len(content), int(pConf.size))]

		ID, err := paste.Store(title, date, content, pConf.directory)

		log.Printf("   Title: %s\n", title)
		log.Printf("   ID: %s\n", ID)

		if err == nil {
			hostname := pConf.server
			port := pConf.port
			fmt.Fprintf(w, "http://%s:%s/%s\n", hostname, port, ID)
			return
		} else {
			fmt.Fprintf(w, "%s\n", err)
		}
	}
}

func reqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		handleGetPaste(w, r)
	case "POST":
		handlePutPaste(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {

	flag.Parse()

	parseConfig(*configFile, &pConf)

	f, err := os.OpenFile(pConf.logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening log file: %s. Exiting\n", pConf.logfile)
		os.Exit(1)
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetPrefix("[curl-paste]: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	log.Println("curl-paste version 0.0.1 -- Starting ")
	log.Printf("  + Config file: %s\n", *configFile)
	log.Printf("  + Serving pastes on: %s:%s\n", pConf.server, pConf.port)
	log.Printf("  + Listening on: %s:%s\n", pConf.address, pConf.port)
	log.Printf("  + Paste directory: %s\n", pConf.directory)
	log.Printf("  + Max size: %d\n", pConf.size)

	http.HandleFunc("/", reqHandler)
	log.Fatal(http.ListenAndServe(pConf.address+":"+pConf.port, nil))
}
