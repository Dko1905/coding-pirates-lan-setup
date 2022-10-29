# Config & such for Coding Pirates lan party
**REMEMBER TO CHANGE IPS AND PORTS IF NECESSARY**

## [mc-ffa-stack](./mc-ffa-stack/mc-ffa-stack.yml)
Uses latest PAPER minecraft version.

## [monitoring-stack](./monitoring-stack/monitoring-stack.yml)
Sets up node-exporter, prometheus and grafana in one stack.
Includes config file which needs to be placed in `prometheus_conf` volume.

If you want to add separate machine that only collects data using node-explorer, then remove
prometheus and grafana from the stack, and configure ports section. 

### Grafana
Grafana dashboard config (for two hosts): [file](./monitoring-stack/Daniel-dashboard-1667037052800.json)

## [lancache-stack](./lancache-stack/lancache-stack.yml)
| name | value |
| ---- | ----- |
| USE_GENERIC_CACHE | true |
| LANCACHE_IP | 10.42.42.100 10.42.42.101 10.42.42.102 10.42.42.103 10.42.42.104 10.42.42.105 10.42.42.106 10.42.42.107 10.42.42.108 10.42.42.109 |
| DNS_BIND_IP | 192.168.42.2 |
| UPSTREAM_DNS | 192.168.42.1 |
| CACHE_ROOT | /tank/lancache-data |
| CACHE_DISK_SIZE | 1500000m |
| CACHE_INDEX_SIZE | 500m |
| CACHE_MAX_AGE | 3650d |
| TZ | Europe/Copenhagen |
