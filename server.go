package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	args := os.Args
	var port string
	if len(args) > 1 {
		port = args[1]
	} else {
		port = "8000"
	}
	fmt.Printf("Running on localhost: " + port + "\n")

	
	http.HandleFunc("/", handler)
	http.ListenAndServe(":" + port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(string(r.URL.Path))
	body, err := ioutil.ReadFile("player.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%s", body)
}
	
