package configuration

import (
	"github.com/jinzhu/copier"
)

//
type UltipaConfig struct {
	Hosts []string
	Username string
	Password string
	Crt []byte
	MaxRecvSize int
	Consistency bool
	CurrentGraph string
	Timeout uint32
}

func (config *UltipaConfig) FillDefault() {
	if config.MaxRecvSize == 0 {
		config.MaxRecvSize = 1024 * 1024 * 10 // 10MB
	}

	if config.CurrentGraph == "" {
		config.CurrentGraph = "default"
	}

	if config.Timeout == 0 {
		config.Timeout = 10
	}
}

func (config *UltipaConfig) MergeRequestConfig(rConfig *RequestConfig) *UltipaConfig{
	newConfig := UltipaConfig{}
	err := copier.Copy(newConfig, *config)
	if err != nil {
		panic(err)
	}

	newConfig.Timeout = rConfig.Timeout
	newConfig.CurrentGraph = rConfig.GraphName
	return &newConfig
}


