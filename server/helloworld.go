package server

import (
	"fmt"
	"net/http"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ytx!")
}
