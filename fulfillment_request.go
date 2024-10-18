package goshopify

import (
    "context"
    "fmt"
)

const (
    fulfillmentRequestBasePath = "fulfillment_orders"
)

// FulfillmentRequestService is an interface for interfacing with the fulfillment request endpoints of the Shopify API.
// https://shopify.dev/docs/api/admin-rest/2023-10/resources/fulfillmentrequest
type FulfillmentRequestService interface {
    Send(context.Context, uint64, FulfillmentRequest) (*FulfillmentOrder, error)
    Accept(context.Context, uint64, FulfillmentRequest) (*FulfillmentOrder, error)
    Reject(context.Context, uint64, FulfillmentRequest) (*FulfillmentOrder, error)
}

type FulfillmentRequest struct {
    Message                   string                       `json:"message,omitempty"`
    FulfillmentOrderLineItems []FulfillmentOrderLineItem   `json:"fulfillment_order_line_items,omitempty"`
    Reason                    string                       `json:"reason,omitempty"`
    LineItems                 []FulfillmentRequestLineItem `json:"line_items,omitempty"`
}

type FulfillmentRequestOrderLineItem struct {
    Id       uint64 `json:"id"`
    Quantity uint64 `json:"quantity"`
}

type FulfillmentRequestLineItem struct {
    FulfillmentOrderLineItemId uint64 `json:"fulfillment_order_line_item_id,omitempty"`
    Message                    string `json:"message,omitempty"`
}

type FulfillmentRequestResource struct {
    FulfillmentOrder         *FulfillmentOrder  `json:"fulfillment_order,omitempty"`
    FulfillmentRequest       FulfillmentRequest `json:"fulfillment_request,omitempty"`
    OriginalFulfillmentOrder *FulfillmentOrder  `json:"original_fulfillment_order,omitempty"`
}

// FulfillmentRequestServiceOp handles communication with the fulfillment request related methods of the Shopify API.
type FulfillmentRequestServiceOp struct {
    client ClientInterface
}

// Send sends a fulfillment request to the fulfillment service of a fulfillment order.
func (s *FulfillmentRequestServiceOp) Send(ctx context.Context, fulfillmentOrderId uint64, request FulfillmentRequest) (*FulfillmentOrder, error) {
    path := fmt.Sprintf("%s/%d/fulfillment_request.json", fulfillmentRequestBasePath, fulfillmentOrderId)
    wrappedData := FulfillmentRequestResource{FulfillmentRequest: request}
    resource := new(FulfillmentRequestResource)
    err := s.client.Post(ctx, path, wrappedData, resource)
    return resource.OriginalFulfillmentOrder, err
}

// Accept accepts a fulfillment request sent to a fulfillment service for a fulfillment order.
func (s *FulfillmentRequestServiceOp) Accept(ctx context.Context, fulfillmentOrderId uint64, request FulfillmentRequest) (*FulfillmentOrder, error) {
    path := fmt.Sprintf("%s/%d/fulfillment_request/accept.json", fulfillmentRequestBasePath, fulfillmentOrderId)
    wrappedData := map[string]interface{}{"fulfillment_request": request}
    resource := new(FulfillmentRequestResource)
    err := s.client.Post(ctx, path, wrappedData, resource)
    return resource.FulfillmentOrder, err
}

// Reject rejects a fulfillment request sent to a fulfillment service for a fulfillment order.
func (s *FulfillmentRequestServiceOp) Reject(ctx context.Context, fulfillmentOrderId uint64, request FulfillmentRequest) (*FulfillmentOrder, error) {
    path := fmt.Sprintf("%s/%d/fulfillment_request/reject.json", fulfillmentRequestBasePath, fulfillmentOrderId)
    wrappedData := map[string]interface{}{"fulfillment_request": request}
    resource := new(FulfillmentRequestResource)
    err := s.client.Post(ctx, path, wrappedData, resource)
    return resource.FulfillmentOrder, err
}
