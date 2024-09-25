package main

import (
	"log"
	"os"
)

type Abons struct {
	TaskID   int `json:"task_id"`
	SectorID int `json:"sector_id"`
}

func main() {
	logFile, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Заполнение структуры Abons
	abons := Abons{
		TaskID:   123, // Пример значения для TaskID
		SectorID: 456, // Пример значения для SectorID
	}

	log.SetOutput(logFile)
	log.Printf("Message received: %+v", abons)
}
