package imap

import "testing"

func TestNewDialer(t *testing.T) {
	tests := []struct {
		name      string
		proxyURL  string
		wantError bool
	}{
		{name: "direct connection", proxyURL: ""},
		{name: "SOCKS5", proxyURL: "socks5://127.0.0.1:1080"},
		{name: "SOCKS5 with remote DNS", proxyURL: "socks5h://proxy.example.com:1080"},
		{name: "SOCKS5 with authentication", proxyURL: "socks5://user:pass@127.0.0.1:1080"},
		{name: "default SOCKS5 port", proxyURL: "socks5://127.0.0.1"},
		{name: "unsupported scheme", proxyURL: "http://127.0.0.1:8080", wantError: true},
		{name: "missing host", proxyURL: "socks5://", wantError: true},
		{name: "invalid URL", proxyURL: "://bad", wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialer, err := NewDialer(tt.proxyURL)
			if (err != nil) != tt.wantError {
				t.Fatalf("NewDialer(%q) error = %v, wantError %v", tt.proxyURL, err, tt.wantError)
			}
			if !tt.wantError && dialer == nil {
				t.Fatal("NewDialer returned a nil dialer")
			}
		})
	}
}
