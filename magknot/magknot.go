/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import ()

type MagKnot struct {
}

func (knot *MagKnot) Init() (err error) {
	return
}

func (knot *MagKnot) Deinit() (err error) {
	return
}

func (knot *MagKnot) Connect(timeout uint32) (err error) {
	return
}

func (knot *MagKnot) Close() (err error) {
	return
}

func (knot *MagKnot) Send(buf []byte, timeout uint32) (err error) {
	return
}

func (knot *MagKnot) Recv(timeout uint32) (data []byte, err error) {
	return
}

func New() *MagKnot {
	knot := new(MagKnot)
	return knot
}
