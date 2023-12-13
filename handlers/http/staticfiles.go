package handlers_http

import(
  "net/http"
)

func StylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "styles.css")
}

func MagicHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "magic.js")
}

func DialogJSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dialog.js")
}

func DialogTSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dialog.ts")
}



