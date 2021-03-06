package ws

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestGetConnectionCount(t *testing.T) {
	hub := NewHub()
	uServer := NewUserCountServer(hub)
	server := httptest.NewServer(uServer)
	t.Cleanup(server.Close)
	wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/usercount"

	t.Run("testing websocket connection", func(t *testing.T) {
		ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		ws.Close()
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsUrl, err)
		}
	})
}
