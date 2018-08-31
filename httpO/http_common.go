package httpO

import (
	"fmt"
	"net/http"
)

func weHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am ok.")
}
