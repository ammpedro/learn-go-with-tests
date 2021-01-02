package concurrency

// WebsiteChecker accepts a url and returns a boolean
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLS and returns a map
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// use channels to solve data race issue
	// makes a channel of result
	resultChannel := make(chan result)

	for _, url := range urls {
		// an anonymous goroutine
		// pass the url so that the value of the url is fixed when wc is called
		go func(u string) {
			//send result struct into the resultChannel for later
			resultChannel <- result{u, wc(u)}
		}(url) //execute right away
	}

	for i := 0; i < len(urls); i++ {
		//receive value from the channel to a variable one at a time
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
