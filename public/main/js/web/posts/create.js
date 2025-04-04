
$(document).ready(function () {

    var lessonID = $('input[name="lesson_id"]').val()

    $(".create-assignment-modal").on("submit", function (event) {
        event.preventDefault();

        var formData = {
            "lesson_id": lessonID,
            "post_title": $('#title').val(),
            "post_body": $('#post-body').val(),
        }
        $.ajax({
            url: "/api/v1/posts",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify(formData),
        })
            .done(function (response) {
                if (response.code === 200) {
                    // window.location.href = "/lessons/" + lessonID + "/posts";
                    $('#postTable').DataTable().ajax.reload();
                    $('#title').val("");
                    $('#post-body').val("");
                    $(".create-assignment-modal").modal("hide");
                }

            })
            .fail(function (xhr) {
                showModalError(xhr.responseJSON.message)
            });
    });
});
