package main

import (
    "client-server/internal"
    "fmt"
    "net/http"
    "os"
    "runtime"
    "strconv"
)

func usage() {
    fmt.Printf("Usage: %s [port]\n", os.Args[0])
    os.Exit(1)
}

func main() {
    if len(os.Args) > 2 {
        usage()
    }

    port := internal.Port
    var err error
    if len(os.Args) == 2 {
        if port, err = strconv.Atoi(os.Args[1]); err != nil {
            usage()
        }
    }
    fmt.Printf("Starting Golang server on port %d\n", port)
    countryServer := http.NewServeMux()

    count := 0
    handlerFunc := func(writer http.ResponseWriter, request *http.Request) {
        if request.Method != "GET" {
            writer.WriteHeader(http.StatusBadRequest)
            return
        }
        fmt.Printf("[%d] Received %s request from: %s\n", count, request.URL.Path,
            request.UserAgent())
        count++
        writer.Header().Add("Server", runtime.Version())
        writer.WriteHeader(http.StatusOK)
        country := "Unknown"
        if request.URL.Path == internal.AsiaPath {
            country = internal.AsiaCountry
        } else if request.URL.Path == internal.AmericaPath {
            country = internal.AmericaCountry
        }
        _, _ = writer.Write([]byte(country))
    }
    countryServer.HandleFunc("/", handlerFunc)

    _ = http.ListenAndServe(":"+strconv.Itoa(port), countryServer)
}
