package a

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func ExampleParallel() {

	var mu sync.Mutex

	a := Parallel(
		AddVal(&mu, "A"), AddVal(&mu, "B"), AddVal(&mu, "C"),
	)

	job := []string{}
	a.Run(&job)

	sort.Strings(job)

	fmt.Println("Hard to test parallel code:", strings.Join(job, "") == "ABC")

	// Output:
	// Hard to test parallel code: true
}

func ExampleParallel_error() {

	var mu sync.Mutex

	a := Parallel(
		AddVal(&mu, "A"),
		AddValErr("B"),
		AddVal(&mu, "C"),
		AddValErr("D"),
	)

	job := []string{}
	err := a.Run(&job)

	sort.Strings(job)

	fmt.Println("Error:", strings.ContainsAny(err.Error(), "BD"))
	fmt.Println("Job:", strings.Join(job, "") == "AC")

	// Output:
	// Error: true
	// Job: true
}

func ExampleSemaphore() {

	var mu sync.Mutex

	a := Semaphore(2, AddVal(&mu, "A"), AddVal(&mu, "B"), AddVal(&mu, "C"))

	job := []string{}
	a.Run(&job)

	sort.Strings(job)

	fmt.Println("Job:", job)

	// Output:
	// Job: [A B C]
}

func ExampleSemaphore_n1_behaves_like_All() {

	var mu sync.Mutex

	// N = 1 makes Semaphore() behave like All().
	a := Semaphore(1, AddVal(&mu, "A"), AddVal(&mu, "B"), AddVal(&mu, "C"))

	job := []string{}
	a.Run(&job)

	fmt.Println("Job:", job)

	// Output:
	// Job: [A B C]
}

func AddVal(mu *sync.Mutex, v string) Func[*[]string] {
	return func(a *[]string) error {
		mu.Lock()
		*a = append(*a, v)
		mu.Unlock()
		return nil
	}
}

func AddValErr(v string) Func[*[]string] {
	return func(a *[]string) error {
		return fmt.Errorf(v)
	}
}
