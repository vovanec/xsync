# xsync

`Once` is an object that behaves exactly like sync.Once with only exception 
that it does not save state on panic or function failure.

  * See [sync.Once](http://golang.org/pkg/sync/#Once)

## Example

Here's an example how `xsync.Once` could be used.

```go
package main

import (
	"errors"
	"fmt"

	"github.com/vovanec/xsync"
)

func main() {

	var (
		once    xsync.Once
		counter = 0
	)

	// This function fails with error, so Once's state is not going to be changed.
	once.Do(func() error {
		counter += 1
		return errors.New("error")
	})
	fmt.Printf("counter value: %d\n", counter) // prints 1

	// This function doesn't fail, so Once's going to change its state exactly one time.
	for i := 0; i < 10; i++ {
		once.Do(func() error {
			counter += 1
			return nil
		})
	}

	fmt.Printf("counter value: %d\n", counter) // prints 2

}
```
