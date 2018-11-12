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
