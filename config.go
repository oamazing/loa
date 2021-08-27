package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var config Config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	env := os.Getenv("GOENV")
	if env == `` {
		env = `dev`
	}
	file := filepath.Join(dir, fmt.Sprintf("release/envs/%s.yml", env))
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}
	url, err := url.Parse(config.Mysql)
	if err != nil {
		panic(err)
	}
	fmt.Println(`Host` + url.Host)
	fmt.Println(`User` + url.User.Username())
	// fmt.Println(`User` + url.User.Password())
	fmt.Println(`Port` + url.Port())
	fmt.Printf("%+v", url)
}

type Config struct {
	Mysql string `yaml:"mysql"`
}

func GetConfig() *Config {
	return &config
}
