package handlers

import (
	"context"
	"gitlab.com/nikchabanyk/blober/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxBlobQ(q data.Blobs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobQCtxKey, q)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BlobQ(r *http.Request) data.Blobs {
	return r.Context().Value(blobQCtxKey).(data.Blobs).New()
}
