package notifier

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//FireCallback call a http / https endpoint
func FireCallback(success bool, callbackURL string, body map[string]string, header map[string]string) error {

	// Simulating a form post is done like this:
	urlValue := url.Values{}
	for key, val := range body {
		urlValue.Add(key, val)
	}
	urlValue.Add("success", strconv.FormatBool(success))

	req, err := http.NewRequest("POST", callbackURL, strings.NewReader(urlValue.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, val := range header {
		req.Header.Set(key, val)
	}
	rs, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if rs.StatusCode == 200 {
		return nil
	}
	fmt.Println(rs)
	return errors.New(rs.Status)
}
