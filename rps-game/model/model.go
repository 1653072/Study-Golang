package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"golangdemo/rps-game/configs/db"
	"golangdemo/rps-game/configs/log-conf"
	"golangdemo/rps-game/configs/structs"
	"golangdemo/rps-game/configs/system-code"
	"golangdemo/rps-game/configs/system-path"
	"golangdemo/rps-game/helpers/logging"
)

/*
	==> [GORM] LƯU Ý VỀ CÁCH ĐẶT TÊN BIẾN TRONG "type struct" ĐỂ MAPPING VỚI TABLE TRONG DATABASE <==
	+ Nếu tên biến trong struct là "StartDate" thì dưới database tên sẽ là start_date. Nếu sử dụng
	GORM với Raw Query, thì tên biến/bảng trong câu truy vấn phải là "start_date", tên biến để hứng
	kết quả truy vấn sẽ là "StartDate".
	+ Nếu tên biến trong struct là "Startdate" thì dưới database tên sẽ là startdate.
	+ Nếu tên struct được đặt là số ít (Account) thì sau khi AutoMigrate xuống database, tên struct
	sẽ trở thành tên table và tên dưới dạng số nhiều (accounts).
*/

var MysqlInstance *gorm.DB = nil

func loadDbInfo() (*structs.MysqlConfigStruct, error) {
	mysqlConfigPath := SystemPath.MysqlFilePath
	viper.SetConfigFile(mysqlConfigPath)
	fErr := viper.ReadInConfig()
	if fErr != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Error, SystemCode.ErrLoadFileFail, fErr.Error()))
		return nil, SystemCode.ErrLoadFileFail
	}

	mysqlConfigFile := structs.MysqlConfigStruct{}
	mysqlConfigFile.DbName = fmt.Sprintf("%s", viper.GetString("dbName"))
	mysqlConfigFile.Host = fmt.Sprintf("%s", viper.GetString("host"))
	mysqlConfigFile.Port = fmt.Sprintf("%s", viper.GetString("port"))
	mysqlConfigFile.User = fmt.Sprintf("%s", viper.GetString("user"))
	mysqlConfigFile.Password = fmt.Sprintf("%s", viper.GetString("password"))
	return &mysqlConfigFile, nil
}

func autoMigrateDb() {
	MysqlInstance.AutoMigrate(&Account{}, &Game{}, &Gameturn{})
}

func configDbConn() {
	MysqlInstance.DB().SetMaxOpenConns(db.MysqlMaxOpenConns)
	MysqlInstance.DB().SetMaxIdleConns(db.MysqlMaxIdleConns)
	MysqlInstance.DB().SetConnMaxLifetime(db.MysqlConnMaxLifetime)
}

func InitializeModel() error {
	mysqlConfigFile, loadError := loadDbInfo()
	if loadError != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Error, SystemCode.ErrConnectToDbFail, "Connect to MySQL database fail"))
		return SystemCode.ErrConnectToDbFail
	}
	connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfigFile.User, mysqlConfigFile.Password, mysqlConfigFile.Host, mysqlConfigFile.DbName)

	var conErr error
	MysqlInstance, conErr = gorm.Open("mysql", connectionString)
	if conErr != nil || MysqlInstance.DB().Ping() != nil {
		logging.SysLog.Println(logging.FormatResult(LogConf.Error, SystemCode.ErrConnectToDbFail, "Connect to MySQL database fail"))
		return SystemCode.ErrConnectToDbFail
	}

	configDbConn()
	autoMigrateDb()
	logging.SysLog.Println(logging.FormatResult(LogConf.Info,nil, "Connect to MySQL database successfully"))
	return nil
}

func CloseModel() {
	if MysqlInstance != nil {
		_ = MysqlInstance.Close()
	}
}