package goshopify

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestWithVersion(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithVersion(testApiVersion))
	expected := fmt.Sprintf("admin/api/%s", testApiVersion)
	if c.ApiClient.GetPathPrefix() != expected {
		t.Errorf("WithVersion client.ApiClient.GetPathPrefix() = %s, expected %s", c.ApiClient.GetPathPrefix(), expected)
	}
}

func TestWithVersionNoVersion(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithVersion(""))
	expected := "admin"
	if c.ApiClient.GetPathPrefix() != expected {
		t.Errorf("WithVersion client.ApiClient.GetPathPrefix() = %s, expected %s", c.ApiClient.GetPathPrefix(), expected)
	}
}

func TestWithoutVersionInInitiation(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd")
	expected := "admin"
	if c.ApiClient.GetPathPrefix() != expected {
		t.Errorf("WithVersion client.ApiClient.GetPathPrefix() = %s, expected %s", c.ApiClient.GetPathPrefix(), expected)
	}
}

func TestWithVersionInvalidVersion(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithVersion("9999-99b"))
	expected := "admin"
	if c.ApiClient.GetPathPrefix() != expected {
		t.Errorf("WithVersion client.ApiClient.GetPathPrefix() = %s, expected %s", c.ApiClient.GetPathPrefix(), expected)
	}
}

func TestWithRetry(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithRetry(5))
	expected := 5
	if c.ApiClient.GetRetries() != expected {
		t.Errorf("WithRetry client.retries = %d, expected %d", c.ApiClient.GetRetries(), expected)
	}
}

func TestWithLogger(t *testing.T) {
	logger := &LeveledLogger{Level: LevelDebug}
	c := MustNewClient(app, "fooshop", "abcd", WithLogger(logger))

	if c.ApiClient.GetLogger() != logger {
		t.Errorf("WithLogger expected logs to match %v != %v", c.ApiClient.GetLogger(), logger)
	}
}

func TestWithUnstableVersion(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithVersion(UnstableApiVersion))
	expected := fmt.Sprintf("admin/api/%s", UnstableApiVersion)
	if c.ApiClient.GetPathPrefix() != expected {
		t.Errorf("WithVersion client.ApiClient.GetPathPrefix() = %s, expected %s", c.ApiClient.GetPathPrefix(), expected)
	}
}

func TestWithHTTPClient(t *testing.T) {
	c := MustNewClient(app, "fooshop", "abcd", WithHTTPClient(&http.Client{Timeout: 30 * time.Second}))
	expected := 30 * time.Second

	if c.ApiClient.GetHttpClient().Timeout.String() != expected.String() {
		t.Errorf("WithVersion client.Client = %s, expected %s", c.ApiClient.GetHttpClient().Timeout, expected)
	}
}
