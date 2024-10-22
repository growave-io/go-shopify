package goshopify

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ShopFullName returns the full shop name, including .myshopify.com
func ShopFullName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Trim(name, ".")
	if strings.Contains(name, "myshopify.com") {
		return name
	}
	return name + ".myshopify.com"
}

// ShopShortName returns the short shop name, excluding .myshopify.com
func ShopShortName(name string) string {
	// Convert to fullname and remove the myshopify part. Perhaps not the most
	// performant solution, but then we don't have to repeat all the trims here
	// :-)
	return strings.Replace(ShopFullName(name), ".myshopify.com", "", -1)
}

// ShopBaseUrl returns the Shop's base url.
func ShopBaseUrl(name string) string {
	name = ShopFullName(name)
	return fmt.Sprintf("https://%s", name)
}

// MetafieldPathPrefix returns the prefix for a metafield path
func MetafieldPathPrefix(resource string, resourceId uint64) string {
	prefix := "metafields"
	if resource != "" {
		prefix = fmt.Sprintf("%s/%d/metafields", resource, resourceId)
	}
	return prefix
}

// FulfillmentPathPrefix returns the prefix for a fulfillment path
func FulfillmentPathPrefix(resource string, resourceId uint64) string {
	prefix := "fulfillments"
	if resource != "" {
		prefix = fmt.Sprintf("%s/%d/fulfillments", resource, resourceId)
	}
	return prefix
}

type OnlyDate struct {
	time.Time
}

func (c *OnlyDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		*c = OnlyDate{time.Time{}}
		return nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return err
	}
	*c = OnlyDate{t}
	return nil
}

func (c *OnlyDate) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}

// EncodeValues it seems shopify accepts both the date with double-quotes and without them, so we just stick to the double-quotes for now.
func (c *OnlyDate) EncodeValues(key string, v *url.Values) error {
	v.Add(key, c.String())
	return nil
}

func (c *OnlyDate) String() string {
	return `"` + c.Format("2006-01-02") + `"`
}

func TimePtr(v time.Time) *time.Time {
	return &v
}
