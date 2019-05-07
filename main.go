package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r);
	bs, err := ioutil.ReadFile("/content/index.html")

	if err != nil {
		fmt.Printf("Couldn't read index.html: %v", err);
		os.Exit(1)
	}

	io.WriteString(w, string(bs[:]))
}

func main() {
    thishost, err := os.Hostname()
    if err != nil {
        fmt.Printf("Couldn't get Hostname: %v", err);
    }
    fmt.Println("hostname:", thishost)
    input, err := ioutil.ReadFile("/content/index.html")
    if err != nil {
        fmt.Printf("Couldn't read index.html: %v", err);
    }

    lines := strings.Split(string(input), "\n")

    for i, line := range lines {
        if strings.Contains(line, "    <h2>__HOSTNAME__</h2>") {
             lines[i] = "    <h2>" + thishost + "</h2>"
        }
    }
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile("/content/index.html", []byte(output), 0644)
    if err != nil {
        fmt.Printf("Error writing /content/index.html: %v", err);
    }
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port);
	http.ListenAndServe(port, nil)
}
