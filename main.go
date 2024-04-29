package main

import (
    "fmt"
    "net/http"
)
func main (){
    resp, err := http.Get("https://distopia.savi2w.workers.dev/")
    if err != nil {
		fmt.Println("", err)
		return
	}
    defer resp.Body.Close()

    contentType := resp.Header.Get("Distopia")
    fmt.Println("Distopia", contentType)
}