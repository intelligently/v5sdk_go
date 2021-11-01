package wInterface

import . "github.com/intelligently/v5sdk_go/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
