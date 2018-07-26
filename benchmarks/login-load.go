package main

import (
  "fmt"
  "math/rand"
  "os"
  "time"

  "github.com/matrix-org/gomatrix"
)

// numUsers is the number of test users to generate
const numUsers = 100
const testPass = "testing123"

func main() {
  if len(os.Args) < 2 {
    fmt.Println("You must provide a port number.")
    return
  }

  hsPort := os.Args[1]
  gm, err := gomatrix.NewClient("http://localhost:"+hsPort, "", "")
  if err != nil {
    fmt.Println("Error:", err)
  }

  for i := 0; i < numUsers; i++ {
    if err = createAccount(gm); err != nil {
      fmt.Println("Error registering:", err)
      return
    }
    fmt.Println("Registered user", i)
  }
}

// createAccount creates a new dummy user account
func createAccount(gm *gomatrix.Client) (err error) {
  username := "testing-" + randString(5)
  // Get the session token
  req := &gomatrix.ReqRegister{
    Username: username,
    Password: testPass,
  }
  _, respInt, err := gm.Register(req)
  if err != nil {
    return
  }

  // Make a dummy register request
  req = &gomatrix.ReqRegister{
    Username: username,
    Password: testPass,
    Auth: struct {
      Session string
    }{
      Session: respInt.Session,
    },
  }
  _, err = gm.RegisterDummy(req)
  return
}

var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
  letterIdxBits = 6                    // 6 bits to represent a letter index
  letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
  letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randString(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
  }
