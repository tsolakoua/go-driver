package fixtures

import "fmt"

var (
	_    = foo()
	_, _ = foo(a, b)
	_    = a.foo()

	_ = func() int {
		return 1
	}()

	_, _ = fmt.Println()
)
