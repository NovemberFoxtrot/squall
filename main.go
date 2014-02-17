package main

import (
	"fmt"
)

https://api.digitalocean.com/droplets/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Droplets []struct {
		Id int `json:"id"`
		Name string `json:"name"`
		ImageId int `json:"image_id"`
		SizeId int `json:"size_id"`
		RegionId int `json:"region_id"`
		BackupsActive bool `json:"backups_active"`
		IpAddress string `json:"ip_address"`
		PrivateIpAddress interface{} `json:"private_ip_address"`
		Locked bool `json:"locked"`
		Status string `json:"status"`
		CreatedAt string `json:"created_at"`
	} `json:"droplets"`
}

https://api.digitalocean.com/droplets/new?client_id=[your_client_id]&api_key=[your_api_key]&name=[droplet_name]&size_id=[size_id]&image_id=[image_id]&region_id=[region_id]&ssh_key_ids=[ssh_key_id1],[ssh_key_id2]

type Droplets struct {
	Status string `json:"status"`
	Droplet struct {
		Id int `json:"id"`
		Name string `json:"name"`
		ImageId int `json:"image_id"`
		SizeId int `json:"size_id"`
		EventId int `json:"event_id"`
	} `json:"droplet"`
}

https://api.digitalocean.com/droplets/[droplet_id]?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Droplet struct {
		Id int `json:"id"`
		ImageId int `json:"image_id"`
		Name string `json:"name"`
		RegionId int `json:"region_id"`
		SizeId int `json:"size_id"`
		BackupsActive bool `json:"backups_active"`
		Backups []undefined `json:"backups"`
		Snapshots []undefined `json:"snapshots"`
		IpAddress string `json:"ip_address"`
		PrivateIpAddress interface{} `json:"private_ip_address"`
		Locked bool `json:"locked"`
		Status string `json:"status"`
	} `json:"droplet"`
}

https://api.digitalocean.com/droplets/[droplet_id]/reboot/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/power_cycle/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/shutdown/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/power_off/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/power_on/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/password_reset/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/resize/?size_id=[size_id]&client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/snapshot/?name=[snapshot_name]&client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/restore/?image_id=[image_id]&client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/rebuild/?image_id=[image_id]&client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/rename/?client_id=[your_client_id]&api_key=[your_api_key]&name=[name]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/droplets/[droplet_id]/destroy/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/regions/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Regions []struct {
		Id int `json:"id"`
		Name string `json:"name"`
	} `json:"regions"`
}

https://api.digitalocean.com/images/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Images []struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Distribution string `json:"distribution"`
	} `json:"images"`
}

https://api.digitalocean.com/images/[image_id]/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Image struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Distribution string `json:"distribution"`
	} `json:"image"`
}

https://api.digitalocean.com/images/[image_id]/destroy/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
}

https://api.digitalocean.com/images/[image_id]/transfer/?client_id=[your_client_id]&api_key=[your_api_key]&region_id=[region_id]

type Droplets struct {
	Status string `json:"status"`
	EventId int `json:"event_id"`
}

https://api.digitalocean.com/ssh_keys/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	SshKeys []struct {
		Id int `json:"id"`
		Name string `json:"name"`
	} `json:"ssh_keys"`
}

https://api.digitalocean.com/ssh_keys/new/?name=[ssh_key_name]&ssh_pub_key=[ssh_public_key]&client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	SshKey struct {
		Id int `json:"id"`
		Name string `json:"name"`
		SshPubKey string `json:"ssh_pub_key"`
	} `json:"ssh_key"`
}

https://api.digitalocean.com/ssh_keys/[ssh_key_id]/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	SshKey struct {
		Id int `json:"id"`
		Name string `json:"name"`
		SshPubKey string `json:"ssh_pub_key"`
	} `json:"ssh_key"`
}

https://api.digitalocean.com/ssh_keys/[ssh_key_id]/edit/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	SshKey struct {
		Id int `json:"id"`
		Name string `json:"name"`
		SshPubKey string `json:"ssh_pub_key"`
	} `json:"ssh_key"`
}

https://api.digitalocean.com/ssh_keys/[ssh_key_id]/destroy/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
}

https://api.digitalocean.com/sizes/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Sizes []struct {
		Id int `json:"id"`
		Name string `json:"name"`
	} `json:"sizes"`
}

https://api.digitalocean.com/domains?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Domains []struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Ttl int `json:"ttl"`
		LiveZoneFile string `json:"live_zone_file"`
		Error interface{} `json:"error"`
		ZoneFileWithError interface{} `json:"zone_file_with_error"`
	} `json:"domains"`
}

https://api.digitalocean.com/domains/new?client_id=[your_client_id]&api_key=[your_api_key]&name=[domain]&ip_address=[ip_address]

type Droplets struct {
	Status string `json:"status"`
	Domain struct {
		Id int `json:"id"`
		Name string `json:"name"`
	} `json:"domain"`
}

https://api.digitalocean.com/domains/[domain_id]?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Domain struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Ttl int `json:"ttl"`
		LiveZoneFile string `json:"live_zone_file"`
		Error interface{} `json:"error"`
		ZoneFileWithError interface{} `json:"zone_file_with_error"`
	} `json:"domain"`
}

https://api.digitalocean.com/domains/[domain_id]/destroy?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
}

https://api.digitalocean.com/domains/[domain_id]/records?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Records []struct {
		Id int `json:"id"`
		DomainId string `json:"domain_id"`
		RecordType string `json:"record_type"`
		Name string `json:"name"`
		Data string `json:"data"`
		Priority interface{} `json:"priority"`
		Port interface{} `json:"port"`
		Weight interface{} `json:"weight"`
	} `json:"records"`
}

https://api.digitalocean.com/domains/[domain_id]/records/new?client_id=[your_client_id]&api_key=[your_api_key]&record_type=[record_type]&data=[data]

type Droplets struct {
	Status string `json:"status"`
	DomainRecord struct {
		Id int `json:"id"`
		DomainId string `json:"domain_id"`
		RecordType string `json:"record_type"`
		Name string `json:"name"`
		Data string `json:"data"`
		Priority interface{} `json:"priority"`
		Port interface{} `json:"port"`
		Weight interface{} `json:"weight"`
	} `json:"domain_record"`
}

https://api.digitalocean.com/domains/[domain_id]/records/[record_id]?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Record struct {
		Id int `json:"id"`
		DomainId string `json:"domain_id"`
		RecordType string `json:"record_type"`
		Name string `json:"name"`
		Data string `json:"data"`
		Priority interface{} `json:"priority"`
		Port interface{} `json:"port"`
		Weight interface{} `json:"weight"`
	} `json:"record"`
}

https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/edit?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Record struct {
		Id int `json:"id"`
		DomainId string `json:"domain_id"`
		RecordType string `json:"record_type"`
		Name string `json:"name"`
		Data string `json:"data"`
		Priority interface{} `json:"priority"`
		Port interface{} `json:"port"`
		Weight interface{} `json:"weight"`
	} `json:"record"`
}

https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/destroy?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
}

https://api.digitalocean.com/events/[event_id]/?client_id=[your_client_id]&api_key=[your_api_key]

type Droplets struct {
	Status string `json:"status"`
	Event struct {
		Id int `json:"id"`
		ActionStatus string `json:"action_status"`
		DropletId int `json:"droplet_id"`
		EventTypeId int `json:"event_type_id"`
		Percentage string `json:"percentage"`
	} `json:"event"`
}

status: 401

type Droplets struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

status: 404

type Droplets struct {
	Status string `json:"status"`
	Message string `json:"message"`
}


func main() {
}
