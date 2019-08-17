
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}