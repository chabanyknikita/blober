package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/nikchabanyk/blober/internal/data"
	"gitlab.com/nikchabanyk/blober/internal/data/postgres"
	"gitlab.com/nikchabanyk/blober/internal/service/api/requests"
	"gitlab.com/nikchabanyk/blober/resources"
	"net/http"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, err := requests.Blob(request)
	if err != nil {
		Log(r).WithError(err).Warn("invalid blob type")
		ape.RenderErr(w, problems.BadRequest(
			validation.Errors{
				"/data/type": errors.New("invalid blob type"),
			})...)
		return
	}

	err = BlobQ(r).Transaction(func(blobs data.Blobs) error {
		if err := blobs.Create(blob); err != nil {
			return errors.Wrap(err, "failed to create blob")
		}
		return nil
	})
	if err != nil {
		if errors.Cause(err) != postgres.ErrBlobsConflict {
			Log(r).WithError(err).Error("failed to save blob")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	response := resources.BlobResponse{
		Data: requests.NewBlob(blob),
	}
	w.WriteHeader(201)
	ape.Render(w, &response)
}
