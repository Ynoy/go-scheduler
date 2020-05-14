package conf

import (
	"io/ioutil"
	"log"
	"os"
)

type (
	Etcd struct {
		Killer   string   `json:"killer" yaml:"killer" validate:"required"`
		Locker   string   `json:"locker" yaml:"locker" validate:"required"`
		Service  string   `json:"service" yaml:"service" validate:"required"`
		Pipeline string   `json:"pipeline" yaml:"pipeline" validate:"required"`
		Config   string   `json:"config" yaml:"config" validate:"required"`
		Endpoint []string `json:"endpoint" yaml:"endpoint" validate:"required"`
		Timeout  int64    `json:"timeout" yaml:"timeout" validate:"required"`
	}

	Database struct {
		Host string `json:"host" yaml:"host" validate:"required"`
		Port int    `json:"port" yaml:"port" validate:"required"`
		Name string `json:"name" yaml:"name" validate:"required"`
		User string `json:"user" yaml:"user" validate:"required"`
		Pass string `json:"pass" yaml:"pass" validate:"required"`
		Char string `json:"char" yaml:"char" validate:"required"`
	}

	User struct {
		Name    string `json:"name" yaml:"-" validate:"required"`
		Email   string `json:"email" yaml:"-" validate:"required"`
		Pass    string `json:"pass" yaml:"-" validate:"required"`
		Confirm string `json:"confirm" yaml:"-" validate:"required"`
	}

	Auth struct {
		Secret string `json:"secret" yaml:"secret" validate:"required"`
		TTL    int64  `json:"ttl" yaml:"ttl" validate:"required"`
	}

	Notification struct {
		Url        string `json:"url" yaml:"url" validate:"required"`
		Host       string `json:"host" yaml:"host" validate:"required"`
		Port       int    `json:"port" yaml:"port" validate:"required"`
		User       string `json:"user" yaml:"user" validate:"required"`
		Pass       string `json:"pass" yaml:"pass" validate:"required"`
		Name       string `json:"name" yaml:"name" validate:"required"`
		Protocol   string `json:"protocol" yaml:"protocol" validate:"required"`
		Encryption string `json:"encryption" yaml:"encryption" validate:"required"`
	}

	Config struct {
		Database     `json:"database"`
		Auth         `json:"auth"`
		Etcd         `json:"etcd"`
		Notification `json:"notification"`
	}
)

var (
	Conf *Config
	Path string
)

func Init() *Config {
	return &Config{}
}

// 检查配置文件是否存在
func CheckConfigFile(path string) (bool, error) {
	_, err := os.Stat(path)
	exist := !os.IsExist(err)
	return exist, err
}

// 创建配置文件文件夹
func CreateConfigDir(dirPath string) {
	_, err := os.Stat(dirPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Printf("CreateConfigDir is error,error:{%s}", err)
			os.Exit(1)
		}
	}
}

// 检查配置文件文件夹权限
func CheckConfigDirPermissions(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if err != nil {
		log.Println(err)
	}
	mode := info.Mode()
	perm := mode.Perm()
	flag := perm & os.FileMode(493)
	if flag == 493 {
		return true
	}
	return false
}

// 将配置写入文件
func WriteConfig2File(file string, content []byte) bool {
	if err := ioutil.WriteFile(file, content, 0644); err != nil {
		log.Println(os.IsNotExist(err))
		log.Println(err)
		return false
	}
	return true
}
