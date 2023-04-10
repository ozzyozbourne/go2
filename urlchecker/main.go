package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func checkAndSaveBody(url string) {
	if res, err := http.Get(url); err != nil {
		panic(fmt.Sprintf("Error in GET -> %v\n", err))
	} else {
		defer func() {
			err := res.Body.Close()
			if err != nil {
				fmt.Printf("[ERROR IN CLOSING THE RESPONSE BODY] -> %v", err)
			}
		}()
		fmt.Printf("[URL] -> %-20s [STATUS CODE] -> %d", url, res.StatusCode)
	}
}
