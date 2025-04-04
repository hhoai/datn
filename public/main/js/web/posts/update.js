function getFiles(assignmentID) {
    $.get("/api/v1/file-posts/" + assignmentID, function (response) {
        const divFiles = $('#files');
        divFiles.empty();
        response.data.forEach(e => {
            divFiles.append(`<p style="padding: 5px 0; margin: 0" class="file-${e.FilePostID}">
                <span data-id="${e.FilePostID}" class="item-post">
                <i class="fa fa-trash color-trash" aria-hidden="true"></i>
                </span> ${e.file_name}
                </p>
            `);
            if (e.default) {
                $(`.file-${e.FilePostID}`).append(`<input type="checkbox" class="file-checkbox" value="${e.FilePostID}" id="display-${e.FilePostID}" checked/>
                <label for="display-${e.FilePostID}" style="padding-left: 30px; margin-left:40px ">Display</label>`)
            }
            else  {
                $(`.file-${e.FilePostID}`).append(` <input type="checkbox" class="file-checkbox" value="${e.FilePostID}" id="display-${e.FilePostID}"/>
                <label for="display-${e.FilePostID}" style="padding-left: 30px; margin-left:40px ">Display</label>`)
            }
        });
    });
}

let editorInstance; // Khai báo biến toàn cục để lưu trữ instance của CKEditor

// Khởi tạo CKEditor (nếu chưa khởi tạo)
function initializeEditor(content) {
    if (!editorInstance) {
        // Khởi tạo CKEditor
        ClassicEditor
            .create(document.querySelector('#post-body-u'))
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

    $('#postTable').on('click', '.edit-btn', function () {
        var assignmentID = $(this).data('id');
        $.ajax({
            url: "/api/v1/posts/details/" + assignmentID,
            method: "GET",
            success: function (response) {
                var post = response.data.post
                $('input[name="post-title"]').val(post.post_title)
                // $('#post-body-u').val(post.post_body)
                initializeEditor(post.post_body);
                getFiles(assignmentID)
                $('input[name="postID"]').val(assignmentID)
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

        var postID = $('input[name="postID"]').val()
        var lessonID = $('input[name="lesson_id"]').val()

        $('.file-checkbox:checked').each(function () {
            formData.append('fileIds[]', $(this).val()); // Thêm mảng fileIds vào FormData
        });

        formData.append("title",$('#title-body').val());
        formData.append("body", editorInstance.getData());
        // formData.append("body",$('#post-body-u').val());

        formData.append("file",file);

        $('.update-confirm-modal').modal("show");

        $('#confirmUpdate').on('click', function () {
            $.ajax({
                url: "/api/v1/posts/" + postID,
                method: "PUT",
                data: formData,
                contentType: false, // Important for file uploads
                processData: false,
                timeout: 300000,
            })
            .done(function (response) {
                if (response.code === 200) {
                    // window.location.href = "/lessons/" + lessonID + "/posts";
                    $('#postTable').DataTable().ajax.reload();
                    $('.modal-fill').modal('hide');
                    $('.update-confirm-modal').modal("hide");
                }

            })
            .fail(function (xhr) {
                showModalError(xhr.responseJSON.message);
            });
        });


    })
    $('#files').on('click', '.color-trash', function () {
        const icon = $(this);
        const postId = icon.closest('.item-post').data('id'); // Get the `data-id` value
        const paragraph = icon.closest('p'); // Get the parent <p> element

        $('.delete-confirm-modal').modal('show');

        $('#confirmDelete').on('click', function () {
            $.ajax({
                url: '/api/v1/file-posts/' + postId,
                type: 'DELETE',
                success: function (response) {
                    if (response.code === 200) {
                        $('.delete-confirm-modal').modal('hide');
                        paragraph.remove();
                        showModalSuccess('Answer updated successfully');
                    }
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                }
            });
        })


    });
})