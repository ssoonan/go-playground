package main

import (
	"fmt"
	"time"
)

type Watcher interface {
	ResultChan() <- chan Event
	Stop()
}

// Event 타입 정의
type Event struct {
	Type    string
	Message string
}

// 이렇게 채널 정의가 필요, 이후에 make로 만들기
type myWatch chan Event

// Stop() 메서드 구현 (실제 구현에서는 채널을 닫거나 리소스를 정리)
func (w myWatch) Stop() {
	close(w) // 채널 닫기
}

func (w myWatch) ResultChan() <-chan Event {
	return w
}

type mywatch2 struct {
	Channel chan Event
}

func (w mywatch2) Stop() {
}

func (w mywatch2) ResultChan() <-chan Event {
	return w.Channel
}

var w2 Watcher = mywatch2{Channel: make(myWatch)}
var w Watcher = make(myWatch)

func main() {
	// 새로운 myWatch 채널 생성. make로 만드는 건 어렵지 않음
	watch := make(myWatch)

	// 새로운 고루틴을 실행하여 이벤트를 전송
	go func() {
		// <- 를 통해 watch에 값을 넣기. 
		watch <- Event{Type: "ADDED", Message: "New pod created"}
		watch <- Event{Type: "MODIFIED", Message: "Pod updated"}
		time.Sleep(1 * time.Second)
		watch <- Event{Type: "DELETED", Message: "Pod deleted"}
		close(watch) // 채널을 닫아서 더 이상 이벤트를 보내지 않도록 함
	}()

	// ResultChan()을 통해 이벤트 수신
	// channel을 통해 받는 쪽에선? 다른 방법도 있던 거 같은데
	for event := range w.ResultChan() {
		fmt.Printf("Received Event: Type=%s, Message=%s\n", event.Type, event.Message)
	}
}
