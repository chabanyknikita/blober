package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	handlers2 "gitlab.com/nikchabanyk/blober/internal/service/api/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers2.CtxLog(s.log),
			handlers2.CtxBlobQ(s.blobsQ),
		),
	)
	r.Route("/integrations", func(r chi.Router) {
		r.Route("/blobs-svc", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Route("/public", func(r chi.Router) {
					r.Post("/", handlers2.CreateBlob)
					r.Get("/{blob}", handlers2.GetBlob)
				})
				r.Route("/private", func(r chi.Router) {
					r.Get("/", handlers2.GetAllBlobs)
					r.Delete("/{blob}", handlers2.DeleteBlob)
				})
			})
		})
	})

	return r
}
