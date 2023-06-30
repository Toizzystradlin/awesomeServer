package main

import (
	"encoding/json"
	"os"
)

func ParseConfig() Configs {
	bytevalue, _ := os.ReadFile("config.json")
	configs := Configs{}
	err := json.Unmarshal(bytevalue, &configs)
	if err != nil {
		print("ERRRROR UNMARSHALING")
	}
	//for k, v := range configs.ProxyServerConfigs {
	//	fmt.Println(k, v)
	//}
	//fmt.Print(reflect.TypeOf(configs.WebServerConfigs))
	return configs
}

type Configs struct {
	WebServerConfigs   []WebServerConfig   `json:"web_servers"`
	ProxyServerConfigs []ProxyServerConfig `json:"proxy_servers"`
}

type WebServerConfig struct {
	Name         string `json:"name"`
	Port         string `json:"port"`
	DefaultHello string `json:"default_hello"`
}

type ProxyServerConfig struct {
	Name            string `json:"name"`
	Port            string `json:"port"`
	EndPointAddress string `json:"end_point_adress"`
}
