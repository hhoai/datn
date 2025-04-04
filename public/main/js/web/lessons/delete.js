$(document).ready(function () {
    var courseID = $('input[name="course_id"]').val();
    $('#lessonTable').on('click', '.delete-btn', function () {
        lessonID = $(this).data('id');

        $('.delete-confirm-modal').modal('show');

        $('#confirmDelete').on('click', function () {
            $.ajax({
                url: "/api/v1/lessons/" + lessonID,
                method: "DELETE",
                success: function (response) {
                    window.location.href = "/courses/" + courseID;
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                },
            });
        })
    });
})