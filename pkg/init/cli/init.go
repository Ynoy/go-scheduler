package cli

import (
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"go-scheduler/pkg/init"
	"go-scheduler/pkg/util"
	"go-scheduler/pkg/web/conf"
	"go-scheduler/pkg/web/model"
	"go-scheduler/pkg/web/service"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	mode string
	path string

	user = &model.User{
		Id:        uuid.NewV4().String(),
		Manager:   true,
		CreatedAt: util.Time(time.Now()),
		UpdatedAt: util.Time(time.Now()),
	}

	initCmd = &cobra.Command{
		Use:     "init",
		Short:   "cli go-scheduler server",
		Long:    "cli go-scheduler server",
		Example: "go-scheduler init",
		Run: func(cmd *cobra.Command, args []string) {
			switch mode {
			case "server":
				// init web server
				startInitWebServer()
				break
			case "config":
				// init config file
				startInitConfigFile()
				break
			default:
				log.Printf("unkown mode,server is exiting...")
				os.Exit(1)
			}
		},
	}
)

func init() {
	// 设置服务最大线程数
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 初始化配置对象
	conf.Conf = conf.Init()
	init.Cli.AddCommand(initCmd)

	service.Runtime = &service.Instance{
		Version: init.Cli.Version,
	}

	initCmd.Flags().StringVarP(&mode, "mode", "m", "server", "Set cli mode with web ui or json, yaml config file")
	initCmd.Flags().StringVarP(&path, "path", "p", "", "Set config file path")
	initCmd.Flags().StringVarP(&user.Name, "name", "n", "", "Set admin name")
	initCmd.Flags().StringVarP(&user.Email, "email", "e", "", "Set admin email")
	initCmd.Flags().StringVarP(&user.Password, "password", "P", "", "Set admin password")

}
