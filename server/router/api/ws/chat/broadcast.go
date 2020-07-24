package chat

import "sync"

// BoradcastChan 广播消息
type BoradcastChan chan *Message

var rlock sync.RWMutex

// Write Push
func (b BoradcastChan) Write(msg *Message) {
	rlock.Lock()
	b <- msg
	rlock.Unlock()
}
