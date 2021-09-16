package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

/*func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Hello world")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Fprintln(w, "Pas World en fait")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
	}
}*/

func clockHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		currentTime := time.Now()
		fmt.Fprintf(w, "Il est actuellement : %v", currentTime.Format("00h00"))
	default:
		fmt.Fprintf(w, "Pas d'heure dézo pas marcher")
	}
}

func addPostHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Fprintln(w, "Ca ne fonctionne pas")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
		fmt.Fprintf(w, "Info reçu: %v\n", req.PostForm)
		saveFile, err := os.OpenFile("./save.data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		defer saveFile.Close()

		w := bufio.NewWriter(saveFile)
		if err == nil {
			fmt.Fprintf(w, "%v:%v:\n", req.PostForm["entries"][0], req.PostForm["author"][0])
		}
		w.Flush()

	}
}

func saveHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		filerc, err := os.Open("./save.data")
		if err != nil {
			fmt.Fprintln(w, "Pas World en fait")
		}
		defer filerc.Close()

		aByte := new(bytes.Buffer)
		aByte.ReadFrom(filerc)
		contents := aByte.String()

		split := strings.Split(contents, ":")

		for k := range split {
			if k%2 == 0 {
				fmt.Fprintf(w, split[k])
			}
		}
	}
}

func main() {
	//http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", clockHandler)
	http.HandleFunc("/add", addPostHandler)
	http.HandleFunc("/entrie", saveHandler)
	http.ListenAndServe(":4567", nil)
}
