package transmission

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Transmission struct {
	session *string
}

var (
	Rpc       Transmission
	sessionId string
	once      sync.Once
)

func CreateSession() error {
	fmt.Println("Creating a transmission session...")
	res, err := http.Get("http://127.0.0.1:9091/transmission/rpc")
	if err != nil {
		if err = err.(*url.Error).Err; strings.Contains(err.Error(), "connection refused") {
			once.Do(func() {
				fmt.Printf("\n\x1b[31mNo daemon is running!\x1b[0m Starting the transmission daemon...\n")
				err = startDaemon()
			})
			return err
		} else {
			return err
		}
	}
	defer res.Body.Close()

	sessionId = res.Header.Get("x-transmission-session-id")
	Rpc = Transmission{session: &sessionId}
	fmt.Printf("\x1b[32mThe transmission session has been created\x1b[0m\n\n")

	return nil
}

func startDaemon() error {
	err := exec.Command("./transmission-daemon").Run()
	if err != nil {
		return err
	}
	time.Sleep(time.Second)
	fmt.Printf("\x1b[32mTransmission daemon has started!\x1b[0m\n\n")

	return CreateSession()
}

func (t Transmission) SessionClose() error {
	var basicResponse BasicResponse

	res, err := t.client("session-close", map[string]interface{}{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&basicResponse)

	if basicResponse.Result != "success" {
		return fmt.Errorf("error closing session")
	}

	fmt.Printf("\x1b[32m\nTransmission session has been successfully closed\x1b[0m\n\n")
	return nil
}

func (t Transmission) SessionGet() error {
	var sessionResponse SessionResponse

	res, err := t.client("session-get", map[string]interface{}{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&sessionResponse)

	fmt.Printf("\nCurren session ID is: %v\n", sessionResponse.Arguments.SessionID)

	return nil
}

func (t Transmission) client(method string, arguments map[string]interface{}) (*http.Response, error) {
	url := "http://127.0.0.1:9091/transmission/rpc"

	var buf bytes.Buffer
	jsonReq := BasicRequest{
		Method:    method,
		Arguments: arguments,
	}

	json.NewEncoder(&buf).Encode(jsonReq)

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if t.session != nil {
		req.Header.Add("X-Transmission-Session-Id", *t.session)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusConflict {
		sessionId = res.Header.Get("x-transmission-session-id")
		req.Header.Add("X-Transmission-Session-Id", *t.session)

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
