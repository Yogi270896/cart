package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

func Send(method string, endpoint string, body interface{}, auth string, service string) ([]byte, error) {

	request, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(request))

	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Set("Authorization", "Basic "+authEncoded)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rawRes, resErr := client.Do(req)
	resBody, _ := io.ReadAll(rawRes.Body)
	//log.Println(service, "Response", string(resBody))
	if resErr != nil {
		return nil, resErr
	}
	defer rawRes.Body.Close()

	return resBody, nil
}
