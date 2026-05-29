package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/arashrasoulzadeh/devenv/src/consts"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorBold   = "\033[1m"
)

// Level represents a log severity level.
type Level string

const (
	LevelInfo    Level = "info"
	LevelWarn    Level = "warn"
	LevelError   Level = "error"
	LevelSuccess Level = "success"
	LevelDebug   Level = "debug"
)

type message struct {
	MType    Level `json:"mType"`
	MContent any   `json:"mContent"`
}

var (
	out      io.Writer = os.Stdout
	messages []message
)

// SetOutput redirects all text output to w.
func SetOutput(w io.Writer) {
	out = w
}

// Start initialises (or resets) the message buffer.
func Start() {
	messages = make([]message, 0)
}

// logAppend is the single internal writer for all levels.
func logAppend(level Level, msgs ...any) {
	for _, m := range msgs {
		messages = append(messages, message{
			MType:    level,
			MContent: m,
		})
	}
}

// Public logging functions.
func Print(msgs ...any)   { logAppend(LevelInfo, msgs...) }
func Info(msgs ...any)    { logAppend(LevelInfo, msgs...) }
func Warn(msgs ...any)    { logAppend(LevelWarn, msgs...) }
func Error(msgs ...any)   { logAppend(LevelError, msgs...) }
func Success(msgs ...any) { logAppend(LevelSuccess, msgs...) }
func Debug(msgs ...any)   { logAppend(LevelDebug, msgs...) }

// levelColor returns the ANSI escape code for a given level.
func levelColor(level Level) string {
	switch level {
	case LevelInfo:
		return colorBlue
	case LevelWarn:
		return colorYellow
	case LevelError:
		return colorRed
	case LevelSuccess:
		return colorGreen
	case LevelDebug:
		return colorBold
	default:
		return colorReset
	}
}

// Flush writes all buffered messages to the configured output, then clears
// the buffer. Output format and colour are controlled by CLI flags defined
// in consts (OutputJSON / OutputColor).
func Flush() {
	defer func() { messages = messages[:0] }()

	useJSON := slices.Contains(os.Args, consts.OutputJSON)
	useColor := slices.Contains(os.Args, consts.OutputColor)

	if useJSON {
		j, err := json.Marshal(messages)
		if err != nil {
			fmt.Fprintf(out, "log: failed to marshal messages: %v\n", err)
			return
		}
		fmt.Fprintln(out, string(j))
		return
	}

	for _, msg := range messages {
		if useColor {
			fmt.Fprintf(out, "%s%s%s %v\n", levelColor(msg.MType), msg.MType, colorReset, msg.MContent)
		} else {
			fmt.Fprintf(out, "%s %v\n", msg.MType, msg.MContent)
		}
	}
}

// Fatal logs messages at error level, flushes output, then exits with code 1.
func Fatal(msgs ...any) {
	Error(msgs...)
	Flush()
	os.Exit(1)
}
