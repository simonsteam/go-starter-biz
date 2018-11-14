package mdl

import (
	// "github.com/go-pg/pg/orm"
	"time"
)

// Config .
// it is designed to store different type of config value in one table
// 配置
type Config struct {
	tableName struct{} `sql:"config"`
	Base
	Key         string     `json:"key" sql:",unique"`          // 程序用key
	Name        string     `json:"name"  sql:",notnull"`       // 名字，简要描述
	Module      string     `json:"module"  sql:",notnull"`     // used for group view,模块，用于简单的配置分组
	Description string     `json:"description" sql:",notnull"` // 详细描述
	Readonly    bool       `json:"readonly"  sql:",notnull"`   // 即：不可修改性
	Type        ConfigType `json:"type"  sql:",notnull"`       // 配置值的类型

	Options     *[]string  `json:"options"`     // used when ConfigType == ConfigTypeRadio | ConfigTypeCheckboxString,可选选项数组,当值为单选、多选时使用该字段
	BoolVal     *bool      `json:"boolVal"`     // 布尔值
	FloatVal    *float64   `json:"floatVal"`    // 浮点值
	IntVal      *int64     `json:"intVal"`      // 整数值
	StrVal      *string    `json:"strVal"`      // 字符值
	TimeVal     *time.Time `json:"timeVal"`     // 时间值
	IntArrayVal *[]int64   `json:"intArrayVal"` // 整数数组值
	StrArrayVal *[]string  `json:"strArrayVal"` // 字符数组值

	UpdRoles []string          `json:"updRoles"`                // roles needed when update config,更新值需要的权限
	Validate map[string]string `json:"validate"`                // 验证规则
	Visible  bool              `json:"visible"  sql:",notnull"` // when false, invisible for human, 可见性，false时不可见
}

// Simple to simple format
func (cfg Config) Simple() ConfigVal {
	return ConfigVal{
		Key: cfg.Key,
		Val: cfg.Val(),
	}
}

// ConfigVal simple format of config
type ConfigVal struct {
	Key string      `json:"key"`
	Val interface{} `json:"val"`
}

// ConfigType 配置值的类型
type ConfigType string

// some ConfigType enums
const (
	ConfigTypeBool           ConfigType = "bool"
	ConfigTypeCheckboxString ConfigType = "checkbox_string"
	ConfigTypeFloat          ConfigType = "float"
	ConfigTypeImageURL       ConfigType = "image_url"
	ConfigTypeInt            ConfigType = "int"
	ConfigTypeIntArr         ConfigType = "int_arr"
	ConfigTypeRadioString    ConfigType = "radio_string"
	ConfigTypeString         ConfigType = "string"
	ConfigTypeStringArr      ConfigType = "string_arr"
	ConfigTypeTime           ConfigType = "time"
)

// Val type -> typeVal
func (cfg Config) Val() interface{} {
	switch cfg.Type {
	case ConfigTypeBool:
		return cfg.BoolVal
	case ConfigTypeCheckboxString:
		return cfg.StrArrayVal
	case ConfigTypeFloat:
		return cfg.FloatVal
	case ConfigTypeImageURL:
		return cfg.StrVal
	case ConfigTypeInt:
		return cfg.IntVal
	case ConfigTypeIntArr:
		return cfg.IntArrayVal
	case ConfigTypeRadioString:
		return cfg.StrVal
	case ConfigTypeString:
		return cfg.StrVal
	case ConfigTypeStringArr:
		return cfg.StrArrayVal
	case ConfigTypeTime:
		return cfg.TimeVal
	default:
		return "[unknown]"
	}
}
