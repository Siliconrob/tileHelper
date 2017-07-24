package main

import (
	"fmt"
    "log"
    "io/ioutil"
    "net/http"
)

func getImage(tileUrl string) []byte {
	res, err := http.Get(tileUrl)
	 if err != nil {
	 	log.Fatalf("http.Get -> %v", err)
	 }
    data, err2 := ioutil.ReadAll(res.Body)
    if err2 != nil {
        log.Fatalf("ioutil.ReadAll -> %v", err2)
    }
    res.Body.Close()
    return data
}

func main() {
	data := getImage("http://otile1.mqcdn.com/tiles/1.0.0/map/3/1/2.jpg")
	name := fmt.Sprintf("%s-%d-%d-%d.jpg", "base", 3, 1, 2)
    ioutil.WriteFile(name, data, 0666)
    fmt.Println("Done")
}