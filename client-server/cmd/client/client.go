package main

import (
    "client-server/internal"
    "fmt"
    "io"
    "math/rand"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    if len(os.Args) > 2 {
        usage()
    }

    ports := []int{}
    if len(os.Args) == 2 {
        strPorts := strings.Split(os.Args[1], ":")
        if len(strPorts) != 2 {
            usage()
        }
        portStart, _ := strconv.Atoi(strPorts[0])
        portEnd, _ := strconv.Atoi(strPorts[1])
        for portStart <= portEnd {
            ports = append(ports, portStart)
            portStart += 1
        }
    } else {
        ports = append(ports, internal.Port)
    }
    fmt.Printf("Sending requests to ports: %v\n", ports)

    // Keep hitting endpoints indefinitely
    count := 0
    for {
        makeRequest(ports[rand.Intn(len(ports))], internal.AsiaPath, count)
        count++
        makeRequest(ports[rand.Intn(len(ports))], internal.AmericaPath, count)
        count++
    }
}

func usage() {
    fmt.Printf("Usage: %s [port-start[:port-end]]\n", os.Args[0])
    os.Exit(1)
}

func makeRequest(port int, restPath string, count int) {
    response, err := http.Get("http://:" + strconv.Itoa(port) + restPath)
    if err != nil {
        fmt.Printf("[%d] Failed", count)
    }
    body, err := io.ReadAll(response.Body)
    fmt.Printf("[%d] To: %d; Got: %d {%s} : %s\n", count, port, response.StatusCode,
        response.Header.Get("Server"), string(body))
    count++
    _ = response.Body.Close()
    time.Sleep(time.Second * 2)
}
