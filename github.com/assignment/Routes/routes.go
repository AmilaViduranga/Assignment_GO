package Routes

import (
	"encoding/json"
	"github.com/assignment/Constants"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Running Muve Apllication" + Constants.APP_VERSION)
}
