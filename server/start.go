package server

import (
	"net/http"
	pathUtil "path/filepath"

	"github.com/Raincal/rikka/common/logger"
	"github.com/Raincal/rikka/common/util"
	"github.com/Raincal/rikka/server/apiserver"
	"github.com/Raincal/rikka/server/webserver"
)

var (
	l = logger.NewLogger("[Server]")
)

// StartRikka start server part of rikka. Include Web Server and API server.
func StartRikka(socket string, password string, maxSizeByMb float64, https bool, certDir string) {

	l.Info("Start web server...")
	viewPath := webserver.StartRikkaWebServer(maxSizeByMb, l)

	l.Info("Start API server")
	apiserver.StartRikkaAPIServer(viewPath, password, maxSizeByMb, l)

	l.Info("Rikka is listening", socket)

	// real http server function call
	var err error
	if https {
		if !util.IsDir(certDir) {
			l.Fatal("Cert dir argument is not a valid dir")
		}
		err = http.ListenAndServeTLS(
			socket,
			pathUtil.Join(certDir, "cert.pem"),
			pathUtil.Join(certDir, "key.pem"),
			nil,
		)
	} else {
		err = http.ListenAndServe(socket, nil)
	}

	if err != nil {
		l.Fatal("Error when try listening", socket, ":", err)
	}
}
