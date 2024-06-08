package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func Time() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Add(response.ClockOffset))
}
