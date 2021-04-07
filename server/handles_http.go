// Author: yangzq80@gmail.com
// Date: 2021-04-06
//
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type DoClient struct {
	Url     string
	Method  string
	Headers string
	//Data    *json.RawMessage `json:"data,ommitempty"`
	//Data map[string]interface{}
	Data interface{}
}

func (s *Server) doHttp(c *gin.Context) {
	var client DoClient
	if err := c.ShouldBind(&client); err == nil {
		c.JSON(http.StatusOK, getUrl(client))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func getUrl(do DoClient) json.RawMessage {
	var jsonStr = []byte("")

	if r, err := json.Marshal(do.Data); err == nil {
		jsonStr = r
	}

	if do.Method == "" {
		do.Method = "GET"
	} else {
		do.Method = strings.ToUpper(do.Method)
	}

	req, err := http.NewRequest(do.Method, do.Url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
		return []byte("http.get i/o timeout")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return body
}
