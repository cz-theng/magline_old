/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// LogLevel is logger's level
type LogLevel int32

const (
	//LNull is empty
	LNull LogLevel = LogLevel(iota)
	//LDebug is debug log
	LDebug
	//LInfo is info log
	LInfo
	//LWarning is warning log
	LWarning
	//LError is error log
	LError
	//LFatal is fatal log
	LFatal
)

// LoggerHooker is type for logger
type LoggerHooker func(string)

// SetLogHooker set hooker for log
func SetLogHooker(hooker LoggerHooker) {
	logcat.hooker = hooker
}

// SetLogLevel set logger's level
func SetLogLevel(level LogLevel) {
	logcat.setLevel(level)
}

var logcat logger

func fmtHooker(log string) {
	fmt.Println(log)
}

func init() {
	logcat.hooker = fmtHooker
	logcat.setLevel(LDebug)
}

type logger struct {
	hooker LoggerHooker
	level  LogLevel
	mtx    sync.Mutex
	buf    []byte
}

func (l *logger) setLevel(level LogLevel) {
	l.level = level
}

func (l *logger) output(level LogLevel, prefix string, format string, v ...interface{}) (err error) {
	var levelStr string
	if level == LDebug {
		levelStr = "[DEBUG]"
	} else if level == LInfo {
		levelStr = "[INFO]"
	} else if level == LWarning {
		levelStr = "[WARNING]"
	} else if level == LError {
		levelStr = "[ERROR]"
	} else if level == LFatal {
		levelStr = "[FATAL]"
	} else {
		levelStr = "[UNKNOWN LEVEL]"
	}

	var msg string
	if format == "" {
		msg = fmt.Sprintln(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}

	l.mtx.Lock()
	defer l.mtx.Unlock()
	l.buf = l.buf[:0]

	l.buf = append(l.buf, levelStr...)
	l.buf = append(l.buf, prefix...)

	l.buf = append(l.buf, ":"+msg...)
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	l.hooker(string(l.buf))
	return
}

/**
* Change from Golang's log.go
* Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
* Knows the buffer has capacity.
 */
func itoa(i int, wid int) string {
	var u = uint(i)
	if u == 0 && wid <= 1 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	return string(b[bp:])
}

func (l *logger) getTime() string {
	// Time is yyyy-mm-dd hh:mm:ss.microsec
	var buf []byte
	t := time.Now()
	year, month, day := t.Date()
	buf = append(buf, itoa(int(year), 4)+"-"...)
	buf = append(buf, itoa(int(month), 2)+"-"...)
	buf = append(buf, itoa(int(day), 2)+" "...)

	hour, min, sec := t.Clock()
	buf = append(buf, itoa(hour, 2)+":"...)
	buf = append(buf, itoa(min, 2)+":"...)
	buf = append(buf, itoa(sec, 2)...)

	buf = append(buf, '.')
	buf = append(buf, itoa(t.Nanosecond()/1e3, 6)...)

	return string(buf[:])
}

func (l *logger) getFileLine() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	return file + ":" + itoa(line, -1)
}

//Debug is debug log
func (l *logger) Debug(format string, v ...interface{}) error {
	if l.level > LDebug {
		return nil
	}

	err := l.output(LDebug, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

func (l *logger) Info(format string, v ...interface{}) error {
	if l.level > LInfo {
		return nil
	}

	err := l.output(LInfo, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

func (l *logger) Warning(format string, v ...interface{}) error {
	if l.level > LWarning {
		return nil
	}
	err := l.output(LWarning, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

func (l *logger) Error(format string, v ...interface{}) error {
	if l.level > LError {
		return nil
	}
	err := l.output(LError, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

func (l *logger) Fatal(format string, v ...interface{}) error {
	if l.level > LFatal {
		return nil
	}
	err := l.output(LFatal, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}
