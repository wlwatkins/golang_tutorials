package main

import (
  "time"
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/gpio"
  "gobot.io/x/gobot/platforms/firmata"
)

func Arduino() *gobot.Robot {

  fa := firmata.NewAdaptor("/dev/ttyACM0")
  led := gpio.NewLedDriver(fa, "13")

  work := func() {
    gobot.Every(500*time.Millisecond, func() {
          led.Toggle()
        })
  }
  robot := gobot.NewRobot("Blinky",
    []gobot.Connection{fa},
    []gobot.Device{led},
  work,)

  return robot
}
