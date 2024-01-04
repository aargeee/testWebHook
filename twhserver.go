package testwebhook

import (
	"fmt"
	"net/http"
)

func TWHServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "xyz@something.com")
}
