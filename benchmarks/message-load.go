package main

import (
  "fmt"
  "math/rand"
  "os"
  "time"
  "sync"

  "github.com/matrix-org/gomatrix"
)

// numUsers is the number of test users to generate
const (
  numUsers = 20
  numMessages = 50
  testPass = "testing123"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("You must provide a port number.")
    return
  }

  hsPort := os.Args[1]
  hsHost := "localhost"
  hsURL := "http://" + hsHost + ":" + hsPort
  gm, err := gomatrix.NewClient(hsURL, "", "")
  if err != nil {
    fmt.Println("Error:", err)
  }

  // Make clients for later use
  clients := make([]*gomatrix.Client, 0, numUsers)
  for i := 0; i < numUsers; i++ {
    accessToken, userID, err := createAccount(gm)
    if err != nil {
      fmt.Println("Error registering:", err)
      return
    }

    // New client
    c, err := gomatrix.NewClient(hsURL, userID, accessToken)
    if err != nil {
      fmt.Println("Error creating client:", err)
    }
    clients = append(clients, c)
  }

  // Create a room
  roomAliasName := randString(10)
  roomID, err := createRoom(clients[0], roomAliasName)
  if err != nil {
    fmt.Println("Error creating room:", err)
  }

  // Join the clients to the room
  roomAlias := "#" + roomAliasName + ":" + hsHost
  for _, c := range clients {
    content := struct{}{}
    _, err := c.JoinRoom(roomAlias, "", content)
    if err != nil {
      fmt.Println("Error joining room:", err)
      return
    }
  }
  fmt.Println("Room Alias:", roomAlias)

  // Create a message
  msg := &gomatrix.TextMessage{
    MsgType: "m.text",
    Body: "Hey there",
  }

  // Start sending messages in the room
  var wg sync.WaitGroup
  wg.Add(numUsers * numMessages)
  for i := range clients {
    go func(c *gomatrix.Client) {
      for j := 0; j < numMessages; j++ {
        _, err := c.SendMessageEvent(roomID, "m.room.message", msg)
        if err != nil {
          fmt.Println("Error:", err)
        }
        wg.Done()
      }
    }(clients[i])
  }
  wg.Wait()
}

// createRoom creates a new dummy room and returns the room ID
func createRoom(gm *gomatrix.Client, roomAliasName string) (roomID string, err error) {
  req := &gomatrix.ReqCreateRoom{
    RoomAliasName: roomAliasName,
    Preset: "public_chat",
  }

  // Create the room
  resp, err := gm.CreateRoom(req)
  if err != nil {
    return
  }

  roomID = resp.RoomID
  return
}

// createAccount creates a new dummy user account
func createAccount(gm *gomatrix.Client) (accessToken, userID string, err error) {
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
  resp, err := gm.RegisterDummy(req)
  if err != nil {
    return
  }

  // Save the access token and UserID
  accessToken = resp.AccessToken
  userID = resp.UserID
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
