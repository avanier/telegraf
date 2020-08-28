package all

import (
	_ "github.com/aleveille/telegraf/plugins/inputs/activemq"
	_ "github.com/aleveille/telegraf/plugins/inputs/aerospike"
	_ "github.com/aleveille/telegraf/plugins/inputs/amqp_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/apache"
	_ "github.com/aleveille/telegraf/plugins/inputs/apcupsd"
	_ "github.com/aleveille/telegraf/plugins/inputs/aurora"
	_ "github.com/aleveille/telegraf/plugins/inputs/azure_storage_queue"
	_ "github.com/aleveille/telegraf/plugins/inputs/bcache"
	_ "github.com/aleveille/telegraf/plugins/inputs/beanstalkd"
	_ "github.com/aleveille/telegraf/plugins/inputs/bind"
	_ "github.com/aleveille/telegraf/plugins/inputs/bond"
	_ "github.com/aleveille/telegraf/plugins/inputs/burrow"
	_ "github.com/aleveille/telegraf/plugins/inputs/cassandra"
	_ "github.com/aleveille/telegraf/plugins/inputs/ceph"
	_ "github.com/aleveille/telegraf/plugins/inputs/cgroup"
	_ "github.com/aleveille/telegraf/plugins/inputs/chrony"
	_ "github.com/aleveille/telegraf/plugins/inputs/cisco_telemetry_mdt"
	_ "github.com/aleveille/telegraf/plugins/inputs/clickhouse"
	_ "github.com/aleveille/telegraf/plugins/inputs/cloud_pubsub"
	_ "github.com/aleveille/telegraf/plugins/inputs/cloud_pubsub_push"
	_ "github.com/aleveille/telegraf/plugins/inputs/cloudwatch"
	_ "github.com/aleveille/telegraf/plugins/inputs/conntrack"
	_ "github.com/aleveille/telegraf/plugins/inputs/consul"
	_ "github.com/aleveille/telegraf/plugins/inputs/couchbase"
	_ "github.com/aleveille/telegraf/plugins/inputs/couchdb"
	_ "github.com/aleveille/telegraf/plugins/inputs/cpu"
	_ "github.com/aleveille/telegraf/plugins/inputs/dcos"
	_ "github.com/aleveille/telegraf/plugins/inputs/disk"
	_ "github.com/aleveille/telegraf/plugins/inputs/diskio"
	_ "github.com/aleveille/telegraf/plugins/inputs/disque"
	_ "github.com/aleveille/telegraf/plugins/inputs/dmcache"
	_ "github.com/aleveille/telegraf/plugins/inputs/dns_query"
	_ "github.com/aleveille/telegraf/plugins/inputs/docker"
	_ "github.com/aleveille/telegraf/plugins/inputs/docker_log"
	_ "github.com/aleveille/telegraf/plugins/inputs/dovecot"
	_ "github.com/aleveille/telegraf/plugins/inputs/ecs"
	_ "github.com/aleveille/telegraf/plugins/inputs/elasticsearch"
	_ "github.com/aleveille/telegraf/plugins/inputs/ethtool"
	_ "github.com/aleveille/telegraf/plugins/inputs/eventhub_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/exec"
	_ "github.com/aleveille/telegraf/plugins/inputs/execd"
	_ "github.com/aleveille/telegraf/plugins/inputs/fail2ban"
	_ "github.com/aleveille/telegraf/plugins/inputs/fibaro"
	_ "github.com/aleveille/telegraf/plugins/inputs/file"
	_ "github.com/aleveille/telegraf/plugins/inputs/filecount"
	_ "github.com/aleveille/telegraf/plugins/inputs/filestat"
	_ "github.com/aleveille/telegraf/plugins/inputs/fireboard"
	_ "github.com/aleveille/telegraf/plugins/inputs/fluentd"
	_ "github.com/aleveille/telegraf/plugins/inputs/github"
	_ "github.com/aleveille/telegraf/plugins/inputs/gnmi"
	_ "github.com/aleveille/telegraf/plugins/inputs/graylog"
	_ "github.com/aleveille/telegraf/plugins/inputs/haproxy"
	_ "github.com/aleveille/telegraf/plugins/inputs/hddtemp"
	_ "github.com/aleveille/telegraf/plugins/inputs/http"
	_ "github.com/aleveille/telegraf/plugins/inputs/http_listener_v2"
	_ "github.com/aleveille/telegraf/plugins/inputs/http_response"
	_ "github.com/aleveille/telegraf/plugins/inputs/httpjson"
	_ "github.com/aleveille/telegraf/plugins/inputs/icinga2"
	_ "github.com/aleveille/telegraf/plugins/inputs/infiniband"
	_ "github.com/aleveille/telegraf/plugins/inputs/influxdb"
	_ "github.com/aleveille/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/aleveille/telegraf/plugins/inputs/internal"
	_ "github.com/aleveille/telegraf/plugins/inputs/interrupts"
	_ "github.com/aleveille/telegraf/plugins/inputs/ipmi_sensor"
	_ "github.com/aleveille/telegraf/plugins/inputs/ipset"
	_ "github.com/aleveille/telegraf/plugins/inputs/iptables"
	_ "github.com/aleveille/telegraf/plugins/inputs/ipvs"
	_ "github.com/aleveille/telegraf/plugins/inputs/jenkins"
	_ "github.com/aleveille/telegraf/plugins/inputs/jolokia"
	_ "github.com/aleveille/telegraf/plugins/inputs/jolokia2"
	_ "github.com/aleveille/telegraf/plugins/inputs/jti_openconfig_telemetry"
	_ "github.com/aleveille/telegraf/plugins/inputs/kafka_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/kafka_consumer_legacy"
	_ "github.com/aleveille/telegraf/plugins/inputs/kapacitor"
	_ "github.com/aleveille/telegraf/plugins/inputs/kernel"
	_ "github.com/aleveille/telegraf/plugins/inputs/kernel_vmstat"
	_ "github.com/aleveille/telegraf/plugins/inputs/kibana"
	_ "github.com/aleveille/telegraf/plugins/inputs/kinesis_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/kube_inventory"
	_ "github.com/aleveille/telegraf/plugins/inputs/kubernetes"
	_ "github.com/aleveille/telegraf/plugins/inputs/lanz"
	_ "github.com/aleveille/telegraf/plugins/inputs/leofs"
	_ "github.com/aleveille/telegraf/plugins/inputs/linux_sysctl_fs"
	_ "github.com/aleveille/telegraf/plugins/inputs/logparser"
	_ "github.com/aleveille/telegraf/plugins/inputs/logstash"
	_ "github.com/aleveille/telegraf/plugins/inputs/lustre2"
	_ "github.com/aleveille/telegraf/plugins/inputs/mailchimp"
	_ "github.com/aleveille/telegraf/plugins/inputs/marklogic"
	_ "github.com/aleveille/telegraf/plugins/inputs/mcrouter"
	_ "github.com/aleveille/telegraf/plugins/inputs/mem"
	_ "github.com/aleveille/telegraf/plugins/inputs/memcached"
	_ "github.com/aleveille/telegraf/plugins/inputs/mesos"
	_ "github.com/aleveille/telegraf/plugins/inputs/minecraft"
	_ "github.com/aleveille/telegraf/plugins/inputs/modbus"
	_ "github.com/aleveille/telegraf/plugins/inputs/mongodb"
	_ "github.com/aleveille/telegraf/plugins/inputs/monit"
	_ "github.com/aleveille/telegraf/plugins/inputs/mqtt_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/multifile"
	_ "github.com/aleveille/telegraf/plugins/inputs/mysql"
	_ "github.com/aleveille/telegraf/plugins/inputs/nats"
	_ "github.com/aleveille/telegraf/plugins/inputs/nats_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/neptune_apex"
	_ "github.com/aleveille/telegraf/plugins/inputs/net"
	_ "github.com/aleveille/telegraf/plugins/inputs/net_response"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx_plus"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx_plus_api"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx_sts"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx_upstream_check"
	_ "github.com/aleveille/telegraf/plugins/inputs/nginx_vts"
	_ "github.com/aleveille/telegraf/plugins/inputs/nsq"
	_ "github.com/aleveille/telegraf/plugins/inputs/nsq_consumer"
	_ "github.com/aleveille/telegraf/plugins/inputs/nstat"
	_ "github.com/aleveille/telegraf/plugins/inputs/ntpq"
	_ "github.com/aleveille/telegraf/plugins/inputs/nvidia_smi"
	_ "github.com/aleveille/telegraf/plugins/inputs/openldap"
	_ "github.com/aleveille/telegraf/plugins/inputs/openntpd"
	_ "github.com/aleveille/telegraf/plugins/inputs/opensmtpd"
	_ "github.com/aleveille/telegraf/plugins/inputs/openweathermap"
	_ "github.com/aleveille/telegraf/plugins/inputs/passenger"
	_ "github.com/aleveille/telegraf/plugins/inputs/pf"
	_ "github.com/aleveille/telegraf/plugins/inputs/pgbouncer"
	_ "github.com/aleveille/telegraf/plugins/inputs/phpfpm"
	_ "github.com/aleveille/telegraf/plugins/inputs/ping"
	_ "github.com/aleveille/telegraf/plugins/inputs/postfix"
	_ "github.com/aleveille/telegraf/plugins/inputs/postgresql"
	_ "github.com/aleveille/telegraf/plugins/inputs/postgresql_extensible"
	_ "github.com/aleveille/telegraf/plugins/inputs/powerdns"
	_ "github.com/aleveille/telegraf/plugins/inputs/powerdns_recursor"
	_ "github.com/aleveille/telegraf/plugins/inputs/processes"
	_ "github.com/aleveille/telegraf/plugins/inputs/procstat"
	_ "github.com/aleveille/telegraf/plugins/inputs/prometheus"
	_ "github.com/aleveille/telegraf/plugins/inputs/proxmox"
	_ "github.com/aleveille/telegraf/plugins/inputs/puppetagent"
	_ "github.com/aleveille/telegraf/plugins/inputs/rabbitmq"
	_ "github.com/aleveille/telegraf/plugins/inputs/raindrops"
	_ "github.com/aleveille/telegraf/plugins/inputs/redfish"
	_ "github.com/aleveille/telegraf/plugins/inputs/redis"
	_ "github.com/aleveille/telegraf/plugins/inputs/rethinkdb"
	_ "github.com/aleveille/telegraf/plugins/inputs/riak"
	_ "github.com/aleveille/telegraf/plugins/inputs/salesforce"
	_ "github.com/aleveille/telegraf/plugins/inputs/sensors"
	_ "github.com/aleveille/telegraf/plugins/inputs/sflow"
	_ "github.com/aleveille/telegraf/plugins/inputs/smart"
	_ "github.com/aleveille/telegraf/plugins/inputs/snmp"
	_ "github.com/aleveille/telegraf/plugins/inputs/snmp_legacy"
	_ "github.com/aleveille/telegraf/plugins/inputs/snmp_trap"
	_ "github.com/aleveille/telegraf/plugins/inputs/socket_listener"
	_ "github.com/aleveille/telegraf/plugins/inputs/solr"
	_ "github.com/aleveille/telegraf/plugins/inputs/sqlserver"
	_ "github.com/aleveille/telegraf/plugins/inputs/stackdriver"
	_ "github.com/aleveille/telegraf/plugins/inputs/statsd"
	_ "github.com/aleveille/telegraf/plugins/inputs/suricata"
	_ "github.com/aleveille/telegraf/plugins/inputs/swap"
	_ "github.com/aleveille/telegraf/plugins/inputs/synproxy"
	_ "github.com/aleveille/telegraf/plugins/inputs/syslog"
	_ "github.com/aleveille/telegraf/plugins/inputs/sysstat"
	_ "github.com/aleveille/telegraf/plugins/inputs/system"
	_ "github.com/aleveille/telegraf/plugins/inputs/systemd_units"
	_ "github.com/aleveille/telegraf/plugins/inputs/tail"
	_ "github.com/aleveille/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/aleveille/telegraf/plugins/inputs/teamspeak"
	_ "github.com/aleveille/telegraf/plugins/inputs/temp"
	_ "github.com/aleveille/telegraf/plugins/inputs/tengine"
	_ "github.com/aleveille/telegraf/plugins/inputs/tomcat"
	_ "github.com/aleveille/telegraf/plugins/inputs/trig"
	_ "github.com/aleveille/telegraf/plugins/inputs/twemproxy"
	_ "github.com/aleveille/telegraf/plugins/inputs/udp_listener"
	_ "github.com/aleveille/telegraf/plugins/inputs/unbound"
	_ "github.com/aleveille/telegraf/plugins/inputs/uwsgi"
	_ "github.com/aleveille/telegraf/plugins/inputs/varnish"
	_ "github.com/aleveille/telegraf/plugins/inputs/vsphere"
	_ "github.com/aleveille/telegraf/plugins/inputs/webhooks"
	_ "github.com/aleveille/telegraf/plugins/inputs/win_perf_counters"
	_ "github.com/aleveille/telegraf/plugins/inputs/win_services"
	_ "github.com/aleveille/telegraf/plugins/inputs/wireguard"
	_ "github.com/aleveille/telegraf/plugins/inputs/wireless"
	_ "github.com/aleveille/telegraf/plugins/inputs/x509_cert"
	_ "github.com/aleveille/telegraf/plugins/inputs/zfs"
	_ "github.com/aleveille/telegraf/plugins/inputs/zipkin"
	_ "github.com/aleveille/telegraf/plugins/inputs/zookeeper"
)
