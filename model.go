package main

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	name    string
	id      int
	mux     *http.ServeMux
	port    string
	handler http.HandlerFunc
	hello   string
}

type ProxyServer struct {
	name              string
	id                int
	mux               *http.ServeMux
	port              string
	end_point_address string
	handler           http.HandlerFunc
	hello             string
}

func (receiver ProxyServer) Start() {
	err := http.ListenAndServe(receiver.port, receiver.mux)
	if err != nil {
		fmt.Println("error on PROXY server", err)
	}
}

func (receiver WebServer) Start() {
	err := http.ListenAndServe(receiver.port, receiver.mux)
	if err != nil {
		fmt.Printf("error on Origin Server: %s", err)
	}
}
