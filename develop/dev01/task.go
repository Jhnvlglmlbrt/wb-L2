package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {

	// запрос к серверу NTP для получения точного времени
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Точное время используя NTP: ", ntpTime)
}
