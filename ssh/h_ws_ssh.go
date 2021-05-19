// Author: yangzq80@gmail.com
// Date: 2021-05-07
//
package ssh

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

type sshHost struct {
	Host     string `form:"host"` //22
	User     string `form:"user"`
	Port     string `form:"port"`
	PwdType  string `form:"pwdType"`
	Password string `form:"password"`
	KeyFile  string `form:"keyFile"`
	Cols     int    `form:"cols"`
	Rows     int    `form:"rows"`
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handle webSocket connection.
// browser(data) -> webSocket -> ssh connection -> ssh server
// ssh server(data) -> ssh connection -> webSocket -> browser
func WsSsh(c *gin.Context) {
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if handleError(c, err) {
		return
	}
	defer wsConn.Close()

	var client *Client

	host := new(sshHost)
	c.ShouldBind(&host)
	if host.Port == "" {
		host.Port = "22"
	}

	if host.PwdType == "key" {
		client, err = DialWithKey(host.Host+":"+host.Port, host.User, host.KeyFile)
	} else {
		client, err = DialWithPasswd(host.Host+":"+host.Port, host.User, host.Password)
	}

	defer client.Close()

	//startTime := time.Now()
	ssConn, err := NewSshConn(host.Cols, host.Rows, client.client)

	if wshandleError(wsConn, err) {
		return
	}
	defer ssConn.Close()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)

	// most messages are ssh output, not webSocket input
	go ssConn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go ssConn.SendComboOutput(wsConn, quitChan)
	go ssConn.SessionWait(quitChan)

	<-quitChan
	logrus.Info("websocket finished")
}
