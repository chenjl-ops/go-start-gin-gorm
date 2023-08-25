package nacos

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
)

const (
	runEnvKey    = "RUNTIME_ENV"
	runGroupKey  = "RUNTIME_GROUP"
	runTenantKey = "RUNTIME_TENANT"
	// runGROUPIdKey = "RUNTIME_GROUPID"
	appNameKey = "RUNTIME_APP_NAME"
)

var Config *Specification

// GetEnv 获取当前运行环境
func GetEnv() string {
	env := os.Getenv(runEnvKey)
	return env
}

// GetAppName 获取当前运行应用名称
func GetAppName() string {
	appName := os.Getenv(appNameKey)
	return appName
}

// GetGroup 获取当前group
func GetGroup() string {
	groupName := os.Getenv(runGroupKey)
	if groupName == "" {
		groupName = "DEFAULT_GROUP"
	}
	return groupName
}

// GetTenant 获取当前namespace
func GetTenant() string {
	tenant := os.Getenv(runTenantKey)
	if tenant == "" {
		tenant = "public"
	}
	return tenant
}

// GetNacosPath 获取nacos路径
func GetNacosPath() string {
	return fmt.Sprintf("/data/nacos")
}

func GetNacosUrl() string {
	url := fmt.Sprintf("http://nacos.configserver.com")
	return url
}

func NewNacos() (nacos *Nacos, err error) {
	appName := GetAppName()
	if appName == "" {
		return nil, errors.Errorf("Env appName Has Empty")
	}
	env := GetEnv()
	if env == "" {
		return nil, errors.Errorf("Env env Has Empty")
	}
	tenant := GetTenant()
	if tenant == "" {
		return nil, errors.Errorf("Env tenant Has Empty")
	}
	group := GetGroup()
	if group == "" {
		return nil, errors.Errorf("Env group Has Empty")
	}
	//nacosServerUrl := GetNacosUrl()
	nacos = &Nacos{
		Tenant: tenant,
		Group:  group,
		DataId: appName,
		//NacosServerUrl: nacosServerUrl,
	}
	return nacos, nil
}

func GetNacosConfigs() (nacosClient *constant.ClientConfig, nacosServer *constant.ServerConfig, err error) {
	appName := GetAppName()
	if appName == "" {
		return nil, nil, errors.Errorf("Env appName Has Empty")
	}
	env := GetEnv()
	if env == "" {
		return nil, nil, errors.Errorf("Env env Has Empty")
	}
	tenant := GetTenant()
	if tenant == "" {
		return nil, nil, errors.Errorf("Env tenant Has Empty")
	}
	group := GetGroup()
	if group == "" {
		return nil, nil, errors.Errorf("Env group Has Empty")
	}
	nacosServerUrl := GetNacosUrl()
	nacosPath := GetNacosPath()

	nacosServerConfigs := *constant.NewServerConfig(nacosServerUrl, 8848, constant.WithContextPath(nacosPath))

	nacosClientConfigs := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		//constant.WithUpdateCacheWhenEmpty(true),
	)

	return &nacosClientConfigs, &nacosServerConfigs, nil
}

func NewNacosBySdk() (iClient config_client.IConfigClient, err error) {
	nacosClientConfig, nacosServerConfig, err := GetNacosConfigs()
	if err != nil {
		return nil, err
	}
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  nacosClientConfig,
			ServerConfigs: []constant.ServerConfig{*nacosServerConfig},
		},
	)
}

func ReadRemoteConfig() error {
	return ReadRemoteConfigCustom(nil)
}

func ReadRemoteConfigCustom(input *Nacos) error {
	config, err := NewNacos()
	if err != nil {
		return err
	}
	client, err := NewNacosBySdk()
	if err != nil {
		return err
	}
	if input != nil {
		if input.Group != "" {
			config.Group = input.Group
		}
		if input.DataId != "" {
			config.DataId = input.DataId
		}
		if input.Tenant != "" {
			config.Tenant = input.Tenant
		}
	}

	//fmt.Println("config: ", config)
	//fmt.Println("Get Config Start: ===============")
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: config.DataId,
		Group:  config.Group,
	})
	if err != nil {
		return err
	}
	//fmt.Println("nacos Data: ", content)

	err = json.Unmarshal([]byte(content), &Config)
	//fmt.Println("unmarshal done data: ", Config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//func ReadRemoteConfigCustom(input *Nacos) error {
//	config, err := NewNacos()
//	if err != nil {
//		return err
//	}
//
//	if input != nil {
//		if input.NacosServerUrl != "" {
//			config.NacosServerUrl = input.NacosServerUrl
//		}
//		if input.Group != "" {
//			config.Group = input.Group
//		}
//		if input.DataId != "" {
//			config.DataId = input.DataId
//		}
//		if input.Tenant != "" {
//			config.Tenant = input.Tenant
//		}
//	}
//	v := viper.New()
//	v.SetConfigType("prop")
//	// TODO List
//	// 1、实现nacos方法。 a.通过三方库 b. 自己注册nacos provider的实现方式
//	// 参考： https://github.com/yoyofxteam/nacos-viper-remote/tree/main
//	// 参考： https://github.com/nacos-group/nacos-sdk-go
//	err = v.AddRemoteProvider("nacos", config.NacosServerUrl, "/v1/cs/configs")
//	if err != nil {
//		return err
//	}
//	err = v.ReadRemoteConfig()
//	if err != nil {
//		return err
//	}
//	err = v.Unmarshal(&Config)
//	if err != nil {
//		return nil
//	}
//	fmt.Printf("nacos data: %+v", Config)
//	return nil
//}
