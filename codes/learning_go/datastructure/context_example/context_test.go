package context_example

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

//
// context in HTTP request
// see also: net/http
//

// Logic with context.
func logic(ctx context.Context, info string) (string, error) {
	// ...
	return "", nil
}

func MiddlewareExample(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// parse request, invoke logic
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data := r.FormValue("data")
		result, err := logic(ctx, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(result))

		// invoke handler with context
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

// call another HTTP service
type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}

	// do invoke
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	// process response
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	id, err := processResponse(resp.Body)
	return id, err
}

func processResponse(respBody io.ReadCloser) (string, error) {
	data, err := io.ReadAll(respBody)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//
// context value
//
