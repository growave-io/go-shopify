package goshopify

import (
    "context"
    "fmt"
    "time"
)

const (
    pagesBasePath     = "pages"
    pagesResourceName = "pages"
)

// PagesPageService is an interface for interacting with the pages
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/page
type PageService interface {
    List(context.Context, interface{}) ([]Page, error)
    Count(context.Context, interface{}) (int, error)
    Get(context.Context, uint64, interface{}) (*Page, error)
    Create(context.Context, Page) (*Page, error)
    Update(context.Context, Page) (*Page, error)
    Delete(context.Context, uint64) error

    // MetafieldsService used for Pages resource to communicate with Metafields
    // resource
    MetafieldsService
}

// PageServiceOp handles communication with the page related methods of the
// Shopify API.
type PageServiceOp struct {
    client ClientInterface
}

// Page represents a Shopify page.
type Page struct {
    Id             uint64     `json:"id,omitempty"`
    Author         string     `json:"author,omitempty"`
    Handle         string     `json:"handle,omitempty"`
    Title          string     `json:"title,omitempty"`
    CreatedAt      *time.Time `json:"created_at,omitempty"`
    UpdatedAt      *time.Time `json:"updated_at,omitempty"`
    BodyHTML       string     `json:"body_html,omitempty"`
    TemplateSuffix string     `json:"template_suffix,omitempty"`
    PublishedAt    *time.Time `json:"published_at,omitempty"`
    // Published can be set when creating a new page.
    // It's not returned in the response
    Published  *bool       `json:"published,omitempty"`
    ShopId     uint64      `json:"shop_id,omitempty"`
    Metafields []Metafield `json:"metafields,omitempty"`
}

// PageResource represents the result from the pages/X.json endpoint
type PageResource struct {
    Page *Page `json:"page"`
}

// PagesResource represents the result from the pages.json endpoint
type PagesResource struct {
    Pages []Page `json:"pages"`
}

// List pages
func (s *PageServiceOp) List(ctx context.Context, options interface{}) ([]Page, error) {
    path := fmt.Sprintf("%s.json", pagesBasePath)
    resource := new(PagesResource)
    err := s.client.Get(ctx, path, resource, options)
    return resource.Pages, err
}

// Count pages
func (s *PageServiceOp) Count(ctx context.Context, options interface{}) (int, error) {
    path := fmt.Sprintf("%s/count.json", pagesBasePath)
    return s.client.Count(ctx, path, options)
}

// Get individual page
func (s *PageServiceOp) Get(ctx context.Context, pageId uint64, options interface{}) (*Page, error) {
    path := fmt.Sprintf("%s/%d.json", pagesBasePath, pageId)
    resource := new(PageResource)
    err := s.client.Get(ctx, path, resource, options)
    return resource.Page, err
}

// Create a new page
func (s *PageServiceOp) Create(ctx context.Context, page Page) (*Page, error) {
    path := fmt.Sprintf("%s.json", pagesBasePath)
    wrappedData := PageResource{Page: &page}
    resource := new(PageResource)
    err := s.client.Post(ctx, path, wrappedData, resource)
    return resource.Page, err
}

// Update an existing page
func (s *PageServiceOp) Update(ctx context.Context, page Page) (*Page, error) {
    path := fmt.Sprintf("%s/%d.json", pagesBasePath, page.Id)
    wrappedData := PageResource{Page: &page}
    resource := new(PageResource)
    err := s.client.Put(ctx, path, wrappedData, resource)
    return resource.Page, err
}

// Delete an existing page.
func (s *PageServiceOp) Delete(ctx context.Context, pageId uint64) error {
    return s.client.Delete(ctx, fmt.Sprintf("%s/%d.json", pagesBasePath, pageId))
}

// List metafields for a page
func (s *PageServiceOp) ListMetafields(ctx context.Context, pageId uint64, options interface{}) ([]Metafield, error) {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.List(ctx, options)
}

// Count metafields for a page
func (s *PageServiceOp) CountMetafields(ctx context.Context, pageId uint64, options interface{}) (int, error) {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.Count(ctx, options)
}

// Get individual metafield for a page
func (s *PageServiceOp) GetMetafield(ctx context.Context, pageId uint64, metafieldId uint64, options interface{}) (*Metafield, error) {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.Get(ctx, metafieldId, options)
}

// Create a new metafield for a page
func (s *PageServiceOp) CreateMetafield(ctx context.Context, pageId uint64, metafield Metafield) (*Metafield, error) {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.Create(ctx, metafield)
}

// Update an existing metafield for a page
func (s *PageServiceOp) UpdateMetafield(ctx context.Context, pageId uint64, metafield Metafield) (*Metafield, error) {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.Update(ctx, metafield)
}

// Delete an existing metafield for a page
func (s *PageServiceOp) DeleteMetafield(ctx context.Context, pageId uint64, metafieldId uint64) error {
    metafieldService := &MetafieldServiceOp{client: s.client, resource: pagesResourceName, resourceId: pageId}
    return metafieldService.Delete(ctx, metafieldId)
}

func NewPageService(client ClientInterface) PageService {
    return &PageServiceOp{client}
}
