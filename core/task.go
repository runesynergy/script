package core

import (
	"errors"
	"fmt"
)

// ActionResult describes the result of a Yield.
type ActionResult int

const (
	// ActionResultOk means that the yield was and the action entered execution.
	ActionResultOk ActionResult = iota
	// ActionResultComplete means that the action completed during the yield.
	ActionResultComplete
	// ActionResultAlreadyCompleted means the action has already completed and the yield has no effect.
	ActionResultAlreadyCompleted
	// ActionResultError means the previous yield returned an error.
	ActionResultError
)

type ActionContext struct {
	yield  chan bool
	cancel chan bool
	done   chan bool
	error  chan error
}

type Action interface {
	Execute(ActionContext)
}

// HostYield should only be called by the core.
func (c ActionContext) HostYield() (result ActionResult, err error) {
	select {
	case c.yield <- true:
	default: // no one is listening, we must've already finished!
		result = ActionResultAlreadyCompleted
		return
	}

	select {
	case err = <-c.error:
		result = ActionResultError
	case <-c.yield:
		result = ActionResultOk
	case <-c.done:
		result = ActionResultComplete
	}
	return
}

// HostCancel should only be called by the core.
func (c ActionContext) HostCancel() {
	c.cancel <- true
}

// ErrScriptCancelled is used to stop the goroutine.
var ErrScriptCancelled = errors.New("script cancelled")

// ErrScriptPanic is used to report unhandled panics.
var ErrScriptPanic = errors.New("script panic")

// Yield should only be called within the script.
func (c ActionContext) Yield() {
	c.yield <- true // tell core we yield
	c.waitOrCancel()
}

// waitOrCancel should only be called within the script goroutine.
func (c ActionContext) waitOrCancel() {
	select {
	case <-c.yield: // core is yielding to us
	case <-c.cancel: // core is telling us to cancel
		panic(ErrScriptCancelled)
	}
}

func (c ActionContext) Wait(duration int) {
	for duration > 0 {
		c.Yield()
		duration--
	}
}

func (c ActionContext) Errorf(format string, args ...any) {
	c.Error(fmt.Errorf(format, args...))
}

func (c ActionContext) Error(err any) {
	switch e := err.(type) {
	case string:
		err = errors.New(e)
	}
	c.error <- err.(error)
}

func Submit(action Action) (ctx ActionContext) {
	ctx = ActionContext{
		yield:  make(chan bool),
		done:   make(chan bool),
		cancel: make(chan bool),
		error:  make(chan error),
	}

	go func(ctx ActionContext) {
		defer func() {
			if err, ok := recover().(error); ok {
				// send the error if it wasn't the cancel error.
				if !errors.Is(err, ErrScriptCancelled) {
					ctx.error <- errors.Join(ErrScriptPanic, err)
				}
			} else
			// panic was called without a proper error type
			{
				ctx.error <- fmt.Errorf("%+v", err)
			}
		}()

		ctx.waitOrCancel()
		action.Execute(ctx)
		ctx.done <- true
	}(ctx)

	return
}
