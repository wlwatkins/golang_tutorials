package main

import (
  "fmt"
  "log"
  "net/http"

  _ "net/http/pprof"

  "github.com/hybridgroup/mjpeg"
  "gocv.io/x/gocv"
)

var (
  deviceID int
  err error
  webcam *gocv.VideoCapture
  stream *mjpeg.Stream
)

func mjpegCapture() {
  img := gocv.NewMat()
  defer img.Close()

  for {
    if ok := webcam.Read(&img); !ok {
      fmt.Printf("Device closed: %v\n", deviceID)
      return
    }

    if img.Empty() {
      continue
    }

    buf, _ := gocv.IMEncode(".jpg", img)
    stream.UpdateJPEG(buf)
  }
}

func main(){

  deviceID = 0 //  -1 ?
  host := "0.0.0.0:8082"

  webcam, err = gocv.OpenVideoCapture(deviceID)
  if err != nil {
    fmt.Printf("Error opening capture device: %v\n", deviceID)
    return
  }
  defer webcam.Close()

  width := 640
  height := float64(width)*0.75

  webcam.Set(3, float64(width))
  webcam.Set(4, float64(height))
  // webcam.Set(5, fpsValue) // fps

  stream = mjpeg.NewStream()

  go mjpegCapture()

  fmt.Println("Capturing, point your browser to", host)

  http.Handle("/", stream)
  log.Fatal(http.ListenAndServe(host, nil))
}
