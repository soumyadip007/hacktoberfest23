package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"npf.io/gorram/run"
)

func Get(pkg string, fun string, args []string) (string, error) {
	var output bytes.Buffer
	env := run.Env{
		Stdout: io.MultiWriter(os.Stdout, &output),
		Stderr: os.Stderr,
	}
	input := run.Command{
		Package:  pkg,
		Function: fun,
		Env:      env,
		Args:     args,
	}
	err := run.Run(&input)
	return output.String(), err
}

func reqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		handlePut(w, r)
	default:
		http.NotFound(w, r)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	err = r.ParseMultipartForm(int64(2 * 4294967295))
	if err != nil {
		fmt.Fprintf(w, "Failed to parse form %s\n", err)
	}
	reqBody := r.PostForm
	reqIP := r.RemoteAddr
	log.Printf("Received new POST from %s\n", reqIP)
	pkg := reqBody.Get("package")
	fun := reqBody.Get("function")
	args := reqBody.Get("arguments")
	arguments := strings.Split(args, " ")
	if err != nil {
		fmt.Fprintf(w, "Failed to collect data %s\n", err)
	}
	output, err := Get(pkg, fun, arguments)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
	fmt.Fprintf(w, "%s", output)
	return
}

func main() {
	http.HandleFunc("/gorram", reqHandler)
	log.Println("Listening on 0.0.0.0:9900")
	log.Fatal(http.ListenAndServe(":9900", nil))
}
