package servers

import (
	"errors"
	"net/http"
)

func LoadingServerS() error {

	mux := http.NewServeMux()

	mux.HandleFunc(homePagePath, MainHandler)

	errLoadingServer := http.ListenAndServe(portAddressNumber, mux)
	if errLoadingServer != nil {
		errMsg = "----<-LoadingServerS------<--ListenAndServe----" + errLoadingServer.Error()
		logMsg = "Error in running server"
		logger.Error(logMsg + errMsg)
		return errors.New(errMsg)
	}
	logger.Info("Server running at port 8080")
	return nil
}
