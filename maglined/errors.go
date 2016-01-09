/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"errors"
)

var (
	//ErrAddr Address may be invalied
	ErrAddr = errors.New("Address may be invalied")
)

var (
	//ErrURL is Invaliad URL
	ErrURL = errors.New("Invaliad URL!")

	//ErrNewBuffer is New a Buffer Error
	ErrNewBuffer = errors.New("New a Buffer Error")

	//ErrNetworkType is Unknown Network Type error
	ErrNetworkType = errors.New("Unknown Network Type!")
)

var (
	// ErrArg is argument's error such as nil
	ErrArg = errors.New("Argument's error")
)

var (
	//ErrNewConn is New Connection Error
	ErrNewConn = errors.New("New Connection Error!")
)

var (
	//ErrNewAgent is New a Agent Error
	ErrNewAgent = errors.New("New a Agent Error!")
	//ErrRemove is "Remove from List Error!
	ErrRemove = errors.New("Remove from List Error!")
	//ErrIndex is A Invalied Agent Index!
	ErrIndex = errors.New("A Invalied Agent Index!")
	//ErrrIdleAgent is v
	ErrrIdleAgent = errors.New("It is a Idle Agent!")
	//ErrNoAgent is Don't Have Such a Agent error
	ErrNoAgent = errors.New("Don't Have Such a Agent")
)
