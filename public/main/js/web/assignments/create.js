
$(document).ready(function () {

    var lessonID = $('input[name="lesson_id"]').val()

    $(".create-assignment-modal").on("submit", function (event) {
        event.preventDefault();

        var formData = {
            "lesson_id": lessonID,
            "title": $('#title').val(),
            "description": $('#description').val(),
            "type_assignment_id": $('#type_assignment_id').val()
        }

        $.ajax({
            url: "/api/v1/assignments",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(formData),
        })
            .done(function (response) {
                if (response.code === 200) {
                    window.location.href = "/lessons/" + lessonID + "/assignments";
                }

            })
            .fail(function (xhr) {
                showModalError(xhr.responseJSON.message);
            });
    });
});
