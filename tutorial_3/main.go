package main

import (
  "fmt"
  "gobot.io/x/gobot"
)

func main() {
  master := gobot.NewMaster()

  master.AddRobot(Arduino())
  master.AddRobot(Camera())

  fmt.Println("Starting master")
  master.Start()
}
