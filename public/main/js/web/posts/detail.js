$(document).ready(function () {
    function formatTime (timestamp) {
        // Convert the timestamp to a Date object
        const date = new Date(timestamp);

        // Format the date to "Jan 2, 2006"
        const formatter = new Intl.DateTimeFormat('en-GB', {
            day: 'numeric',
            month: 'numeric',
            year: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            hour12: false
        });
        return formatter.format(date);
    }

    function viewFile(file, src, block){
        const encodedSrc = src.replace(/ /g, '%20');
        const fileExtension = src.split('.').pop().toLowerCase();
        // Get file extension in lowercase
        if (["jpg", "png", "jpeg", "gif"].includes(fileExtension)) {
            block.append(`
        <img src="${encodedSrc}" alt="img" style="width: 100%">
    `);
        }
        else if (file === "mp4" || file === ".webm") {
            block.append(`<video controls width="100%" height="600px"><source src=${encodedSrc} type="video/mp4">Your browser does not support the video tag.</video>`);
        }
        else {
            switch (file) {
                case "pdf":
                    block.append(`<embed src=${encodedSrc} width="100%" height="600px" type="application/pdf">`);
                    break
                case "doc":
                case "docx":
                    const domain = "https://5b28-2402-800-6106-a935-ad3d-c9d0-4399-8331.ngrok-free.app"
                    encodedSrc = "https://view.officeapps.live.com/op/embed.aspx?src=" + domain + encodedSrc
                    block.append(`<iframe src="${encodedSrc}" width="100%" height="600px" frameborder="0"></iframe>`);
                    // block.append(`<iframe
                    //    src="https://docs.google.com/gview?url=${domain}${src}/to/document.doc&embedded=true"></iframe>`)
                    break
                default:
                    block.append(`<p data-i18n="download-to-view-file"></p>`);
                    break
            }
        }
    }

    var postID = $('#PostID').val()

    const apiGetPost = "/api/v1/student_courses/posts/" + postID + "/details"

    $.get(apiGetPost, function (response) {
        var data = response.data
        const lessonID = data.LessonID

        const imgFile = $('#img-file')
        const mp4File = $('#mp4-file')
        const fileContainer = $('#file')
        const postContainer = $('#post-content');
        const commentContent = $("#comment-content")
        const preview = $('#preview')

        const fileDefault = $(".file-default")

        postContainer.empty();

        $('.page-title').append(`: ${data.Title}`)

        postContainer.append(
            `<div class="align-items-center" style="display: flex; gap: 20px">
                <a class="avatar avatar-lg status-success" href="#">
                    <img src="/images/avatar/avatar-10.png" class="bg-success-light" alt="...">
                </a>
                <div class="media-body">
                    <p class="font-size-16">
                        <a class="hover-primary" href="#"><strong>${data.CreatedBy}</strong></a>
                    </p>
                     ${formatTime(data.CreatedAt)}
                </div>
            </div>`);

        $(".post-header").append(`<b style="font-size: 24px">${data.PostTitle}</b>`)

        $(".post-body").append(`<div>${data.PostBody}</div>`)

        fileContainer.empty()
        data.Files.forEach(file => {
            let src = "/posts/" + lessonID + "/" + file.file_name
            const encodedSrc = src.replace(/ /g, '%20');
            let ext = file.file_name.split('.').pop().toLowerCase();
            if (["jpg", "png", "jpeg", "gif"].includes(ext)) {
                imgFile.append(`
                    <div class="media" style="align-items: center; padding: 10px 5px ; width: 400px">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${encodedSrc}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href="${encodedSrc}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                `)
            }
            else {
                if (ext === "mp4" || ext === ".webm") {
                    mp4File.append(`
                    <div class="media" style="align-items: center; padding: 10px 5px ; width: 400px">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${encodedSrc}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href=${encodedSrc} download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                    `)
                }
                else {
                    fileContainer.append(`
                    <div class="media" style="align-items: center; padding: 10px 5px ; width: 400px">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${encodedSrc}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href="/posts/${lessonID}/${file.file_name}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                    `)

                }
            }

            if (file.default) {
                viewFile(ext, src, fileDefault)
            }
        });
        fileContainer.on("click", ".eye", function() {
            preview.empty()
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src, preview);
        });
        mp4File.on("click", ".eye", function() {
            preview.empty()
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src, preview);
        });
        imgFile.on("click", ".eye", function() {
            preview.empty()
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src, preview);
        });

        commentContent.empty()
        data.Comments.forEach((comment, index) => {
            commentContent.append(`<div class="media" id="${index}">
                  <a class="avatar" href="#">
                <img src="/images/avatar/9.jpg" alt="...">
                  </a>
                  <div class="media-body" style="margin: 0">
                        <p>
                          <a href="#"><strong>${comment.User.name}</strong></a>
                          <time class="float-end text-fade">${formatTime(comment.CreatedAt)}</time>
                        </p>
                        <p>${comment.comment_content}</p>
                  </div>
                </div>`)
            if (data.Username === comment.User.name) {
                $(`#${index}`).append(` 
                    <span type="button" data-toggle="dropdown" aria-expanded="true"><span class="glyphicon glyphicon-option-vertical" ></span></span>
                    <div class="dropdown-menu" style="left: -140px;">
                    <button class="delete-comment" data-id="${comment.CommentID}" data-i18n="delete" style="
                    border: none;
                    background: transparent;
                    padding: 0 20px;
                    color: #2a2a2a;"></button>
                    </div>
                `)
            }
        })

        $("#send-comment").on("click", function(){
            var comment = {
                "comment_content" : $("#message").val(),
                "post_id": $('#PostID').val()
            }
            $.ajax({
                url: "/api/v1/comments",
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify(comment),
                success: function (respone) {
                    let data = respone.data
                    commentContent.empty()
                    data.Comments.forEach((comment, index) => {
                        commentContent.append(`<div class="media" id="${index}">
                              <a class="avatar" href="#">
                            <img src="/images/avatar/9.jpg" alt="...">
                              </a>
                              <div class="media-body" style="margin: 0">
                            <p>
                              <a href="#"><strong>${comment.User.name}</strong></a>
                              <time class="float-end text-fade">${formatTime(comment.CreatedAt)}</time>
                            </p>
                            <p>${comment.comment_content}</p>
                              </div>
                            </div>`)
                        if (data.Username === comment.User.name) {
                            $(`#${index}`).append(` 
                            <span type="button" data-toggle="dropdown" aria-expanded="true"><span class="glyphicon glyphicon-option-vertical" ></span></span>
                            <div class="dropdown-menu" style="left: -140px;">
                            <button class="delete-comment" data-id="${comment.CommentID}" data-i18n="delete" style="
                            border: none;
                            background: transparent;
                            padding: 0 20px;
                            color: #2a2a2a;"></button>
                            </div>
                        `)
                        }
                    })
                    $("#message").val("")
                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                },
            });
        })

        commentContent.on("click", ".delete-comment", function (){
            var commentID = $(this).data('id');
            $('.delete-confirm-modal').modal('show');

            $('#confirmDelete').on('click', function () {
                $('.delete-confirm-modal').modal('hide');
                $.ajax({
                    url: "/api/v1/comments/" + commentID,
                    method: "DELETE",
                    success: function (respone) {
                        let data = respone.data
                        commentContent.empty()
                        data.Comments.forEach((comment, index) => {
                            commentContent.append(`<div class="media" id="${index}">
                                  <a class="avatar" href="#">
                                <img src="/images/avatar/9.jpg" alt="...">
                                  </a>
                                  <div class="media-body" style="margin: 0">
                                <p>
                                  <a href="#"><strong>${comment.User.name}</strong></a>
                                  <time class="float-end text-fade">${formatTime(comment.CreatedAt)}</time>
                                </p>
                                <p>${comment.comment_content}</p>
                                  </div>
                                </div>`)
                            if (data.Username === comment.User.name) {
                                $(`#${index}`).append(` 
                                <span type="button" data-toggle="dropdown" aria-expanded="true"><span class="glyphicon glyphicon-option-vertical" ></span></span>
                                <div class="dropdown-menu" style="left: -140px;">
                                <button class="delete-comment" data-id="${comment.CommentID}" data-i18n="delete" style="
                                border: none;
                                background: transparent;
                                padding: 0 20px;
                                color: #2a2a2a;"></button>
                                </div>
                                `)
                            }
                        })
                    },
                    error: function (xhr) {
                        showModalError(xhr.responseJSON.message);
                    },
                });
            })
        })
    });
})