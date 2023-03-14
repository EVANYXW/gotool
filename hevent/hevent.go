package hevent

import (
	"sync"
	"time"
)

type HEvent struct {
	Data  interface{}
	Topic string
}

type HEventChan chan HEvent
type HEventChanArray []HEventChan //一个topic 可以有多个消费者

type HEventBus struct {
	sub map[string]HEventChanArray
	rm  sync.RWMutex
}

var h *HEventBus

func init() {
	h = &HEventBus{
		sub: make(map[string]HEventChanArray),
	}
}

type HEventRepo interface {
	Sub(topic string, ch HEventChan)
	Push(topic string, data interface{})
}

func HEventSrv() *HEventBus {
	return h
}

func (h *HEventBus) Sub(topic string, ch HEventChan) {
	h.rm.Lock()
	if chanEvent, ok := h.sub[topic]; ok {
		h.sub[topic] = append(chanEvent, ch)
	} else {
		h.sub[topic] = append([]HEventChan{}, ch)
	}
	defer h.rm.Unlock()
}

func (h *HEventBus) Push(topic string, data interface{}) {
	h.rm.RLock()
	defer h.rm.RUnlock()
	if chanEvent, ok := h.sub[topic]; ok {
		for _, ch := range chanEvent {
			ch <- HEvent{
				Data:  data,
				Topic: topic,
			}
		}
	}
}

func (h *HEventBus) PushFullDrop(topic string, data interface{}) {
	h.rm.RLock()
	defer h.rm.RUnlock()
	if chanEvent, ok := h.sub[topic]; ok {
		for _, ch := range chanEvent {
			select {
			case ch <- HEvent{
				Data:  data,
				Topic: topic,
			}:
			case <-time.After(time.Second):
			}
		}
	}
}
