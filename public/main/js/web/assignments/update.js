
function getFiles(assignmentID) {
    // Load Roles
    $.get("/api/v1/file-assignments/" + assignmentID, function (response) {
        const divFiles = $('#files');

        divFiles.empty()
        response.data.forEach(e => {
            divFiles.append(`<p style="padding: 2px 0"><span data-id="${e.FileAssignmentID}" class="item-assignment"><i class="fa fa-trash color-trash" aria-hidden="true"></i></span> ${e.file_name}</p>`);
        });
    });
}


let editorInstance; // Khai báo biến toàn cục để lưu trữ instance của CKEditor

// Khởi tạo CKEditor (nếu chưa khởi tạo)
function initializeEditor(content) {
    if (!editorInstance) {
        // Khởi tạo CKEditor
        ClassicEditor
            .create(document.querySelector('#body'))
            .then(editor => {
                editorInstance = editor;
                editorInstance.setData(content); // Gán nội dung ban đầu
            })
            .catch(error => {
                console.error('Error initializing CKEditor:', error);
            });
    } else {
        // CKEditor đã tồn tại, chỉ cần cập nhật nội dung
        editorInstance.setData(content);
    }
}


$(document).ready(function () {
    $('#assignmentTable').on('click', '.edit-btn', function () {
        assignmentID = $(this).data('id');
        $.ajax({
            url: "/api/v1/assignments/details/" + assignmentID,
            method: "GET",
            success: function (response) {
                var assignment = response.data.assignment
                const DateTime = assignment.due_date.slice(0, 16);
                $('input[name="title"]').val(assignment.title)
                initializeEditor(assignment.assignment_body);
                getFiles(assignmentID)
                $('#due_date').val(DateTime)
                $('#score').val(assignment.score)
                $('#edit_type_assignment_id').val(assignment.TypeAssignment.type_assignment_id)
                $('input[name="assignmentId"]').val(assignmentID)
                $('.modal-fill').modal('show');

            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            },
        });

    });


    $("#save-assignment").on("click", function (event) {
        event.preventDefault();

        const fileInput = $('#file')[0];
        const file = fileInput.files[0];
        const formData = new FormData();
        formData.append("title", $('#title-assignment').val());
        formData.append("body", editorInstance.getData());
        formData.append("file", file);
        formData.append("due_date", $('#due_date').val());
        formData.append("score", $('#score').val());
        formData.append("type_assignment_id", parseInt($('#edit_type_assignment_id').val()));

        var assignmentID = $('input[name="assignmentId"]').val();
        var lessonID = $('input[name="lesson_id"]').val();

        $('.update-confirm-modal').modal("show");

        $('#confirmUpdate').on('click', function () {
            $.ajax({
                url: "/api/v1/assignments/upload/" + assignmentID,
                method: "PUT",
                data: formData,
                contentType: false,
                processData: false,
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


    $('#files').on('click', '.color-trash', function () {
        const icon = $(this);
        const assignmentId = icon.closest('.item-assignment').data('id'); // Get the `data-id` value
        const paragraph = icon.closest('p'); // Get the parent <p> element

        $('.delete-confirm-modal').modal('show');

        $('#confirmDelete').on('click', function () {
            $.ajax({
                url: '/api/v1/file-assignments/' + assignmentId,
                type: 'DELETE',
                success: function (response) {
                    if (response.code === 200) {
                        paragraph.remove();
                    } else {
                    }
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                }
            });
        })


    });
})