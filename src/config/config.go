package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg  map[string]map[string]any
	once sync.Once
)

const fileName = "config.toml"
const defaultContent = ""

func Parse(path string) {
	if path == "" {
		path = fileName
	}

	data, err := load(path)
	if err != nil {
		handleErr(err, path)
		return
	}

	once.Do(func() {
		cfg = data
	})
}

func Get() map[string]map[string]any {
	return cfg
}

func load(path string) (map[string]map[string]any, error) {
	raw := make(map[string]any)
	_, err := toml.DecodeFile(path, &raw)
	if err != nil {
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

func handleErr(err error, path string) {
	switch {
	case os.IsNotExist(err):
		if create(path) {
			fmt.Printf("created config: %s\n", path)
		}
	case os.IsPermission(err):
		fmt.Printf("no permission: %s\n", path)
	default:
		fmt.Printf("decode error '%s': %v\n", path, err)
	}
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
