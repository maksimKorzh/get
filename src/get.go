package main

import (
  "os"
  "context"
  "fmt"
  "io/ioutil"
  "net"
  "net/http"
  "time"
  "crypto/tls" // Import the tls package
)

func main() {
  // Create a custom DNS resolver with a specific server address
  resolver := &net.Resolver{
    PreferGo: true,
    Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
      d := net.Dialer{
        Timeout: time.Second * 10,
      }
      return d.DialContext(ctx, "udp", "8.8.8.8:53") // Use Google's public DNS server
    },
  }

  // Create an HTTP client with the custom DNS resolver and insecure skip verification
  client := &http.Client{
    Transport: &http.Transport{
      DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
        d := net.Dialer{
          Timeout:   time.Second * 10,
          Resolver:  resolver,
        }
        return d.DialContext(ctx, network, addr)
      },
      TLSClientConfig: &tls.Config{
        InsecureSkipVerify: true,
      },
    },
  }

  // URL to make the GET request to
  url := ""
  if len(os.Args) > 1 {
    url = os.Args[1]
  } else {
    fmt.Scanf("%s", &url)
  }

  // Send the GET request using the client with custom DNS resolver and insecure skip verification
  response, err := client.Get(url)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  defer response.Body.Close()

  // Read the response body
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("Error reading response body:", err)
    return
  }

  // Print the HTML response
  fmt.Println(string(body))
}
