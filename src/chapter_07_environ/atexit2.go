package main

import (
	"fmt"
	"os"
	"sync"
)

func AcquireResource() {
	fmt.Println("进入AcquireResource函数")
}
func AcquireResource2() {
	fmt.Println("进入AcquireResource2函数")
}
func main() {
	defer Do(func() {
		AcquireResource()
	})()
	defer Do(func() {
		AcquireResource2()
	})()
	fmt.Println("Main函数执行")
	Exit(1)

}

type deferred struct {
	f    func()
	m    sync.Mutex
	done uint32
}

// runAtDefer calls the deferred function during standard deferred mechanism.
func (d *deferred) runAtDefer() {
	d.m.Lock()
	defer d.m.Unlock()

	if d.done == 0 {
		defer func() {
			d.done = 1
		}()

		d.f()
	}
}

// runAtExit calls the deferred function during an abnormal program exit.
func (d *deferred) runAtExit() {
	// No need to unlock mutex and update done field.
	d.m.Lock()

	if d.done == 0 {
		d.f()
	}
}

// Do registers the function f to be called in case of abnormal program
// termination, caused by calling the Exit function, and returns a function
// that can be used in a defer statement.
//
// Do SHOULD only be used to release resources that are not automatically
// released by the operating system at program termination.
// The idiomatic use of Do is
//  AcquireResource(...)
//  ...
//  defer atexit.Do(func() {
//      ReleaseResource(...)
//  })()
//
// The function f is called exactly once, either during standard deferred
// mechanism or as part of atexit termination.
//
// If f panics, atexit considers it to have returned; it will not be called again.
func Do(f func()) func() {
	d := &deferred{f: f}
	// TODO(mperillo): protect with a Mutex.
	dl = append(dl, d)

	return d.runAtDefer
}

// Registered deferred functions.
var dl = make([]*deferred, 0, 10)

// Exit runs all registered deferred functions, in FIFO order, and then causes
// the current program to exit with the given status code.
// In case one of the deferred functions panics, the exit status is ignored and
// control passes to Go runtime.
func Exit(code int) {
	func() {
		for _, d := range dl {
			defer d.runAtExit()
		}
	}()

	os.Exit(code)
}
