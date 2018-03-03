package main

import (
  "fmt"
  "net"
  "io/ioutil"
  "os"
)

const PACKETSIZE = 20000

func main() {
  
  myAddress, _ := net.ResolveUDPAddr("udp", "155.246.202.19:3000")

  conn, _ := net.ListenUDP("udp", myAddress)

  buffer := make([]byte, PACKETSIZE)

  var finalBuffer []byte
  var fileHasEnded = false
  var toWriteToFile []byte

  for fileHasEnded == false {
    conn.ReadFromUDP(buffer)
    //fmt.Println("Data:", buffer)
    toWriteToFile, fileHasEnded = unfillByteArray(buffer, PACKETSIZE)
    for _, value := range toWriteToFile {
      finalBuffer = append(finalBuffer, value)
    }
  }

  //fmt.Println("Output:", finalBuffer)
  _ = ioutil.WriteFile("video.mp4", []byte(finalBuffer), os.ModeAppend)
  fmt.Println("DONE")
}

func unfillByteArray(inBuf []byte, fromLength int) ([]byte, bool) {

  finalOutBuf := make([]byte, PACKETSIZE)
  outBuf := make([]byte, PACKETSIZE)
  fileHasEnded := false

  for i:=0; i < len(inBuf); i++ {
      if ((i+6) < len(inBuf)) && (string(inBuf[i:i+6]) == "::::::") {
        fmt.Println("FOUND IT")
        outBuf2 := make([]byte, i)
        copy(outBuf2[:], inBuf[:i])
        finalOutBuf = outBuf2
        fileHasEnded = true
        break
      } else {
        outBuf[i] = inBuf[i]
        finalOutBuf = outBuf
        fileHasEnded = false
      }
    }

    return finalOutBuf, fileHasEnded
  }


func CheckError(err error) {
  if err != nil {
    fmt.Println("Err:", err)
    os.Exit(1)
  }
}
