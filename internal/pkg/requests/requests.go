package requests

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"slices"
	"time"
)

var METHODS = []string{"GET", "POST", "PUT", "DELETE"}

func Request(url string, method string, headers map[string]string, requestData interface{}, responseData interface{}) error {
	//log.Info("Request start: ", url, method, headers, requestData, responseData)
	c := &http.Client{Timeout: time.Duration(5) * time.Second}

	if slices.Contains(METHODS, method) == true {
		var requestDataJson io.Reader
		if requestData != nil {
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				log.Error("Request load requestData error: ", err)
			}
			requestDataJson = bytes.NewReader(jsonData)
		} else {
			requestDataJson = nil
		}

		//log.Info("Requests requestData: ", requestDataJson)
		req, err := http.NewRequest(method, url, requestDataJson)
		//log.Info("Request NewRequest start: ", req)

		if err != nil {
			log.Error("Request NewRequest error: ", err)
			return err
		}
		// add headers
		for k, v := range headers {
			//log.Info("Request add header: ", k, v)
			req.Header.Add(k, v)
		}
		resp, err := c.Do(req)

		if err != nil {
			log.Error("http request error: ", err)
			return err
		}
		defer resp.Body.Close()

		jsonErr := json.NewDecoder(resp.Body).Decode(responseData)
		if jsonErr != nil {
			log.Error("解析失败: ", jsonErr)
			return jsonErr
		}
	} else {
		errors.Errorf("method: %s currently not supported, please use supported method in: %v", method, METHODS)
	}

	return nil
}
