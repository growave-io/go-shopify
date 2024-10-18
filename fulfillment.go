package goshopify

import (
    "context"
    "fmt"
    "time"
)

// FulfillmentService is an interface for interfacing with the fulfillment endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentService interface {
    List(context.Context, interface{}) ([]Fulfillment, error)
    Count(context.Context, interface{}) (int, error)
    Get(context.Context, uint64, interface{}) (*Fulfillment, error)
    Create(context.Context, Fulfillment) (*Fulfillment, error)
    Update(context.Context, Fulfillment) (*Fulfillment, error)
    Complete(context.Context, uint64) (*Fulfillment, error)
    Transition(context.Context, uint64) (*Fulfillment, error)
    Cancel(context.Context, uint64) (*Fulfillment, error)
}

// FulfillmentsService is an interface for other Shopify resources
// to interface with the fulfillment endpoints of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentsService interface {
    ListFulfillments(context.Context, uint64, interface{}) ([]Fulfillment, error)
    CountFulfillments(context.Context, uint64, interface{}) (int, error)
    GetFulfillment(context.Context, uint64, uint64, interface{}) (*Fulfillment, error)
    CreateFulfillment(context.Context, uint64, Fulfillment) (*Fulfillment, error)
    UpdateFulfillment(context.Context, uint64, Fulfillment) (*Fulfillment, error)
    CompleteFulfillment(context.Context, uint64, uint64) (*Fulfillment, error)
    TransitionFulfillment(context.Context, uint64, uint64) (*Fulfillment, error)
    CancelFulfillment(context.Context, uint64, uint64) (*Fulfillment, error)
}

// FulfillmentServiceOp handles communication with the fulfillment
// related methods of the Shopify API.
type FulfillmentServiceOp struct {
    client     ClientInterface
    resource   string
    resourceId uint64
}

// Fulfillment represents a Shopify fulfillment.
type Fulfillment struct {
    Id                          uint64                       `json:"id,omitempty"`
    OrderId                     uint64                       `json:"order_id,omitempty"`
    LocationId                  uint64                       `json:"location_id,omitempty"`
    Status                      string                       `json:"status,omitempty"`
    CreatedAt                   *time.Time                   `json:"created_at,omitempty"`
    Service                     string                       `json:"service,omitempty"`
    UpdatedAt                   *time.Time                   `json:"updated_at,omitempty"`
    TrackingCompany             string                       `json:"tracking_company,omitempty"`
    ShipmentStatus              string                       `json:"shipment_status,omitempty"`
    TrackingInfo                FulfillmentTrackingInfo      `json:"tracking_info,omitempty"`
    TrackingNumber              string                       `json:"tracking_number,omitempty"`
    TrackingNumbers             []string                     `json:"tracking_numbers,omitempty"`
    TrackingUrl                 string                       `json:"tracking_url,omitempty"`
    TrackingUrls                []string                     `json:"tracking_urls,omitempty"`
    Receipt                     Receipt                      `json:"receipt,omitempty"`
    LineItems                   []LineItem                   `json:"line_items,omitempty"`
    LineItemsByFulfillmentOrder []LineItemByFulfillmentOrder `json:"line_items_by_fulfillment_order,omitempty"`
    NotifyCustomer              bool                         `json:"notify_customer"`
}

// FulfillmentTrackingInfo represents the tracking information used to create a Fulfillment.
// https://shopify.dev/docs/api/admin-rest/2023-01/resources/fulfillment#post-fulfillments
type FulfillmentTrackingInfo struct {
    Company string `json:"company,omitempty"`
    Number  string `json:"number,omitempty"`
    Url     string `json:"url,omitempty"`
}

// LineItemByFulfillmentOrder represents the FulfillmentOrders (and optionally the items) used to create a Fulfillment.
// https://shopify.dev/docs/api/admin-rest/2023-01/resources/fulfillment#post-fulfillments
type LineItemByFulfillmentOrder struct {
    FulfillmentOrderId        uint64                                   `json:"fulfillment_order_id,omitempty"`
    FulfillmentOrderLineItems []LineItemByFulfillmentOrderItemQuantity `json:"fulfillment_order_line_items,omitempty"`
}

// LineItemByFulfillmentOrderItemQuantity represents the quantity to fulfill for one item.
type LineItemByFulfillmentOrderItemQuantity struct {
    Id       uint64 `json:"id"`
    Quantity uint64 `json:"quantity"`
}

// Receipt represents a Shopify receipt.
type Receipt struct {
    TestCase      bool   `json:"testcase,omitempty"`
    Authorization string `json:"authorization,omitempty"`
}

// FulfillmentResource represents the result from the fulfillments/X.json endpoint
type FulfillmentResource struct {
    Fulfillment *Fulfillment `json:"fulfillment"`
}

// FulfillmentsResource represents the result from the fullfilments.json endpoint
type FulfillmentsResource struct {
    Fulfillments []Fulfillment `json:"fulfillments"`
}

// List fulfillments
func (s *FulfillmentServiceOp) List(ctx context.Context, options interface{}) ([]Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s.json", prefix)
    resource := new(FulfillmentsResource)
    err := s.client.Get(ctx, path, resource, options)
    return resource.Fulfillments, err
}

// Count fulfillments
func (s *FulfillmentServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/count.json", prefix)
    return s.client.Count(ctx, path, options)
}

// Get individual fulfillment
func (s *FulfillmentServiceOp) Get(ctx context.Context, fulfillmentId uint64, options interface{}) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentId)
    resource := new(FulfillmentResource)
    err := s.client.Get(ctx, path, resource, options)
    return resource.Fulfillment, err
}

// Create a new fulfillment
func (s *FulfillmentServiceOp) Create(ctx context.Context, fulfillment Fulfillment) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s.json", prefix)
    wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
    resource := new(FulfillmentResource)
    err := s.client.Post(ctx, path, wrappedData, resource)
    return resource.Fulfillment, err
}

// Update an existing fulfillment
func (s *FulfillmentServiceOp) Update(ctx context.Context, fulfillment Fulfillment) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/%d.json", prefix, fulfillment.Id)
    wrappedData := FulfillmentResource{Fulfillment: &fulfillment}
    resource := new(FulfillmentResource)
    err := s.client.Put(ctx, path, wrappedData, resource)
    return resource.Fulfillment, err
}

// Complete an existing fulfillment
func (s *FulfillmentServiceOp) Complete(ctx context.Context, fulfillmentId uint64) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/%d/complete.json", prefix, fulfillmentId)
    resource := new(FulfillmentResource)
    err := s.client.Post(ctx, path, nil, resource)
    return resource.Fulfillment, err
}

// Transition an existing fulfillment
func (s *FulfillmentServiceOp) Transition(ctx context.Context, fulfillmentId uint64) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/%d/open.json", prefix, fulfillmentId)
    resource := new(FulfillmentResource)
    err := s.client.Post(ctx, path, nil, resource)
    return resource.Fulfillment, err
}

// Cancel an existing fulfillment
func (s *FulfillmentServiceOp) Cancel(ctx context.Context, fulfillmentId uint64) (*Fulfillment, error) {
    prefix := FulfillmentPathPrefix(s.resource, s.resourceId)
    path := fmt.Sprintf("%s/%d/cancel.json", prefix, fulfillmentId)
    resource := new(FulfillmentResource)
    err := s.client.Post(ctx, path, nil, resource)
    return resource.Fulfillment, err
}

func NewFulfillmentService(client ClientInterface) FulfillmentService {
    return &FulfillmentServiceOp{client: client}
}
