package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
    ServerName string `json:"serverName"`
    ServerIP string `json:"serverIP"`
}

type ServerSlice struct {
    Servers []Server `json:"servers"`
}

func main() {
    var s ServerSlice
    s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    b, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    }
    fmt.Println(string(b))
}
