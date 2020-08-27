package types

type Request_Property = struct {
	Dataset DBType;
}

type Request_Common struct {
	GraphSetName string
	TimeoutSeconds uint32
	Retry *Retry
	UseHost string
}