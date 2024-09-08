package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)

func main() {
	server := gin.Default()

	// Use CORS middleware to allow all origins
	server.Use(cors.Default())

	server.GET("/metrics", getMetrics)

	server.Run(":3043")
}

func getMetrics(context *gin.Context) {
	metrics := `
	# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="get", handler="/api/v1/users", status="200"} 13452
http_requests_total{method="post", handler="/api/v1/users", status="201"} 1034
http_requests_total{method="get", handler="/api/v1/orders", status="200"} 8921
http_requests_total{method="get", handler="/api/v1/users", status="404"} 232
http_requests_total{method="delete", handler="/api/v1/users", status="200"} 43
http_requests_total{method="post", handler="/api/v1/orders", status="201"} 421
http_requests_total{method="get", handler="/api/v1/orders", status="500"} 12
http_requests_total{method="put", handler="/api/v1/products", status="200"} 345
http_requests_total{method="get", handler="/api/v1/products", status="200"} 6789
http_requests_total{method="post", handler="/api/v1/login", status="200"} 9201
http_requests_total{method="post", handler="/api/v1/login", status="401"} 34

# HELP memory_usage_bytes Memory usage in bytes.
# TYPE memory_usage_bytes gauge
memory_usage_bytes{instance="server1"} 345234234
memory_usage_bytes{instance="server2"} 423456789
memory_usage_bytes{instance="server3"} 298765432
memory_usage_bytes{instance="server4"} 567890123
memory_usage_bytes{instance="server5"} 789012345

# HELP cpu_usage_seconds_total Total CPU time consumed in seconds.
# TYPE cpu_usage_seconds_total counter
cpu_usage_seconds_total{instance="server1", mode="user"} 45234.67
cpu_usage_seconds_total{instance="server1", mode="system"} 1234.56
cpu_usage_seconds_total{instance="server2", mode="user"} 37890.89
cpu_usage_seconds_total{instance="server2", mode="system"} 987.45
cpu_usage_seconds_total{instance="server3", mode="user"} 12045.32
cpu_usage_seconds_total{instance="server3", mode="system"} 234.23
cpu_usage_seconds_total{instance="server4", mode="user"} 58034.67
cpu_usage_seconds_total{instance="server4", mode="system"} 2004.56
cpu_usage_seconds_total{instance="server5", mode="user"} 47890.11
cpu_usage_seconds_total{instance="server5", mode="system"} 1345.89

# HELP disk_io_operations_total Total disk IO operations.
# TYPE disk_io_operations_total counter
disk_io_operations_total{instance="server1", device="sda"} 10567890
disk_io_operations_total{instance="server1", device="sdb"} 5674321
disk_io_operations_total{instance="server2", device="sda"} 12903456
disk_io_operations_total{instance="server3", device="sda"} 34001234
disk_io_operations_total{instance="server3", device="sdb"} 23456000
disk_io_operations_total{instance="server4", device="sda"} 41234567
disk_io_operations_total{instance="server5", device="sda"} 13450000

# HELP network_receive_bytes_total Total network bytes received.
# TYPE network_receive_bytes_total counter
network_receive_bytes_total{instance="server1", interface="eth0"} 1234567890
network_receive_bytes_total{instance="server2", interface="eth0"} 987654321
network_receive_bytes_total{instance="server3", interface="eth0"} 567890123
network_receive_bytes_total{instance="server4", interface="eth0"} 234567890
network_receive_bytes_total{instance="server5", interface="eth0"} 789012345

# HELP network_transmit_bytes_total Total network bytes transmitted.
# TYPE network_transmit_bytes_total counter
network_transmit_bytes_total{instance="server1", interface="eth0"} 345678901
network_transmit_bytes_total{instance="server2", interface="eth0"} 234567890
network_transmit_bytes_total{instance="server3", interface="eth0"} 987654321
network_transmit_bytes_total{instance="server4", interface="eth0"} 678901234
network_transmit_bytes_total{instance="server5", interface="eth0"} 543210987

# HELP app_errors_total The total number of application errors.
# TYPE app_errors_total counter
app_errors_total{type="database", severity="critical"} 12
app_errors_total{type="api", severity="critical"} 8
app_errors_total{type="authentication", severity="warning"} 34
app_errors_total{type="database", severity="warning"} 15
app_errors_total{type="api", severity="warning"} 21

# HELP active_sessions_total The total number of active user sessions.
# TYPE active_sessions_total gauge
active_sessions_total{region="us-east"} 132
active_sessions_total{region="us-west"} 98
active_sessions_total{region="eu-central"} 65
active_sessions_total{region="ap-south"} 110
active_sessions_total{region="ap-northeast"} 43

# HELP service_up Status of the service (1 = up, 0 = down).
# TYPE service_up gauge
service_up{service="db", instance="db-server1"} 1
service_up{service="api", instance="api-server1"} 1
service_up{service="auth", instance="auth-server1"} 1
service_up{service="cache", instance="cache-server1"} 1
service_up{service="db", instance="db-server2"} 0
service_up{service="api", instance="api-server2"} 1
service_up{service="auth", instance="auth-server2"} 1

# HELP queue_length_total The total number of items in processing queues.
# TYPE queue_length_total gauge
queue_length_total{queue="email"} 23
queue_length_total{queue="notifications"} 45
queue_length_total{queue="payments"} 12
queue_length_total{queue="jobs"} 67

# HELP http_response_time_seconds Histogram of response time for HTTP requests.
# TYPE http_response_time_seconds histogram
http_response_time_seconds_bucket{le="0.05", method="get", handler="/api/v1/users"} 240
http_response_time_seconds_bucket{le="0.1", method="get", handler="/api/v1/users"} 532
http_response_time_seconds_bucket{le="0.2", method="get", handler="/api/v1/users"} 894
http_response_time_seconds_bucket{le="0.5", method="get", handler="/api/v1/users"} 1334
http_response_time_seconds_bucket{le="1", method="get", handler="/api/v1/users"} 1521
http_response_time_seconds_bucket{le="2.5", method="get", handler="/api/v1/users"} 1823
http_response_time_seconds_bucket{le="5", method="get", handler="/api/v1/users"} 2024
http_response_time_seconds_bucket{le="+Inf", method="get", handler="/api/v1/users"} 2500
http_response_time_seconds_sum{method="get", handler="/api/v1/users"} 9230.4
http_response_time_seconds_count{method="get", handler="/api/v1/users"} 2500

# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="get", handler="/api/v1/users", status="200"} 13452
http_requests_total{method="post", handler="/api/v1/users", status="201"} 1034
http_requests_total{method="get", handler="/api/v1/orders", status="200"} 8921
http_requests_total{method="get", handler="/api/v1/users", status="404"} 232
http_requests_total{method="delete", handler="/api/v1/users", status="200"} 43
http_requests_total{method="post", handler="/api/v1/orders", status="201"} 421
http_requests_total{method="get", handler="/api/v1/orders", status="500"} 12
http_requests_total{method="put", handler="/api/v1/products", status="200"} 345
http_requests_total{method="get", handler="/api/v1/products", status="200"} 6789
http_requests_total{method="post", handler="/api/v1/login", status="200"} 9201
http_requests_total{method="post", handler="/api/v1/login", status="401"} 34

# HELP memory_usage_bytes Memory usage in bytes.
# TYPE memory_usage_bytes gauge
memory_usage_bytes{instance="server1"} 345234234
memory_usage_bytes{instance="server2"} 423456789
memory_usage_bytes{instance="server3"} 298765432
memory_usage_bytes{instance="server4"} 567890123
memory_usage_bytes{instance="server5"} 789012345

# HELP cpu_usage_seconds_total Total CPU time consumed in seconds.
# TYPE cpu_usage_seconds_total counter
cpu_usage_seconds_total{instance="server1", mode="user"} 45234.67
cpu_usage_seconds_total{instance="server1", mode="system"} 1234.56
cpu_usage_seconds_total{instance="server2", mode="user"} 37890.89
cpu_usage_seconds_total{instance="server2", mode="system"} 987.45
cpu_usage_seconds_total{instance="server3", mode="user"} 12045.32
cpu_usage_seconds_total{instance="server3", mode="system"} 234.23
cpu_usage_seconds_total{instance="server4", mode="user"} 58034.67
cpu_usage_seconds_total{instance="server4", mode="system"} 2004.56
cpu_usage_seconds_total{instance="server5", mode="user"} 47890.11
cpu_usage_seconds_total{instance="server5", mode="system"} 1345.89

# HELP disk_io_operations_total Total disk IO operations.
# TYPE disk_io_operations_total counter
disk_io_operations_total{instance="server1", device="sda"} 10567890
disk_io_operations_total{instance="server1", device="sdb"} 5674321
disk_io_operations_total{instance="server2", device="sda"} 12903456
disk_io_operations_total{instance="server3", device="sda"} 34001234
disk_io_operations_total{instance="server3", device="sdb"} 23456000
disk_io_operations_total{instance="server4", device="sda"} 41234567
disk_io_operations_total{instance="server5", device="sda"} 13450000

# HELP network_receive_bytes_total Total network bytes received.
# TYPE network_receive_bytes_total counter
network_receive_bytes_total{instance="server1", interface="eth0"} 1234567890
network_receive_bytes_total{instance="server2", interface="eth0"} 987654321
network_receive_bytes_total{instance="server3", interface="eth0"} 567890123
network_receive_bytes_total{instance="server4", interface="eth0"} 234567890
network_receive_bytes_total{instance="server5", interface="eth0"} 789012345

# HELP network_transmit_bytes_total Total network bytes transmitted.
# TYPE network_transmit_bytes_total counter
network_transmit_bytes_total{instance="server1", interface="eth0"} 345678901
network_transmit_bytes_total{instance="server2", interface="eth0"} 234567890
network_transmit_bytes_total{instance="server3", interface="eth0"} 987654321
network_transmit_bytes_total{instance="server4", interface="eth0"} 678901234
network_transmit_bytes_total{instance="server5", interface="eth0"} 543210987

# HELP app_errors_total The total number of application errors.
# TYPE app_errors_total counter
app_errors_total{type="database", severity="critical"} 12
app_errors_total{type="api", severity="critical"} 8
app_errors_total{type="authentication", severity="warning"} 34
app_errors_total{type="database", severity="warning"} 15
app_errors_total{type="api", severity="warning"} 21

# HELP active_sessions_total The total number of active user sessions.
# TYPE active_sessions_total gauge
active_sessions_total{region="us-east"} 132
active_sessions_total{region="us-west"} 98
active_sessions_total{region="eu-central"} 65
active_sessions_total{region="ap-south"} 110
active_sessions_total{region="ap-northeast"} 43

# HELP service_up Status of the service (1 = up, 0 = down).
# TYPE service_up gauge
service_up{service="db", instance="db-server1"} 1
service_up{service="api", instance="api-server1"} 1
service_up{service="auth", instance="auth-server1"} 1
service_up{service="cache", instance="cache-server1"} 1
service_up{service="db", instance="db-server2"} 0
service_up{service="api", instance="api-server2"} 1
service_up{service="auth", instance="auth-server2"} 1

# HELP queue_length_total The total number of items in processing queues.
# TYPE queue_length_total gauge
queue_length_total{queue="email"} 23
queue_length_total{queue="notifications"} 45
queue_length_total{queue="payments"} 12
queue_length_total{queue="jobs"} 67

# HELP http_response_time_seconds Histogram of response time for HTTP requests.
# TYPE http_response_time_seconds histogram
http_response_time_seconds_bucket{le="0.05", method="get", handler="/api/v1/users"} 240
http_response_time_seconds_bucket{le="0.1", method="get", handler="/api/v1/users"} 532
http_response_time_seconds_bucket{le="0.2", method="get", handler="/api/v1/users"} 894
http_response_time_seconds_bucket{le="0.5", method="get", handler="/api/v1/users"} 1334
http_response_time_seconds_bucket{le="1", method="get", handler="/api/v1/users"} 1521
http_response_time_seconds_bucket{le="2.5", method="get", handler="/api/v1/users"} 1823
http_response_time_seconds_bucket{le="5", method="get", handler="/api/v1/users"} 2024
http_response_time_seconds_bucket{le="+Inf", method="get", handler="/api/v1/users"} 2500
http_response_time_seconds_sum{method="get", handler="/api/v1/users"} 9230.4
http_response_time_seconds_count{method="get", handler="/api/v1/users"} 2500
# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="get", handler="/api/v1/users", status="200"} 13452
http_requests_total{method="post", handler="/api/v1/users", status="201"} 1034
http_requests_total{method="get", handler="/api/v1/orders", status="200"} 8921
http_requests_total{method="get", handler="/api/v1/users", status="404"} 232
http_requests_total{method="delete", handler="/api/v1/users", status="200"} 43
http_requests_total{method="post", handler="/api/v1/orders", status="201"} 421
http_requests_total{method="get", handler="/api/v1/orders", status="500"} 12
http_requests_total{method="put", handler="/api/v1/products", status="200"} 345
http_requests_total{method="get", handler="/api/v1/products", status="200"} 6789
http_requests_total{method="post", handler="/api/v1/login", status="200"} 9201
http_requests_total{method="post", handler="/api/v1/login", status="401"} 34

# HELP memory_usage_bytes Memory usage in bytes.
# TYPE memory_usage_bytes gauge
memory_usage_bytes{instance="server1"} 345234234
memory_usage_bytes{instance="server2"} 423456789
memory_usage_bytes{instance="server3"} 298765432
memory_usage_bytes{instance="server4"} 567890123
memory_usage_bytes{instance="server5"} 789012345

# HELP cpu_usage_seconds_total Total CPU time consumed in seconds.
# TYPE cpu_usage_seconds_total counter
cpu_usage_seconds_total{instance="server1", mode="user"} 45234.67
cpu_usage_seconds_total{instance="server1", mode="system"} 1234.56
cpu_usage_seconds_total{instance="server2", mode="user"} 37890.89
cpu_usage_seconds_total{instance="server2", mode="system"} 987.45
cpu_usage_seconds_total{instance="server3", mode="user"} 12045.32
cpu_usage_seconds_total{instance="server3", mode="system"} 234.23
cpu_usage_seconds_total{instance="server4", mode="user"} 58034.67
cpu_usage_seconds_total{instance="server4", mode="system"} 2004.56
cpu_usage_seconds_total{instance="server5", mode="user"} 47890.11
cpu_usage_seconds_total{instance="server5", mode="system"} 1345.89

# HELP disk_io_operations_total Total disk IO operations.
# TYPE disk_io_operations_total counter
disk_io_operations_total{instance="server1", device="sda"} 10567890
disk_io_operations_total{instance="server1", device="sdb"} 5674321
disk_io_operations_total{instance="server2", device="sda"} 12903456
disk_io_operations_total{instance="server3", device="sda"} 34001234
disk_io_operations_total{instance="server3", device="sdb"} 23456000
disk_io_operations_total{instance="server4", device="sda"} 41234567
disk_io_operations_total{instance="server5", device="sda"} 13450000

# HELP network_receive_bytes_total Total network bytes received.
# TYPE network_receive_bytes_total counter
network_receive_bytes_total{instance="server1", interface="eth0"} 1234567890
network_receive_bytes_total{instance="server2", interface="eth0"} 987654321
network_receive_bytes_total{instance="server3", interface="eth0"} 567890123
network_receive_bytes_total{instance="server4", interface="eth0"} 234567890
network_receive_bytes_total{instance="server5", interface="eth0"} 789012345

# HELP network_transmit_bytes_total Total network bytes transmitted.
# TYPE network_transmit_bytes_total counter
network_transmit_bytes_total{instance="server1", interface="eth0"} 345678901
network_transmit_bytes_total{instance="server2", interface="eth0"} 234567890
network_transmit_bytes_total{instance="server3", interface="eth0"} 987654321
network_transmit_bytes_total{instance="server4", interface="eth0"} 678901234
network_transmit_bytes_total{instance="server5", interface="eth0"} 543210987

# HELP app_errors_total The total number of application errors.
# TYPE app_errors_total counter
app_errors_total{type="database", severity="critical"} 12
app_errors_total{type="api", severity="critical"} 8
app_errors_total{type="authentication", severity="warning"} 34
app_errors_total{type="database", severity="warning"} 15
app_errors_total{type="api", severity="warning"} 21

# HELP active_sessions_total The total number of active user sessions.
# TYPE active_sessions_total gauge
active_sessions_total{region="us-east"} 132
active_sessions_total{region="us-west"} 98
active_sessions_total{region="eu-central"} 65
active_sessions_total{region="ap-south"} 110
active_sessions_total{region="ap-northeast"} 43

# HELP service_up Status of the service (1 = up, 0 = down).
# TYPE service_up gauge
service_up{service="db", instance="db-server1"} 1
service_up{service="api", instance="api-server1"} 1
service_up{service="auth", instance="auth-server1"} 1
service_up{service="cache", instance="cache-server1"} 1
service_up{service="db", instance="db-server2"} 0
service_up{service="api", instance="api-server2"} 1
service_up{service="auth", instance="auth-server2"} 1

# HELP queue_length_total The total number of items in processing queues.
# TYPE queue_length_total gauge
queue_length_total{queue="email"} 23
queue_length_total{queue="notifications"} 45
queue_length_total{queue="payments"} 12
queue_length_total{queue="jobs"} 67

# HELP http_response_time_seconds Histogram of response time for HTTP requests.
# TYPE http_response_time_seconds histogram
http_response_time_seconds_bucket{le="0.05", method="get", handler="/api/v1/users"} 240
http_response_time_seconds_bucket{le="0.1", method="get", handler="/api/v1/users"} 532
http_response_time_seconds_bucket{le="0.2", method="get", handler="/api/v1/users"} 894
http_response_time_seconds_bucket{le="0.5", method="get", handler="/api/v1/users"} 1334
http_response_time_seconds_bucket{le="1", method="get", handler="/api/v1/users"} 1521
http_response_time_seconds_bucket{le="2.5", method="get", handler="/api/v1/users"} 1823
http_response_time_seconds_bucket{le="5", method="get", handler="/api/v1/users"} 2024
http_response_time_seconds_bucket{le="+Inf", method="get", handler="/api/v1/users"} 2500
http_response_time_seconds_sum{method="get", handler="/api/v1/users"} 9230.4
http_response_time_seconds_count{method="get", handler="/api/v1/users"} 2500 
# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="get", handler="/api/v1/users", status="200"} 13452
http_requests_total{method="post", handler="/api/v1/users", status="201"} 1034
http_requests_total{method="get", handler="/api/v1/orders", status="200"} 8921
http_requests_total{method="get", handler="/api/v1/users", status="404"} 232
http_requests_total{method="delete", handler="/api/v1/users", status="200"} 43
http_requests_total{method="post", handler="/api/v1/orders", status="201"} 421
http_requests_total{method="get", handler="/api/v1/orders", status="500"} 12
http_requests_total{method="put", handler="/api/v1/products", status="200"} 345
http_requests_total{method="get", handler="/api/v1/products", status="200"} 6789
http_requests_total{method="post", handler="/api/v1/login", status="200"} 9201
http_requests_total{method="post", handler="/api/v1/login", status="401"} 34

# HELP memory_usage_bytes Memory usage in bytes.
# TYPE memory_usage_bytes gauge
memory_usage_bytes{instance="server1"} 345234234
memory_usage_bytes{instance="server2"} 423456789
memory_usage_bytes{instance="server3"} 298765432
memory_usage_bytes{instance="server4"} 567890123
memory_usage_bytes{instance="server5"} 789012345

# HELP cpu_usage_seconds_total Total CPU time consumed in seconds.
# TYPE cpu_usage_seconds_total counter
cpu_usage_seconds_total{instance="server1", mode="user"} 45234.67
cpu_usage_seconds_total{instance="server1", mode="system"} 1234.56
cpu_usage_seconds_total{instance="server2", mode="user"} 37890.89
cpu_usage_seconds_total{instance="server2", mode="system"} 987.45
cpu_usage_seconds_total{instance="server3", mode="user"} 12045.32
cpu_usage_seconds_total{instance="server3", mode="system"} 234.23
cpu_usage_seconds_total{instance="server4", mode="user"} 58034.67
cpu_usage_seconds_total{instance="server4", mode="system"} 2004.56
cpu_usage_seconds_total{instance="server5", mode="user"} 47890.11
cpu_usage_seconds_total{instance="server5", mode="system"} 1345.89

# HELP disk_io_operations_total Total disk IO operations.
# TYPE disk_io_operations_total counter
disk_io_operations_total{instance="server1", device="sda"} 10567890
disk_io_operations_total{instance="server1", device="sdb"} 5674321
disk_io_operations_total{instance="server2", device="sda"} 12903456
disk_io_operations_total{instance="server3", device="sda"} 34001234
disk_io_operations_total{instance="server3", device="sdb"} 23456000
disk_io_operations_total{instance="server4", device="sda"} 41234567
disk_io_operations_total{instance="server5", device="sda"} 13450000

# HELP network_receive_bytes_total Total network bytes received.
# TYPE network_receive_bytes_total counter
network_receive_bytes_total{instance="server1", interface="eth0"} 1234567890
network_receive_bytes_total{instance="server2", interface="eth0"} 987654321
network_receive_bytes_total{instance="server3", interface="eth0"} 567890123
network_receive_bytes_total{instance="server4", interface="eth0"} 234567890
network_receive_bytes_total{instance="server5", interface="eth0"} 789012345

# HELP network_transmit_bytes_total Total network bytes transmitted.
# TYPE network_transmit_bytes_total counter
network_transmit_bytes_total{instance="server1", interface="eth0"} 345678901
network_transmit_bytes_total{instance="server2", interface="eth0"} 234567890
network_transmit_bytes_total{instance="server3", interface="eth0"} 987654321
network_transmit_bytes_total{instance="server4", interface="eth0"} 678901234
network_transmit_bytes_total{instance="server5", interface="eth0"} 543210987

# HELP app_errors_total The total number of application errors.
# TYPE app_errors_total counter
app_errors_total{type="database", severity="critical"} 12
app_errors_total{type="api", severity="critical"} 8
app_errors_total{type="authentication", severity="warning"} 34
app_errors_total{type="database", severity="warning"} 15
app_errors_total{type="api", severity="warning"} 21

# HELP active_sessions_total The total number of active user sessions.
# TYPE active_sessions_total gauge
active_sessions_total{region="us-east"} 132
active_sessions_total{region="us-west"} 98
active_sessions_total{region="eu-central"} 65
active_sessions_total{region="ap-south"} 110
active_sessions_total{region="ap-northeast"} 43

# HELP service_up Status of the service (1 = up, 0 = down).
# TYPE service_up gauge
service_up{service="db", instance="db-server1"} 1
service_up{service="api", instance="api-server1"} 1
service_up{service="auth", instance="auth-server1"} 1
service_up{service="cache", instance="cache-server1"} 1
service_up{service="db", instance="db-server2"} 0
service_up{service="api", instance="api-server2"} 1
service_up{service="auth", instance="auth-server2"} 1

# HELP queue_length_total The total number of items in processing queues.
# TYPE queue_length_total gauge
queue_length_total{queue="email"} 23
queue_length_total{queue="notifications"} 45
queue_length_total{queue="payments"} 12
queue_length_total{queue="jobs"} 67

# HELP http_response_time_seconds Histogram of response time for HTTP requests.
# TYPE http_response_time_seconds histogram
http_response_time_seconds_bucket{le="0.05", method="get", handler="/api/v1/users"} 240
http_response_time_seconds_bucket{le="0.1", method="get", handler="/api/v1/users"} 532
http_response_time_seconds_bucket{le="0.2", method="get", handler="/api/v1/users"} 894
http_response_time_seconds_bucket{le="0.5", method="get", handler="/api/v1/users"} 1334
http_response_time_seconds_bucket{le="1", method="get", handler="/api/v1/users"} 1521
http_response_time_seconds_bucket{le="2.5", method="get", handler="/api/v1/users"} 1823
http_response_time_seconds_bucket{le="5", method="get", handler="/api/v1/users"} 2024
http_response_time_seconds_bucket{le="+Inf", method="get", handler="/api/v1/users"} 2500
http_response_time_seconds_sum{method="get", handler="/api/v1/users"} 9230.4
http_response_time_seconds_count{method="get", handler="/api/v1/users"} 2500`

	context.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(metrics))
}
