//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"container/list"
	"sync"
)

//LineMgr is connection poll
type LineMgr struct {
	mtx   sync.Mutex
	lines list.List
}

//Init is initialize
func (lm *LineMgr) Init() error {
	lm.mtx.Lock()
	defer lm.mtx.Unlock()
	return nil
}

//Alloc is a allocater
func (lm *LineMgr) Alloc() (line *Line, err error) {
	lm.mtx.Lock()
	defer lm.mtx.Unlock()
	line, err = NewLine()
	if err != nil {
		return
	}
	err = line.Init()
	if err != nil {
		return
	}
	lm.lines.PushBack(line)
	return
}

//Release will reuse a connection
func (lm *LineMgr) Release(line *Line) (err error) {
	lm.lines.Remove(line.Elem)
	err = nil
	return
}

//NewMLLineMgr is ConnPoll creater
func NewMLLineMgr(size int) (lm *LineMgr, err error) {
	lm = new(LineMgr)
	err = lm.Init()
	return
}
