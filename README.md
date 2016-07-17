# haloob
halloob is an Out Of Box (oob) implementation of the Hal interface for Bellina.

> Hal stands for the `Hardware Abstraction Layer`

Through an object that implements the Hal interface, Bellina can interact with the OS window and be aware of window resize, mouse, and keyboard events.
Hal is how Bellina creates an initial window with a given dimension and title.

# usage

In your application, call `bl.Start(...)` to start Bellina

See example/haloob-example.go for a complete application

> NOTE: haloob's implementation uses GLFW, which requires `runtime.LockOSThread()` in `init()`

```go
import "github.com/amortaza/go-bellina"
import "github.com/amortaza/go-hal-oob"

func init() {
	runtime.LockOSThread()
}

func main() {
    hal := haloob.New()
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