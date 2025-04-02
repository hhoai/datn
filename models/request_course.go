package models

type RequestCourse struct {
	RequestCourseID uint32 `json:"request_course_id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID          uint32 `json:"user_id"`
	User            User   `gorm:"foreignKey:UserID;references:UserID"`
	CourseID        uint32 `json:"course_id"`
	Course          Course `gorm:"foreignKey:CourseID;references:CourseID"`
}
