package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    protoc "github.com/jeffotoni/go.protobuffer.customer"
)

func main() {
    body := &protoc.Customer{
        Id:   12304,
        Name: "Carlos",
    }

    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(body)
    req, err := http.NewRequest("POST", "localhost:8080/customer/proto", buf)
    req.Header.Set("Content-Type", "application/proto")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        log.Println("Error Do:", err)
        return
    }

    defer res.Body.Close()
    fmt.Println("response Status:", res.Status)
    io.Copy(os.Stdout, res.Body)
}
