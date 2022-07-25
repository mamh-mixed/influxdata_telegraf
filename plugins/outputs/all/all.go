package all

import (
	//Blank imports for plugins to register themselves
	_ "github.com/influxdata/telegraf/plugins/outputs/amon"
	_ "github.com/influxdata/telegraf/plugins/outputs/amqp"
	_ "github.com/influxdata/telegraf/plugins/outputs/application_insights"
	_ "github.com/influxdata/telegraf/plugins/outputs/azure_data_explorer"
	_ "github.com/influxdata/telegraf/plugins/outputs/azure_monitor"
	_ "github.com/influxdata/telegraf/plugins/outputs/bigquery"
	_ "github.com/influxdata/telegraf/plugins/outputs/cloud_pubsub"
	_ "github.com/influxdata/telegraf/plugins/outputs/cloudwatch"
	_ "github.com/influxdata/telegraf/plugins/outputs/cloudwatch_logs"
	_ "github.com/influxdata/telegraf/plugins/outputs/cratedb"
	_ "github.com/influxdata/telegraf/plugins/outputs/datadog"
	_ "github.com/influxdata/telegraf/plugins/outputs/discard"
	_ "github.com/influxdata/telegraf/plugins/outputs/dynatrace"
	_ "github.com/influxdata/telegraf/plugins/outputs/elasticsearch"
	_ "github.com/influxdata/telegraf/plugins/outputs/event_hubs"
	_ "github.com/influxdata/telegraf/plugins/outputs/exec"
	_ "github.com/influxdata/telegraf/plugins/outputs/execd"
	_ "github.com/influxdata/telegraf/plugins/outputs/file"
	_ "github.com/influxdata/telegraf/plugins/outputs/graphite"
	_ "github.com/influxdata/telegraf/plugins/outputs/graylog"
	_ "github.com/influxdata/telegraf/plugins/outputs/groundwork"
	_ "github.com/influxdata/telegraf/plugins/outputs/health"
	_ "github.com/influxdata/telegraf/plugins/outputs/http"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb_v2"
	_ "github.com/influxdata/telegraf/plugins/outputs/instrumental"
	_ "github.com/influxdata/telegraf/plugins/outputs/kafka"
	_ "github.com/influxdata/telegraf/plugins/outputs/kinesis"
	_ "github.com/influxdata/telegraf/plugins/outputs/librato"
	_ "github.com/influxdata/telegraf/plugins/outputs/logzio"
	_ "github.com/influxdata/telegraf/plugins/outputs/loki"
	_ "github.com/influxdata/telegraf/plugins/outputs/mongodb"
	_ "github.com/influxdata/telegraf/plugins/outputs/mqtt"
	_ "github.com/influxdata/telegraf/plugins/outputs/nats"
	_ "github.com/influxdata/telegraf/plugins/outputs/newrelic"
	_ "github.com/influxdata/telegraf/plugins/outputs/nsq"
	_ "github.com/influxdata/telegraf/plugins/outputs/opentelemetry"
	_ "github.com/influxdata/telegraf/plugins/outputs/opentsdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/influxdata/telegraf/plugins/outputs/redistimeseries"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann_legacy"
	_ "github.com/influxdata/telegraf/plugins/outputs/sensu"
	_ "github.com/influxdata/telegraf/plugins/outputs/signalfx"
	_ "github.com/influxdata/telegraf/plugins/outputs/socket_writer"
	_ "github.com/influxdata/telegraf/plugins/outputs/sql"
	_ "github.com/influxdata/telegraf/plugins/outputs/stackdriver"
	_ "github.com/influxdata/telegraf/plugins/outputs/stomp"
	_ "github.com/influxdata/telegraf/plugins/outputs/sumologic"
	_ "github.com/influxdata/telegraf/plugins/outputs/syslog"
	_ "github.com/influxdata/telegraf/plugins/outputs/timestream"
	_ "github.com/influxdata/telegraf/plugins/outputs/warp10"
	_ "github.com/influxdata/telegraf/plugins/outputs/wavefront"
	_ "github.com/influxdata/telegraf/plugins/outputs/websocket"
	_ "github.com/influxdata/telegraf/plugins/outputs/yandex_cloud_monitoring"
)
