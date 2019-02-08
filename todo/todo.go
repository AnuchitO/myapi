package todo

import "net/http"

func todosHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte(`hello GET todos`))
}
