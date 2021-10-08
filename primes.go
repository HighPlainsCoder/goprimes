package main

import (
    "fmt"
    "os"
    "os/signal"
    "runtime"
    "syscall"
    "time"
)

func main() {
  fmt.Println("I will now generate primes forever")
  fmt.Printf("%d ",2) // the first prime, taken as given
  candidate := 3 // the next possible prime

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  done := make(chan struct{},1)

  go func() {
    for {
      // this gag keeps showing up where channels & goroutines are
      select {
      case <- done:
        fmt.Printf("\ngot a done so exiting\n")
        return
      default:
      }

      time.Sleep(1*time.Second)
      isprime:=true // until shown otherwise
      for ii:=2;ii<candidate;ii++ {
        time.Sleep(1*time.Second)
        
        select {
        case <-done:
          fmt.Printf("\ngot done in inner loop. not even finishing this check of %v\n",candidate)
          time.Sleep(1*time.Second)
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
  }()


  <- sigs
  fmt.Printf("\n\ninterrupted by user ...\n\n")
  signal.Stop(sigs)
  close(sigs)
  close(done)

  // 2 routines left, main() and the signal handler
  for n:=runtime.NumGoroutine();n>2;n=runtime.NumGoroutine() {
    time.Sleep(100*time.Millisecond)
  }

  fmt.Printf("\nNow I can store state and wrap up here.\n")

  fmt.Println("Really done")

}

