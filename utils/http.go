// Author: yangzq80@gmail.com
// Date: 2021/8/17
// RESTfull-APIs通用HTTP/JSON请求操作
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Http struct {
	Url     string
	Method  string //GET,POST,PUT,DELETE
	Header  map[string]string
	Payload any
}

func GetUrl(url string) map[string]interface{} {
	return DoRequest(http.MethodGet, url, nil, nil)
}

func Get(url string, header map[string]string) map[string]interface{} {
	return DoRequest(http.MethodGet, url, nil, header)
}
func Post(url string, body any, header map[string]string) map[string]interface{} {
	return DoRequest(http.MethodPost, url, body, header)
}
func Put(url string, body any, header map[string]string) map[string]interface{} {
	return DoRequest(http.MethodPut, url, body, header)
}
func Delete(url string, header map[string]string) map[string]interface{} {
	return DoRequest(http.MethodDelete, url, "", header)
}

func DoRequest(method string, url string, body any, header map[string]string) map[string]interface{} {
	http := &Http{
		Url:     url,
		Method:  method,
		Header:  header,
		Payload: body,
	}
	return http.Do()
}

func (h *Http) Do() map[string]interface{} {
	client := &http.Client{}

	var payLoad io.Reader
	if h.Payload != nil {
		// 字符串
		if str, ok := h.Payload.(string); ok && len(h.Payload.(string)) > 0 {
			payLoad = bytes.NewBufferString(str)
		} else {
			// json encoding
			b, err := json.Marshal(h.Payload)
			if err != nil {
				fmt.Println("json.marshal err:" + err.Error())
			}
			payLoad = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequest(h.Method, h.Url, payLoad)

	req.Header.Set("Content-Type", "application/json")
	for k, v := range h.Header {
		req.Header.Set(k, v)
	}

	respJ := make(map[string]interface{})
	resp, err := client.Do(req)

	if err != nil {
		log.Errorf("ExecReq-request-error:%v", err.Error())
		respJ["error"] = err.Error()
		respJ["StatusCode"] = resp.StatusCode
		return respJ
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("ExecReq-ReadResp-error:%v", err.Error())
		respJ["error"] = err.Error()
		respJ["StatusCode"] = resp.StatusCode
		return respJ
	}
	if len(body) == 0 {
		respJ["StatusCode"] = resp.StatusCode
		return respJ
	}
	err = json.Unmarshal(body, &respJ)
	if err != nil {
		respJ["body"] = string(body)
		return respJ
	}
	return respJ
}
