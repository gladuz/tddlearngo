package concurrency

import "sync"

type WebsiteChecker func(string) bool

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			results[u] = checker(u)
		}(url)
	}
	wg.Wait()
	return results
}
