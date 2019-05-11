// docker run -p 80:8080 -p 40000:40000 --security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE -d hello_server

package main

import (
    "hello"
    "log"
    "net/http"
)

func main() {
    log.Println("starting hello...")
    http.HandleFunc("/", hello.Timestamp)
    log.Fatal(http.ListenAndServe(":8080", nil))
}