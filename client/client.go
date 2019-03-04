package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	emby "github.com/encabulators/emby-go"
)

// EmbyClient is a client of the Emby API
type EmbyClient struct {
	hostURL string
	apiKey  string
}

// New creates a new Emby client pointed at a specific server
func New(hostURL string, apiKey string) *EmbyClient {
	return &EmbyClient{
		hostURL, apiKey,
	}
}

// GetSystemInfo retrieves server information corresponding to the System/Info API endpoint
func (c *EmbyClient) GetSystemInfo() (*emby.SystemInfo, error) {
	raw, err := c.makeRequest("GET", "System/Info", "")
	if err != nil {
		return nil, err
	}
	sysInfo := &emby.SystemInfo{}
	err = json.Unmarshal(raw, sysInfo)
	if err != nil {
		return nil, err
	}
	return sysInfo, nil
}

// GetViewsForUser retrieves a list of media views for a given user.
func (c *EmbyClient) GetViewsForUser(userID string) (*emby.UserViewsResponse, error) {
	raw, err := c.makeRequest("GET", fmt.Sprintf("Users/%s/Views", userID), "")
	if err != nil {
		return nil, err
	}
	views := &emby.UserViewsResponse{}
	err = json.Unmarshal(raw, views)
	if err != nil {
		return nil, err
	}
	return views, nil
}

// GetViewItems retrieves the list of media items contained within a given user view
func (c *EmbyClient) GetViewItems(userID string, viewID string) (*emby.MediaItemList, error) {
	raw, err := c.makeRequest("GET", fmt.Sprintf("Users/%s/Items?parentId=%s", userID, viewID), "")
	if err != nil {
		return nil, err
	}
	items := &emby.MediaItemList{}
	err = json.Unmarshal(raw, items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetItemDetail retrieves all of the available details for a specific media item
func (c *EmbyClient) GetItemDetail(userID string, itemID string) (*emby.MediaItemDetail, error) {
	raw, err := c.makeRequest("GET", fmt.Sprintf("Users/%s/Items/%s", userID, itemID), "")
	if err != nil {
		return nil, err
	}
	detail := &emby.MediaItemDetail{}
	err = json.Unmarshal(raw, detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func (c *EmbyClient) makeRequest(method string, URL string, body string) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.hostURL, URL), strings.NewReader(body))
	req.Header.Set("X-Emby-Token", c.apiKey)
	req.Header.Set("Application-Type", "application/json")

	if len(body) > 0 {
		bodybytes := []byte(body)
		buf := bytes.NewBuffer(bodybytes)
		req.Body = ioutil.NopCloser(buf)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respbody, _ := ioutil.ReadAll(resp.Body)
	return respbody, nil
}
