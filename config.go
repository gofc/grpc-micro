package micro

import (
	"context"
	"flag"
	"github.com/gofc/grpc-micro/code"
	"github.com/gofc/grpc-micro/locale"
	"github.com/gofc/grpc-micro/localizer"
	"github.com/gofc/grpc-micro/logger"
	"github.com/jinzhu/configor"
	"path/filepath"
)

var (
	envName    = flag.String("env", "local", "environment name")
	configPath = flag.String("conf", "", "the folder path of config files")
)

//ServerConfig サーバー設定
type ServerConfig struct {
	BindAddress  string `yaml:"bind_address" default:"0.0.0.0"`
	Port         int    `yaml:"port" default:"50000"`
	HTTPPort     int    `yaml:"http_port" default:"31000"`
	Kuberesolver bool   `yaml:"kuberesolver"`
}

//LoggerInterface logger config interface
type LoggerInterface interface {
	GetLoggerConfig() *logger.Config
}

//LanguageInterface config interface
type LanguageInterface interface {
	GetLanguageConfig() *locale.Languages
}

//LocalizerInterface config interface
type LocalizerInterface interface {
	GetLocalizerConfig() *localizer.Config
}

// service config
func Init(code code.ServiceCode, conf interface{}) (closeFunc func(), err error) {
	flag.Parse()

	var funcs []func()
	var filePath string
	if *configPath == "" {
		filePath = filepath.Clean(
			filepath.Join("./configs/app", *envName, code.Name()+".yml"),
		)
	} else {
		filePath = filepath.Clean(
			filepath.Join(*configPath, *envName, code.Name()+".yml"),
		)
	}
	if err := configor.Load(conf, filePath); err != nil {
		return nil, err
	}

	//init logger
	if lc, ok := conf.(LoggerInterface); ok {
		logger.InitLogger(lc.GetLoggerConfig())
		funcs = append(funcs, func() {
			logger.Close(context.Background())
		})
	}

	//init Language
	if lc, ok := conf.(LanguageInterface); ok {
		lc.GetLanguageConfig().SetGlobal()
	}

	//init Localizer
	if lc, ok := conf.(LocalizerInterface); ok {
		localizer.Init(lc.GetLocalizerConfig())
	}

	return func() {
		for _, fn := range funcs {
			fn()
		}
	}, nil
}
