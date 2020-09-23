package receivers

import (
	"github.com/vectorman1/saruman/consts"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, consts.RedirectUrl, http.StatusMovedPermanently)
}
