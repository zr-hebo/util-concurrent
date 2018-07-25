package concurrent

import (
	"errors"
	"strings"
	"sync"
)

// ErrorCollector error collector
type ErrorCollector struct {
	errMsgs []string
	lock    *sync.Mutex
}

// NewErrorCollector create a new error collector
func NewErrorCollector() (ec *ErrorCollector) {
	return &ErrorCollector{
		errMsgs: make([]string, 0, defaultSliceSize),
		lock:    new(sync.Mutex),
	}
}

// CollectError deal error returned
func (ec *ErrorCollector) CollectError(err error) {
	ec.lock.Lock()
	defer ec.lock.Unlock()

	if err != nil {
		ec.errMsgs = append(ec.errMsgs, err.Error())
	}
}

// SumUp sum up all error
func (ec *ErrorCollector) SumUp() (err error) {
	return sumUpError(ec.errMsgs)
}

func sumUpError(errMsgs []string) (err error) {
	distinctErrMsgs := deduplicateStrings(errMsgs)
	if len(distinctErrMsgs) > 0 {
		err = errors.New(strings.Join(distinctErrMsgs, ";"))
	}
	return
}

func deduplicateStrings(strs []string) (newStrs []string) {
	strDict := make(map[string]struct{})
	for _, str := range strs {
		strDict[str] = struct{}{}
	}

	newStrs = make([]string, 0, len(strs))
	for str := range strDict {
		newStrs = append(newStrs, str)
	}

	return
}
