package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg map[string]map[string]any
	mu  sync.RWMutex
)

const fileName = "config.toml"
const defaultContent = ""

func Parse(path string) error {
	if path == "" {
		path = fileName
	}

	data, err := load(path)
	if err != nil {
		return handleErr(err, path)
	}

	mu.Lock()
	cfg = data
	mu.Unlock()

	return nil
}

func Get() map[string]map[string]any {
	mu.RLock()
	defer mu.RUnlock()

	if cfg == nil {
		return nil
	}

	// return copy to prevent mutation bugs
	out := make(map[string]map[string]any, len(cfg))
	for k, v := range cfg {
		cp := make(map[string]any, len(v))
		for k2, v2 := range v {
			cp[k2] = v2
		}
		out[k] = cp
	}

	return out
}

func Reset() {
	mu.Lock()
	defer mu.Unlock()
	cfg = nil
}

func load(path string) (map[string]map[string]any, error) {
	raw := make(map[string]any)

	if _, err := toml.DecodeFile(path, &raw); err != nil {
		return nil, err
	}

	out := make(map[string]map[string]any)

	for k, v := range raw {
		if m, ok := v.(map[string]any); ok {
			out[k] = m
		}
	}

	return out, nil
}

func handleErr(err error, path string) error {
	if os.IsNotExist(err) {
		if create(path) {
			fmt.Printf("[INFO] created config: %s\n", path)
		}
		return err
	}

	if os.IsPermission(err) {
		fmt.Printf("[ERROR] no permission: %s\n", path)
		return err
	}

	fmt.Printf("[ERROR] decode error '%s': %v\n", path, err)
	return err
}

func create(path string) bool {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	_, err = f.WriteString(defaultContent)
	return err == nil
}
