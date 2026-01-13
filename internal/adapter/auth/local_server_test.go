package auth_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"meegle-cli/internal/adapter/auth"
)

func TestLocalServer(t *testing.T) {
	server, err := auth.NewLocalServer()
	if err != nil {
		t.Fatalf("Failed to create local server: %v", err)
	}
	defer server.Shutdown(context.Background())

	// Test obtaining the port
	if server.Port == 0 {
		t.Error("Server port should not be 0")
	}

	// Simulate callback in a goroutine
	go func() {
		// Wait a bit for server to be ready (though Listen is already done)
		time.Sleep(50 * time.Millisecond)
		url := fmt.Sprintf("http://localhost:%d/?code=test_code", server.Port)
		resp, err := http.Get(url)
		if err != nil {
			// t.Error in goroutine is tricky, but we can log
			fmt.Printf("HTTP Get failed: %v\n", err)
			return
		}
		defer resp.Body.Close()
	}()

	// Wait for code
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	code, err := server.WaitForCode(ctx)
	if err != nil {
		t.Fatalf("WaitForCode failed: %v", err)
	}
	if code != "test_code" {
		t.Errorf("Expected 'test_code', got '%s'", code)
	}
}

func TestLocalServer_Timeout(t *testing.T) {
	server, err := auth.NewLocalServer()
	if err != nil {
		t.Fatalf("Failed to create local server: %v", err)
	}
	defer server.Shutdown(context.Background())

	// Wait for code with short timeout, no request sent
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err = server.WaitForCode(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("Expected DeadlineExceeded error, got %v", err)
	}
}
