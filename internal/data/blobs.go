package data

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Blob struct {
	ID    string   `db:"id" structs:"id"`
	Type  BlobType `db:"type" structs:"type"`
	Value string   `db:"value" structs:"value"`
}

//go:generate mockery -case underscore -name Blobs
type Blobs interface {
	New() Blobs
	Transaction(fn func(Blobs) error) error
	Create(blob *Blob) error
	Get(id string) (*Blob, error)
	DeleteByID(iD string) error
	GetAll() ([]*Blob, error)
}

const (
	BlobTypeAssetDescription BlobType = 1 << iota
	BlobTypeFundOverview
	BlobTypeFundUpdate
	BlobTypeNavUpdate
	BlobTypeFundDocument
	BlobTypeAlpha
	BlobTypeBravo
	BlobTypeCharlie
	BlobTypeDelta
	BlobTypeTokenTerms
	BlobTypeTokenMetrics
	BlobTypeKYCForm
	BlobTypeKYCIdDocument
	BlobTypeKYCPoa
	BlobTypeIdentityMindReject
)

var (
	ErrAddressInvalid = errors.New("address is invalid")
)

var _BlobTypeValueToName = map[BlobType]string{
	BlobTypeAssetDescription:   "asset_description",
	BlobTypeFundOverview:       "fund_overview",
	BlobTypeFundUpdate:         "fund_update",
	BlobTypeNavUpdate:          "nav_update",
	BlobTypeFundDocument:       "fund_document",
	BlobTypeAlpha:              "alpha",
	BlobTypeBravo:              "bravo",
	BlobTypeCharlie:            "charlie",
	BlobTypeDelta:              "delta",
	BlobTypeTokenTerms:         "token_terms",
	BlobTypeTokenMetrics:       "token_metrics",
	BlobTypeKYCForm:            "kyc_form",
	BlobTypeKYCIdDocument:      "kyc_id_document",
	BlobTypeKYCPoa:             "kyc_poa",
	BlobTypeIdentityMindReject: "identity_mind_reject",
}

var _BlobTypeNameToValue = map[string]BlobType{
	"asset_description":    BlobTypeAssetDescription,
	"fund_overview":        BlobTypeFundOverview,
	"fund_update":          BlobTypeFundUpdate,
	"nav_update":           BlobTypeNavUpdate,
	"fund_document":        BlobTypeFundDocument,
	"alpha":                BlobTypeAlpha,
	"bravo":                BlobTypeBravo,
	"charlie":              BlobTypeCharlie,
	"delta":                BlobTypeDelta,
	"token_terms":          BlobTypeTokenTerms,
	"token_metrics":        BlobTypeTokenMetrics,
	"kyc_form":             BlobTypeKYCForm,
	"kyc_id_document":      BlobTypeKYCIdDocument,
	"kyc_poa":              BlobTypeKYCPoa,
	"identity_mind_reject": BlobTypeIdentityMindReject,
}

//go:generate jsonenums -tprefix=false -transform=snake -type=BlobType
type BlobType int32

func GetBlobType(v string) (b BlobType, err error) {
	err = b.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, v)))
	if err != nil {
		return b, errors.Wrap(err, "failed to unmarshal blob type")
	}
	return b, nil
}

func (r *BlobType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BlobType should be a string, got %s", data)
	}
	v, ok := _BlobTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BlobType %q", s)
	}
	*r = v
	return nil
}

func (r BlobType) String() string {
	s, ok := _BlobTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("BlobType(%d)", r)
	}
	return s
}
