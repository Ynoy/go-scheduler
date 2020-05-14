package cli

import (
	"encoding/json"
	client "github.com/coreos/etcd/clientv3"
	"go-scheduler/pkg/scheduler/registry"
	"go-scheduler/pkg/web/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func startInitConfigFile() {
	// 校验文件路径
	validateFilePath()
	// 校验管理员用户相关参数
	validateUserParams()

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	switch mode {
	case "json":
		if err := json.Unmarshal(buffer, conf.Conf); err != nil {
			log.Fatal(err)
		}
		break
	case "yaml":
		if err := yaml.Unmarshal(buffer, conf.Conf); err != nil {
			log.Fatal(err)
		}
		break
	default:
		log.Fatal("not support this mode,select 'json' or 'yaml'...")
		break
	}
	// New Etcd V3 Client
	registry.NewClient()

	buffer, err = json.Marshal(conf.Conf)
	if err != nil {
		log.Fatal(err)
	}

	if res, err := registry.Client.Put(context.TODO(), conf.Conf.Etcd.Config, string(buffer), client.WithPrevKV()); err != nil {
		log.Fatal(err)
	} else {
		if len(res.PrevKv.Value) > 0 {
			log.Printf("%s %s \n", "OLD CONFIG IS", string(res.PrevKv.Value))
		}
	}

	if err := util.CreateDatabase(); err != nil {
		log.Fatal("Failed to create database", err)
	}
}

func validateUserParams() {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Fatal("Please enter admin user info")
	}
}

// 校验配置文件路径正确性
func validateFilePath() {
	if path == "" {
		log.Fatal("Please enter your config file path or use --mode=server")
	}
	if exist, err := conf.CheckConfigFile(path); err != nil {
		log.Fatal(err)
	} else if !exist {
		log.Fatal("Config file does not exist")
	}

}

