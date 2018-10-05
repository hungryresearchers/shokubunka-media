package config

import (
	"context"

	"github.com/google/go-cloud/blob"
	"github.com/google/go-cloud/blob/gcsblob"
	"github.com/google/go-cloud/gcp"
)

func Setup(ctx context.Context, service string) (*blob.Bucket, error) {
	switch service {
	case "gcp":
		return SetupGCP(ctx, "shokubunka-media-image")
	default:
		return SetupGCP(ctx, "shokubunka-media-image")
	}
}

func SetupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
	if err != nil {
		return nil, err
	}
	return gcsblob.OpenBucket(ctx, bucket, c)
}
