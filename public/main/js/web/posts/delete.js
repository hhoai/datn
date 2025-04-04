$(document).ready(function () {
    // delete
    var lessonID = $('input[name="lesson_id"]').val();
    $('#postTable').on('click', '.delete-btn', function () {
        var postID = $(this).data('id');
        $('.delete-confirm-modal').modal('show');

        $('#confirmDelete').on('click', function () {
            $.ajax({
                url: "/api/v1/posts/" + postID,
                method: "DELETE",
                success: function (response) {
                    // window.location.href = "/lessons/" + lessonID + "/posts";
                    $('.delete-confirm-modal').modal('hide');
                    $('#postTable').DataTable().ajax.reload();
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                },
            });
        })
    });
})