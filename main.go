package main

import (
	"flag"
)

func main() {
	//Config Path
	conf := flag.String("c", "Configuration", "Config File Path")
	confPath := *conf
	DoConfig(confPath)

	AliKey := ReadAliKey(config.AliyunKeySecretPath)
	AliSecret := ReadAliSecret(config.AliyunSecretSecretPath)

	ac := AliClient(AliKey, AliSecret)
}