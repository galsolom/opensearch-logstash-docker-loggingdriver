package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type testCase struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Labels  []string `json:"labels"`
	Domains []string `json:"domains"`
}

type testCases []struct {
	testCase
}

// misc methods
func Contains(array []string, input string) bool {
	for _, v := range array {
		if v == input {
			return true
		}
	}
	return false
}

func ContainsAnyOfList(array []string, inputArray []string) bool {
	for _, k := range inputArray {
		if Contains(array, k) {
			return true
		}
	}

	for _, v := range array {
		if Contains(inputArray, v) {
			return true
		}
	}

	return false
}

func MatchesParams(test testCase, params []string) bool {

	var matchDomains = ContainsAnyOfList(test.Domains, params)
	var matchLabels = ContainsAnyOfList(test.Labels, params)

	return matchDomains || matchLabels
}

func main() {
	var params []string
	var tests []string
	var testsList testCases

	var delimiter = "|tc="
	var dotnetFolder = `..\testproject`
	var testFilter = `--filter="tc=`

	for _, arg := range os.Args[1:] {
		params = append(params, arg)
	}

	content, err := ioutil.ReadFile("labels.json")
	if err != nil {
		log.Fatal("Error when loading file: ", err)
	}

	err = json.Unmarshal(content, &testsList)
	if err != nil {
		log.Fatal("Error during Unmarshal():", err)
	}

	for _, testCase := range testsList {
		if MatchesParams(testCase.testCase, params) {
			fmt.Println(testCase.Name)
			tests = append(tests, testCase.ID)
		}

	}
	if len(tests) == 0 {
		log.Print("found no tests to run using params:")
		log.Print(params)
		return
	}
	var dotnetSyntaxedFilter = strings.Join(tests, delimiter)
	var finalCommand = testFilter + dotnetSyntaxedFilter + `"`

	os.Chdir(dotnetFolder)
	cmd := exec.Command(`C:\Program Files\dotnet\dotnet.exe`, "test", finalCommand)

	stderr, error := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()

	if error != nil || err != nil {
		log.Fatal(err)
		log.Fatal(error)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)
	errdata, error := ioutil.ReadAll(stderr)

	if error != nil || err != nil {
		log.Fatal(err)
		log.Fatal(error)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("ERRORS: %s\n", (string(errdata)))
		fmt.Println("failed here4")
		log.Fatal(err)
	}

	fmt.Printf("%s\n", (string(data)))
}
