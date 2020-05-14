package service

type Instance struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Mode        string `json:"mode"`
	Status      string `json:"status"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

var (
	ConfigKey string
	Runtime   *Instance
	Endpoint  []string
)

func InitInstance() {
	ConfigKey = "/data/config"
	Runtime = &Instance{}
	Endpoint = []string{}
}
