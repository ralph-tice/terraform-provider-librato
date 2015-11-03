package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/underarmour/terraform-provider-librato/provider"
)

func DoRequest(
	method, path string,
	prov *provider.Provider,
	bodyStruct, respStruct interface{},
	expectedCode int,
) (int, error) {
	var bodyReader io.Reader

	// marshall request body
	if bodyStruct != nil {
		bodyBytes, err := json.Marshal(bodyStruct)
		if err != nil {
			return -1, fmt.Errorf("doRequest failed to marshal: %v %#v", err, bodyStruct)
		}
		log.Printf("[DEBUG] doRequest marshalled body: %#v %s", bodyStruct, bodyBytes)
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// build request
	req, err := prov.NewRequest(method, path, bodyReader)
	if err != nil {
		return -1, fmt.Errorf("doRequest failed to build request: %v %v", err, req)
	}
	log.Printf("[DEBUG] doRequest built request: %s %s %v", method, path, req)

	// make request
	resp, err := prov.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return -1, fmt.Errorf("doRequest failed request: %v %v", err, req)
	}
	log.Printf("[DEBUG] doRequest made request: %v %v", req, resp)

	// check status code
	if resp.StatusCode != expectedCode {
		return resp.StatusCode, fmt.Errorf("doRequest unexpected status: %v %v", expectedCode, resp.StatusCode)
	}

	// read response body
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("doRequest failed read response body: %v", err)
	}
	log.Printf("[DEBUG] doRequest read response body: %s", respBytes)

	// unmarshal response
	err = json.Unmarshal(respBytes, respStruct)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("doRequest failed unmarshal: %v", err)
	}
	log.Printf("[DEBUG] doRequest marshal response: %#v", respStruct)

	return resp.StatusCode, nil
}
