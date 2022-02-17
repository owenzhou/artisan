/*
Copyright © 2021 artisan
*/
package cmd

import (
	"os"

	"github.com/owenzhou/artisan/config"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	db      *gorm.DB
)

var rootCmd = &cobra.Command{
	Use:   "github.com/owenzhou/artisan",
	Short: "代码自动生成器",
	Long:  `可自动生成：项目，控制器，模型，事件，监听器等代码，方便开发`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile = "config.yaml"
		//没有文件直接终止
		if _, err := os.Stat(cfgFile); err != nil {
			return
		}
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&config.Config); err != nil {
			panic(err)
		}
	}

	initMysql()
}

func initMysql() {
	var err error
	cfg := config.Config.Mysql
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Dbname + "?charset=" + cfg.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		panic(err)
	}
}
