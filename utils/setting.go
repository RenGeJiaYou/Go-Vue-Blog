package utils

//抓取 ini 文件中的配置项
import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbName     string
	DbPassWord string
	DbUser     string
	DbPort     string
	DbHost     string
	Db         string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查配置文件路径", err)
	}
	LoadServer(file)
	LoadData(file)
}

// 加载 [server] 配置文件
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	JwtKey = file.Section("server").Key("JwtKey").String()
}

// 加载 [database] 配置文件
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
}
