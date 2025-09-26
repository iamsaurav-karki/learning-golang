package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func postJSON() error {
	p := Person{Name: "Saurav", Age: 54}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// Simple helper: http.Post (convenient)
	resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("POST JSON (http.Post) status:", resp.Status)
	fmt.Println("response:", string(body))
	return nil
}

func postJSONWithClient() error {
	p := Person{Name: "ClientExample", Age: 30}
	b, _ := json.Marshal(p)
	req, err := http.NewRequest("POST", "http://localhost:8080/submit", bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// you can add custom headers too: req.Header.Set("X-My-Header","value")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("POST JSON (client) status:", resp.Status)
	fmt.Println("response:", string(body))
	return nil
}

func postForm() error {
	data := url.Values{}
	data.Set("username", "saurav")
	data.Set("password", "hunter2")

	resp, err := http.PostForm("http://localhost:8080/submit", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println("POST Form status:", resp.Status)
	fmt.Println("response:", string(b))
	return nil
}

func runClient() {
	if err := postJSON(); err != nil {
		fmt.Println("postJSON err:", err)
	}
	if err := postJSONWithClient(); err != nil {
		fmt.Println("postJSONWithClient err:", err)
	}
	if err := postForm(); err != nil {
		fmt.Println("postForm err:", err)
	}
}
