package configuration

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

//
type UltipaConfig struct {
	Hosts            []string
	Username         string
	Password         string
	DefaultGraph     string `yaml:"default_graph"`
	Crt              []byte
	MaxRecvSize      int `yaml:"max_recv_size"`
	Consistency      bool
	CurrentGraph     string `yaml:"current_graph"`
	CurrentClusterId string `yaml:"current_cluster_id"`
	Timeout          uint32
	Debug            bool
	HeartBeat        int `yaml:"heart_beat"` // frequency:second,  if 0 means no heart beat
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

func LoadConfigFromYAML(file string) (*UltipaConfig, error) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}
	config := &UltipaConfig{}

	err = yaml.Unmarshal(content, config)

	if err != nil {
		return nil, err
	}

	config = NewUltipaConfig(config)

	return config, nil
}
