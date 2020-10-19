package concurrency

// WebsiteChecker is the type that stores CheckWebsite function
type WebsiteChecker func(string) bool

type result struct {
	url    string
	status bool
}

// CheckWebsites checks status of websites
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.status
	}

	return results
}
