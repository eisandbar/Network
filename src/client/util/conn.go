package conn

const (
	CONN_HOST = "network_haproxy_1"
	CONN_PORT = "5100"
)

var ConnStr = "http://" + CONN_HOST + ":" + CONN_PORT