package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type coin struct {
	Name string
	CurrentValue []byte
	YesterdayVal []byte
}

func (c coin) String() string {
	return fmt.Sprintf("%s:\n  Yesterday:  %s\n  Today:      %s\n", c.Name, c.YesterdayVal, c.CurrentValue)
}

func main() {

	url := "https://cbr.ru"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	startIndex := bytes.Index(body, []byte("main-indicator_rate"))
	eurIndex := bytes.Index(body[startIndex:], []byte("EUR"))
	endIndex := bytes.Index(body[startIndex + eurIndex:], []byte("Официальный курс Банка России"))

	reg, err := regexp.Compile("\\d+,\\d+")
	if err != nil {
		fmt.Println(err)
		return
	}
	matches := reg.FindAll(body[startIndex : startIndex + eurIndex + endIndex], 4)
	if len(matches) != 4 {
		fmt.Println("Something went wrong, abort...")
		return
	}
	usb := coin{
		Name: "USB",
		YesterdayVal: matches[0],
		CurrentValue: matches[1],
	}
	eur := coin{
		Name: "EUR",
		YesterdayVal: matches[2],
		CurrentValue: matches[3],
	}
	fmt.Println(usb)
	fmt.Println(eur)

	return
}
