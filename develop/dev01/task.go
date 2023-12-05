package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения времени с сервера NTP:", err)
		os.Exit(1)
	}

	fmt.Println("Точное время:", ntpTime.Format(time.RFC3339))
}
