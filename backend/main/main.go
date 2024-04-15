package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/cassaram/ece1896/backend/core"
	"github.com/cassaram/ece1896/backend/hardware"
)

func main() {
	// Setup main process log
	l := log.New(os.Stdout, "MAIN: ", log.Ldate|log.Ltime)
	// Load environment varaibles
	address := os.Getenv("ADDRESS")
	channels, err := strconv.ParseUint(os.Getenv("CHANNELS"), 10, 64)
	if err != nil {
		l.Fatalln(err)
		return
	}
	busses, err := strconv.ParseUint(os.Getenv("BUSSES"), 10, 64)
	if err != nil {
		l.Fatalln(err)
		return
	}
	cfgPath := os.Getenv("CFG_PATH")

	// Get hardware for I2C
	hardware := hardware.NewHardware(l)

	// Start program
	c := core.NewCore(address, channels, busses, l, cfgPath, hardware)
	go c.Run()

	// Hold until close
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	l.Println("Running until interrupt")
	<-done
}
