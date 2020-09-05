# Package `haloob`
`haloob` is an Out Of Box (oob) implementation of the `hal` interface for Bellina.

> HAL stands for the `Hardware Abstraction Layer`

It uses the `g5` package for its graphics implementation.

Through an object that implements the `hal` interface, Bellina can interact with the OS window and be aware of window resize, mouse, and keyboard events.

`hal` is how Bellina creates an initial window with a given dimension and title.

# Installation

`go get -u github.com/amortaza/go-hal-oob`

# usage

In your application, call `bl.Start(...)` to start Bellina.  `bl.Start(...)` requires a `hal` interface implementation.

`haloob.NewHal()` provides an object that implements the `hal` interface.

See example/haloob-example.go for a complete application

> NOTE: `haloob`'s implementation uses GLFW, which requires `runtime.LockOSThread()` in `init()`

```go
import "github.com/amortaza/go-bellina"
import "github.com/amortaza/go-hal-oob"

func init() {
	runtime.LockOSThread()
}

func main() {
    hal := hal_g5.NewHal()()
    bl.Start(hal, 800, 600, "Title", onInit, onTick, onUninit))
}

func onInit() {
    // Initialize my app
}

func onTick() {
    // Update app state
}

func onUninit() {
    // Free memory in my app
}

```