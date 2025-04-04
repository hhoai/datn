
$(document).ready(function () {
    $("#create-lesson-form").on("submit", function (event) {
        event.preventDefault();

        var courseID = $('input[name="course_id"]').val()

        var formData = {
            "title": $('input[name="title"]').val(),
            "level_id": $('#level_id').val(),
            "course_id": courseID,
            "lesson_category_id": $('#lesson_category_id').val()
        }

        $.ajax({
            url: "/api/v1/lessons",
            method: "POST",
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
    });


});
