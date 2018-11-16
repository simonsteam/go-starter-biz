package utl

import (
	"github.com/go-pg/pg"
	"log"
)

// AnyNoneNil return first none nil error
func AnyNoneNil(errors ...error) error {
	for _, e := range errors {
		if e != nil {
			return e
		}
	}
	return nil
}

func FnDBLogger(event *pg.QueryProcessedEvent) {
	query, err := event.FormattedQuery()
	if err != nil {
		panic(err)
	}

	log.Printf("[DBG] sql: %s \n", query)
}

// SliceDistinct remove same item in slice
func SliceDistinct(slc *[]interface{}) *[]interface{} {
	m := map[interface{}]bool{}

	var results []interface{}
	for _, ele := range *slc {
		if _, ok := m[ele]; !ok {
			m[ele] = true
			results = append(results, ele)
		}
	}
	return &results
}
