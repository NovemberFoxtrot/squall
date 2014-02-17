squall
======

~~~
squall client_id=CLIENT_ID api_key=API_KEY droplets
~~~

// https://api.digitalocean.com/droplets
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

// https://api.digitalocean.com/images
// https://api.digitalocean.com/images/[image_id]/
// https://api.digitalocean.com/images/[image_id]/destroy/
// https://api.digitalocean.com/images/[image_id]/transfer/?region_id=[region_id]

// https://api.digitalocean.com/ssh_keys/new/?name=[ssh_key_name]&ssh_pub_key=[ssh_public_key]
// https://api.digitalocean.com/ssh_keys
// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/
// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/edit/
// https://api.digitalocean.com/ssh_keys/[ssh_key_id]/destroy/

// https://api.digitalocean.com/domains/new?name=[domain]&ip_address=[ip_address]

// https://api.digitalocean.com/domains
// https://api.digitalocean.com/domains/[domain_id]
// https://api.digitalocean.com/domains/[domain_id]/destroy
// https://api.digitalocean.com/domains/[domain_id]/records
// https://api.digitalocean.com/domains/[domain_id]/records/new?record_type=[record_type]&data=[data]
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/edit
// https://api.digitalocean.com/domains/[domain_id]/records/[record_id]/destroy

// https://api.digitalocean.com/events/[event_id]/
