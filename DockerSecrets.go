package main

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

func ReadAliKey(f string) string {
	r, err := ioutil.ReadFile(f)
	if err != nil {
		log.WithFields(log.Fields{
			"Method": "Open AliKey Docker Secret File",
			"Error":  err,
		}).Fatal("AliKeyReader")
	}

	return string(r)
}

func ReadAliSecret(f string) string {
	r, err := ioutil.ReadFile(f)
	if err != nil {
		log.WithFields(log.Fields{
			"Method": "Open AliSecret Docker Secret File",
			"Error":  err,
		}).Fatal("AliKeyReader")
	}

	return string(r)
}