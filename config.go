package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration of program.
// It can be loaded from a YAML file.
type Config struct {
	// General network settings. (wosu related)
	Network struct {
		// The port the DNS server should listen on.
		ListenPort int `yaml:"listen_port"`
		// The IP address of the DNS server to use for resolving [ALLOWED] requests.
		Resolver string `yaml:"resolver"`
	} `yaml:"network"`
	// Logging settings. (wosu related)
	Logging struct {
		// If true, blocked requests will be logged.
		Blocked bool `yaml:"blocked"`
		// If true, allowed requests will be logged.
		Allowed bool `yaml:"allowed"`
	} `yaml:"logging"`
	// Request filtering and blocking settings. (webOS related)
	Blocking struct {
		SmartAd          bool `yaml:"lg_smart_ad"`
		HomeDashboard    bool `yaml:"home_dashboard"`
		Sports           bool `yaml:"sports"`
		AppStore         bool `yaml:"app_store"`
		InternetChannels bool `yaml:"internet_channels"`
		LGIOT            bool `yaml:"lg_iot"`
		Amazon           bool `yaml:"amazon"`
		PhilipsHue       bool `yaml:"philips_hue"`
		SoftwareUpdates  bool `yaml:"software_updates"`
	} `yaml:"blocking"`
}

// Gets the default config settings.
func DefaultConfig() *Config {
	cfg := &Config{}
	cfg.Network.ListenPort = 53
	cfg.Network.Resolver = "8.8.8.8:53"
	cfg.Logging.Blocked = true
	cfg.Logging.Allowed = true
	cfg.Blocking.SmartAd = true
	cfg.Blocking.HomeDashboard = false
	cfg.Blocking.Sports = true
	cfg.Blocking.AppStore = false
	cfg.Blocking.InternetChannels = false
	cfg.Blocking.LGIOT = false
	cfg.Blocking.Amazon = false
	cfg.Blocking.PhilipsHue = true
	cfg.Blocking.SoftwareUpdates = false
	return cfg
}

// Checks if the config file with the given name exists.
func (cfg *Config) Exists(name string) bool {
	_, err := os.Stat(name)
	return !errors.Is(err, os.ErrNotExist)
}

// Sets the content of the config object to the content of the file with the given name.
// Returns an error if the file does not exist or if the file is not a valid YAML file.
func (cfg *Config) Load(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	return err
}
