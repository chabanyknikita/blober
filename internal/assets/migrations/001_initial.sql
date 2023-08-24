-- +migrate Up

CREATE TABLE blobs (
    id TEXT PRIMARY KEY,
    "type" BIGINT,
    "value" TEXT
);

-- +migrate Down

DROP TABLE blobs;