// Author: yangzq80@gmail.com
// Date: 2021-04-19
//
package ssh

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type sshHost struct {
	Host     string `form:"host"` //22
	User     string `form:"user"`
	Port     string `form:"port"`
	PwdType  string `form:"pwdType"`
	Password string `form:"password"`
	KeyFile  string `form:"keyFile"`
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
	host := new(sshHost)
	c.ShouldBind(&host)

	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if handleError(c, err) {
		return
	}
	defer wsConn.Close()
	if handleWsError(wsConn, err) {
		return
	}

	var client *Client

	if host.Port == "" {
		host.Port = "22"
	}

	if host.PwdType == "key" {
		client, err = DialWithKey(host.Host+":"+host.Port, host.User, host.KeyFile)
	} else {
		client, err = DialWithPasswd(host.Host+":"+host.Port, host.User, host.Password)
	}

	defer client.Close()

	if handleError(c, err) {
		return
	}

	session, _ := NewSshWsSession(client, wsConn)
	defer session.Close()

	session.Start()
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		logrus.WithError(err).Error("gin context http handler error")
		c.AbortWithStatusJSON(200, gin.H{"ok": false, "msg": err.Error()})
		return true
	}
	return false
}

func handleWsError(ws *websocket.Conn, err error) bool {
	if err != nil {
		logrus.WithError(err).Error("handler ws ERROR:")
		dt := time.Now().Add(time.Second)
		if err := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); err != nil {
			logrus.WithError(err).Error("websocket writes control message failed:")
		}
		return true
	}
	return false
}
