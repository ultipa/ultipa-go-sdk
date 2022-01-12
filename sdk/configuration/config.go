package configuration

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/copier"
	"strings"
)

//
type UltipaConfig struct {
	Hosts            []string
	Username         string
	Password         string
	DefaultGraph     string
	Crt              []byte
	MaxRecvSize      int
	Consistency      bool
	CurrentGraph     string
	CurrentClusterId string
	Timeout          uint32
	HeartBeat        int // if 0 means no heart beat
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

	if config.DefaultGraph != "" {
		config.CurrentGraph = config.DefaultGraph
	}

	if config.CurrentGraph == "" {
		config.CurrentGraph = "default"
	}

	if config.Timeout == 0 {
		config.Timeout = 1000
	}
}

func (config *UltipaConfig) MergeRequestConfig(rConfig *RequestConfig) *UltipaConfig {

	newConfig := &UltipaConfig{}

	copier.Copy(newConfig, config)

	if rConfig.Timeout > 0 {
		newConfig.Timeout = rConfig.Timeout
	}

	if rConfig.GraphName != "" {
		newConfig.CurrentGraph = rConfig.GraphName
	}
	if rConfig.ClusterId != "" {
		newConfig.CurrentClusterId = rConfig.ClusterId
	}

	return newConfig
}

func (config *UltipaConfig) ToContextKV(rConfig *RequestConfig) []string {

	graphName := config.CurrentGraph

	if rConfig != nil && rConfig.GraphName != "" {
		graphName = rConfig.GraphName
	}

	return []string{
		"user",
		config.Username,
		"password",
		config.Password,
		"graph_name",
		graphName,
		//"cluster_id",
		//config.CurrentClusterId,
	}
}
