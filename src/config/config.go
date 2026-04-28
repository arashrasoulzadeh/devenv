package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const fileName = "config.toml"
const defaultContent = ""

type Config struct {
	data map[string]map[string]any
}

// constructor
func New() *Config {
	return &Config{
		data: make(map[string]map[string]any),
	}
}

// Load reads and parses config file into THIS instance
func (c *Config) Load(path string) error {
	if path == "" {
		path = fileName
	}

	raw, err := load(path)
	if err != nil {
		return handleErr(err, path)
	}

	c.data = raw
	return nil
}

// Get returns a SAFE copy (no mutation leaks)
func (c *Config) Get() map[string]map[string]any {
	if c.data == nil {
		return nil
	}

	out := make(map[string]map[string]any, len(c.data))

	for k, v := range c.data {
		cp := make(map[string]any, len(v))
		for k2, v2 := range v {
			cp[k2] = v2
		}
		out[k] = cp
	}

	return out
}

// ----------------- internal helpers -----------------

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
		_ = create(path)
		return errors.New(fmt.Sprintf("file %s not found,so created %s with empty data", path, path))
	}

	if os.IsPermission(err) {
		return err
	}

	return err
}

func create(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(defaultContent)
	return err
}
