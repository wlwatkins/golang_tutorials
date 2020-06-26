package main

import (
  "fmt"
  // "time"
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/gpio"
  "gobot.io/x/gobot/platforms/firmata"
)

func stepperManager(stepper *gpio.StepperDriver, fa *firmata.Adaptor) {

  fmt.Println("Init stepper")
  for {
    hit, err := fa.DigitalRead("4")
    if err != nil {
      panic("Could not read pin 4")
    }

    if hit == 0 {
      stepper.Move(-100)
      break
    } else {
      stepper.Move(10)
    }
  }

  for j:= 0 ; j< 5 ; j++{
  for i := 0; i < 10; i++ {
    stepper.Move(-50)
  }

  for i := 0; i < 10; i++ {
    stepper.Move(50)
  }
}

}

func Arduino() *gobot.Robot {

  fa := firmata.NewAdaptor("/dev/ttyACM0")
  stepper := gpio.NewStepperDriver(fa, [4]string{"9", "8", "7", "10"}, gpio.StepperModes.DualPhaseStepping, 2048)

  work := func() {
    stepperManager(stepper, fa)
  }

  robot := gobot.NewRobot("Stepper",
    []gobot.Connection{fa},
    []gobot.Device{stepper},
    work,
  )

  return robot
}
