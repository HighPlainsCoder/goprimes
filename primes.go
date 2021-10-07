package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
  fmt.Println("I will now generate primes forever")
  fmt.Printf("%d ",2) // the first prime, taken as given
  candidate := 3 // the next possible prime

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  done := make(chan bool,1)

  go func() {
    for {
      select {
      case <- done:
        fmt.Printf("\ngot a done so exiting\n")
        done <- false
        return
      default:
        isprime:=true // until shown otherwise
        for ii:=2;ii<candidate;ii++ {
          select {
          case <-done:
            fmt.Printf("\ngot done true in inner loop. not even finishing this check of %v\n",candidate)
            done <- true
            return
          default:
          }
          if candidate % ii == 0 {
            isprime=false
            break
          }
        }  
        if isprime {
          fmt.Printf("%d ",candidate)
        }
        candidate+=1
      }
    }
  }()


  sig := <- sigs
  fmt.Printf("\n\ninterrupting because %v\n\n",sig)
  done <- true

  inner := <- done
  fmt.Printf("\nDone acked %v, so can store state and stuff here.\n",inner)

  fmt.Println("Really done")

}

