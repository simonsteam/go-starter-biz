package biz

import (
	"fmt"
	"time"

	"go.uber.org/dig"
)

// NewEnv .
func NewEnv(modules ...Module) *Env {
	var slc []Module
	for _, m := range modules {
		slc = append(slc, m)
	}
	env := Env{
		modules:   &slc,
		Container: dig.New(),
	}
	err := env.Container.Provide(func() *Env {
		return &env
	})
	if err != nil {
		panic(err)
	}
	return &env
}

// Env .
// any module runs im env
// lang:zh_CN 代表运行的环境，包括模块，启动信息
type Env struct {
	Container    *dig.Container
	modules      *[]Module
	okConditions *[]BootCondition
	bootFlag     bool
}

func (e *Env) conditionsMeet(conditions []BootCondition) bool {
	if len(conditions) == 0 {
		return true
	}
	m := map[BootCondition]bool{}
	for _, okCondition := range *e.okConditions {
		m[okCondition] = true
	}
	for _, c := range conditions {
		if _, exists := m[c]; !exists {
			return false
		}
	}
	return true
}

// Boot .
func (e *Env) Boot() {
	for _, m := range *e.modules {
		if arr, ok := m.Provider.([]interface{}); ok {
			for _, itfc := range arr {
				err := e.Container.Provide(itfc)
				if err != nil {
					panic(err)
				}
			}
		} else {
			err := e.Container.Provide(m.Provider)
			if err != nil {
				panic(err)
			}
		}
	}

	//3.boot modules
	loopCount := 0
	for {
		invokedCount := 0
		for _, m := range *e.modules {
			if m.BootFn == nil || m.BootFn.invoked {
				invokedCount++
				continue
			}

			if e.conditionsMeet(m.BootFn.Preconditions) {
				err := e.Container.Invoke(m.BootFn.Fn)
				if err != nil {
					panic(err)
				}
				m.BootFn.invoked = true
				invokedCount++
			}
		}
		if invokedCount == len(*e.modules) {
			break
		}
		loopCount++
		if loopCount > len(*e.modules) {
			panic("There may be condition will never meet!\\理论上，循环次数应该小于模块数，可能有的启动条件无法达成")
		}
	}
}

// ConditionOK some boot condition meet \\lang:zh_CN 启动条件达成
func (e *Env) ConditionOK(condition BootCondition) {
	conditions := append(*e.okConditions, condition)
	e.okConditions = &conditions
}

// Close when env come to end,call every Close function in modules(if not nil)
// lang:zh_CN 关闭环境时，挨个调用模块的关闭方法
func (e *Env) Close() []error {
	var errors []error
	for _, m := range *e.modules {
		if m.CloseFn != nil {
			err := e.Container.Invoke(m.CloseFn)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

// BootCondition any declared info, some module may need to boot
type BootCondition string

// Module repesent a module,a module should declare what it provides(that is Provider) and what it need to do to start(that is Bootstrap)
// dependencies should as less as possible,to keep module simple,also outputs should be simple too, That is 'High Cohesion & Low Coupling'
// lang:zh_CN 业务模块，模块应该声明所需要注入的条件以及可以提供的类型，另外还有启动函数
// 模块的依赖应该尽可能的少，最好也输出简单的对象，即：高内聚低耦合
type Module struct {
	Name         string
	Introduction string
	Provider     interface{} // func, params declare what it needs, return what it provide, lang:zh_CN 参数声明需要注入的类型，返回值表示模块能够提供的类型
	BootFn       *BootFunc   // func, params declare what it needs, it will be invoked before using, lang:zh_CN 模块启动前需要执行的逻辑（是这样设定的）
	CloseFn      interface{} // invoked when close env
}

// BootFunc boot function, will be invoked
type BootFunc struct {
	Preconditions []BootCondition
	Fn            interface{}
	invoked       bool
}

// Err error with code,message,time
type Err struct {
	Code uint32    `json:"code"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

func (e Err) Error() string {
	return fmt.Sprintf("Error %d, %s %v", e.Code, e.Msg, e.Time)
}
