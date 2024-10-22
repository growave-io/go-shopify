package goshopify

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAssetList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/themes/1/assets.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewStringResponder(
			200,
			`{"assets": [{"key":"assets\/1.liquid"},{"key":"assets\/2.liquid"}]}`,
		),
	)

	assets, err := client.Asset.List(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("Asset.List returned error: %v", err)
	}

	expected := []Asset{{Key: "assets/1.liquid"}, {Key: "assets/2.liquid"}}
	if !reflect.DeepEqual(assets, expected) {
		t.Errorf("Asset.List returned %+v, expected %+v", assets, expected)
	}
}

func TestAssetGet(t *testing.T) {
	setup()
	defer teardown()
	params := map[string]string{
		"asset[key]": "foo/bar.liquid",
		"theme_id":   "1",
	}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/themes/1/assets.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewStringResponder(
			200,
			`{"asset": {"key":"foo\/bar.liquid"}}`,
		),
	)

	asset, err := client.Asset.Get(context.Background(), 1, "foo/bar.liquid")
	if err != nil {
		t.Errorf("Asset.Get returned error: %v", err)
	}

	expected := &Asset{Key: "foo/bar.liquid"}
	if !reflect.DeepEqual(asset, expected) {
		t.Errorf("Asset.Get returned %+v, expected %+v", asset, expected)
	}
}

func TestAssetUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"PUT",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/themes/1/assets.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(
			200,
			loadFixture("asset.json"),
		),
	)

	asset := Asset{
		Key:   "templates/index.liquid",
		Value: "content",
	}

	returnedAsset, err := client.Asset.Update(context.Background(), 1, asset)
	if err != nil {
		t.Errorf("Asset.Update returned error: %v", err)
	}
	if returnedAsset == nil {
		t.Errorf("Asset.Update returned nil")
	}
}

func TestAssetDelete(t *testing.T) {
	setup()
	defer teardown()

	params := map[string]string{"asset[key]": "foo/bar.liquid"}
	httpmock.RegisterResponderWithQuery(
		"DELETE",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/themes/1/assets.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewStringResponder(200, "{}"),
	)

	err := client.Asset.Delete(context.Background(), 1, "foo/bar.liquid")
	if err != nil {
		t.Errorf("Asset.Delete returned error: %v", err)
	}
}
