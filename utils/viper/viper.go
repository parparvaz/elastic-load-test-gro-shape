package viper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	confs = Config{}
	lock  = sync.Mutex{}
)

type ElasticSearch struct {
	Address string `yaml:"address" required:"true"`
	Port    int    `yaml:"port" required:"true"`
}

type Database struct {
	ElasticSearch ElasticSearch `yaml:"elasticsearch" required:"true"`
}

type LoadTest struct {
	WholeTime    int `yaml:"wholeTime"`
	Interval     int `yaml:"interval"`
	RequestCount int `yaml:"requestCount"`
}

type Config struct {
	Database Database            `yaml:"database" required:"true"`
	LoadTest map[string]LoadTest `yaml:"loadTest"`
}

func InitConfigs() Config {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	conf := loadConfigs()

	viper.OnConfigChange(func(in fsnotify.Event) {
		lock.Lock()
		defer lock.Unlock()
		lastUpdate := viper.GetTime("fsnotify")
		if time.Since(lastUpdate) < time.Second {
			return
		}
		viper.Set("fsnotify", time.Now())
		loadConfigs()
		log.Println("config file changed. restarting...")
	})
	viper.WatchConfig()

	return conf
}

func Validate(c any) error {
	errmsg := ""
	numFields := reflect.TypeOf(c).NumField()
	for i := 0; i < numFields; i++ {
		fieldType := reflect.TypeOf(c).Field(i)
		tagval, ok := fieldType.Tag.Lookup("required")
		isRequired := ok && tagval == "true"
		if !isRequired {
			continue
		}
		fieldVal := reflect.ValueOf(c).Field(i)
		if fieldVal.Kind() == reflect.Struct {
			if err := Validate(fieldVal.Interface()); err != nil {
				errmsg += fmt.Sprintf("%s > [%v], ", fieldType.Name, err)
			}
		} else {
			if fieldVal.IsZero() {
				errmsg += fmt.Sprintf("%s, ", fieldType.Name)
			}
		}
	}
	if errmsg == "" {
		return nil
	}
	return errors.New(errmsg)
}

func loadConfigs() Config {
	must(viper.Unmarshal(&confs),
		"could not unmarshal config file")
	must(Validate(confs),
		"some required configurations are missing")
	log.Printf("config loaded from file successfully \n")

	return confs
}

func must(err error, logMsg string) {
	if err != nil {
		log.Println(logMsg)
		panic(err)
	}
}

func C() *Config {
	return &confs
}
