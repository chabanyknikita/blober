package requests

import (
	"encoding/json"
	"github.com/google/uuid"
	"gitlab.com/nikchabanyk/blober/internal/data"
	"gitlab.com/nikchabanyk/blober/resources"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type BlobRequest struct {
	BlobID string `json:"-"`
}

func Blob(r resources.BlobRequest) (*data.Blob, error) {

	blob, err := data.GetBlobType(string(r.Type))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create blob")
	}

	return &data.Blob{
		ID:    uuid.NewString(),
		Type:  blob,
		Value: r.Attributes.Value,
	}, nil
}

func NewBlob(blob *data.Blob) resources.Blob {
	b := resources.Blob{
		Key: resources.Key{
			ID:   blob.ID,
			Type: resources.ResourceType(blob.Type.String()),
		},
		Attributes: resources.BlobAttributes{
			Value: blob.Value,
		},
	}
	return b
}

func NewCreateBlobRequest(r *http.Request) (resources.BlobRequest, error) {
	var request resources.BlobRequestResponse

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}

	if request.Data.Relationships.Owner.Data == nil {
		request.Data.Relationships.Owner.Data = &resources.Key{ID: chi.URLParam(r, "address")}
	}

	return request.Data, ValidateCreateBlobRequest(request.Data)
}

func NewGetBlobRequest(r *http.Request) BlobRequest {
	request := BlobRequest{
		BlobID: chi.URLParam(r, "blob"),
	}
	return request
}

func NewDeleteBlobRequest(r *http.Request) (BlobRequest, error) {
	request := BlobRequest{
		BlobID: chi.URLParam(r, "blob"),
	}
	return request, nil
}

func ValidateCreateBlobRequest(r resources.BlobRequest) error {
	return validation.Errors{
		"/data/type":                        validation.Validate(&r.Type, validation.Required),
		"/data/attributes/value":            validation.Validate(&r.Attributes.Value, validation.Required),
		"/data/relationships/owner/data/id": validation.Validate(&r.Relationships.Owner.Data.ID, validation.Required),
	}.Filter()
}
