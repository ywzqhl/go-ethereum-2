package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	NodeID        string
	Endpoint      string
	PeersChan     chan []string
	BroadcastChan chan string
	ConsumeChan   chan string
	RelayChan     chan string
	DoneChan      chan string
}

func NewClient(nodeID string) *Client {
	return &Client{
		NodeID:        nodeID,
		Endpoint:      "http://localhost:8585/event", // Default. Override if you'd like to change,
		PeersChan:     make(chan []string),
		BroadcastChan: make(chan string),
		ConsumeChan:   make(chan string),
		RelayChan:     make(chan string),
		DoneChan:      make(chan string),
	}
}

func (self *Client) LogPeers(peers []string) {
	self.PeersChan <- peers
	//data := self.initData("peers")
	//data["peers"] = peers
	//self.postEvent(data)
}

func (self *Client) LogBroadcast(streamID string) {
	self.BroadcastChan <- streamID
}

func (self *Client) LogConsume(streamID string) {
	self.ConsumeChan <- streamID
}

func (self *Client) LogRelay(streamID string) {
	self.RelayChan <- streamID
}

func (self *Client) LogDone(streamID string) {
	self.DoneChan <- streamID
}

func (self *Client) InitData(eventName string) (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["name"] = eventName
	data["node"] = self.NodeID
	return
}

func (self *Client) PostEvent(data map[string]interface{}) {
	enc, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", self.Endpoint, bytes.NewBuffer(enc))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Couldn't connect to the event server", err)
		return
	}
	defer resp.Body.Close()
}
