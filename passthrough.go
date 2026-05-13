package sa

import "github.com/SeaVerseAI/sea-sdk-go/internal/transport"

// PassthroughService provides vendor-compatible passthrough APIs.
type PassthroughService struct {
	client *transport.Client
}
