package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "http://169.254.169.254/latest/meta-data"

func getToken() (string, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("PUT", "http://169.254.169.254/latest/api/token", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-aws-ec2-metadata-token-ttl-seconds", "300")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	return string(token), err
}

func getMetadata(path string, token string) (interface{}, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseURL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-aws-ec2-metadata-token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	content := string(body)

	if strings.HasSuffix(content, "/") || strings.Contains(content, "\n") {
		// Directory or multi-key
		lines := strings.Split(strings.TrimSpace(content), "\n")
		result := make(map[string]interface{})
		for _, line := range lines {
			key := strings.TrimSuffix(line, "/")
			val, err := getMetadata(path + "/" + key, token)
			if err != nil {
				result[key] = fmt.Sprintf("error: %v", err)
			} else {
				result[key] = val
			}
		}
		return result, nil
	}

	return content, nil
}

func main() {
	var path string
	if len(os.Args) > 1 {
		path = strings.Trim(os.Args[1], "/")
	}

	token, err := getToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get token: %v\n", err)
		os.Exit(1)
	}

	if path == "" {
		// Get top-level keys to recurse over
		resp, err := getMetadata("", token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get metadata: %v\n", err)
			os.Exit(1)
		}
		json.NewEncoder(os.Stdout).Encode(resp)
	} else {
		// Get specific key
		resp, err := getMetadata(path, token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get key '%s': %v\n", path, err)
			os.Exit(1)
		}
		json.NewEncoder(os.Stdout).Encode(resp)
	}
}
