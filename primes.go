package main
import "fmt"
func main() {
  fmt.Println("I will now generate primes forever")
  fmt.Printf("%d ",2) // the first prime, taken as given
  candidate := 3 // the next possible prime

  for true {
    isprime:=true // until shown otherwise
    for ii:=2;ii<candidate;ii++ {
      if candidate % ii == 0 {
        isprime=false
        break // from the for loop?
      }
    }  
    if isprime {
      fmt.Printf("%d ",candidate)
    }
    candidate+=1
  }
}

