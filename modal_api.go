package sa

import (
	"context"
	"net/http"

	mmservice "github.com/seaart/sa-go/internal/multimodal/service"
	mmtypes "github.com/seaart/sa-go/internal/multimodal/types"
	"github.com/seaart/sa-go/internal/transport"
)

func (m *ModalService) Create(ctx context.Context, body JSONMap, opts ...RequestOption) (*Task, error) {
	resp, err := mmservice.CreateTask(m.client, ctx, body, buildRequestOptions(opts).headers)
	if err != nil {
		return nil, err
	}
	return &Task{
		ID:     resp.ID,
		Status: resp.Status,
		Model:  resp.Model,
		Error:  resp.Error,
		client: m.client,
	}, nil
}

// ListModels searches multimodal model skills via GET /v1/models/skill/search.
//
// Supported params:
//   - Query maps to q
//   - Input maps to input
//   - Output maps to output
//   - Type maps to type
//   - Provider maps to provider
//   - Limit maps to limit
func (m *ModalService) ListModels(ctx context.Context, params ModalModelSearchParams, opts ...RequestOption) (*ModalModelSearchResponse, error) {
	return mmservice.SearchModels(m.client, ctx, mmtypes.ModelSearchParams(params), buildRequestOptions(opts).headers)
}

// SearchModels searches multimodal model skills via GET /v1/models/skill/search.
//
// Supported params:
//   - Query maps to q
//   - Input maps to input
//   - Output maps to output
//   - Type maps to type
//   - Provider maps to provider
//   - Limit maps to limit
func (m *ModalService) SearchModels(ctx context.Context, params ModalModelSearchParams, opts ...RequestOption) (*ModalModelSearchResponse, error) {
	return m.ListModels(ctx, params, opts...)
}

func (m *ModalService) GetModelSkill(ctx context.Context, model string, opts ...RequestOption) (string, error) {
	return mmservice.GetModelSkill(m.client, ctx, model, buildRequestOptions(opts).headers)
}

func getTask(client *transport.Client, ctx context.Context, taskID string, headers http.Header) (*Task, error) {
	resp, err := mmservice.GetTask(client, ctx, taskID, headers)
	if err != nil {
		return nil, err
	}
	return newTaskFromResponse(client, resp), nil
}

func newTaskFromResponse(client *transport.Client, resp *mmtypes.TaskResponse) *Task {
	if resp == nil {
		return nil
	}
	return &Task{
		ID:       resp.ID,
		Status:   resp.Status,
		Model:    resp.Model,
		Progress: resp.Progress,
		Output:   resp.Output,
		Usage:    resp.Usage,
		Error:    resp.Error,
		client:   client,
	}
}
