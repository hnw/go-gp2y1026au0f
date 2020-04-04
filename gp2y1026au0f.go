// https://github.com/sharpsensoruser/sharp-sensor-demos/blob/master/sharp_gp2y1026au0f_demo/sharp_gp2y1026au0f_demo.ino
package main

import (
	"bufio"
	"flag"
	"fmt"

	"github.com/tarm/serial"
)

var (
	serialLine  = flag.String("l", "/dev/serial0", "line")
	serialSpeed = flag.Int("s", 2400, "speed")
	debug       = flag.Bool("debug", false, "debug")
)

func main() {
	flag.Parse()

	c := &serial.Config{
		Name:     *serialLine,
		Baud:     *serialSpeed,
		Size:     8,
		StopBits: 1,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		panic(fmt.Sprintf("serial.OpenPort(): %v", err))
	}

	reader := bufio.NewReader(s)
	for {
		b, err := reader.ReadBytes(0xff)
		if err != nil {
			break
		}
		if len(b) != 7 {
			continue
		}
		if b[0] != 0xaa || b[6] != 0xff {
			continue
		}
		VoutH := b[1]
		VoutL := b[2]
		VrefH := b[3]
		VrefL := b[4]
		testSum := VoutH + VoutL + VrefH + VrefL
		if b[5] != testSum {
			continue
		}

		Vout := (float64(VoutH)*256 + float64(VoutL)) / 1024.0 * 5.0
		dustDensity := 100.0 / 0.35 * Vout

		fmt.Printf("VoutH=%v, VoutL=%v, VrefH=%v, VrefL=%v, Vout=%vmV, dustDensity=%vug/m3\n",
			VoutH, VoutL, VrefH, VrefL, Vout*1000.0, dustDensity)
	}
	s.Close()
}
