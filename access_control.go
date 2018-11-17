package biz

import (
	"fmt"
)

// Res some resources that will be access (do)
type Res struct {
	ID      *uint32
	Name    string // used for log or message
	Type    string
	Domains []uint32
	Owner   uint32
}

// SubType .
type SubType string

// some SubTypes
const (
	SubTypeSys   = "system"
	SubTypeHuman = "human"
)

// Sub subject, human or system
type Sub struct {
	ID          uint32
	Name        string // for log or message
	Type        SubType
	Permissions []string
	Domains     []uint32
}

func (s Sub) String() string {
	return fmt.Sprintf("Sub %s<%d>(type:%s), with Permissions: %s", s.Name, s.ID, s.Type, s.Permissions)
}

// some actions
const (
	CREATE = "create"
	READ   = "read"
	UPDATE = "update"
)

// Do some action verb
type Do struct {
	Action string
}

// To .
func (do Do) To(res Res) DoActToRes {
	return DoActToRes{do, res}
}

// DoActToRes do action to resource
type DoActToRes struct {
	Do
	Res
}

// Should add checker
func (d DoActToRes) Should(checkers ...Checker) Rule {
	return Rule{
		DoActToRes: d,
		Checkers:   checkers,
	}
}

// Rule do{ACTION}.To(res).Should(checkers)
type Rule struct {
	DoActToRes
	Checkers []Checker
}

// Check do check on each Checker, return nil if all Checker return true
func (r Rule) Check(sub Sub) error {
	if len(r.Checkers) == 0 {
		return nil
	}
	for idx, ck := range r.Checkers {
		if rs := ck(sub, r.Res); !rs {
			return fmt.Errorf("failed on chcker %d", idx)
		}
	}
	return nil
}

// Checker check subject and resource, return true if allowed
type Checker func(Sub, Res) bool

// HasPermission return true if sub has permission ,or false
func HasPermission(permisson string) Checker {
	return func(sub Sub, res Res) bool {
		for _, perm := range sub.Permissions {
			if perm == permisson {
				return true
			}
		}
		return false
	}
}

// InDomain .
func InDomain(domain uint32) Checker {
	return func(sub Sub, res Res) bool {
		for _, dm := range sub.Domains {
			if dm == domain {
				return true
			}
		}
		return false
	}
}

// BeOwner .
var BeOwner Checker = func(sub Sub, res Res) bool {
	return res.Owner == sub.ID
}
