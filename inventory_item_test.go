package goshopify

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func inventoryItemTests(t *testing.T, item *InventoryItem) {
	if item == nil {
		t.Errorf("InventoryItem is nil")
		return
	}

	expectedInt := uint64(808950810)
	if item.Id != expectedInt {
		t.Errorf("InventoryItem.Id returned %+v, expected %+v", item.Id, expectedInt)
	}

	expectedSKU := "new sku"
	if item.SKU != expectedSKU {
		t.Errorf("InventoryItem.SKU sku is %+v, expected %+v", item.SKU, expectedSKU)
	}

	if item.Cost == nil {
		t.Errorf("InventoryItem.Cost is nil")
		return
	}

	expectedCost := 25.00
	costFloat, _ := item.Cost.Float64()
	if costFloat != expectedCost {
		t.Errorf("InventoryItem.Cost (float) is %+v, expected %+v", costFloat, expectedCost)
	}

	expectedOrigin := "US"
	if *item.CountryCodeOfOrigin != expectedOrigin {
		t.Errorf("InventoryItem.CountryCodeOfOrigin returned %+v, expected %+v", item.CountryCodeOfOrigin, expectedOrigin)
	}

	// strings.Join is used to compare slices since package's go.mod is set to 1.13
	// which predates the experimental slices package that has a Compare() func.
	expectedCountryHSCodes := strings.Join([]string{"8471.70.40.35", "8471.70.50.35"}, ",")
	if strings.Join(item.CountryHarmonizedSystemCodes, ",") != expectedCountryHSCodes {
		t.Errorf("InventoryItem.CountryHarmonizedSystemCodes returned %+v, expected %+v", item.CountryHarmonizedSystemCodes, expectedCountryHSCodes)
	}

	expectedHSCode := "8471.70.40.35"
	if *item.HarmonizedSystemCode != expectedHSCode {
		t.Errorf("InventoryItem.HarmonizedSystemCode returned %+v, expected %+v", item.CountryHarmonizedSystemCodes, expectedHSCode)
	}

	expectedProvince := "ON"
	if *item.ProvinceCodeOfOrigin != expectedProvince {
		t.Errorf("InventoryItem.ProvinceCodeOfOrigin returned %+v, expected %+v", item.ProvinceCodeOfOrigin, expectedHSCode)
	}
}

func inventoryItemsTests(t *testing.T, items []InventoryItem) {
	expectedLen := 3
	if len(items) != expectedLen {
		t.Errorf("InventoryItems list lenth is %+v, expected %+v", len(items), expectedLen)
	}
}

func TestInventoryItemsList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/inventory_items.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("inventory_items.json")))

	items, err := client.InventoryItem.List(context.Background(), nil)
	if err != nil {
		t.Errorf("InventoryItems.List returned error: %v", err)
	}

	inventoryItemsTests(t, items)
}

func TestInventoryItemsListWithIds(t *testing.T) {
	setup()
	defer teardown()

	params := map[string]string{
		"ids": "1,2",
	}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/inventory_items.json", client.ApiClient.GetPathPrefix()),
		params,
		httpmock.NewBytesResponder(200, loadFixture("inventory_items.json")),
	)

	options := ListOptions{
		Ids: []uint64{1, 2},
	}

	items, err := client.InventoryItem.List(context.Background(), options)
	if err != nil {
		t.Errorf("InventoryItems.List returned error: %v", err)
	}

	inventoryItemsTests(t, items)
}

func TestInventoryItemGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/inventory_items/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("inventory_item.json")))

	item, err := client.InventoryItem.Get(context.Background(), 1, nil)
	if err != nil {
		t.Errorf("InventoryItem.Get returned error: %v", err)
	}

	inventoryItemTests(t, item)
}

func TestInventoryItemUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/inventory_items/1.json", client.ApiClient.GetPathPrefix()),
		httpmock.NewBytesResponder(200, loadFixture("inventory_item.json")))

	item := InventoryItem{
		Id: 1,
	}

	updatedItem, err := client.InventoryItem.Update(context.Background(), item)
	if err != nil {
		t.Errorf("InentoryItem.Update returned error: %v", err)
	}

	inventoryItemTests(t, updatedItem)
}
