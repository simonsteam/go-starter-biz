package utl

// AnyNoneNil return first none nil error
func AnyNoneNil(errors ...error) error {
	for _, e := range errors {
		if e != nil {
			return e
		}
	}
	return nil
}