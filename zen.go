package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
)

// Bezen prints a zen advice
func Bezen() {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("https://api.github.com/zen")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("%s\n", color.YellowString(string(bytes)))
}
