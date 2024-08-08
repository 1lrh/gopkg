package metadatax

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

var MetadataNotFound = errors.New("metadata not found")
var EmptyMetadata = errors.New("empty metadata")

func GetMetadata(ctx context.Context, key string) ([]string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, MetadataNotFound
	}
	return md.Get(fmt.Sprintf("gateway-%s", key)), nil
}

func GetApp(ctx context.Context) (string, error) {
	md, err := GetMetadata(ctx, "app")
	if err != nil {
		return "", err
	}

	if len(md) == 0 || md[0] == "" || len(md[0]) == 0 {
		return "", EmptyMetadata
	}
	return md[0], nil
}
