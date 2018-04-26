package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	/*
		The make built-in function allocates and initializes an object
		of type slice, map, or chan (only).
	*/
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/*
			A Goroutine is an non-blocking operation, with sintax as seen just below.
			It's a good practice to use anonymous functions to call a goroutine.
		*/
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
