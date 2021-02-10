package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//resp, err := http.Get("https://www.imooc.com/")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com/", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {

			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	s, err := httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)
}
