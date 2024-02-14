package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://skillissue.com" {
		return false
	}
	return true
}

func slowWebsiteCheker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://skillissue.com",
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	want := map[string]bool{
		"https://google.com":     true,
		"https://yahoo.com":      true,
		"https://skillissue.com": false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = fmt.Sprintf("some url %d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteCheker, urls)
	}
}
