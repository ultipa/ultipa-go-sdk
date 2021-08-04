package configuration

type RequestConfig struct {
	GraphName string
	Timeout uint32
	ClusterId string
	Host string // set for force host test
}
