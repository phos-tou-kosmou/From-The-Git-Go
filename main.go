package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(os.Args), os.Args)

	if len(os.Args) > 5 || len(os.Args) < 5 {
		fmt.Println("You are using too many or too few cli arguments.\n This tool only supports:\n\t -u -username \nand \n\t-r -reponame\nand both must be used to complete the operation.")
		return
	}

	var userCliArg string
	var repoCliArg string

	for i := range os.Args {
		if os.Args[i] == "-u" || os.Args[i] == "-username" {
			userCliArg = os.Args[i+1]
		} else if os.Args[i] == "-r" || os.Args[i] == "-reponame" {
			repoCliArg = os.Args[i+1]
		}

	}

	buildAndSendURL(userCliArg, repoCliArg)

	fmt.Println(userCliArg, repoCliArg)

}

func buildAndSendURL(uca string, rca string) {
	body := strings.NewReader(`-u ` + "'" + uca + "'" + `  -d ` + "'" + `{"name":"` + string(rca) + `"}` + "'")
	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	defer resp.Body.Close()
}
