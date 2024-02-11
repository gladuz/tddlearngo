package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(checker WebsiteChecker, urls []string) map[string] bool{
	results := make(map[string]bool)

	for _, url := range urls{
		go func(u string){
			results[u] =  checker(u)
		}(url)
	}
	return results
}