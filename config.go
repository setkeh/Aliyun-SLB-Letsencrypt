package main

import (
	"encoding/json"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

var (
	config Config
)

//Config Main Configuration Struct
type Config struct {
	//CertName For Aliyun Cert Identification
	CertName string
	//CertDirectory LetsEncrypt Cert Output Directory
	CertDirectory string
	//AliyunKeySecretPath Location of the Docker Secret File for Aliyun Key
	AliyunKeySecretPath string
	//AliyunSecretSecretPath Location of Docker Secret File for Aliyun Secret
	AliyunSecretSecretPath string
}

func DoConfig(c string){
	rawjson, jsoperr := ioutil.ReadFile(c)
	if jsoperr != nil {
		log.WithFields(log.Fields{
			"Method": "Open Json File",
			"Error":  jsoperr,
		}).Fatal("ConfigBuilder")
	}

	UmError := json.Unmarshal(rawjson, &config)
	if UmError != nil {
		log.WithFields(log.Fields{
			"Method": "Open Json File",
			"Error":  UmError,
		}).Fatal("ConfigBuilder")
	}
}