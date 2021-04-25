// Author: yangzq80@gmail.com
// Date: 2021-04-20
// SSH协议数据<->Websocket协议<->浏览器xterm.js
package ssh

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"sync"
	"time"
)

type SshWsSession struct {
	sshSession *ssh.Session
	sshPipe    io.WriteCloser
	sshOutput  *safeBuffer
	wsConn     *websocket.Conn
}

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
}

func NewSshWsSession(client *Client, wsConn *websocket.Conn) (*SshWsSession, error) {
	sshSession, err := client.client.NewSession()
	if err != nil {
		return nil, err
	}
	stdinPipe, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout := new(safeBuffer)

	sshSession.Stdout = stdout
	sshSession.Stderr = stdout

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := sshSession.RequestPty("xterm", 100, 880, modes); err != nil {
		return nil, err
	}

	// Start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SshWsSession{
		sshPipe:    stdinPipe,
		sshOutput:  stdout,
		sshSession: sshSession,
		wsConn:     wsConn,
	}, nil
}

func (s *SshWsSession) Start() {
	c := make(chan bool, 3)
	go s.wsToSsh(c)
	go s.sshToWs(c)
	s.Wait(c)
	<-c
	logrus.Info("Stopped sshWsSession.")
}

func (s *SshWsSession) wsToSsh(c chan bool) {
	logrus.Info("Staring accept ws data...")
	wsConn := s.wsConn
	//tells other go routine quit
	defer setQuit(c)
	for {
		select {
		case <-c:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if err != nil {
				logrus.WithError(err).Error("reading webSocket message failed")
				return
			}
			//unmashal bytes into struct
			msg := wsMsg{}
			if err := json.Unmarshal(wsData, &msg); err != nil {
				logrus.WithError(err).WithField("wsData", string(wsData)).Error("unmarshal websocket message failed")
			}
			switch msg.Type {
			case "cmd":
				if _, err := s.sshPipe.Write([]byte(msg.Cmd)); err != nil {
					logrus.WithError(err).Error("ws cmd bytes write to ssh.stdin pipe failed")
				}
			}
		}
	}
	logrus.Info("Stropped wsToSsh...")
}

func (s *SshWsSession) sshToWs(c chan bool) {

	logrus.Info("Stating accept ssh data...")

	wsCon := s.wsConn
	//tells other go routine quit
	defer setQuit(c)

	//every 120ms write combine output bytes into websocket response
	tick := time.NewTicker(time.Millisecond * time.Duration(60))
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if s.sshOutput.Bytes() == nil {
				return
			}
			bs := s.sshOutput.Bytes()
			if len(bs) > 0 {
				err := wsCon.WriteMessage(websocket.TextMessage, bs)
				if err != nil {
					logrus.WithError(err).Error("ssh sending combo output to webSocket failed")
				}
				s.sshOutput.Reset()
			}
		case <-c:
			return
		}
	}
	logrus.Info("Stopped sshToWs...")
}

func (s *SshWsSession) Close() {
	if s.sshSession != nil {
		s.sshSession.Close()
	}
	if s.sshOutput != nil {
		s.sshOutput = nil
	}
}

func (sws *SshWsSession) Wait(quitChan chan bool) {
	if err := sws.sshSession.Wait(); err != nil {
		logrus.WithError(err).Error("ssh session wait failed")
		setQuit(quitChan)
	}
}

func setQuit(c chan bool) {
	c <- true
}

type safeBuffer struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *safeBuffer) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}
func (w *safeBuffer) Bytes() []byte {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Bytes()
}
func (w *safeBuffer) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.buffer.Reset()
}
