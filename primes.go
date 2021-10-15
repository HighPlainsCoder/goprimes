package main

import (
    "fmt"
    "os"
    "os/signal"
    "runtime"
    "syscall"
    "time"
)

var candidate int = 3 // the next possible prime

func getNextCandidate() {
  candidate += 2
}

func testPossible(applicant int, done chan struct{}) bool {
  isprime:=true // until shown otherwise
  for ii:=2;ii<applicant;ii++ {
    time.Sleep(1*time.Second)

    select {
    case <-done:
      fmt.Printf("\ngot done in inner loop. not even finishing this check of %v\n",applicant)
      time.Sleep(1*time.Second)
      return false
    default:
    }

    if applicant % ii == 0 {
      isprime=false
      break
    }
  }
  return isprime
}

func looper(done chan struct{}) {
  for {
    // this gag keeps showing up where channels & goroutines are
    select {
    case <- done:
      fmt.Printf("\ngot a done so exiting\n")
      return
    default:
    }

    time.Sleep(1*time.Second)
    isprime := testPossible(candidate, done)
    if isprime {
      fmt.Printf("%d ",candidate)
    }

    getNextCandidate()
  }
}




func main() {
  fmt.Println("I will now generate primes forever")
  fmt.Printf("%d ",2) // the first prime, taken as given

  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
  done := make(chan struct{},1)

  go looper(done)

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

