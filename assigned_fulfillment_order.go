package goshopify

import (
    "context"
    "fmt"
)

const (
    assignedFulfillmentOrderBasePath = "assigned_fulfillment_orders"
)

// AssignedFulfillmentService is an interface for interfacing with the assigned fulfillment orders
// of the Shopify API.
// https://shopify.dev/docs/api/admin-rest/2023-07/resources/assignedfulfillmentorder
type AssignedFulfillmentOrderService interface {
    Get(context.Context, interface{}) ([]AssignedFulfillmentOrder, error)
}

type AssignedFulfillmentOrder struct {
    Id                 uint64                              `json:"id,omitempty"`
    AssignedLocationId uint64                              `json:"assigned_location_id,omitempty"`
    Destination        AssignedFulfillmentOrderDestination `json:"destination,omitempty"`
    LineItems          []AssignedFulfillmentOrderLineItem  `json:"line_items,omitempty"`
    OrderId            uint64                              `json:"order_id,omitempty"`
    RequestStatus      string                              `json:"request_status,omitempty"`
    ShopId             uint64                              `json:"shop_id,omitempty"`
    Status             string                              `json:"status,omitempty"`
}

// AssignedFulfillmentOrderDestination represents a destination for a AssignedFulfillmentOrder
type AssignedFulfillmentOrderDestination struct {
    Id        uint64 `json:"id,omitempty"`
    Address1  string `json:"address1,omitempty"`
    Address2  string `json:"address2,omitempty"`
    City      string `json:"city,omitempty"`
    Company   string `json:"company,omitempty"`
    Country   string `json:"country,omitempty"`
    Email     string `json:"email,omitempty"`
    FirstName string `json:"first_name,omitempty"`
    LastName  string `json:"last_name,omitempty"`
    Phone     string `json:"phone,omitempty"`
    Province  string `json:"province,omitempty"`
    Zip       string `json:"zip,omitempty"`
}

// AssignedFulfillmentOrderLineItem represents a line item for a AssignedFulfillmentOrder
type AssignedFulfillmentOrderLineItem struct {
    Id                  uint64 `json:"id,omitempty"`
    ShopId              uint64 `json:"shop_id,omitempty"`
    FulfillmentOrderId  uint64 `json:"fulfillment_order_id,omitempty"`
    LineItemId          uint64 `json:"line_item_id,omitempty"`
    InventoryItemId     uint64 `json:"inventory_item_id,omitempty"`
    Quantity            uint64 `json:"quantity,omitempty"`
    FulfillableQuantity uint64 `json:"fulfillable_quantity,omitempty"`
    VariantId           uint64 `json:"variant_id,omitempty"`
}

// AssignedFulfillmentOrderResource represents the result from the assigned_fulfillment_order.json endpoint
type AssignedFulfillmentOrdersResource struct {
    AssignedFulfillmentOrders []AssignedFulfillmentOrder `json:"fulfillment_orders,omitempty"`
}

type AssignedFulfillmentOrderOptions struct {
    AssignmentStatus string `url:"assignment_status,omitempty"`
    LocationIds      string `url:"location_ids,omitempty"`
}

// AssignedFulfillmentOrderServiceOp handles communication with the AssignedFulfillmentOrderService
// related methods of the Shopify API
type AssignedFulfillmentOrderServiceOp struct {
    client ClientInterface
}

// Gets a list of all the fulfillment orders that are assigned to an app at the shop level
func (s *AssignedFulfillmentOrderServiceOp) Get(ctx context.Context, options interface{}) ([]AssignedFulfillmentOrder, error) {
    path := fmt.Sprintf("%s.json", assignedFulfillmentOrderBasePath)
    resource := new(AssignedFulfillmentOrdersResource)
    err := s.client.Get(ctx, path, resource, options)
    return resource.AssignedFulfillmentOrders, err
}
