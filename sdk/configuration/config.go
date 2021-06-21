package configuration

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/copier"
	"strings"
)

//
type UltipaConfig struct {
	Hosts []string
	Username string
	Password string
	DefaultGraph string
	Crt []byte
	MaxRecvSize int
	Consistency bool
	CurrentGraph string
	Timeout uint32
}

func NewUltipaConfig(config *UltipaConfig) *UltipaConfig {
	config.FillDefault()

	h := md5.New()
	h.Write([]byte(config.Password))
	config.Password = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	return config
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

func (config *UltipaConfig) ToMetaKV() []string{
	return []string{
		"user",
		config.Username,
		"password",
		config.Password,
		"graph_name",
		config.CurrentGraph,
	}
}


