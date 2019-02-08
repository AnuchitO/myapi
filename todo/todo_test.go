package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodosHandler(t *testing.T) {
	t.Run("GET /todos", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/todos", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(todosHandler)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if http.StatusOK != status {
			t.Errorf("should response status code %v but got %v \n", http.StatusOK, status)
		}

		expected := `hello GET todos`
		resp := rr.Body.String()
		if expected != resp {
			t.Errorf("should response body %q but got %q \n", expected, resp)
		}
	})

	t.Run("should not allow DELETE /todos", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/todos", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(todosHandler)

		handler.ServeHTTP(rr, req)

		status := rr.Code
		if http.StatusNotFound != status {
			t.Errorf("should response status code %v but got %v \n", http.StatusOK, status)
		}
	})
}
