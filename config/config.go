package config

import (
	"os"
	"path/filepath"

	"github.com/bhoriuchi/dossier/util"
)

// Default locations
const (
	DefaultConfig  = "dossier.yaml"
	DefaultDatadir = "data"
	DefaultMapdir  = "map"
	ConfigEnvVar   = "DOSSIER_CONFIG"
)

// Options configuration options
type Options struct {
	Config    string
	Variables []Variable
	Paths     []string
}

// Variable a variable
type Variable struct {
	Name  string
	Value string
}

// Config a configuration
type Config struct {
	Version   string        `yaml:"version" json:"version"`
	Defaults  *Defaults     `yaml:"defaults" json:"defaults"`
	Hierarchy []*Datasource `yaml:"hierarchy" json:"hierarchy"`
	path      string
	dataRoot  string
	mapRoot   string
}

// Defaults configuration defaults
type Defaults struct {
	Datadir string `yaml:"datadir" json:"datadir"`
	Mapdir  string `yaml:"mapdir" json:"mapdir"`
}

// Datasource source for data
type Datasource struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
}

// Read reads the configuration
func (c *Config) Read(opts Options) (err error) {
	var wd string

	if wd, err = os.Getwd(); err != nil {
		return
	}

	configFile := filepath.Join(wd, DefaultConfig)

	if opts.Config != "" {
		configFile = opts.Config
	} else if os.Getenv(ConfigEnvVar) != "" {
		configFile = os.Getenv(ConfigEnvVar)
	}

	if !filepath.IsAbs(configFile) {
		configFile = filepath.Join(wd, configFile)
	}

	if err = util.ReadFile(configFile, c); err != nil {
		return
	}

	c.path = configFile
	rootDir := filepath.Dir(configFile)
	c.dataRoot = filepath.Join(rootDir, DefaultDatadir)
	c.mapRoot = filepath.Join(rootDir, DefaultMapdir)

	if c.Defaults != nil {
		if c.Defaults.Datadir != "" {
			c.dataRoot = c.Defaults.Datadir
		}
		if c.Defaults.Mapdir != "" {
			c.mapRoot = c.Defaults.Mapdir
		}
	}

	if !filepath.IsAbs(c.dataRoot) {
		c.dataRoot = filepath.Join(rootDir, c.dataRoot)
	}
	if !filepath.IsAbs(c.mapRoot) {
		c.mapRoot = filepath.Join(rootDir, c.mapRoot)
	}

	return
}
