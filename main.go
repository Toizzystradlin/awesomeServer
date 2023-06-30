package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	config := ParseConfig()
	for k, serverConfig := range config.WebServerConfigs {
		wg.Add(1)
		fmt.Println(serverConfig.Port)
		go startWebServer(serverConfig.Name, k, serverConfig.Port, serverConfig.DefaultHello, &wg)
	}

	//var wg_proxy sync.WaitGroup
	for k, serverConfig := range config.ProxyServerConfigs {
		wg.Add(1)
		fmt.Println(serverConfig.Port)
		go StartProxyServer(serverConfig.Name, k, serverConfig.Port, serverConfig.EndPointAddress, &wg)
	}
	wg.Wait()
	fmt.Println("Done!")

	//go func() { StartProxyServer(":8080") }()
	//select {}
}

func startWebServer(name string, id int, port string, hello string, wg *sync.WaitGroup) {
	defer wg.Done()
	newWebServer := WebServer{
		name:  name,
		id:    id,
		mux:   http.NewServeMux(),
		port:  port,
		hello: hello,
	}
	fmt.Println("Starting new Web Server : ", newWebServer)
	newWebServer.mux.HandleFunc("/", HelloHandler)
	newWebServer.Start()

	return
}

func StartProxyServer(name string, id int, port string, end_point_adress string, wg *sync.WaitGroup) {
	defer wg.Done()
	newProxyServer := ProxyServer{
		name:              name,
		id:                id,
		mux:               http.NewServeMux(),
		port:              port,
		handler:           Redirect,
		end_point_address: end_point_adress,
	}
	newProxyServer.mux.HandleFunc("/", Redirect)
	fmt.Println("Starting new Proxy Server : ", newProxyServer)
	newProxyServer.Start()
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s got new request\n")
	w.Header().Add("test-headerXXX", "test-header-valueXXX")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("This is my websiteXXX"))
}

func Redirect(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	newReq, err := http.NewRequest(req.Method, "http://127.0.0.1:9001", req.Body)
	fmt.Println(newReq)
	if err != nil {
		print("ERRROR")
	}
	//newReq.Header = req.Header.Clone()
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(newReq)
	if err != nil {
		print(err)
	}
	//w.Write([]byte("aaaaaaaasssssssssssssssss"))
	//print(resp.Body)
	io.Copy(w, resp.Body)
}
