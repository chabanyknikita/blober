package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/nikchabanyk/blober/internal/service/api/requests"
	"gitlab.com/nikchabanyk/blober/resources"
	"net/http"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	log := Log(r).WithField("tag", "blob_performance")
	log.Info("Request started")
	request := requests.NewGetBlobRequest(r)

	log.Info("Try to get blob from DB")
	blob, err := BlobQ(r).Get(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	log.Info("Render response")
	response := resources.BlobResponse{
		Data: requests.NewBlob(blob),
	}

	ape.Render(w, &response)
}
