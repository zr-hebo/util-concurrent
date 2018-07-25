package concurrent

// GroupController concurrent group controller
type GroupController struct {
	counter chan struct{}
	maxNum  int
}

// NewGroupController create a new result collector
func NewGroupController(num int) (gc *GroupController) {
	return &GroupController{
		counter: make(chan struct{}, num),
		maxNum:  num,
	}
}

// Acquire Acquire semaphore
func (gc *GroupController) Acquire() {
	gc.counter <- struct{}{}
}

// Release Release semaphore
func (gc *GroupController) Release() {
	<-gc.counter
}

// WaitFinish wait all channel empty
func (gc *GroupController) WaitFinish() {
	for i := 0; i < gc.maxNum; i++ {
		gc.counter <- struct{}{}
	}
}
