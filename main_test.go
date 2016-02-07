package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testGin = Gin()

func TestCreatePdfImplicitFormat(t *testing.T) {
	ts := httptest.NewServer(testGin)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/process", strings.NewReader(`\documentclass[12pt]{article}
	  \begin{document}
	  Hello world!
	  $Hello world!$ %math mode
	  \end{document}`))

	fmt.Println(ts.URL + "/process")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected a status of 200, Actual: %v", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/pdf" {
		t.Fatalf("Expected a content type of application/pdf, Actual: %v", contentType)
	}
}

func TestCreateInvalidTex(t *testing.T) {
	ts := httptest.NewServer(testGin)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/process", strings.NewReader(`\documentclass[12pt]{article}
		\begin{document}
		Hello world!
		$Hello world!$ %math mode
		`))

	fmt.Println(ts.URL + "/process")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 500 {
		t.Fatalf("Expected a status of 500, Actual: %v", res.Status)
	}

	defer res.Body.Close()
	var message Message

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(body, &message)

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(message.Value, "This is pdfTeX") {
		t.Fatalf("Expected Error Message to contain log output', Actual: \"%v\"", message.Value)
	}

}

func TestCreateInvalidFormat(t *testing.T) {
	ts := httptest.NewServer(testGin)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/process?format=invalidFormat", strings.NewReader(`\documentclass[12pt]{article}
	  \begin{document}
	  Hello world!
	  $Hello world!$ %math mode
	  \end{document}`))

	fmt.Println(ts.URL + "/process")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 400 {
		t.Fatalf("Expected a status of 400, Actual: %v", res.Status)
	}

	defer res.Body.Close()
	var message Message

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(body, &message)

	if err != nil {
		t.Fatal(err)
	}

	if message.Value != "\"invalidFormat\" was an invalid format" {
		t.Fatalf("Expected Error Message equal to '\"invalidFormat\" was an invalid format\"', Actual: %v", message.Value)
	}
}

func TestCreatePdfEmptyBody(t *testing.T) {
	ts := httptest.NewServer(testGin)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/process", strings.NewReader(``))

	fmt.Println(ts.URL + "/process")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 400 {
		t.Fatalf("Expected a status of 400, Actual: %v", res.Status)
	}

	defer res.Body.Close()
	var message Message

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(body, &message)

	if err != nil {
		t.Fatal(err)
	}

	if message.Value != "Empty Input is Not Valid Form Submission" {
		t.Fatalf("Expected Error Message equal to 'Empty Input is Not Valid Form Submission', Actual: %v", message.Value)
	}

}

func TestCreatePng(t *testing.T) {
	ts := httptest.NewServer(testGin)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/process?format=png", strings.NewReader(`\documentclass[12pt]{article}
	  \begin{document}
	  Hello world!
	  $Hello world!$ %math mode
	  \end{document}`))

	fmt.Println(ts.URL + "/process")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected a status of 200, Actual: %v", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "image/png" {
		t.Fatalf("Expected a content type of image/png, Actual: %v", contentType)
	}

}
