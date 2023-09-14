package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ostrich/pkg/log"
	"net/http"
)

type GinServer struct {
	*gin.Engine
	Address string
}

// Run .
func (s *GinServer) Run() error {
	srv := &http.Server{
		Addr:    s.Address,
		Handler: s,
	}

	log.Infof("Start to listening the incoming requests on http address: %s", s.Address)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	log.Infof("Server on %s stopped", s.Address)

	return nil
}
