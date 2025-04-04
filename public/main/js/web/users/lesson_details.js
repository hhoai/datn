
$(document).ready(function () {
    const url = window.location.href.split('/');
    const userId = url[url.indexOf('users') + 1];
    const courseId = url[url.indexOf('courses') + 1];
    const apiUrl = `/api/v1/users/${userId}/courses/${courseId}/lessons`;

    $.ajax({
        url: apiUrl,
        method: 'GET',
        dataType: 'json',
        success: function (response) {
            if (response && response.data) {
                response.data.forEach(lesson => {
                    const lessonHtml = `
                    <div data-lesson-id="${lesson.LessonID}">
                        <div class="bg-primary-light px-20 py-40 rounded20 mb-20 lesson-card" style="max-height: 132px; margin-right: 20px">
                            <a class="flexbox flex-grow gap-items text-primary" href="/users/${userId}/lessons/${lesson.LessonID}/assignments" data-toggle="quickview">
                                <span class="icon-Equalizer d-block font-size-40">
                                    <span class="path1"></span><span class="path2"></span><span class="path3"></span><span class="path4"></span>
                                </span>
                                <div class="media-body">
                                    <div class="d-flex">
                                        <b>${lesson.title}</b>
                                        <span> 
                                            <small class="sidetitle">
                                                <span data-i18n="level"></span>: ${lesson.Level.name}
                                            </small>
                                        </span>
                                    </div>
                                    <div style="display: flex; height: 25px">
                                        <div>
                                            <p style="color: #0052cc">
                                            <span data-i18n="category"></span>: ${lesson.LessonCategory.name}
                                            </p>
                                        </div>
                                        <div style="padding-left: 30px">
                                            <p class="assignment-progress" style="background-color: #00ec8d"></p> 
                                         </div>                                        
                                    </div>
                                    <p class="status"></p> 
                                </div>
                            </a>
                        </div>
                    </div>`;
                    $("#lessons").append(lessonHtml);
                    fetchLessonStatistics(userId,lesson.LessonID)

                });

                $.ajax({
                    url: `/api/v1/users/${userId}/lesson_process`,
                    method: 'GET',
                    dataType: 'json',
                    success: function (processResponse) {
                        if (processResponse.code === 200 && processResponse.data) {
                            const completedLessons = processResponse.data;
                            $(".lesson-card").each(function () {
                                const lessonID = $(this).parent().data("lesson-id");
                                const matchingLesson = completedLessons.find(lesson => lesson.lesson_id === lessonID);

                                if (matchingLesson) {
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
            }
        },
        error: function (xhr, status, error) {
            console.error("Error fetching lessons:", error);
        }
    });

    function fetchLessonStatistics(userId, lessonId) {
        $.ajax({
            url: `/api/v1/users/${userId}/lessons/${lessonId}/statisticsLesson`,
            method: 'GET',
            dataType: 'json',
            success: function (statisticsResponse) {
                if (statisticsResponse.code === 200 && statisticsResponse.data) {
                    const { completedAssignments, totalAssignments } = statisticsResponse.data;
                    $(`[data-lesson-id="${lessonId}"]`).find(".assignment-progress")
                        .text(`${completedAssignments}/${ totalAssignments} Assignment`);


                } else {
                    console.error(`Failed to fetch statistics for lessonID ${lessonId}`);
                }
            },
            error: function (xhr, status, error) {
                console.error(`Error occurred while fetching statistics for course ID ${lessonId}:`, error);
            }
        });
    }

});
