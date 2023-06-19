# Countable Wait Group

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

The Countable Wait Group is a Go package providing a wait group with a counter. It's a simple wrapper around the `sync.WaitGroup` type from the standard library, which keeps track of the number of goroutines it's waiting for. In addition, it provides an ability to prevent new tasks from being added.

## Usage

In your project you'll need to add the package.

```shell
go get -u github.com/andreimerlescu/go-countable-waitgroup
```

Here's an example program that uses this package to show you what its capable of. Of course, you should be reviewing the `countable-waitgroup_test.go` to see the implementation details that work, but this will start you on the right path.

```go
import "github.com/andreimerlescu/go-countable-waitgroup"

func main() {
    var wg countable_waitgroup.CountableWaitGroup

    wg.Add(1)
    go func() {
        // do something
        wg.Done()
    }()

    // Prevent further tasks from being added
    wg.PreventAdd()

    // Wait for all tasks to complete
    wg.Wait()
}
```

## Functions
### `Add(i int)`
The `Add` function increments the WaitGroup counter by i. If PreventAdd has been called, this function does nothing.

### `Done()`
The `Done` function decrements the WaitGroup counter by one.

### `Count() int64`
The `Count` function returns the current counter value.

### `IsPending() bool`
The `IsPending` function returns true if the counter is greater than 0, false otherwise.

### `PreventAdd()`
The `PreventAdd` function sets a flag that prevents further increments of the WaitGroup counter.

### `Wait()`
The `Wait` function blocks until the WaitGroup counter is zero.

### `CanAdd() bool`
The `CanAdd` function is responsible for indicating whether the `PreventAdd()` has been invoked or not.

## Tests
Run the tests with `go test .`

## License
This project is licensed under the MIT License. See the LICENSE file for details.
