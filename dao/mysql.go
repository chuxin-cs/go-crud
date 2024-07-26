package dao

//
//import (
//	"fmt"
//	"gopkg.in/yaml.v2"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"io/ioutil"
//)
//
//// DRIVER 指定数据库驱动
//const DRIVER = "mysql"
//
//var db *gorm.DB
//
//// 定义 config 结构体
//type config struct {
//	Url      string `yaml:"url"`
//	UserName string `yaml:"userName"`
//	Password string `yaml:"password"`
//	DbName   string `yaml:"dbName"`
//	Port     string `yaml:"port"`
//}
//
//// GetConfig 获取 application.yaml 配置文件 并且将内容注入到 config 中
//func (c *config) getConfig() *config {
//	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
//	//若出现错误，打印错误提示
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	err = yaml.Unmarshal(yamlFile, c)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	return c
//}
//
//func InitMysql() (err error) {
//	var c config
//	var conf = c.getConfig()
//	// 将yaml配置参数拼接成连接数据库的url
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		conf.UserName,
//		conf.Password,
//		conf.Url,
//		conf.Port,
//		conf.DbName,
//	)
//
//	// 连接数据库
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//	return db
//}
//
//func Close() {
//	db.Close()
//}
