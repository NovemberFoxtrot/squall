package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

func timeoutdialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)

		if err != nil {
			return nil, err
		}

		conn.SetDeadline(time.Now().Add(rwTimeout))

		return conn, nil
	}
}

func fetchURL(theurl string) string {
	var client *http.Client

	if proxy := os.Getenv("http_proxy"); proxy != `` {
		proxyURL, err := url.Parse(proxy)

		if err != nil {
			fmt.Println(err)
		}

		transport := http.Transport{
			Dial:  timeoutdialer(5*time.Second, 5*time.Second),
			Proxy: http.ProxyURL(proxyURL),
		}

		client = &http.Client{Transport: &transport}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest(`GET`, theurl, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

type droplet struct {
	Backups          []interface{} `json:"backups"`
	BackupsActive    bool          `json:"backups_active"`
	CreatedAt        string        `json:"created_at"`
	EventID          int           `json:"event_id"`
	ID               int           `json:"id"`
	ImageID          int           `json:"image_id"`
	IPAddress        string        `json:"ip_address"`
	Locked           bool          `json:"locked"`
	Name             string        `json:"name"`
	PrivateIPAddress interface{}   `json:"private_ip_address"`
	RegionID         int           `json:"region_id"`
	SizeID           int           `json:"size_id"`
	Snapshots        []interface{} `json:"snapshots"`
	Status           string        `json:"status"`
}

type record struct {
	Data       string      `json:"data"`
	DomainID   string      `json:"domain_id"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Port       interface{} `json:"port"`
	Priority   interface{} `json:"priority"`
	RecordType string      `json:"record_type"`
	Weight     interface{} `json:"weight"`
}

type domain struct {
	Error             interface{} `json:"error"`
	ID                int         `json:"id"`
	LiveZoneFile      string      `json:"live_zone_file"`
	Name              string      `json:"name"`
	Ttl               int         `json:"ttl"`
	ZoneFileWithError interface{} `json:"zone_file_with_error"`
}

type event struct {
	ActionStatus string `json:"action_status"`
	DropletID    int    `json:"droplet_id"`
	EventTypeID  int    `json:"event_type_id"`
	ID           int    `json:"id"`
	Percentage   string `json:"percentage"`
}

type image struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
}

type sshkey struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SshPubKey string `json:"ssh_pub_key"`
}

type domainrecord struct {
	ID         int         `json:"id"`
	DomainID   string      `json:"domain_id"`
	RecordType string      `json:"record_type"`
	Name       string      `json:"name"`
	Data       string      `json:"data"`
	Priority   interface{} `json:"priority"`
	Port       interface{} `json:"port"`
	Weight     interface{} `json:"weight"`
}

type regions []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type sizes []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type squall struct {
	Domain       domain       `json:"domain"`
	DomainRecord domainrecord `json:"domain_record"`
	Domains      []domain     `json:"domains"`
	Droplet      droplet      `json:"droplet"`
	Droplets     []droplet    `json:"droplets"`
	Event        event        `json:"event"`
	EventID      int          `json:"event_id"`
	Image        image        `json:"image"`
	Images       []image      `json:"images"`
	Message      string       `json:"message"`
	Record       record       `json:"record"`
	Records      []record     `json:"records"`
	Regions      []regions    `json:"regions"`
	Sizes        []sizes      `json:"sizes"`
	SshKey       sshkey       `json:"ssh_key"`
	SshKeys      []sshkey     `json:"ssh_keys"`
	Status       string       `json:"status"`
}

func main() {
	var path []string
	var opts []string
	var theURL string

	path = append(path, `https://api.digitalocean.com`)

	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, `=`) {
			opts = append(opts, arg)
		} else {
			path = append(path, arg)
		}
	}

	theURL = strings.Join(path, `/`) + `?` + strings.Join(opts, `&`)

	theString := fetchURL(theURL)

	var s squall

	json.Unmarshal([]byte(theString), &s)

	fmt.Println(s)
}
