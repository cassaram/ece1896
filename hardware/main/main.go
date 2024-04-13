package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cassaram/ece1896/hardware/backend"
	"github.com/cassaram/ece1896/hardware/gpiowrappers"
	"github.com/cassaram/ece1896/hardware/mcp23s17"
	"github.com/cassaram/ece1896/hardware/mcp3008"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func main() {
	// Generate logger
	logger := log.New(os.Stdout, "Ctrl: ", log.Ldate|log.Ltime)

	// Configure backend connection
	backendWs := backend.NewBackendConnection("ws://localhost:8080/api/v1/ws", logger)

	// Initialize SPI busses
	host.Init()

	// Configure GPIO Expander SPI bus
	gpio_spi, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		logger.Fatalf(err.Error())
	}
	gpio_spi_con, err := gpio_spi.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Reset MCP GPIO expander chips
	gpio_rst_port := gpioreg.ByName("GPIO4")
	gpio_rst_port.Out(gpio.High)
	time.Sleep(time.Millisecond)
	gpio_rst_port.Out(gpio.Low)
	time.Sleep(2 * time.Microsecond)
	gpio_rst_port.Out(gpio.High)
	time.Sleep(time.Millisecond)

	txData := []byte{0x40, 0x0A, 0x18}
	rxData := make([]byte, 3)
	gpio_spi_con.Tx(txData, rxData)
	txData = []byte{0x4E, 0x0A, 0x18}
	gpio_spi_con.Tx(txData, rxData)
	txData = []byte{0x41, 0x05, 0x00}
	gpio_spi_con.Tx(txData, rxData)
	fmt.Println("IOCTRL (0x05) DEVICE 0:", rxData)
	txData = []byte{0x43, 0x05, 0x00}
	gpio_spi_con.Tx(txData, rxData)
	fmt.Println("IOCTRL (0x05) DEVICE 1:", rxData)
	txData = []byte{0x41, 0x0A, 0x00}
	gpio_spi_con.Tx(txData, rxData)
	fmt.Println("IOCTRL (0x0A) DEVICE 0:", rxData)
	txData = []byte{0x43, 0x0A, 0x00}
	gpio_spi_con.Tx(txData, rxData)
	fmt.Println("IOCTRL (0x0A) DEVICE 1:", rxData)

	// Configure ADC SPI bus
	adc_spi, err := spireg.Open("/dev/spidev0.0")
	if err != nil {
		logger.Fatalf(err.Error())
	}
	adc_spi_con, err := adc_spi.Connect(2340*physic.KiloHertz, spi.Mode3, 8)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// Initialize driver register
	_, err = driverreg.Init()
	if err != nil {
		logger.Fatal(err)
	}

	// Get MCP23S17 chips
	ledgpio_mcp := mcp23s17.NewMCP23S17(&gpio_spi_con, 0)
	btngpio_mcp := mcp23s17.NewMCP23S17(&gpio_spi_con, 1)

	ledgpio_mcp.Init()
	btngpio_mcp.Init()

	// Configure LED GPIO Expander
	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			ledgpio_mcp.ConfigurePin(uint8(i), uint8(j), true, false, false, false)
			ledgpio_mcp.SetPin(uint8(i), uint8(j), true) // Disable LED
		}
	}

	ledgpio := gpiowrappers.NewLEDWrapper(ledgpio_mcp, backendWs)
	ledgpio.Start()

	// Configure Mute/PFL/AFL GPIO expander
	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			btngpio_mcp.ConfigurePin(uint8(i), uint8(j), false, false, true, true)
		}
	}

	btngpio_intA := gpioreg.ByName("GPIO22")
	err = btngpio_intA.In(gpio.PullUp, gpio.FallingEdge)
	if err != nil {
		logger.Fatal(err)
	}
	btngpio_intB := gpioreg.ByName("GPIO23")
	err = btngpio_intB.In(gpio.PullUp, gpio.FallingEdge)
	if err != nil {
		logger.Fatal(err)
	}

	btngpio := gpiowrappers.NewSwitchesWrapper(btngpio_mcp, backendWs, btngpio_intA, btngpio_intB, logger)
	btngpio.Start()

	// Configure ADC
	faderadc_mcp := mcp3008.NewMCP3008(&adc_spi_con)
	faderadc_mcp.ReadPort(0) // TEMP

	// Connect
	backendWs.Connect()

	testTimer := time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case <-testTimer.C:
			val, err := ledgpio_mcp.ReadPort(0)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("0 Val: %08b\n", val)

			val, err = btngpio_mcp.ReadPort(0)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("1 Val: %08b\n", val)

		}
	}

	// Hold until close
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	logger.Println("Running until interrupt")
	<-done
}
