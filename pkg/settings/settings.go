package settings

import (
    "github.com/spf13/cast"
    "gopkg.in/ini.v1"
    "log"
    "strings"
    "time"
)

var Conf *ini.File

var (
    buildTime    string
    LastModified string
)

type Server struct {
    HttpHost          string `json:"http_host"`
    HttpPort          string `json:"http_port"`
    RunMode           string `json:"run_mode"`
    JwtSecret         string `json:"jwt_secret"`
    NodeSecret        string `json:"node_secret"`
    HTTPChallengePort string `json:"http_challenge_port"`
    Email             string `json:"email"`
    Database          string `json:"database"`
    StartCmd          string `json:"start_cmd"`
    CADir             string `json:"ca_dir"`
    Demo              bool   `json:"demo"`
    PageSize          int    `json:"page_size"`
    GithubProxy       string `json:"github_proxy"`
}

type Node struct {
    List string `json:"list"`
    Ips  string `json:"ips"`
}

type Nginx struct {
    MasterIp      string `json:"master_ip"`
    NodePort      string `json:"node_port"`
    MasterUrl     string `json:"master_url"`
    Via           string `json:"via"`
    AccessLogPath string `json:"access_log_path"`
    ErrorLogPath  string `json:"error_log_path"`
    ConfigDir     string `json:"config_dir"`
    CertDir       string `json:"cert_dir"`
    PIDPath       string `json:"pid_path"`
    ReloadCmd     string `json:"reload_cmd"`
    RestartCmd    string `json:"restart_cmd"`
}

type OpenAI struct {
    BaseUrl string `json:"base_url"`
    Token   string `json:"token"`
    Proxy   string `json:"proxy"`
    Model   string `json:"model"`
}

type Database struct {
    Dsn string `json:"dsn"`
}

type Redis struct {
    Addr     string `json:"Addr"` // localhost:6379,
    Password string `json:"password"`
    Db       int    `json:"db"`
}

var ServerSettings = Server{
    HttpHost:          "0.0.0.0",
    HttpPort:          "9000",
    RunMode:           "debug",
    HTTPChallengePort: "9180",
    Database:          "database",
    StartCmd:          "login",
    Demo:              false,
    PageSize:          10,
    CADir:             "",
    GithubProxy:       "",
}

var NginxSettings = Nginx{
    AccessLogPath: "",
    ErrorLogPath:  "",
}

var OpenAISettings = OpenAI{}
var (
    DbSettings = Database{
        Dsn: "default",
    }
    RedisSettings = Redis{}
    //Master = MasterConf{}
)

var nodeConf Node
var ConfPath string

var sections = map[string]any{
    "node":   &nodeConf,
    "db":     &DbSettings,
    "redis":  &RedisSettings,
    "server": &ServerSettings,
    "nginx":  &NginxSettings,
    "openai": &OpenAISettings,
}

func init() {
    t := time.Unix(cast.ToInt64(buildTime), 0)
    LastModified = strings.ReplaceAll(t.Format(time.RFC1123), "UTC", "GMT")
}

func Init(confPath string) (err error) {
    ConfPath = confPath
    return Setup()
}

func Setup() (err error) {
    Conf, err = ini.LooseLoad(ConfPath)
    if err != nil {
        log.Fatalf("setting.Setup: %v\n", err)
    }

    if err != nil {
        return
    }

    MapTo()
    return
}

func MapTo() {
    for k, v := range sections {
        mapTo(k, v)
    }
}

func ReflectFrom() {
    for k, v := range sections {
        reflectFrom(k, v)
    }
}

func mapTo(section string, v any) {
    err := Conf.Section(section).MapTo(v)
    if err != nil {
        log.Fatalf("Cfg.MapTo %s err: %v", section, err)
    }
}

func reflectFrom(section string, v interface{}) {
    log.Print(section, v)
    err := Conf.Section(section).ReflectFrom(v)
    if err != nil {
        log.Fatalf("Cfg.ReflectFrom %s err: %v", section, err)
    }
}

func Save() (err error) {
    err = Conf.SaveTo(ConfPath)
    if err != nil {
        return
    }
    Setup()
    return
}

func NginxConfigDir() string {
    thisDir := NginxSettings.ConfigDir
    if "/" != thisDir[len(thisDir)-1:] {
        thisDir = thisDir + "/"
    }

    return thisDir
}

func GetNodeList() []string {
    return strings.Split(nodeConf.List, ",")
}

func GetNodeIps() []string {
    return strings.Split(nodeConf.Ips, ",")
}
