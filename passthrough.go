package sa

import "github.com/SeaVerseAI/sa-go/internal/transport"

// PassthroughService provides vendor-compatible passthrough APIs.
type PassthroughService struct {
	client *transport.Client
}
