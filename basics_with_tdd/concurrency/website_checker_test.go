package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://bulbapedia.bulbagarden.net/",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":                  true,
		"http://bulbapedia.bulbagarden.net/": true,
		"waat://furhurterwe.geds":            false,
	}

	/*
		DeepEqual reports whether x and y are “deeply equal,”. Two values of identical type are deeply equal
		if one of the following cases applies. Values of distinct types are never deeply equal.

		1. Array values are deeply equal when their corresponding elements are deeply equal.
		2. Struct values are deeply equal if their corresponding fields, both exported and unexported, are deeply equal.
		3. Func values are deeply equal if both are nil; otherwise they are not deeply equal.
		4. Interface values are deeply equal if they hold deeply equal concrete values.
	*/
	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}
