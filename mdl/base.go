package mdl

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// Base general base type for model
type Base struct {
	ID      int        `json:"id" pg:",pk"`
	AddTime time.Time  `json:"addTime"`
	UpdTime *time.Time `json:"updTime"`
	DelTime *time.Time `json:"delTime"`
}

// BeforeInsert orm hook
func (bs *Base) BeforeInsert(_ orm.DB) error {
	bs.AddTime = time.Now()
	return nil
}

// BeforeUpdate orm hook
func (bs *Base) BeforeUpdate(_ orm.DB) error {
	now := time.Now()
	bs.UpdTime = &now
	return nil
}
