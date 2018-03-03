package main

import (
  "fmt"
  "os"
  "net"
  "time"
  "io/ioutil"
)

const PACKETSIZE = 1500

func main()  {

  /* 4 Arguments (3 plus Arg[0]) needed to run the program (1) local address (2) blink node address
     (3) destination address */
  if len(os.Args) != 3 {
    fmt.Println("Incorrect number of arguments!!!")
    os.Exit(1)
  }

  localAddr, err := net.ResolveUDPAddr("udp", os.Args[1])
  CheckError(err)

  destAddr, err := net.ResolveUDPAddr("udp", os.Args[2])
  CheckError(err)

  filename := "1.jpg"

  videoBuffer, err := ioutil.ReadFile(filename)
  CheckError(err)

  conn, _ := net.ListenUDP("udp", localAddr)

  outputBuf := make([]byte, PACKETSIZE)

  for i:=0; i < len(videoBuffer); i += PACKETSIZE {
    if i + PACKETSIZE + 1 < len(videoBuffer) {
      copy(outputBuf[:], videoBuffer[i:i+PACKETSIZE])
    } else {
      finalValue := len(videoBuffer)-1
      outputBuf = fillByteArray(videoBuffer[i:finalValue], PACKETSIZE)
    }


    _, err = conn.WriteToUDP(outputBuf, destAddr)
    CheckError(err)
    fmt.Println("Buffer:", outputBuf)
    time.Sleep(1 * time.Millisecond)
  }
}


func fillByteArray(buffer []byte, toLength int) []byte {
 outBuf := make([]byte, toLength)
 copy(outBuf[:], buffer[:])
    for i:=len(buffer); i < toLength; i++ {
        outBuf[i] = ":"[0]
    }

    return outBuf
}

func CheckError(err error)  {
  if err != nil {
  		fmt.Println("Error: ", err)
      os.Exit(1)
  	}
}
