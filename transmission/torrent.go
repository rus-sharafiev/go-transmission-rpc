package transmission

import (
	"encoding/json"
	"fmt"
)

func (t Transmission) TorrentGet() error {
	var sessionResponse SessionResponse

	res, err := t.client("torrent-get", map[string]interface{}{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&sessionResponse)

	fmt.Printf("\nCurren session ID is: %v\n", sessionResponse.Arguments.SessionID)

	return nil
}

func (t Transmission) TorrentAdd() error {
	var sessionResponse SessionResponse

	res, err := t.client("torrent-add", map[string]interface{}{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&sessionResponse)

	fmt.Printf("\nCurren session ID is: %v\n", sessionResponse.Arguments.SessionID)

	return nil
}

func (t Transmission) TorrentRemove() error {
	var sessionResponse SessionResponse

	res, err := t.client("torrent-remove", map[string]interface{}{})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&sessionResponse)

	fmt.Printf("\nCurren session ID is: %v\n", sessionResponse.Arguments.SessionID)

	return nil
}
