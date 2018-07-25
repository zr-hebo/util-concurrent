package concurrent

import "sync"

const defaultSliceSize = 8

type resultPair struct {
	data interface{}
	err  error
}

// ResultCollector error collector
type ResultCollector struct {
	results []resultPair
	lock    *sync.Mutex
}

// NewResultCollector create a new result collector
func NewResultCollector() (rc *ResultCollector) {
	return &ResultCollector{
		results: make([]resultPair, 0, defaultSliceSize),
		lock:    new(sync.Mutex),
	}
}

// Collect collect result returned
func (rc *ResultCollector) Collect(data interface{}, err error) {
	rc.lock.Lock()
	defer rc.lock.Unlock()

	rc.results = append(rc.results, resultPair{data: data, err: err})
}

// SumUp sum up
func (rc *ResultCollector) SumUp() (allData []interface{}, err error) {
	allData = make([]interface{}, 0, defaultSliceSize)
	errMsgs := make([]string, 0, defaultSliceSize)
	for _, r := range rc.results {
		if r.data != nil {
			allData = append(allData, r.data)
		}

		if r.err != nil {
			errMsgs = append(errMsgs, r.err.Error())
		}
	}

	err = sumUpError(errMsgs)
	return
}
