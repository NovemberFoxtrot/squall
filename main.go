package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)

		if err != nil {
			return nil, err
		}

		conn.SetDeadline(time.Now().Add(rwTimeout))

		return conn, nil
	}
}

func FetchURL(theurl string) string {
	var client *http.Client

	if proxy := os.Getenv("http_proxy"); proxy != `` {
		proxyUrl, err := url.Parse(proxy)

		if err != nil {
			fmt.Println(err)
		}

		transport := http.Transport{
			Dial:  TimeoutDialer(5*time.Second, 5*time.Second),
			Proxy: http.ProxyURL(proxyUrl),
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

type Droplet struct {
	Backups          []interface{} `json:"backups"`
	BackupsActive    bool          `json:"backups_active"`
	CreatedAt        string        `json:"created_at"`
	EventId          int           `json:"event_id"`
	Id               int           `json:"id"`
	ImageId          int           `json:"image_id"`
	IpAddress        string        `json:"ip_address"`
	Locked           bool          `json:"locked"`
	Name             string        `json:"name"`
	PrivateIpAddress interface{}   `json:"private_ip_address"`
	RegionId         int           `json:"region_id"`
	SizeId           int           `json:"size_id"`
	Snapshots        []interface{} `json:"snapshots"`
	Status           string        `json:"status"`
}

type Droplets struct {
	Status     string    `json:"status"`
	Droplets   []Droplet `json:"droplets"`
	TheDroplet Droplet   `json:"droplet"`
	Message    string    `json:"message"`
	EventId    int       `json:"event_id"`
	Regions    []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"regions"`
	Images []struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		Distribution string `json:"distribution"`
	} `json:"images"`
	Image struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		Distribution string `json:"distribution"`
	} `json:"image"`
	SshKeys []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"ssh_keys"`
	SshKey struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		SshPubKey string `json:"ssh_pub_key"`
	} `json:"ssh_key"`
	Sizes []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"sizes"`
	Domains []struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Ttl               int         `json:"ttl"`
		LiveZoneFile      string      `json:"live_zone_file"`
		Error             interface{} `json:"error"`
		ZoneFileWithError interface{} `json:"zone_file_with_error"`
	} `json:"domains"`
	Domain struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Ttl               int         `json:"ttl"`
		LiveZoneFile      string      `json:"live_zone_file"`
		Error             interface{} `json:"error"`
		ZoneFileWithError interface{} `json:"zone_file_with_error"`
	} `json:"domain"`
	Record struct {
		Data       string      `json:"data"`
		DomainId   string      `json:"domain_id"`
		Id         int         `json:"id"`
		Name       string      `json:"name"`
		Port       interface{} `json:"port"`
		Priority   interface{} `json:"priority"`
		RecordType string      `json:"record_type"`
		Weight     interface{} `json:"weight"`
	} `json:"record"`
	Records []struct {
		Id         int         `json:"id"`
		DomainId   string      `json:"domain_id"`
		RecordType string      `json:"record_type"`
		Name       string      `json:"name"`
		Data       string      `json:"data"`
		Priority   interface{} `json:"priority"`
		Port       interface{} `json:"port"`
		Weight     interface{} `json:"weight"`
	} `json:"records"`
	DomainRecord struct {
		Id         int         `json:"id"`
		DomainId   string      `json:"domain_id"`
		RecordType string      `json:"record_type"`
		Name       string      `json:"name"`
		Data       string      `json:"data"`
		Priority   interface{} `json:"priority"`
		Port       interface{} `json:"port"`
		Weight     interface{} `json:"weight"`
	} `json:"domain_record"`
	Event struct {
		Id           int    `json:"id"`
		ActionStatus string `json:"action_status"`
		DropletId    int    `json:"droplet_id"`
		EventTypeId  int    `json:"event_type_id"`
		Percentage   string `json:"percentage"`
	} `json:"event"`
}

// https://api.digitalocean.com/droplets/new?name=[droplet_name]&size_id=[size_id]&image_id=[image_id]&region_id=[region_id]&ssh_key_ids=[ssh_key_id1],[ssh_key_id2]

// https://api.digitalocean.com/droplets/[droplet_id]
// https://api.digitalocean.com/droplets/[droplet_id]/reboot/
// https://api.digitalocean.com/droplets/[droplet_id]/power_cycle/
// https://api.digitalocean.com/droplets/[droplet_id]/shutdown/
// https://api.digitalocean.com/droplets/[droplet_id]/power_off/
// https://api.digitalocean.com/droplets/[droplet_id]/power_on/
// https://api.digitalocean.com/droplets/[droplet_id]/password_reset/
// https://api.digitalocean.com/droplets/[droplet_id]/resize/?size_id=[size_id]
// https://api.digitalocean.com/droplets/[droplet_id]/snapshot/?name=[snapshot_name]
// https://api.digitalocean.com/droplets/[droplet_id]/restore/?image_id=[image_id]
// https://api.digitalocean.com/droplets/[droplet_id]/rebuild/?image_id=[image_id]
// https://api.digitalocean.com/droplets/[droplet_id]/rename/?name=[name]
// https://api.digitalocean.com/droplets/[droplet_id]/destroy/

// https://api.digitalocean.com/images/[image_id]/
// https://api.digitalocean.com/images/[image_id]/destroy/
// https://api.digitalocean.com/images/[image_id]/transfer/?region_id=[region_id]

// https://api.digitalocean.com/ssh_keys/new/?name=[ssh_key_name]&ssh_pub_key=[ssh_public_key]

// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/
// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/edit/
// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/destroy/

// https://api.digitalocean.com/domains/new?name=[domain]&ip_address=[ip_address]

// https://api.digitalocean.com/domains/[domain_id]
// https://api.digitalocean.com/domains/[domain_id]/destroy
// https://api.digitalocean.com/domains/[domain_id]/records
// https://api.digitalocean.com/domains/[domain_id]/records/new?record_type=[record_type]&data=[data]
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/edit
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/destroy

// https://api.digitalocean.com/events/[event_id]/

func main() {
	clientID := os.Args[1]
	apiKey := os.Args[2]

	url := `https://api.digitalocean.com/`

	switch os.Args[3] {
	case `domains`, `droplets`, `events`, `images`, `regions`, `sizes`, `ssh_keys`:
		url = url + os.Args[3]
	}

	theString := FetchURL(url + `?client_id=` + clientID + `&api_key=` + apiKey)

	var droplets Droplets

	json.Unmarshal([]byte(theString), &droplets)

	fmt.Println(droplets)
}
