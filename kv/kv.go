// Create by Yale 2018/10/25 17:15
package kv

type EventType int32

const (
	EVENT_PUT  EventType = 0
	EVENT_DELETE EventType = 1
	EVENT_UPDATE EventType = 2
	EVENT_UNKOWN EventType = 3
)

type Node struct {
	Key string
	Value string
	Event EventType
}

type KV interface {

	Init()error
	Get(key string) (*Node,error)
	Delete(key string) (error)
	GetPrefix(key string) ([]*Node,error)
	Put(key ,value string) error
	Watch(key string,ch chan  *Node)
	WatchPrefix(key string,ch  chan *Node)
	Lock(key string)(error,string)
	UnLock(key string)error
	Close()

}
