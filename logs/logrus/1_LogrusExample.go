package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND| os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"genre": "metal",
		"name": "Rammstein",
	}).Info("Немецкая метал-группа, образованная в январе 1994 года в Берлине.")

	log.WithFields(log.Fields{
        "omg": true,
        "name": "Garbage",
    }).Warn("В 2021 году вышел новый альбом No Gods No Masters.")

	log.WithFields(log.Fields{
        "omg": true,
        "name": "Linkin Park",
    }).Fatal("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
}