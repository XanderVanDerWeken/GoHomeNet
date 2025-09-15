package events

import (
	"sync/atomic"
	"testing"
	"time"
)

type TestEvent struct {
	name string
}

func (e TestEvent) Name() string {
	return e.name
}

func TestRegisterAndPublish(t *testing.T) {
	eb := NewEventBus()

	var called int32
	eb.Register("test", func(e Event) {
		atomic.AddInt32(&called, 1)
	})

	eb.Publish(TestEvent{name: "test"})
	eb.Wait()

	if called != 1 {
		t.Errorf("expected handler to be called once, got %d", called)
	}
}

func TestMultipleHandlers(t *testing.T) {
	eb := NewEventBus()

	var called int32
	for i := 0; i < 5; i++ {
		eb.Register("multi", func(e Event) {
			atomic.AddInt32(&called, 1)
		})
	}

	eb.Publish(TestEvent{name: "multi"})
	eb.Wait()

	if called != 5 {
		t.Errorf("expected 5 handlers to be called, got %d", called)
	}
}

func TestWaitBlocksUntilHandlersFinish(t *testing.T) {
	eb := NewEventBus()

	var called int32
	eb.Register("slow", func(e Event) {
		time.Sleep(100 * time.Millisecond)
		atomic.AddInt32(&called, 1)
	})

	start := time.Now()
	eb.Publish(TestEvent{name: "slow"})
	eb.Wait()
	elapsed := time.Since(start)

	if called != 1 {
		t.Errorf("expected handler to be called once, got %d", called)
	}
	if elapsed < 100*time.Millisecond {
		t.Errorf("Wait did not block long enough")
	}
}

func TestHandlerPanicRecovery(t *testing.T) {
	eb := NewEventBus()

	var called int32
	eb.Register("panic", func(e Event) {
		panic("boom")
	})
	eb.Register("panic", func(e Event) {
		atomic.AddInt32(&called, 1)
	})

	eb.Publish(TestEvent{name: "panic"})
	eb.Wait()

	if called != 1 {
		t.Errorf("expected second handler to be called despite panic, got %d", called)
	}
}
