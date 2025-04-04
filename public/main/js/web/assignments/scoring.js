$(document).ready(function () {
    var url = window.location.pathname;
    var assignmentID = url.split('/')[3];
    var apiGetUserStudents = "/api/v1/user-assignments/" + assignmentID

    const listStudents = $("#list-students")
    const studentAssignment = $("#assignment-content")

    function viewFile(file, src){
        const encodedSrc = src.replace(/ /g, '%20');
        const fileExtension = src.split('.').pop().toLowerCase();
        const preview = $('#preview')
        // Get file extension in lowercase
        preview.empty()
        if (["jpg", "png", "jpeg", "gif"].includes(fileExtension)) {
            preview.append(`
        <img src="${encodedSrc}" alt="img" style="width: 800px">
    `);
        }
        else if (file === "mp4" || file === "webm") {
            preview.append(`<video controls width="100%" height="600px"><source src=${src} type="video/mp4">Your browser does not support the video tag.</video>`);
        }
        else {
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
                    preview.append(`
                    <p>download-to-view-file</p>
                    `);
                    break
            }
        }
    }

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

    $.get(apiGetUserStudents, function (response) {
        var data = response.data.UserAssignments
        var countCompleted = response.data.IsCompleted
        var count = response.data.CountAssignment

        const maxScore = $(".max-score")
        const isSubmitted = $(".count-submitted")
        const header = $(".header")

        data.forEach((data, index) => {
            if (index === 0) {
                header.append(`<span>${data.Assignment.title}</span>`)

                maxScore.append(`<h3 style="margin: 0; text-align: center; padding-bottom: 10px">${data.Assignment.score}</h3>`)

                isSubmitted.append(`<h3 style="margin: 0; text-align: center; padding-bottom: 10px"><span style="color: #ff0000">${countCompleted}</span>/${count}</h3>`)
            }
        })

        data.forEach((data) => {
            if(data.status) {
                listStudents.append(`
                <div class="d-flex align-items-center mb-30">
                    <div class="mr-15">
                        <img src="/images/avatar/avatar-11.png" class="avatar avatar-lg rounded10 bg-primary-light" alt="">
                    </div>
                    <div class="d-flex flex-column flex-grow-1" style="margin-bottom: 0">
                        <button class="text-dark hover-success mb-1 font-size-16 student-assignment" data-id="${data.user_id}" style="border: none; background-color: transparent; text-align: left">${data.User.name}</button>
                        <p style="padding-left: 5px; margin: 0; color: #ff0000; display: flex; flex-direction: column; gap: 2px"><span data-i18n="assignment_submitted"></span><span id="score"></span></p>
                    </div>
                </div>
                `)
                if (data.score > 0) {
                    $("#score").append(`<span data-i18n="score"></span>`)
                    $("#score").append(`: ${data.score}`)
                }
            }
            else {
                listStudents.append(`
                <div class="d-flex align-items-center mb-30">
                    <div class="mr-15">
                        <img src="/images/avatar/avatar-11.png" class="avatar avatar-lg rounded10 bg-primary-light" alt="">
                    </div>
                    <div class="d-flex flex-column flex-grow-1 font-weight-500">
                        <button class="text-dark hover-success mb-1 font-size-16 student-assignment" data-id="${data.user_id}" style="border: none; background-color: transparent; text-align: left">${data.User.name}</button>
                        <span class="text-fade" style="padding-left: 5px;" data-i18n="not-submitted-yet"></span>
                    </div>
                </div>
                `)
            }

            $('.student-assignment').on('click', function () {
                var userID = $(this).data('id');
                var assignmentID = data.assignment_id

                $.ajax({
                    url: "/api/v1/user-assignments/" + assignmentID + "/" + userID,
                    method: "GET",
                    success: function (response) {
                        var data = response.data
                        studentAssignment.empty()
                        $(".student").empty()

                        $(".student").append(`<p><span>Student</span>: ${data.Username}</p>`);
                        data.FileAssignments.forEach(file => {
                            let src = "/student/" + userID + "/" + assignmentID + "/" + file.file_name
                            src = src.replace(/ /g, '%20');
                            let ext = file.file_name.split('.').pop().toLowerCase();

                            studentAssignment.append(`
                            <div style="display: flex; justify-content: space-between">
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
                                <p><span>Submitted At</span>: <span class="text-danger">${formatTime(file.CreatedAt)}</span></p>
                            </div>
                            `)
                        });

                        studentAssignment.on("click", ".eye", function () {
                            const ext = $(this).data("ext");
                            const src = $(this).data("src");
                            viewFile(ext, src);
                        });

                        $("#comment").empty()
                        $("#comment").append(`
                            <div style="display: flex; justify-content: space-between;">
                            <div class="input-group" style="width: 12%;">
                            <span class="input-group-prepend">
                                <span class="input-group-text"><i class="fa fa-pencil-square-o"></i></span>
                            </span>
                            <input type="number" class="form-control" id="score" min="0" value="${data.UserAssignments.score}" max="${data.UserAssignments.Assignment.score}" style="border: 1px solid #86a4c3;">
                            </div>
                            <div class="input-group" style="width: 60%;">
                                <span class="input-group-prepend">
                                    <span class="input-group-text"><i class="fa fa-commenting-o"></i></span>
                                </span>
                                <input type="text" class="form-control" id="comment-content" value="${data.UserAssignments.comment}" style="border: 1px solid #86a4c3;">
                            </div>
                            <button class="btn btn-primary" id="comment-btn">Save</button>
                            <button class="btn btn-danger" id="cancel-submission">Cancel submission</button>
                            </div>
                            <div class="err"></div>
                        `)

                        let topic = response.data.Topic;
                        if (topic) {
                            $("#file-topic").empty();
                            $("#file-topic").append(`
                            <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                                 <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                                    <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                                 </div>
                                 <span class="title font-weight-500 font-size-16">${topic.name}</span>
                                
                                 <a class="font-size-18 text-gray hover-info" href="/courses/${response.data.UserAssignments.user_id}/assignments/${response.data.UserAssignments.assignment_id}/topics/${topic.topic_id}"><i class="fa fa-eye"></i></a>
                             </div>
                        `);
                        }

                        $("#comment-btn").on("click", function () {
                            var formData = {
                                "score" : $("#score").val(),
                                "comment": $("#comment-content").val(),
                                "student_id": data.UserAssignments.user_id,
                                "assignment_id": data.UserAssignments.assignment_id
                            }

                            let score = $("#score").val()

                            $(".err").empty()
                            if (+score <= 0) {
                                $(".err").append(`<span style="color:#ff0000;">* Score must be greater than 0 </span>`)
                            }
                            else if (+score > data.UserAssignments.Assignment.score) {
                                $(".err").append(`<span style="color:#ff0000;">* Please re-enter the score</span>`)
                            }
                            else {
                                $.ajax({
                                    url: "/api/v1/user-assignments/comment",
                                    method: "PUT",
                                    contentType: "application/json",
                                    data: JSON.stringify(formData),
                                    success: function (response) {
                                        let data = response.data

                                        $("#score").val(data.score)
                                        $("#comment-content").val(data.comment)
                                        showModalSuccess("Save successfully")
                                    },
                                    error: function (xhr) {
                                        showModalError(xhr.responseJSON.message);
                                    },
                                })
                            }
                        })

                        $("#cancel-submission").on("click", function (){
                            var formData = {
                                "student_id": data.UserAssignments.user_id,
                                "assignment_id": data.UserAssignments.assignment_id
                            }

                            $(".delete-confirm-modal").modal("show")

                            $("#confirmDelete").on("click", ()=>{
                                $.ajax({
                                    url: "/api/v1/user-assignments/cancel-submission",
                                    method: "PUT",
                                    contentType: "application/json",
                                    data: JSON.stringify(formData),
                                    success: function (response) {
                                        window.location.href = "/lessons/assignment/" + data.UserAssignments.assignment_id + "/scoring"
                                    },
                                    error: function (xhr) {
                                        showModalError(xhr.responseJSON.message);
                                    },
                                })
                            })
                        })
                    },
                    error: function (xhr) {
                        showModalError(xhr.responseJSON.message);
                    },
                });
            })
        })
    })
})