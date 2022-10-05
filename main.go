package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	data, err := ioutil.ReadFile("File.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	res := string(data)                         // convertn content to a 'string'
	index := getbefore(backline(res), "before") // get the index of the word before
	fmt.Println("Fragment 1 :", backline(res)[0])
	fmt.Println("Fragment 2 :", backline(res)[len(backline(res))-1])
	tempin := strings.TrimSpace(backline(res)[index])
	in, _ := strconv.Atoi(tempin)
	fmt.Println("Fragment 3 :", backline(res)[in-1])
	index2 := getbefore(backline(res), "now") - 2
	temp2 := backline(res)[index2][1] / 37
	fmt.Println("Fragment 4 :", backline(res)[int(temp2)-1])
	fmt.Print("Un numéro aleatoire entre 0 et 100 : ", rand.Intn(max-min)+min)

}

func getbefore(content []string, word string) int { //function that get the index of the word before
	for i, v := range content {
		if strings.Contains(v, word) {
			return i + 1
		}
	}
	return 0
}

func backline(content string) []string { //function to split the data in array of string
	var res []string
	var word string
	for _, c := range content {
		if c == '\n' {
			res = append(res, word)
			word = ""
		} else {
			word += string(c)
		}
	}
	res = append(res, "a")
	return res
}

func random(min, max int) int {
	url := "http://www.randomnumberapi.com/api/v1.0/random?min="
	newurl := (url + strconv.Itoa(min) + "&max=" + strconv.Itoa(max) + "&count=1")
	resp, err := http.Get(newurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	res := string(body)
	newres := string(res[1])
	value, _ := strconv.Atoi(string(newres))
	return value
}
