package service

import (
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/nikchabanyk/blober/internal/config"
	"gitlab.com/nikchabanyk/blober/internal/data"
	"gitlab.com/nikchabanyk/blober/internal/data/postgres"
	"golang.org/x/net/context"
	"net"
	"net/http"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	config   config.Config
	ctx      context.Context
	blobsQ   data.Blobs
}

func (s *service) Config() config.Config {
	return s.config
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		blobsQ:   postgres.NewBlobs(cfg.DB()),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
