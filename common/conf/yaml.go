package conf

import (
	"apiBook/common/log"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"apiBook/common/utils"

	"gopkg.in/yaml.v3"
)

var Conf *Configs = &Configs{}

type Configs struct {
	YamlPath string
	YamlData map[string]interface{}
	Default  *DefaultConf
}

type DefaultConf struct {
	App          *App        `yaml:"app"`
	HttpServer   *HttpServer `yaml:"httpServer"`
	GrpcServer   *GrpcServer `yaml:"grpcServer"`
	GrpcClient   *GrpcClient `yaml:"grpcClient"`
	TcpServer    *TcpServer  `yaml:"tcpServer"`
	TcpClient    *TcpClient  `yaml:"tcpClient"`
	UdpServer    *UdpServer  `yaml:"udpServer"`
	UdpClient    *UdpClient  `yaml:"udpClient"`
	Redis        []*Redis    `yaml:"redis"`
	Mysql        []*Mysql    `yaml:"mysql"`
	MqType       string      `yaml:"mqType"`
	Nsq          *Nsq        `yaml:"nsq"`
	Rabbit       *Rabbit     `yaml:"rabbit"`
	Kafka        *Kafka      `yaml:"kafka"`
	Mongo        []*Mongo    `yaml:"mongo"`
	TTF          string      `yaml:"ttf"`
	Cluster      *Cluster    `yaml:"cluster"`
	Jwt          *Jwt        `yaml:"jwt"`
	Minio        *Minio      `yaml:"minio"`
	Mq           string      `yaml:"mq"`
	Etcd         *Etcd       `yaml:"etcd"`
	ShortLinkUrl string      `yaml:"shortLinkUrl"`
	AdminSign    string      `yaml:"adminSign"`
	Domain       string      `yaml:"domain"`
	ShopDomain   string      `yaml:"shopDomain"`
}

// App app相关基础信息
type App struct {
	Name    string `yaml:"name"`
	RunType string `yaml:"runType"` // 项目昵称
}

// HttpServer http服务
type HttpServer struct {
	Open   bool   `yaml:"open"`
	Prod   string `yaml:"prod"`
	Domain string `yaml:"domain"`
}

// GrpcServer grpc服务
type GrpcServer struct {
	Open bool   `yaml:"open"`
	Prod string `yaml:"prod"`
	Log  bool   `yaml:"log"`
}

// GrpcClient grpc客户端
type GrpcClient struct {
	Prod string `yaml:"prod"`
}

// TcpServer tcp服务
type TcpServer struct {
	Open bool   `yaml:"open"`
	Prod string `yaml:"prod"`
}

// TcpClient tcp客户端
type TcpClient struct {
	Prod string `yaml:"prod"`
}

// UdpServer udp服务
type UdpServer struct {
	Open bool   `yaml:"open"`
	Prod string `yaml:"prod"`
}

// UdpClient udp客户端
type UdpClient struct {
	Prod string `yaml:"prod"`
}

// Redis redis配置
type Redis struct {
	Name      string `yaml:"name"` // 自定义一个昵称
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	Password  string `yaml:"password"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

// Mysql mysql配置
type Mysql struct {
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

// MqType 消息队列类型
type MqType struct {
}

// Nsq 消息队列nsq配置
type Nsq struct {
	Producer string `yaml:"producer"`
	Consumer string `yaml:"consumer"`
}

// Rabbit 消息队列rabbit配置
type Rabbit struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Kafka 消息队列kafka配置
type Kafka struct {
	Addr string `yaml:"addr"`
}

// Mongo mongo配置
type Mongo struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Databases string `yaml:"databases"`
}

// Cluster 集群使用 主要用于 ServiceTable
type Cluster struct {
	Open        bool   `yaml:"open"`
	MyAddr      string `yaml:"myAddr"`
	InitCluster string `yaml:"initCluster"`
}

// Jwt jwt配置
type Jwt struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// Minio 对象存储 minio配置
type Minio struct {
	Host   string `yaml:"host"`
	Access string `yaml:"access"`
	Secret string `yaml:"secret"`
}

// Etcd  etcd
type Etcd struct {
	Addr []string `yaml:"addr"`
}

type WechatSDK struct {
	Scheme                string `yaml:"scheme"`
	Host                  string `yaml:"host"`
	AccessTokenPath       string `yaml:"accessTokenPath"`
	AppID                 string `yaml:"appid"`
	Secret                string `yaml:"secret"`
	TemplateMessagePath   string `yaml:"templateMessagePath"`
	TemplatePaySuccess    string `yaml:"templatePaySuccess"`
	TemplateRefundSuccess string `yaml:"templateRefundSuccess"`
	TemplateRefundRefuse  string `yaml:"templateRefundRefuse"`
}

type Sms struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}

type Smtp struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

// InitConf 读取yaml文件 获取配置, 常用于 func init() 中
func InitConf(path string) {
	workPath, _ := os.Getwd()
	appConfigPath := filepath.Join(workPath, path)
	if !utils.FileExists(appConfigPath) {
		panic("【启动失败】 未找到配置文件!" + appConfigPath)
	}
	log.Info("[启动]读取配置文件:", appConfigPath)
	//读取yaml文件到缓存中
	config, err := ioutil.ReadFile(appConfigPath)
	if err != nil {
		panic("【启动失败】" + err.Error())
	}

	Conf.YamlPath = path
	Conf.YamlData = make(map[string]interface{})
	err = yaml.Unmarshal(config, Conf.YamlData)
	if err != nil {
		panic("【启动失败】" + err.Error())
	}
	Conf.Default = &DefaultConf{}
	err = yaml.Unmarshal(config, Conf.Default)
	if err != nil {
		panic("【启动失败】" + err.Error())
	}
	if Conf.Default.Jwt == nil {
		Conf.Default.Jwt = &Jwt{}
	}
	if Conf.Default.Jwt.Secret == "" {
		Conf.Default.Jwt.Secret = "default secret"
	}
	if Conf.Default.Jwt.Expire == 0 {
		Conf.Default.Jwt.Expire = 3600 * 24 * 7 // 默认7天
	}
}

// YamlGet :: 区分 每一级
func YamlGet(key string) (interface{}, bool) {
	var (
		d  interface{}
		ok bool
	)
	keyList := strings.Split(key, "::")
	temp := make(map[string]interface{})
	temp = Conf.YamlData
	for _, v := range keyList {
		d, ok = temp[v]
		if !ok {
			break
		}
		temp = utils.AnyToMap(d)
	}
	return d, ok
}

// GetString  获取配置,配置不存在返回 "", false
// ex: conf.GetString("jwt::secret")
func GetString(key string) (string, bool) {
	data, ok := YamlGet(key)
	if ok {
		return utils.AnyToString(data), true
	}
	return "", false
}

// GetInt64  获取配置,配置不存在返回  0, false
// ex: conf.GetInt64("jwt::secret")
func GetInt64(key string) (int64, bool) {
	data, ok := YamlGet(key)
	if ok {
		return utils.AnyToInt64(data), true
	}
	return 0, false
}

// YamlGetString  区分
func YamlGetString(key string) (string, error) {
	data, ok := YamlGet(key)
	if ok {
		return utils.AnyToString(data), nil
	}
	return "", fmt.Errorf("配置文件没有找到参数 %s", key)
}

func YamlGetInt64(key string) (int64, error) {
	data, ok := YamlGet(key)
	if ok {
		return utils.AnyToInt64(data), nil
	}
	return 0, fmt.Errorf("配置文件没有找到参数 %s", key)
}
