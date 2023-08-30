package configuration

import (
	"crypto/md5"
	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
	"time"
)

//
type UltipaConfig struct {
	Hosts            []string // hosts with ports
	Username         string   // ultipa graph username
	Password         string   // ultipa graph password
	PasswordEncrypt  string   // method of encrypt password, MD5, LDAP, NOTHING
	DefaultGraph     string   `yaml:"default_graph"` // default graph when connection established
	Crt              []byte   // certification file for encrypt messages
	MaxRecvSize      int      `yaml:"max_recv_size"` // grpc max receive size
	Consistency      bool     // if consistency, reading query will send to master
	CurrentGraph     string   `yaml:"current_graph"`      // the current graph, used when user what get the connection's current graph name
	CurrentClusterId string   `yaml:"current_cluster_id"` // used for name server only
	Timeout          int32    // timeout - seconds
	Debug            bool     // debug, print more logs
	HeartBeat        int      `yaml:"heart_beat"` // frequency:second,  if 0 means no heart beat, to make sure the connection is alive
}

var DefaultTimeout int32 = 1000

func NewUltipaConfig(config *UltipaConfig) (*UltipaConfig, error) {
	config.FillDefault()

	h := md5.New()
	h.Write([]byte(config.Password))
	encryptedPwd, err := Encrypt(config.PasswordEncrypt, config.Password)
	if err != nil {
		return config, err
	}
	config.Password = encryptedPwd
	return config, nil
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
		config.Timeout = DefaultTimeout
	}
	if config.PasswordEncrypt == "" {
		config.PasswordEncrypt = "MD5"
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

	headers := []string{
		"user",
		config.Username,
		"password",
		config.Password,
		"graph_name",
		graphName,
		//"cluster_id",
		//config.CurrentClusterId,
	}
	if rConfig == nil || (rConfig.TimezoneOffset == 0 && rConfig.Timezone == "") {
		_, offset := time.Now().Zone()
		headers = append(headers, "tz_offset", strconv.Itoa(offset))
	} else if rConfig.TimezoneOffset != 0 {
		headers = append(headers, "tz_offset", strconv.FormatInt(rConfig.TimezoneOffset, 10))
	} else if rConfig.Timezone != "" {
		headers = append(headers, "tz", rConfig.Timezone)
	}

	return headers
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

	config, err = NewUltipaConfig(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
