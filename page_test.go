package goshopify

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func pageTests(t *testing.T, page Page) {
	// Check that Id is assigned to the returned page
	expectedInt := uint64(1)
	if page.Id != expectedInt {
		t.Errorf("Page.Id returned %+v, expected %+v", page.Id, expectedInt)
	}
}

func TestPageList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"pages": [{"id":1,"published_at": "2008-07-15T20:00:00Z"},{"id":2,"published_at": "2008-07-15T21:00:00Z"}]}`))

	pages, err := client.Page.List(context.Background(), nil)
	if err != nil {
		t.Errorf("Page.List returned error: %v", err)
	}

	expected := []Page{
		{Id: 1, PublishedAt: TimePtr(time.Date(2008, 7, 15, 20, 0, 0, 0, time.UTC))},
		{Id: 2, PublishedAt: TimePtr(time.Date(2008, 7, 15, 21, 0, 0, 0, time.UTC))},
	}
	if !reflect.DeepEqual(pages, expected) {
		t.Errorf("Page.List returned %+v, expected %+v", pages, expected)
	}
}

func TestPageCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/count.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/count.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.Count(context.Background(), nil)
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.Count(context.Background(), CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}
}

func TestPageGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"page": {"id":1,"published_at": "2008-07-15T20:00:00Z"}}`))

	page, err := client.Page.Get(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("Page.Get returned error: %v", err)
	}

	expected := &Page{Id: 1, PublishedAt: TimePtr(time.Date(2008, 7, 15, 20, 0, 0, 0, time.UTC))}
	if !reflect.DeepEqual(page, expected) {
		t.Errorf("Page.Get returned %+v, expected %+v", page, expected)
	}
}

func TestPageCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		Title:    "404",
		BodyHTML: "<strong>NOT FOUND!<\\/strong>",
	}

	returnedPage, err := client.Page.Create(context.Background(), page)
	if err != nil {
		t.Errorf("Page.Create returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		Id: 1,
	}

	returnedPage, err := client.Page.Update(context.Background(), page)
	if err != nil {
		t.Errorf("Page.Update returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Page.Delete returned error: %v", err)
	}
}

func TestPageListMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"metafields": [{"id":1},{"id":2}]}`))

	metafields, err := client.Page.ListMetafields(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("Page.ListMetafields() returned error: %v", err)
	}

	expected := []Metafield{{Id: 1}, {Id: 2}}
	if !reflect.DeepEqual(metafields, expected) {
		t.Errorf("Page.ListMetafields() returned %+v, expected %+v", metafields, expected)
	}
}

func TestPageCountMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields/count.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields/count.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.CountMetafields(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("Page.CountMetafields() returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.CountMetafields() returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.CountMetafields(context.Background(), 1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.CountMetafields() returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.CountMetafields() returned %d, expected %d", cnt, expected)
	}
}

func TestPageGetMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields/2.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"metafield": {"id":2}}`))

	metafield, err := client.Page.GetMetafield(context.Background(), 1, 2, nil)
	if err != nil {
		t.Errorf("Page.GetMetafield() returned error: %v", err)
	}

	expected := &Metafield{Id: 2}
	if !reflect.DeepEqual(metafield, expected) {
		t.Errorf("Page.GetMetafield() returned %+v, expected %+v", metafield, expected)
	}
}

func TestPageCreateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		Key:       "app_key",
		Value:     "app_value",
		Type:      MetafieldTypeSingleLineTextField,
		Namespace: "affiliates",
	}

	returnedMetafield, err := client.Page.CreateMetafield(context.Background(), 1, metafield)
	if err != nil {
		t.Errorf("Page.CreateMetafield() returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestPageUpdateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields/2.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		Id:        2,
		Key:       "app_key",
		Value:     "app_value",
		Type:      MetafieldTypeSingleLineTextField,
		Namespace: "affiliates",
	}

	returnedMetafield, err := client.Page.UpdateMetafield(context.Background(), 1, metafield)
	if err != nil {
		t.Errorf("Page.UpdateMetafield() returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestPageDeleteMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metafields/2.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.DeleteMetafield(context.Background(), 1, 2)
	if err != nil {
		t.Errorf("Page.DeleteMetafield() returned error: %v", err)
	}
}
