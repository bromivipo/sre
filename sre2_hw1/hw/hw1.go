package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)
   
func GetData(url string) (string, error){
	var bodyString string
	gotResp := false
	for sleepSec := time.Second; sleepSec.Seconds() <= 2; sleepSec += time.Second { 
		resp, err := http.Get(url)
		if err != nil {
			return bodyString, err
		}
		if resp.StatusCode < 500 {
			fmt.Printf("Got response with status %d\n", resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			bodyString = string(body)
			gotResp = true
			break
		}
		fmt.Printf("Error response, retrying in %d second\n", int(sleepSec.Seconds()))
		time.Sleep(sleepSec)
	}
	var err error
	if !gotResp {
		err = fmt.Errorf("Got no response with status < 500")
	}
	return bodyString, err

}
   
func main() {
	s, err := GetData("http://localhost:8080/target-handler")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(s)
	}
}
