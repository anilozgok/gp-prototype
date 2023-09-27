package config

import (
	"fmt"
	"github.com/anilozgok/gp-prototype/internal/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"strings"
)

var (
	configWatcher = viper.New()
	secretWatcher = viper.New()
)

func Get(configFolder string) (*Config, error) {
	config := New()
	err := watchFile[Config](config, configFolder, "config", configWatcher)
	if err != nil {
		return nil, err
	}

	err = watchFile[Secrets](config.Secrets, configFolder, "secrets", secretWatcher)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func watchFile[T any](bindTo *T, configFolder, file string, watcher *viper.Viper) error {
	watcher.AddConfigPath(configFolder)
	watcher.SetConfigName(file)
	watcher.SetConfigType("json")
	watcher.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := watcher.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read in config. err: %s", err)
	}

	if err := watcher.Unmarshal(bindTo); err != nil {
		return fmt.Errorf("failed to unmarshal config. err: %s", err)
	}

	watcher.WatchConfig()
	watcher.OnConfigChange(func(event fsnotify.Event) {
		if err := watcher.Unmarshal(bindTo); err == nil {
			log.Logger().Info(fmt.Sprintf("config %s updated", file))
		} else {
			log.Logger().Error(fmt.Sprintf("could not unmarshal %s config.", file), zap.Error(err))
		}
	})

	return nil
}
