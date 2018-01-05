package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/gorilla/mux"
)

const (
	// DefaultPort to listen if no PORT declared
	DefaultPort = "9000"
)

var appEnv, _ = cfenv.Current()

// GetRootHandler Get /
func GetRootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Info, Getting root information.")
	w.Write([]byte("Hello World!"))
	//json.NewEncoder(w).Encode("Hello World!")
}

// GetCfEnvHandler Get /v1/cf_envs
func GetCfEnvHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Info, Getting cf environments.")
	//json.NewEncoder(w).Encode(os.Getenv("VCAP_APPLICATION"))
	w.Write([]byte(os.Getenv("VCAP_APPLICATION")))
}

// GetAppNameHandler Get /v1/app_name
func GetAppNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Info, Getting application name.")
	w.Write([]byte(appEnv.Name))
}

// GetSpaceNameHandler Get /v1/space_name
func GetSpaceNameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Info, Getting space name.")
	w.Write([]byte(appEnv.SpaceName))
}

func main() {
	var port string
	// Cloud Foundry assigns random ports
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v\n", DefaultPort)
		port = DefaultPort
	}

	router := mux.NewRouter()
	router.HandleFunc("/", GetRootHandler).Methods("GET")
	router.HandleFunc("/v1/cf_envs", GetCfEnvHandler).Methods("GET")
	router.HandleFunc("/v1/app_name", GetAppNameHandler).Methods("GET")
	router.HandleFunc("/v1/space_name", GetSpaceNameHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
