package testwebhook_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	testwebhook "github.com/Rahul-NITD/testWebHook"
)

func TestWHServer(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/newtunnel", nil)
	res := httptest.NewRecorder()
	testwebhook.TWHServer(res, req)
	got := res.Body.String()
	want := "xyz@something.com"
	if got != want {
		t.Errorf("got %s != %s", got, want)
	}
}
