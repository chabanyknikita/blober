package handlers

import (
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/nikchabanyk/blober/internal/service/api/requests"
	"gitlab.com/nikchabanyk/blober/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func GetAllBlobs(w http.ResponseWriter, r *http.Request) {
	log := Log(r).WithField("tag", "blob_performance")
	log.Info("Request started")

	blobs, err := BlobQ(r).GetAll()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blobs")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var blobResponses []resources.BlobResponse
	for _, blob := range blobs {
		blobResponses = append(blobResponses, resources.BlobResponse{
			Data: requests.NewBlob(blob),
		})
	}

	log.Info("Render response")

	ape.Render(w, &blobResponses)
}
