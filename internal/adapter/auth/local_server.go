package auth

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// LocalServer handles the temporary HTTP server for OAuth callbacks.
type LocalServer struct {
	server   *http.Server
	CodeChan chan string
	ErrChan  chan error
	Port     int
}

// NewLocalServer creates a new LocalServer.
// It attempts to listen on port 8888 by default.
func NewLocalServer() (*LocalServer, error) {
	const defaultPort = 8888

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", defaultPort))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on localhost:%d: %w", defaultPort, err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	codeChan := make(chan string)
	errChan := make(chan error)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Basic HTML response
		w.Header().Set("Content-Type", "text/html")

		code := r.URL.Query().Get("code")
		if code != "" {
			fmt.Fprintf(w, "<h1>Login Successful</h1><p>You can close this window now.</p>")
			codeChan <- code
		} else {
			// Check for error query param if relevant, or just timeout
			fmt.Fprintf(w, "<h1>Login Failed</h1><p>No code received.</p>")
			errChan <- fmt.Errorf("no code received in callback")
		}
	})

	srv := &http.Server{
		Handler: mux,
	}

	// Start server in goroutine
	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	return &LocalServer{
		server:   srv,
		CodeChan: codeChan,
		ErrChan:  errChan,
		Port:     port,
	}, nil
}

// Shutdown stops the local server.
func (s *LocalServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// WaitForCode blocks until a code is received or the context times out.
func (s *LocalServer) WaitForCode(ctx context.Context) (string, error) {
	select {
	case code := <-s.CodeChan:
		return code, nil
	case err := <-s.ErrChan:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
