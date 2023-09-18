package facility

import (
	"context"
)

type IFacility interface {
	SendRequest(ctx context.Context, command string, payload map[string]interface{}) ([]byte, error)
}
