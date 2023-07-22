package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Authorization", "ApiKey YWxhZGRpbjpvcGVuc2VzYW1l")

	//fmt.Println(req.Header.Get("Authorization"))

	apiKey, err := GetAPIKey(req.Header)

	if err != nil {
		t.Errorf("Not expecting any error %v ", err)
	}

	expected := "YWxhZGRpbjpvcGVuc2VzYW1l"
	if apiKey != expected {
		t.Errorf("Actual %s, expected %s", "", apiKey)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Add("Authorization", "ApiKey")
	apiKey2, err2 := GetAPIKey(req2.Header)
	if apiKey2 != "" {
		t.Errorf("Expecting empty API Key")
	}
	if err2.Error() != "malformed authorization header" {
		t.Errorf("Expecting error as %s but received %s", "malformed authorization header", err2)
	}

	req3 := httptest.NewRequest(http.MethodGet, "/", nil)
	apiKey3, err3 := GetAPIKey(req3.Header)
	if apiKey3 != "" {
		t.Errorf("Expecting empty API Key")
	}
	if err3.Error() != "no authorization header included" {
		t.Errorf("Expecting error as %s but received %s", "no authorization header included", err3)
	}

}
