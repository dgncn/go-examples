package dynowebportal

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunWebPortal(serverAddress string) error {
	http.HandleFunc("/", handleUrl)
	err := http.ListenAndServe(serverAddress, nil)
	return err
}

func handleUrl(respWriter http.ResponseWriter, httpReq *http.Request) {
	log.Println("server started")
	fmt.Fprintf(respWriter, "Hoş geldin Dogancan saat: %v", time.Now())
}
