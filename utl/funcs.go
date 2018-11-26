package utl

import (
	"log"
	"strconv"

	"github.com/go-pg/pg"
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

// FnItoaPtr convert int to *string
func FnItoaPtr(i int) *string {
	str := strconv.FormatUint(uint64(i), 10)
	return &str
}

// IntsToStrings convert int array to string array
func IntsToStrings(numbers []int) (strs []string) {
	for _, n := range numbers {
		strs = append(strs, strconv.FormatUint(uint64(n), 10))
	}
	return
}

// FnErrorString return error.Error() if not nil
func FnErrorString(e error) string {
	if e == nil {
		return "<nil error>"
	}
	return e.Error()
}
