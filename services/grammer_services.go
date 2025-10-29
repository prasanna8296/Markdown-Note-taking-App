package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// CheckGrammar sends text to the LanguageTool API for grammar checking
func CheckGrammar(text string) (map[string]interface{}, error) {
	apiURL := "https://api.languagetool.org/v2/check"
	data := "language=en-US&text=" + text

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
