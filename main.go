package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Creamos nuestras estructuras
type serverGeneral struct {
	Host              string              `json:"host"`
	Port              int                 `json:"port"`
	Protocol          string              `json:"protocol"`
	IsPublic          bool                `json:"isPublic"`
	Status            string              `json:"status"`
	StartTime         int                 `json:"startTime"`
	EngineVersion     float32             `json:"engineVersion"`
	CriteriaVersion   complex128          `json:"criteriaVersion"`
	Endpointslevelone []Endpointslevelone `json:"endpoints"`
	//Endpointsleveltwo Endpointsleveltwo
}

//endpoints level-one
type Endpointslevelone struct {
	IpAddress            float64           `json:"ipAddress"`
	ServerName           string            `json:"serverName"`
	StatusMessage        string            `json:"statusMessage"`
	StatusDetails        string            `json:"statusDetails"`
	StatusDetailsMessage string            `json:"statusDetailsMessage"`
	Delegation           int               `json:"delegation"`
	SecondServer         Endpointsleveltwo `json:"endpoints"`
}

//endpoints level-two
type Endpointsleveltwo struct {
	IpAddress     float64 `json:"ipAddress"`
	ServerName    string  `json:"serverName"`
	StatusMessage string  `json:"statusMessage"`
	Delegation    int     `json:"delegation"`
}

// Nuestro query url de la API
func main() {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=truora.com")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//Unmarshalling our JSON
	var responseObject serverGeneral
	json.Unmarshal(responseData, &responseObject)
	//impresiones
	//fmt.Println(responseObject.Host)
	fmt.Println(len(responseObject.Endpointslevelone))
	//Listamos servidores
	for i := 0; i < len(responseObject.Endpointslevelone); i++ {
		fmt.Println(responseObject.Endpointslevelone[i].SecondServer.ServerName)
	}
}
