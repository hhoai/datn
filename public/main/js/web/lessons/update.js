
function loadDropdowns(selectedLessonCategoryID) {
    // Load Roles
    $.get("/api/v1/lesson-categories", function (response) {
        const roleSelect = $('#lesson_category_id-u');
        roleSelect.empty();

        response.data.forEach(e => {
            roleSelect.append(`<option value="${e.lesson_category_id}" ${e.lesson_category_id == selectedLessonCategoryID ? 'selected' : ''}>${e.name}</option>`);
        });
    });

}

function loadLevelDropdowns(selectedLevelID) {
    // Load Roles
    $.get("/api/v1/levels", function (response) {
        const roleSelect = $('#lesson_level_id-u');
        roleSelect.empty();

        response.data.forEach(e => {
            roleSelect.append(`<option value="${e.level_id}" ${e.level_id == selectedLevelID ? 'selected' : ''}>${e.name}</option>`);
        });
    });

}

$(document).ready(function () {
    $('#lessonTable').on('click', '.edit-btn', function () {
        lessonID = $(this).data('id');
        $.ajax({
            url: "/api/v1/lessons/details/" + lessonID,
            method: "GET",
            success: function (response) {
                var lesson = response.data.lesson
                $('input[name="lesson_title"]').val(lesson.title)
                $('input[name="lesson_level"]').val(lesson.level)
                loadDropdowns(lesson.lesson_category_id)
                loadLevelDropdowns(lesson.level_id);
                $('input[name="lesson_id"]').val(lesson.LessonID)
                $('.setting-lesson-modal').modal('show');
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            },
        });

    });



    $("#update-lessons").on("click", function (event) {
        event.preventDefault();

        var formData = {
            "title": $('input[name="lesson_title"]').val(),
            "level_id": $('#lesson_level_id-u').val(),
            "lesson_category_id": $('#lesson_category_id-u').val()
        }

        var lessonID = $('input[name="lesson_id"]').val()
        var courseID = $('input[name="course_id"]').val()


        $('.update-confirm-modal').modal("show");

        $('#confirmUpdate').on('click', () => {
            $.ajax({
                url: "/api/v1/lessons/" + lessonID,
                method: "PUT",
                contentType: "application/json",
                data: JSON.stringify(formData),
            })
            .done(function (response) {
                if (response.code === 200) {
                    window.location.href = "/courses/" + courseID;
                }

            })
            .fail(function (xhr) {
                showModalError(xhr.responseJSON.message);
            });
        })
    });
})