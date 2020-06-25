package main

import (
  "fmt"
  "time"
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/gpio"
  "gobot.io/x/gobot/platforms/firmata"
)

func servoManager(s *gpio.ServoDriver) {
  s.Center()
  time.Sleep(200*time.Millisecond)

  fmt.Println("Initiql Hello")
  for i := 0; i < 5 ; i++ {
    s.Move(20)
    time.Sleep(500*time.Millisecond)
    s.Move(160)
    time.Sleep(500*time.Millisecond)
  }

  s.Center()
  time.Sleep(200*time.Millisecond)

  for i := 0; i < 5 ; i++ {

    fmt.Println("Going up")
    for n := 20; n<160;n++{
      s.Move(uint8(n))
      time.Sleep(50*time.Millisecond)
    }

    time.Sleep(200*time.Millisecond)
    fmt.Println("Going down")
    for n := 160; n>20;n--{
      s.Move(uint8(n))
      time.Sleep(50*time.Millisecond)
    }
  }

}

func Arduino() *gobot.Robot {

  fa := firmata.NewAdaptor("/dev/ttyACM0")
  servoY := gpio.NewServoDriver(fa, "2")

  work := func() {
    servoManager(servoY)
  }
  robot := gobot.NewRobot("Blinky",
    []gobot.Connection{fa},
    []gobot.Device{servoY},
    work,
  )

  return robot
}
