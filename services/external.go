package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Classifier struct {
}

func (c Classifier) Classify(line string) (category string, err error) {
	// Call classifier API with line in JSON and extract cateogry from JSON response

	body := []byte(`{ "line": "` + line + `"}`)
	req, err := http.NewRequest("POST", "http://localhost:8000/classify", bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if category, ok := response["category"]; ok {
		return category, nil
	} else {
		return "", err
	}
}
