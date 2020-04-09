package main

import "chick/pkg/log"

func main() {
	//log.SetPrefix("User")
	testLogger()
	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
	log.Debug("Debug")
	log.Notice("Notice")
	log.Critical("Critical")
	log.Fatal("Fatal")
	log.Panic("Panic")
}

func testLogger() {
	log.Debug("Debug")
}
