package biz

import (
	"fmt"
	"time"
)

// Err error with code,message,time
type Err struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

func (e Err) Error() string {
	return fmt.Sprintf("Error %d, %s %v", e.Code, e.Msg, e.Time)
}

// UsrType .
type UsrType byte

// some user types that may invoke business methods
const (
	UsrTypeSys   = 0 // system
	UsrTypeHuman = 2 //
)

func (t UsrType) String() string {
	switch t {
	case UsrTypeSys:
		return "system"
	case UsrTypeHuman:
		return "human"
	default:
		return "unknown"
	}
}

// Usr repsent for some "user" that call busines logic
type Usr struct {
	ID     uint32
	Name   string // for log or message
	Type   UsrType
	Groups []string
	Roles  []string
}

func (u Usr) String() string {
	return fmt.Sprintf("Usr %d<%s>(%d) in %v, with roles: %s", u.ID, u.Name, u.Type, u.Groups, u.Roles)
}

// Module repesent a module,a module should declare what it provides(that is Provider) and what it need to do to start(that is Bootstrap)
// dependencies should as less as possible,to keep module simple,also outputs should be simple too, That is 'High Cohesion & Low Coupling'
// lang:zh_CN 业务模块，模块应该声明所需要注入的条件以及可以提供的类型，另外还有启动函数
// 模块的依赖应该尽可能的少，最好也输出简单的对象，即：高内聚低耦合
type Module struct {
	Provider  interface{} // func, params declare what it needs, return what it provide, lang:zh_CN 参数声明需要注入的类型，返回值表示模块能够提供的类型
	Bootstrap interface{} // func, params declare what it needs, it will be invoked before using, lang:zh_CN 模块启动前需要执行的逻辑（是这样设定的）
}
