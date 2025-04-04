$(document).ready(function () {
    var url = window.location.pathname;
    var assignmentID = url.split('/')[3];
    const courseID = getId('courses');
    console.log(courseID);
    const apiGetAssignment = "/api/v1/student_courses/courses/" + courseID + "/assignments/" + assignmentID + "/details"
    const imgFile = $('#img-file')
    const mp4File = $('#mp4-file')
    const fileContainer = $('#file')

    let userAssignmentID = 0;

    function formatTime(timestamp) {
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

    function viewFile(file, src) {
        const encodedSrc = src.replace(/ /g, '%20');
        const fileExtension = src.split('.').pop().toLowerCase();
        const preview = $('#preview')
        // Get file extension in lowercase
        preview.empty()
        if (["jpg", "png", "jpeg", "gif"].includes(fileExtension)) {
            preview.append(`
        <img src="${encodedSrc}" alt="img" style="width: 800px">
    `);
        } else if (file === "mp4" || file === "webm") {
            preview.append(`<video controls width="100%" height="600px"><source src=${encodedSrc} type="video/mp4">Your browser does not support the video tag.</video>`);
        } else {
            switch (file) {
                case "pdf":
                    preview.append(`<embed src=${encodedSrc} width="100%" height="600px" type="application/pdf">`);
                    break
                case "doc":
                case "docx":
                    const domain = "https://5b28-2402-800-6106-a935-ad3d-c9d0-4399-8331.ngrok-free.app"
                    encodedSrc = "https://view.officeapps.live.com/op/embed.aspx?src=" + domain + encodedSrc
                    preview.append(`<iframe src="${encodedSrc}" width="100%" height="600px" frameborder="0"></iframe>`);
                    // preview.append(`<iframe
                    //    src="https://docs.google.com/gview?url=${domain}${src}/to/document.doc&embedded=true"></iframe>`)
                    break
                default:
                    preview.append(`<p data-i18n="download-to-view-file"></p>`);
                    break
            }
        }
    }

    $.get(apiGetAssignment, function (response) {
        var data = response.data
        console.log(data)
        $("#page-title").text(data.Lesson.title);
        userAssignmentID = data.UserAssignment.user_assignment_id;
        $("#assignment-content").append(
            `<div class="media align-items-center">
                <a class="avatar avatar-lg status-success" href="#">
                    <img src="/images/avatar/avatar-10.png" class="bg-success-light" alt="...">
                </a>
                <div class="media-body">
                    <p class="font-size-16">
                        <a class="hover-primary" href="#"><strong>${data.Assignment.User.name}</strong></a>
                    </p>
                     ${formatTime(data.Assignment.created_at)}
                </div>
                <div class="media-right">
                    <span class="badge badge-danger-light" style="font-size: 14px;
                                padding: 10px 15px;
                                border-radius: 8px;">${formatTime(data.Assignment.due_date)}</span>
                </div>
                </div>
                <div style="padding: 5px 105px;">
                    <b>${data.Assignment.title}</b>
                    <p>${data.Assignment.assignment_body} </p>
                </div>
            `);
        fileContainer.empty()
        data.FileAssignment.forEach(file => {
            let src = "/assignments/" + data.Assignment.lesson_id + "/" + file.file_name
            src = src.replace(/ /g, '%20');
            let ext = file.file_name.split('.').pop().toLowerCase();
            if (["jpg", "png", "jpeg", "gif"].includes(ext)) {
                imgFile.append(`
                    <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href="${src}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                `)
            } else {
                if (ext === "mp4" || ext === ".webm") {
                    mp4File.append(`
                    <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16" data-id="${file.FileAssignmentID}" class="item-assignment">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href="${src}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                          <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                             <i class="fa fa-trash color-trash" aria-hidden="true"></i>
                         </button>
                     </div>
                    `)
                } else {
                    fileContainer.append(`
                    <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info"  href="${src}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                    `)
                }
            }
        });
        fileContainer.on("click", ".eye", function () {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });
        mp4File.on("click", ".eye", function () {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });
        imgFile.on("click", ".eye", function () {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });
        if (data.Assignment.type_assignment_id === 1) {


            const fileAssignment = $("#file-student")
            fileAssignment.empty()
            data.FileStudentAssignment.forEach(file => {
                let src = "/student/" + file.created_by + "/" + data.Assignment.AssignmentID + "/" + file.file_name
                src = src.replace(/ /g, '%20');
                let ext = file.file_name.split('.').pop().toLowerCase();
                fileAssignment.append(`
                <div class="media" style="align-items: center; width: 450px; padding: 5px;">
                    <div class="mr-15 bg-danger-light rounded text-center l-h-60" style="width: 50px; height: 50px;">
                        <span class="icon-Write font-size-24"><span class="path1"></span><span class="path2"></span></span>
                    </div>
                    <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                    <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                        <i class="fa fa-eye"></i>
                    </button>
                    <a class="font-size-18 text-gray hover-info" href="${src}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                    <span data-id="${file.FileAssignmentID}" class="item-assignment"><i class="fa fa-trash color-trash" aria-hidden="true"></i></span>
                </div>
            `)
            });
            fileAssignment.on("click", ".eye", function () {
                const ext = $(this).data("ext");
                const src = $(this).data("src");
                viewFile(ext, src);
            });

            if (data.UserAssignment.status) {
                var submittedTime = new Date(data.UserAssignment.completed_at); // Get current time

                // Convert deadline to a Date object
                var deadlineDate = new Date(data.Assignment.due_date);

                // Compare current time with deadline
                if (submittedTime > deadlineDate) {
                    $(".status").append(`<p data-i18n="late-submission" style="color:red"></p>`)
                } else {
                    $(".status").append(`<p data-i18n="assignment_submitted" style="color: #0d6efd"></p>`)
                }
                $("#submit-assignment").css("display", "none");
                $(".item-assignment").css("display", "none")
                $(".add-file").css("display", "none");
                $(".cancel-submit-assignment").css("display", "block");
            } else {
                $(".status").append(`<p data-i18n="not-submitted-yet" style="color: #0d6efd"></p>`);
                $(".cancel-submit-assignment").css("display", "none");
            }

            if (data.UserAssignment.score > 0) {
                $(".score").text(data.UserAssignment.score);
                $("#comment").text(": " + data.UserAssignment.comment);
                $(".not-graded-yet").css("display", "none");
                $(".score").css("display", "block");
                // $(".item-assignment").css("display", "none");
                $(".add-file").css("display", "none");
            }
            $("#add-file").on("change", () => {
                const filePath = $("#add-file").val();
                const fileName = filePath.split('\\').pop(); // lay phan tu cuoi trong mang
                $(".new-file").empty()
                $(".new-file").append(
                    `<p style="margin-top: 0; margin-bottom: 0; padding-left: 20px;">${fileName}</p>
             <button style="background-color: transparent; border: none" class="remove">
                <i class="fa fa-trash" style="color: red"></i>
             </button>
             `
                )
                $(".remove").on("click", () => {
                    $("#add-file").val("")
                    $(".new-file").empty()
                })
            })

            $(".remove").on("click", () => {
                $("#add-file").val("")
                $(".new-file").empty()
            })
            // delete file
            fileAssignment.on('click', '.color-trash', function () {
                const icon = $(this);
                const assignmentID = icon.closest('.item-assignment').data('id'); // Get the `data-id` value
                const paragraph = icon.closest('.media'); // Get the parent <p> element
                $('.delete-confirm-modal').modal('show');

                $('#confirmDelete').on('click', function () {
                    $.ajax({
                        url: '/api/v1/file-assignments/' + assignmentID,
                        type: 'DELETE',
                        success: function (response) {
                            if (response.code === 200) {
                                $(".delete-confirm-modal").modal("hide")
                                paragraph.remove();
                                $("#submit-assignment").css("display", "none");
                                $(".item-assignment").css("display", "none");
                                $(".add-file").css("display", "none");
                                $(".cancel-submit-assignment").css("display", "block");
                                showModalSuccess("Delete file success!")
                            } else {
                            }
                        },
                        error: function (xhr) {
                            showModalError(xhr.responseJSON.message);
                        }
                    });
                })
            })
            // upload file
            $("#submit-assignment").on("click", function (event) {
                event.preventDefault();
                const fileInput = $('#add-file')[0];
                const file = fileInput.files[0];
                const formData = new FormData();
                formData.append("assignment_id", assignmentID);
                formData.append("file", file);

                $.ajax({
                    url: "/api/v1/student_courses/assignment",
                    method: "POST",
                    data: formData,
                    contentType: false, // Important for file uploads
                    processData: false,
                })
                    .done(function (response) {
                        if (response.code === 200) {
                            fileAssignment.empty()
                            var data = response.data
                            userAssignmentID = data.UserAssignment.user_assignment_id;
                            data.FileAssignments.forEach(file => {
                                let src = "/student/" + file.created_by + "/" + file.assignment_id + "/" + file.file_name
                                src = src.replace(/ /g, '%20');
                                let ext = file.file_name.split('.').pop().toLowerCase();
                                fileAssignment.append(`
                                <div class="media" style="align-items: center; width: 450px; padding: 5px;">
                                    <div class="mr-15 bg-danger-light rounded text-center l-h-60" style="width: 50px; height: 50px;">
                                        <span class="icon-Write font-size-24"><span class="path1"></span><span class="path2"></span></span>
                                    </div>
                                    <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                                    <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                                        <i class="fa fa-eye"></i>
                                    </button>
                                    <a class="font-size-18 text-gray hover-info" href="${src}" download="${file.file_name}"><i class="fa fa-download"></i></a>
                                    <span data-id="${file.FileAssignmentID}" class="item-assignment"><i class="fa fa-trash color-trash" aria-hidden="true"></i></span>
                                </div>
                                `)


                            });
                            if (data.UserAssignment.status) {
                                var currentTime = new Date(data.UserAssignment.completed_at); // Get current time

                                // Convert deadline to a Date object
                                var deadlineDate = new Date(data.UserAssignment.Assignment.due_date);
                                // Compare current time with deadline
                                $(".status").empty();
                                if (currentTime > deadlineDate) {

                                    $(".status").append(`<p data-i18n="late-submission" style="color:red"></p>`)
                                } else {
                                    $(".status").append(`<p data-i18n="assignment_submitted" style="color: #0d6efd"></p>`)
                                }

                                $("#submit-assignment").css("display", "none");
                                $(".cancel-submit-assignment").css("display", "block");
                                $(".item-assignment").css("display", "none");
                                $(".add-file").css("display", "none");
                            }
                            $(".new-file").empty()
                            showModalSuccess("Assignment Submitted!")
                        }
                    })
                    .fail(function (xhr) {
                        showModalError(xhr.responseJSON.message);
                    });
            })

            // cancel submit

            $(".cancel-submit-assignment").on("click", function () {
                $.ajax({
                    url: "/api/v1/student_courses/courses/" + courseID + "/assignment/" + userAssignmentID,
                    method: "POST",
                    success: function () {
                        $("#submit-assignment").css("display", "block");
                        $(".item-assignment").css("display", "block");
                        $(".add-file").css("display", "flex");
                        $(".cancel-submit-assignment").css("display", "none");
                        $(".status").append(`<p data-i18n="not-submitted-yet" style="color: #0d6efd"></p>`);
                    },
                    error: function (xhr) {
                        const errorMessage = xhr.responseJSON?.message || "This course is Closed..";
                        showModalError(errorMessage);
                    },
                });

            })
        } else {
            if (data.UserAssignment.score > 0) {
                $(".score").text(data.UserAssignment.score);
                $("#comment").text(": " + data.UserAssignment.comment);
                $(".not-graded-yet").css("display", "none");
                $(".score").css("display", "block");
                // $(".item-assignment").css("display", "none");
                // $(".add-file").css("display", "none");
            }

            if (data.UserAssignment.status) {
                var submittedTime = new Date(data.UserAssignment.completed_at); // Get current time

                // Convert deadline to a Date object
                var deadlineDate = new Date(data.Assignment.due_date);

                // Compare current time with deadline
                if (submittedTime > deadlineDate) {
                    $(".status").append(`<p data-i18n="late-submission" style="color:red"></p>`)
                } else {
                    $(".status").append(`<p data-i18n="assignment_submitted" style="color: #0d6efd"></p>`)
                }
                $("#submit-assignment").css("display", "none");
                $(".item-assignment").css("display", "none")
                $(".add-file").css("display", "none");
                $(".cancel-submit-assignment").css("display", "none");
            } else {
                $("#submit-assignment").css("display", "none");
                $(".status").append(`<p data-i18n="not-submitted-yet" style="color: #0d6efd"></p>`);
                $(".cancel-submit-assignment").css("display", "none");
            }
            // $("#file-student-assignment").css("display","none");
            $("#submit-assignment").css("display", "none");
            const fileAssignment = $("#file-student")
            fileAssignment.empty();
            $(".add-file").css("display", "none");
            $.ajax({
                url: `/api/v1/assignments/${assignmentID}/topic`,
                method: 'GET',
                success: function (response) {
                    if (response.code === 200) {
                                const topic = response.data;
                                fileAssignment.append(`
                            <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                                 <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                                    <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                                 </div>
                                 <span class="title font-weight-500 font-size-16">${topic.name}</span>
                                
                                 <a class="font-size-18 text-gray hover-info" href="/student-courses/assignments/${assignmentID}/topics/${topic.topic_id}"><i class="fa fa-eye"></i></a>
                             </div>
                        `);
                    } else {
                                fileAssignment.append(`
                            <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                                 <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                                    <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                                 </div>
                                 <span class="title font-weight-500 font-size-16">No Topics found in this assignment</span>
                                
                                 <a class="font-size-18 text-gray hover-info"><i class="fa fa-eye"></i></a>
                             </div>
                        `);
                    }
                },
                error: function () {
                    fileAssignment.append(`
                    <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">No Topics found in this assignment</span>
                        
                         <a class="font-size-18 text-gray hover-info"><i class="fa fa-eye"></i></a>
                     </div>
                `);
                }
            });
        }
    })
})

function getId(key) {
    var url = window.location.href;
    var parts = url.split('/');
    var keyIndex = parts.indexOf(key);
    if (keyIndex !== -1 && keyIndex + 1 < parts.length) {
        return parseInt(parts[keyIndex + 1], 10);
    }
    return null;
}

