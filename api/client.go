package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchJSON sends a GET request to the given URL
// and decodes the JSON response into target.
func FetchJSON(url string, target any) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}