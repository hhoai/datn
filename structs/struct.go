package structs

import "lms/models"

type FormLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FormSignup struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Name       string `json:"name"`
}

type CreateLessonForm struct {
	CourseID           string `json:"course_id"`
	Title              string `json:"title"`
	LevelID            string `json:"level_id"`
	LessonCategoriesID string `json:"lesson_category_id"`
}

type CreateAssignmentForm struct {
	LessonID         string `json:"lesson_id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	TypeAssignmentID string `json:"type_assignment_id"`
}

type UpdateAssignmentForm struct {
	Title            string `json:"title"`
	Body             string `json:"body"`
	File             string `json:"file"`
	DueDate          string `json:"due_date"`
	Score            string `json:"score"`
	TypeAssignmentID uint32 `json:"type_assignment_id"`
}

type FormCreateCourse struct {
	Title                string `json:"title"`
	Description          string `json:"description"`
	LevelID              uint32 `json:"level_id"`
	ProgramID            uint32 `json:"program_id"`
	CourseCategoriesID   uint32 `json:"course_categories_id"`
	Image                string `json:"image"`
	Amount               uint32 `json:"amount"`
	StartTime            string `json:"start_time"`
	EndTime              string `json:"end_time"`
	Status               bool   `json:"status"`
	PrerequisiteCourseID uint32 `json:"prerequisite_course_id"`
}
type FormCreateQuestion struct {
	Content        string `json:"content"`
	TypeQuestionID uint32 `json:"type_question_id"`
	Score          uint32 `json:"score"`
	ProgramID      uint32 `json:"program_id"`
	LevelID        uint32 `json:"level_id"`
	ChallengeID    uint32 `json:"challenge_id"`
	SkillID        uint32 `json:"skill_id"`
}
type CreatePostForm struct {
	LessonID string `json:"lesson_id"`
	Title    string `json:"post_title"`
	Body     string `json:"post_body"`
}

type UpdatePostForm struct {
	Title   string `json:"post_title"`
	Body    string `json:"post_body"`
	File    string `json:"file"`
	Default string `json:"default"`
}

type CreateNewsForm struct {
	Title string `json:"news_title"`
	Body  string `json:"news_body"`
}

type UpdateOptionForm struct {
	OptionContent string `json:"content"`
	IsCorrect     bool   `json:"is_correct"`
}

type Filters struct {
	ProgramID      uint32   `json:"program_id"`
	SkillID        uint32   `json:"skill_id"`
	ChallengeID    uint32   `json:"challenge_id"`
	TypeQuestionID uint32   `json:"type_question_id"`
	LevelID        uint32   `json:"level_id"`
	QuestionIDs    []uint32 `json:"question_ids"`
}

type AnswerResult struct {
	TopicQuestionID uint32                          `json:"topic_question_id"`
	TopicID         uint32                          `json:"topic_id"`
	QuestionID      uint32                          `json:"question_id"`
	TypeQuestionID  uint32                          `json:"type_question_id"`
	Content         string                          `json:"content"`
	Score           uint32                          `json:"score"`
	IsCorrect       bool                            `json:"is_correct"`
	Options         []models.OptionWithoutIsCorrect `json:"options"`
}

type HighchartsData struct {
	Date  string `json:"date"`
	Count uint32 `json:"count"`
}

type CourseDetails struct {
	CourseID        uint32 `json:"course_id"`
	CourseCode      string `json:"course_code"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	LessonCompleted uint32 `json:"lesson_completed"`
	Lesson          uint32 `json:"lesson"`
	Status          bool   `json:"status"`
	Width           string `json:"width"`
}
