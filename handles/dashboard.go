package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"time"
)

func GetDashboard(c *fiber.Ctx) error {
	repo.OutPutDebug("GetDashboard")
	user := GetInfoUser(c)
	//courseUsers := utilS.CourseUserRepo.FindCourseByUserID(user.UserID, user.TypeUserID)
	//
	//var courses []models.Course
	//// foreach error
	//for i, course := range courseUsers {
	//	if i >= 4 {
	//		break
	//	}
	//	courseTemp, _ := utilS.CourseRepo.FindByID(course.CourseID)
	//	courseTemp.Status = course.Status
	//	courses = append(courses, *courseTemp)
	//}
	//
	//news := utilS.NewsRepo.Find4News()
	//
	//type NewsRs struct {
	//	NewsID    uint32
	//	Title     string
	//	Body      template.HTML
	//	CreatedBy string
	//	CreatedAt time.Time `json:"created_at"`
	//}
	//
	//var newsRs []NewsRs
	//for _, n := range news {
	//	newsRs = append(newsRs, NewsRs{
	//		NewsID:    n.NewsID,
	//		Title:     n.Title,
	//		Body:      template.HTML(n.Body),
	//		CreatedBy: n.User.Name,
	//		CreatedAt: n.CreatedAt,
	//	})
	//}
	//
	//type bannersRs struct {
	//	BannersID   uint32
	//	FileName    string
	//	TypeFile    string
	//	Description string
	//}
	//
	//var rs []bannersRs
	//banners := utilS.BannerRepo.Find3Banner()
	//for _, b := range banners {
	//	rs = append(rs, bannersRs{
	//		BannersID:   b.BannerID,
	//		FileName:    b.FileName,
	//		Description: b.Description,
	//		TypeFile:    filepath.Ext(b.FileName),
	//	})
	//}

	return c.Render("index", fiber.Map{
		"Ctx":      c,
		"Username": user.Name,
		//"Courses":  courses,
		//"News":     newsRs,
		//"Banners":  rs,
	}, "layouts/main")
}

func ApiGetDashboard(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetDashboard")

	countUser, _ := utilS.UserRepo.CountUser()
	countCourse, _ := utilS.CourseRepo.CountCourse()
	countTopic, _ := utilS.TopicRepo.CountTopic()
	countQuestion, _ := utilS.QuestionBankRepo.CountQuestion()

	courses := utilS.CourseRepo.FindAll()

	statsNewUser := utilS.UserRepo.FindNewUsersLast7Days()

	var rs []structs.HighchartsData

	for i := 6; i >= 0; i-- {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		rs = append(rs, structs.HighchartsData{
			Date:  day,
			Count: 0,
		})
	}

	for _, stat := range statsNewUser {
		for i, r := range rs {
			if stat.Date == r.Date {
				//res = append(res, stat)
				rs[i].Count = stat.Count
				break
			}
		}
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "get data dashboard successfully", fiber.Map{
		"CountUser":     countUser,
		"CountCourse":   countCourse,
		"CountTopic":    countTopic,
		"CountQuestion": countQuestion,
		"Courses":       courses,
		"CountNewUser":  rs,
	})
}

func ApiSearchPrograms(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiSearchPrograms")
	search := c.Params("q")

	programs, err := utilS.ProgramRepo.SearchProgram(search)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "search programs error", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "search programs successfully", programs)
}
