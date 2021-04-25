// Author: yangzq80@gmail.com
// Date: 2021-04-07
//
package server

import (
	"net/http"
	"strings"
	"testing"
)

func BenchmarkHttpDo(b *testing.B) {
	body := `{"url":"http://localhost:9999/api/kv/chaos/designer","method":"post","data":{"a":"1"}}`
	b.ResetTimer()
	client := http.DefaultClient
	for i := 0; i < b.N; i++ {
		client.Post("http://localhost:9999/api/http/do", "application/json", strings.NewReader(body))
	}
}

func BenchmarkCreate(b *testing.B) {
	body := `{"url":"http://localhost:9999/api/kv/chaos/designer","method":"post","data":{"a":"1"}}`
	b.ResetTimer()
	client := http.DefaultClient
	for i := 0; i < b.N; i++ {
		client.Post("http://localhost:9999/api/http/do", "application/json", strings.NewReader(body))
	}
}
