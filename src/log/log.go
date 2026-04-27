package log

import (
	"fmt"
	"os"
)

func Start() {
}

func Info(msgs ...any) {
	fmt.Print("[INFO] ")
	fmt.Println(msgs...)
}

func Fatal(msgs ...any) {
	fmt.Print("[ERROR] ")
	fmt.Println(msgs...)
	os.Exit(1)
}
