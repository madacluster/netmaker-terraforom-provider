package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Generated by https://quicktype.io

type Network struct {
	Addressrange        string        `json:"addressrange"`
	Addressrange6       string        `json:"addressrange6"`
	Displayname         string        `json:"displayname"`
	Netid               string        `json:"netid"`
	Nodeslastmodified   int64         `json:"nodeslastmodified"`
	Networklastmodified int64         `json:"networklastmodified"`
	Defaultinterface    string        `json:"defaultinterface"`
	Defaultlistenport   int64         `json:"defaultlistenport"`
	Nodelimit           int64         `json:"nodelimit"`
	Defaultpostup       string        `json:"defaultpostup"`
	Defaultpostdown     string        `json:"defaultpostdown"`
	Keyupdatetimestamp  int64         `json:"keyupdatetimestamp"`
	Defaultkeepalive    int64         `json:"defaultkeepalive"`
	Defaultsaveconfig   string        `json:"defaultsaveconfig"`
	Accesskeys          []interface{} `json:"accesskeys"`
	Allowmanualsignup   string        `json:"allowmanualsignup"`
	Islocal             string        `json:"islocal"`
	Isdualstack         string        `json:"isdualstack"`
	Isipv4              string        `json:"isipv4"`
	Isipv6              string        `json:"isipv6"`
	Isgrpchub           string        `json:"isgrpchub"`
	Localrange          string        `json:"localrange"`
	Checkininterval     int64         `json:"checkininterval"`
	Defaultudpholepunch string        `json:"defaultudpholepunch"`
	Defaultextclientdns string        `json:"defaultextclientdns"`
	Defaultmtu          int64         `json:"defaultmtu"`
}

// GetNetworks - Returns list of coffees (no auth required)
func (c *Client) GetNetworks() ([]Network, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	networks := []Network{}
	err = json.Unmarshal(body, &networks)
	if err != nil {
		return nil, err
	}

	return networks, nil
}

// GetNetworks - Returns list of coffees (no auth required)
func (c *Client) GetNetwork(networkID string) (*Network, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	network := Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// GetNetworks - Returns list of coffees (no auth required)
func (c *Client) CreateNetwork(networkID, addressrange string) (*Network, error) {
	network := Network{
		Addressrange: addressrange,
		Netid:        networkID,
	}
	rb, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/networks", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// UpdateNetwork - Updates a network
func (c *Client) UpdateNetwork(data map[string]string) (*Network, error) {
	network := Network{}
	mapFiels(data, &network)
	rb, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/networks/%s", c.HostURL, network.Netid), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	network = Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

func (c *Client) DeleteNetwork(networkID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/networks/%s", c.HostURL, networkID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func mapFiels(data map[string]string, network *Network) {
	for k, v := range data {
		switch k {
		case "displayname":
			network.Displayname = v
		case "addressrange":
			network.Addressrange = v
		case "netid":
			network.Netid = v
		case "islocal":
			network.Islocal = v
		case "isDualStack":
			network.Isdualstack = v
		case "isIPv4":
			network.Isipv4 = v
		case "isIPv6":
			network.Isipv6 = v
		case "isGRPCHub":
			network.Isgrpchub = v
		case "localrange":
			network.Localrange = v
		case "checkininterval":
			i, _ := strconv.Atoi(v)
			network.Checkininterval = int64(i)
		case "defaultudpholepunch":
			network.Defaultudpholepunch = v
		case "defaultextclientdns":
			network.Defaultextclientdns = v
		case "defaultmtu":
			i, _ := strconv.Atoi(v)
			network.Defaultmtu = int64(i)
		case "defaultkeepalive":
			i, _ := strconv.Atoi(v)
			network.Defaultkeepalive = int64(i)
		case "allowmanualsignup":
			network.Allowmanualsignup = v
		case "nodeslastmodified":
			network.Nodeslastmodified, _ = strconv.ParseInt(v, 10, 64)
		case "networklastmodified":
			network.Networklastmodified, _ = strconv.ParseInt(v, 10, 64)
		case "defaultinterface":
			network.Defaultinterface = v
		case "defaultlistenport":
			network.Defaultlistenport, _ = strconv.ParseInt(v, 10, 64)
		case "defaultsaveconfig":
			network.Defaultsaveconfig = v
		case "nodelimit":
			network.Nodelimit, _ = strconv.ParseInt(v, 10, 64)
		case "defaultpostup":
			network.Defaultpostup = v
		case "defaultpostdown":
			network.Defaultpostdown = v
		case "keyupdatetimestamp":
			network.Keyupdatetimestamp, _ = strconv.ParseInt(v, 10, 64)
		}
	}
}

func mapFielsRevert(network *Network) map[string]string {
	data := make(map[string]string)
	data["displayname"] = network.Displayname
	data["addressrange"] = network.Addressrange
	data["netid"] = network.Netid
	data["nodeslastmodified"] = strconv.FormatInt(network.Nodeslastmodified, 10)
	data["networklastmodified"] = strconv.FormatInt(network.Networklastmodified, 10)
	data["defaultinterface"] = network.Defaultinterface
	data["defaultlistenport"] = strconv.FormatInt(network.Defaultlistenport, 10)
	data["nodelimit"] = strconv.FormatInt(network.Nodelimit, 10)
	data["defaultpostup"] = network.Defaultpostup
	data["defaultpostdown"] = network.Defaultpostdown
	data["keyupdatetimestamp"] = strconv.FormatInt(network.Keyupdatetimestamp, 10)
	data["defaultsaveconfig"] = network.Defaultsaveconfig
	data["defaultmtu"] = strconv.FormatInt(network.Defaultmtu, 10)
	data["defaultkeepalive"] = strconv.FormatInt(network.Defaultkeepalive, 10)
	data["allowmanualsignup"] = network.Allowmanualsignup
	data["defaultudpholepunch"] = network.Defaultudpholepunch
	data["defaultextclientdns"] = network.Defaultextclientdns
	data["islocal"] = network.Islocal
	data["isDualStack"] = network.Isdualstack
	data["isIPv4"] = network.Isipv4
	data["isIPv6"] = network.Isipv6
	data["isGRPCHub"] = network.Isgrpchub
	data["localrange"] = network.Localrange
	data["checkininterval"] = strconv.FormatInt(network.Checkininterval, 10)
	return data
}
