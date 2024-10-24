package vertex

import (
	"local-test/internal/config"
	"local-test/pkg/apperrors"
	"sync"
)

var (
	once sync.Once
	vertexConfig *config.VertexConfig
)

func InitVertexConfig(c *config.VertexConfig) error {
	// Validate Vertex configuration
	if err := validateVertexConfig(c); err != nil {
		return apperrors.WrapInitError(
			&apperrors.ErrOperationFailed{
				Operation: "validate vertex config",
				Err:       err,
			},
		)
	}

	once.Do(func() {
		vertexConfig = c
	})
	return nil
}