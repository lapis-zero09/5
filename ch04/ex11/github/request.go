package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ReqUser struct {
	token string
}

func NewReqUser(token string) *ReqUser {
	return &ReqUser{token}
}

func (r *ReqUser) newRequest(method, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endPoint+uri, body)
	if err != nil {
		return nil, err
	}
	// req.Header.Set(
	// 	"Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+r.token)
	return req, nil
}

func (r *ReqUser) do(req *http.Request, i interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 2xx Success系ではないもの
	if resp.StatusCode > 299 {
		return fmt.Errorf("failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(i); err != nil {
		return err
	}
	return nil
}
