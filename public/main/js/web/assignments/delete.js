$(document).ready(function () {
    var lessonID = $('input[name="lesson_id"]').val();
    $('#assignmentTable').on('click', '.delete-btn', function () {
        var assignmentID = $(this).data('id');
        $('.delete-confirm-modal').modal('show');

        $('#confirmDelete').on('click', function () {
            $.ajax({
                url: "/api/v1/assignments/" + assignmentID,
                method: "DELETE",
                success: function () {
                    window.location.href = "/lessons/" + lessonID + "/assignments";
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                },
            });
        })
    });
})