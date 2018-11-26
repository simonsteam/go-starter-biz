package ac

import (
	"context"
	"fmt"
	"local/biz"
)

type contextKey int

const (
	// CtxUserIDKey userID key in context
	CtxUserIDKey contextKey = iota
)

// Permission .
type Permission string

// Res some resources that will be access (do)
type Res struct {
	ID      *string
	Name    string // used for log or message
	Type    string
	Domains []string
	Owner   *string
}

// SubType subject type
type SubType string

// some SubTypes
const (
	SubTypeSys   SubType = "system"
	SubTypeHuman SubType = "human"
)

// Sub subject, human or system
type Sub struct {
	ID          string // usually user ID
	Name        string // for log or message
	Type        SubType
	Permissions []string
	Domains     []string
}

func (s Sub) String() string {
	return fmt.Sprintf("Sub %s<%s√>(type:%s), with Permissions: %s", s.Name, s.ID, s.Type, s.Permissions)
}

// some actions
const (
	CREATE = "create"
	READ   = "read"
	UPDATE = "update"
	DELETE = "delete"
)

// Do do action
type Do string

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
	for _, ck := range r.Checkers {
		if err := ck(sub, r.Res); err != nil {
			return err
		}
	}
	return nil
}

// Checker check subject and resource, return nil if allowed
type Checker func(Sub, Res) error

// BeOwnerOrHasPermission is owner or has permission XXX
func BeOwnerOrHasPermission(permission Permission) Checker {
	return func(sub Sub, res Res) error {
		err := BeOwner(sub, res)
		if err == nil {
			return nil
		}
		err = HasPermission(permission)(sub, res)
		return err
	}
}

// HasPermissionAndInResDomains .
func HasPermissionAndInResDomains(permission Permission) Checker {
	return func(sub Sub, res Res) error {
		err := HasPermission(permission)(sub, res)
		if err == nil {
			err = InResDomains(sub, res)
		}
		return err
	}
}

// InResDomainsAndHasPermission and
func InResDomainsAndHasPermission(permission Permission) Checker {
	return func(sub Sub, res Res) error {
		err := BeOwner(sub, res)
		if err == nil {
			err = HasPermission(permission)(sub, res)
		}
		return err
	}
}

// HasPermission return true if sub has permission ,or false
func HasPermission(permisson Permission) Checker {
	return func(sub Sub, res Res) error {
		for _, perm := range sub.Permissions {
			if perm == string(permisson) {
				return nil
			}
		}
		return biz.NewErr(biz.CodeForbidden, fmt.Sprintf("[%s] required!", permisson))
	}
}

func InDomains(domains []string) Checker {
	return func(sub Sub, _ Res) error {
		m := map[string]bool{}
		for _, dm := range sub.Domains {
			m[dm] = true
		}

		for _, dm := range domains {
			if _, exists := m[dm]; !exists {
				return biz.NewErr(biz.CodeForbidden, fmt.Sprintf("[Forbidden] should in domain <%s>", dm))
			}
		}
		return nil
	}
}

// InResDomains res domains should be subset of subject domains
// lang:zh_CN 资源的域需要是访问者所属域的一个子集
var InResDomains Checker = func(sub Sub, res Res) error {
	m := map[string]bool{}
	for _, dm := range sub.Domains {
		m[dm] = true
	}

	for _, dm := range res.Domains {
		if _, exists := m[dm]; !exists {
			return biz.NewErr(biz.CodeForbidden, fmt.Sprintf("[Forbidden] should in domain <%s>", dm))
		}
	}
	return nil
}

// BeOwner .
var BeOwner Checker = func(sub Sub, res Res) error {
	if *res.Owner != sub.ID {
		return biz.NewErr(biz.CodeForbidden, "should be owner")
	}
	return nil
}

// GetSubFromContext .
func GetSubFromContext(ctx context.Context) (Sub, bool) {
	sub, ok := ctx.Value(CtxUserIDKey).(Sub)
	return sub, ok
}

func CreateContextWithSub(sub Sub) context.Context {
	return context.WithValue(
		context.Background(),
		CtxUserIDKey,
		sub,
	)
}
