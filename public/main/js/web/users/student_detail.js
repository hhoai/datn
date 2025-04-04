$(document).ready(function () {
    const userId = getId();
    const apiUrl = '/api/v1/users/' + userId + '/courses';

    $.ajax({
        url: apiUrl,
        method: 'GET',
        dataType: 'json',
        success: function (response) {
            if (response && response.data) {
                response.data.forEach(course => {
                    const courseHtml = `
                    <div  data-course-id = "${course.Course.course_id}" class="col-xl-4 col-md-6 col-12 course-card">
                        <a href="/users/${userId}/courses/${course.Course.course_id}/lessons" class="box bg-secondary-light pull-up" style="background-image: url(/images/svg-icon/color-svg/st-1.svg); background-position: right bottom; background-repeat: no-repeat;">
                            <div class="box-body">
                                <div class="flex-grow-1">
                                    <div class="d-flex align-items-center pr-2 justify-content-between">
                                        <div class="d-flex">
                                            ${course.Course.status ? '<span class="badge badge-primary mr-15">Active</span>' : '<span class="badge badge-danger mr-15">No active</span>'}
                                            <span class="badge badge-primary mr-5"><i class="fa fa-lock"></i></span>
                                            <span class="badge badge-primary"><i class="fa fa-clock-o"></i></span>
                                        </div>
                                    </div>
                                    <h4 class="mt-25 mb-5">${course.Course.title}</h4>
                                    <div class="progress progress-xs" style="border-radius: 5px; height: 8px;">
                                    </div>
                                    <p class="text-fade mb-0 font-size-12">45 Days Left</p>
                                     <p class="lesson-progress"></p>
                                    <p class="status"></p>
                                </div>
                            </div>
                        </a>
                    </div>`;

                    $("#courses").append(courseHtml);
                    fetchCourseStatistics(userId, course.Course.course_id);

                });
            }
        },
        error: function (xhr, status, error) {
            console.error("Error fetching courses:", error);
        }
    });
    $.ajax({
        url: `/api/v1/users/${userId}/course_process`,
        method: 'GET',
        dataType: 'json',
        success: function (processResponse) {
            if (processResponse.code === 200 && processResponse.data) {
                const completedCourses = processResponse.data;
                $(".course-card").each(function () {
                    const courseID = $(this).data("course-id");
                    const matchingCourse = completedCourses.find(course => course.course_id === courseID);

                    if (matchingCourse) {
                        $(this).find(".status").text("Completed").css("color", "green");
                    } else {
                        $(this).find(".status").text("Incomplete").css("color", "red");
                    }
                });
            } else {
                console.error("Failed to fetch completed lessons data");
            }
        },
        error: function (xhr, status, error) {
            console.error("Error occurred while fetching process data:", error);
        }
    });
    function fetchCourseStatistics(userId,courseId) {
        $.ajax({
            url: `/api/v1/users/${userId}/courses/${courseId}/statisticsCourse`,
            method: 'GET',
            dataType: 'json',
            success: function (statisticsResponse) {
                if (statisticsResponse.code === 200 && statisticsResponse.data) {
                    const { completed_lessons, total_lessons } = statisticsResponse.data;
                    $(`[data-course-id="${courseId}"]`).find(".lesson-progress").text(`${completed_lessons}/${total_lessons} Lessons`);
                } else {
                    console.error(`Failed to fetch statistics for course ID ${courseId}`);
                }
            },
            error: function (xhr, status, error) {
                console.error(`Error occurred while fetching statistics for course ID ${courseId}:`, error);
            }
        });
    }


});

function getId() {
    var url = window.location.href;
    var parts = url.split('/');
    var userIdIndex = parts.indexOf("users") + 1;
    return parts[userIdIndex];
}
