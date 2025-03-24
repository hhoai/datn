package repo

import (
	"github.com/joho/godotenv"
	"github.com/zetamatta/go-outputdebug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var DB *gorm.DB

func OutPutDebug(err string) {
	outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [LMS]: ----" + err)
}

func OutPutDebugError(err string) {
	outputdebug.String(time.Now().Format("02-01-2006 15:04:05") + " [LMS]: ERROR: ----" + err)
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		OutPutDebugError(err.Error())
	}
}

func ConnectDB() (*gorm.DB, error) {
	LoadEnvVariables()
	var err error
	dsn := os.Getenv("USER_DB") + ":" + os.Getenv("PASS_DB") + `@tcp(` + os.Getenv("HOST_DB") + `:` + os.Getenv("PORT_DB") + `)/` + os.Getenv("NAME_DB") + `?charset=utf8mb4&parseTime=True&loc=Local`

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		OutPutDebugError(err.Error())
		//return nil, err
	}
	return DB, nil
}
