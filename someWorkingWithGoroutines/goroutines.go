package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)
type Coin interface {
	String() string
}


type coin struct {
	Name string
	CurrentValue []byte
	YesterdayVal []byte
}

func (c coin) String() string {
	return fmt.Sprintf("%s:\n  Yesterday:  %s\n  Today:      %s\n", c.Name, c.YesterdayVal, c.CurrentValue)
}

type coinV2 struct {
	intCode string
	strCode string
	Count int
	Name string
	Value string
}

func (c coinV2) String() string {
	args := []string{
		c.Name + ":",
		"  Int code: " + c.intCode,
		"  Str code: " + c.strCode,
		"  Count:    " + strconv.Itoa(c.Count),
		"  Value:    " + string(c.Value),
	}
	return  strings.Join(args, "\n")
}

func main() {

	coins, err := GetCoin()
	if err != nil {
		fmt.Println(err)
	}
	for _, coin := range coins {
		fmt.Println(coin)
	}

	coinsV2, err := GetAllCoins()
	if err != nil {
		fmt.Println(err)
	}
	for _, coin := range coinsV2 {
		fmt.Println(coin)
	}


	return
}

func GetAllCoins() ([]coinV2, error) {
	url := "https://cbr.ru/currency_base/daily/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(body))

	start := "<table class=\"data\">"
	end := "</table>"
	ind_start := bytes.Index(body, []byte(start))
	ind_end := bytes.Index(body[ind_start:], []byte(end))
	//fmt.Println(string(body[ind_start : ind_start + ind_end + 8]))

	reg := regexp.MustCompile("<td>.*</td>")
	values := reg.FindAll(body[ind_start : ind_start + ind_end + 8], -1)

	coins := make([]coinV2, len(values)/5)

	for ind_val, ind := 0, 0; ind_val < len(values); ind_val = ind_val + 5 {
		coins[ind].intCode = string(bytes.Trim(values[ind_val], "</td>"))
		coins[ind].strCode = string(bytes.Trim(values[ind_val + 1], "</td>"))
		coins[ind].Count, err = strconv.Atoi(string(bytes.Trim(values[ind_val + 2], "</td>")))
		if err != nil {
			return nil, nil
		}
		coins[ind].Name = string(bytes.Trim(values[ind_val + 3], "</td>"))
		coins[ind].Value = string(bytes.Trim(values[ind_val + 4], "</td>"))
		ind++
	}
	return coins, nil
}


func GetCoin() ([]coin, error) {

	url := "https://cbr.ru"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	startIndex := bytes.Index(body, []byte("main-indicator_rate"))
	eurIndex := bytes.Index(body[startIndex:], []byte("EUR"))
	endIndex := bytes.Index(body[startIndex + eurIndex:], []byte("Официальный курс Банка России"))

	reg, err := regexp.Compile("\\d+,\\d+")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	matches := reg.FindAll(body[startIndex : startIndex + eurIndex + endIndex], 4)
	if len(matches) != 4 {
		fmt.Println("Something went wrong, abort...")
		return nil, nil
	}

	coins := make([]coin, 2)

	coins[0] = coin{
		Name: "USB",
		YesterdayVal: matches[0],
		CurrentValue: matches[1],
	}
	coins[1] = coin{
		Name: "EUR",
		YesterdayVal: matches[2],
		CurrentValue: matches[3],
	}
	return coins, nil
}
