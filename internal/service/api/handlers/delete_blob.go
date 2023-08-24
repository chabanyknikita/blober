package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/nikchabanyk/blober/internal/data"
	"gitlab.com/nikchabanyk/blober/internal/service/api/requests"
	"net/http"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = BlobQ(r).Transaction(func(blobs data.Blobs) error {
		err := blobs.DeleteByID(request.BlobID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		Log(r).WithError(err).Error("failed to delete blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(201)
	ape.Render(w, "Successful deleted blob choosing blob")
}
