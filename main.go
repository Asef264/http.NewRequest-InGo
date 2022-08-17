package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const serverport = 3333

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
			fmt.Fprintf(w, `{"message": "hello!"}`)
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverport),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("err running http server : %s \n", err)
		}

	}()
	time.Sleep(100 * time.Millisecond)

	requestURL := fmt.Sprintf("http://localhost:%d", serverport)
	//http.NewRequest creates the request but dose not send it to server right away
	//so we can make any changes to it
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	//http.DefaultClient.Do sends the req to the server
	//DefaultClient value is Go's default http client. the same we've been using with http.Get()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response! \n")
	fmt.Printf("client : Status code : %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}
