package test_utils

import (
	"io"

	"github.com/samber/lo"
)

// Reads a http response body and unmarshals it into the specified type T.
func ParseResponseBody[T any](body io.ReadCloser) T {
	return lo.FromBytes[T](lo.Ok(io.ReadAll(body)))
}
