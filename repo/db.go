package repo

import (
	"github.com/joho/godotenv"
	"github.com/zetamatta/go-outputdebug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lms/models"
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

func MigrateDB() error {
	if err := DB.AutoMigrate(&models.UserProgram{}); err != nil {
		OutPutDebugError(err.Error())
	}
	if err := DB.AutoMigrate(&models.UserLesson{}); err != nil {
		OutPutDebugError(err.Error())
	}
	if err := DB.AutoMigrate(&models.Comment{}); err != nil {
		OutPutDebugError(err.Error())
	}
	//if err := DB.AutoMigrate(&models.RequestCourse{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.FileAssignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.UserAssignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Role{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Permission{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.RolePermission{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.CourseCategory{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.TypeUser{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Level{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Program{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Challenge{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Skill{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.TypeAssignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.TypeQuestion{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Assignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.CourseUser{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Course{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.FilePost{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Lesson{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.LessonCategory{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.Option{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		OutPutDebugError(err.Error())
	}
	//if err := DB.AutoMigrate(&models.Question{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.LessonCategory{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.UserAnswer{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	if err := DB.AutoMigrate(&models.FileNews{}); err != nil {
		OutPutDebugError(err.Error())
	}

	if err := DB.AutoMigrate(&models.News{}); err != nil {
		OutPutDebugError(err.Error())
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		OutPutDebugError(err.Error())
	}
	//if err := DB.AutoMigrate(&models.QuestionAssignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	if err := DB.AutoMigrate(&models.Topic{}); err != nil {
		OutPutDebugError(err.Error())
	}
	//if err := DB.AutoMigrate(&models.TopicAssignment{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.TopicQuestion{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	//if err := DB.AutoMigrate(&models.UserOption{}); err != nil {
	//	OutPutDebugError(err.Error())
	//}
	if err := DB.AutoMigrate(&models.Banner{}); err != nil {
		OutPutDebugError(err.Error())
	}
	if err := DB.AutoMigrate(&models.Feedback{}); err != nil {
		OutPutDebugError(err.Error())
	}
	if err := DB.AutoMigrate(&models.Faqs{}); err != nil {
		OutPutDebugError(err.Error())
	}
	return nil
}
