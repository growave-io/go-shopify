package goshopify

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func MetafieldTests(t *testing.T, metafield Metafield) {
	// Check that Id is assigned to the returned metafield
	expectedInt := uint64(1)
	if metafield.Id != expectedInt {
		t.Errorf("Metafield.Id returned %+v, expected %+v", metafield.Id, expectedInt)
	}
}

func TestMetafieldList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"metafields": [{"id":1},{"id":2}]}`))

	metafields, err := client.Metafield.List(context.Background(), nil)
	if err != nil {
		t.Errorf("Metafield.List returned error: %v", err)
	}

	expected := []Metafield{{Id: 1}, {Id: 2}}
	if !reflect.DeepEqual(metafields, expected) {
		t.Errorf("Metafield.List returned %+v, expected %+v", metafields, expected)
	}
}

func TestMetafieldCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/count.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/count.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Metafield.Count(context.Background(), nil)
	if err != nil {
		t.Errorf("Metafield.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Metafield.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Metafield.Count(context.Background(), CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Metafield.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Metafield.Count returned %d, expected %d", cnt, expected)
	}
}

func TestMetafieldGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield, err := client.Metafield.Get(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("Metafield.Get returned error: %v", err)
	}

	createdAt := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC)
	expected := &Metafield{
		Id:                1,
		Key:               "app_key",
		Value:             "app_value",
		Type:              MetafieldTypeSingleLineTextField,
		Namespace:         "affiliates",
		Description:       "some amaaazing app's value",
		OwnerId:           1,
		CreatedAt:         &createdAt,
		UpdatedAt:         &updatedAt,
		OwnerResource:     "shop",
		AdminGraphqlApiId: "gid://shopify/Metafield/1",
	}
	if !reflect.DeepEqual(metafield, expected) {
		t.Errorf("Metafield.Get returned %+v, expected %+v", metafield, expected)
	}
}

func TestMetafieldCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		Namespace: "inventory",
		Key:       "warehouse",
		Value:     25,
		Type:      MetafieldTypeNumberInteger,
	}

	returnedMetafield, err := client.Metafield.Create(context.Background(), metafield)
	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestMetafieldUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		Id:    1,
		Value: "something new",
		Type:  MetafieldTypeSingleLineTextField,
	}

	returnedMetafield, err := client.Metafield.Update(context.Background(), metafield)
	if err != nil {
		t.Errorf("Metafield.Update returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestMetafieldDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Metafield.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Metafield.Delete returned error: %v", err)
	}
}
