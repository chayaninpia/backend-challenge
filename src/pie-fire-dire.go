package src

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Result struct {
	Beef map[string]int `json:"beef"`
}

func PieFireDire(w http.ResponseWriter, r *http.Request) {

	inputPath := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	input, err := reqInput(inputPath)
	if err != nil {
		panic(err)
	}

	result := filterBeef(input)

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func reqInput(path string) (string, error) {

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	inputText := string(b)
	return inputText, nil
}

func filterBeef(input string) Result {
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, "\n", "")
	inputArr := strings.Split(input, " ")
	result := Result{
		Beef: map[string]int{},
	}
	for _, v := range inputArr {
		if v != "" {
			result.Beef[v]++
		}
	}
	return result
}
