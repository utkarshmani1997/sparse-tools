package rest

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/openebs/sparse-tools/sparse"
)

type SyncServer struct {
	filePath string
	fileIo   sparse.FileIoProcessor
}

// TestServer daemon serves only one connection for each test then exits
func TestServer(port string, filePath string, timeout int) {
	Server("", port, filePath)
}

func Server(hostIP string, port string, filePath string) error {
	log.Infof("Creating Ssync service")
	router := NewRouter(&SyncServer{filePath: filePath})
	return http.ListenAndServe(hostIP+":"+port, router)
}
