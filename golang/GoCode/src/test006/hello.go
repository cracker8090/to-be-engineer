package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	var str string = "abc"
	sli := make([]string, 10)
	sli = append(sli, str)
	log.Info(str)
	log.Info(sli)
	event := "event"
	topic := "topic"
	key := 100
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	log.WithFields(log.Fields{
		"event": event,
		"topic": topic,
		"key":   key,
	}).Fatal("Failed to send event")

}

func init() {
	// default ASCII text
	//log.SetFormatter(&log.JSONFormatter{})
	logFile := "log.txt"
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	//log.SetOutput(os.Stdout)
	log.SetOutput(file)
	//log.SetLevel(log.WarnLevel)

}
