$(document).ready(function () {
    var url = window.location.href;
    var parts = url.split('/');
    var courseId = parseInt(parts[parts.length - 1], 10);
    var id = parseInt(parts[parts.indexOf("assignments") + 1], 10);

    const apiGetPost = "/api/v1/student_courses/courses/" + courseId + "/assignments/" + id + "/details";
    const imgFile = $('#img-file');
    const mp4File = $('#mp4-file');
    const fileContainer = $('#file');
    const apiTopic = '/api/v1/topics';


    function loadDropdowns(selectedTopicId = null) {
        $.get(apiTopic, function (response) {
            const topicSelect = $('#topic_id');
            if (response.data && Array.isArray(response.data)) {
                response.data.forEach(topic => {
                    topicSelect.append(
                        `<option value="${topic.topic_id}" ${topic.topic_id == selectedTopicId ? 'selected' : ''}>${topic.name}</option>`
                    );
                });
            } else {
            }
        })
            .fail(function (xhr) {
            });
    }

    function loadTopic(fileAssignment) {
        $.ajax({
            url: `/api/v1/assignments/${id}/topic`,
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
                                
                                 <a class="font-size-18 text-gray hover-info" href="/topics/${topic.topic_id}"><i class="fa fa-eye"></i></a>
                                 <button class="btn delete-topic" style="color: red" ><i class="fa fa-trash-o"></i></button>
                             </div>
                        `);

                    $('.delete-topic').on("click", function (){
                        $('.delete-confirm-topic').modal('show');

                        $('.confirmDelete').on('click', function () {
                            $.ajax({
                                url: `/api/v1/topics/` + topic.topic_id + "/" + id,
                                method: 'DELETE',
                                success: function (response) {
                                    if (response.code === 200) {
                                        showModalSuccess("Delete Successfully");
                                        fileAssignment.reload()
                                    } else {
                                        showModalError(response.message);
                                    }
                                    $('.delete-confirm-topic').modal('hide');
                                },
                                error: function (xhr) {
                                    showModalError("Fail to delete");
                                    $('.delete-confirm-topic').modal('hide');
                                }
                            });
                        });
                    })

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
            preview.append(`<video controls width="100%" height="600px"><source src=${encodedSrc} type="video/mp4">Your browser does not support the video tag.</video>`);
        }
        else {
            switch (file) {
                case "pdf":
                    preview.append(`<embed src=${encodedSrc} width="100%" height="600px" type="application/pdf">`);
                    break
                case "doc":
                case "docx":
                    const domain = "https://5b28-2402-800-6106-a935-ad3d-c9d0-4399-8331.ngrok-free.app"
                    src = "https://view.officeapps.live.com/op/embed.aspx?src=" + domain + src
                    preview.append(`<iframe src="${src}" width="100%" height="600px" frameborder="0"></iframe>`);
                    // preview.append(`<iframe
                    //    src="https://docs.google.com/gview?url=${domain}${src}/to/document.doc&embedded=true"></iframe>`)
                    break
                default:
                    preview.append(`<p data-i18n="download-to-view-file"></p>`);
                    break
            }
        }
    }

    $.get(apiGetPost, function (response) {
        var data = response.data

        $("#page-title").text(data.Lesson.title);
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
            }
            else {
                if (ext === "mp4" || ext === ".webm") {
                    mp4File.append(`
                    <div class="media" style="align-items: center; width: 450px; padding: 5px ;">
                         <div class="ml-0 mr-15 bg-primary-light l-h-50 rounded text-center" style="height: 50px; width: 50px;">
                            <span class="font-size-24 text-primary"><i class="fa fa-file-text"></i></span>
                         </div>
                         <span class="title font-weight-500 font-size-16">${file.file_name}</span>
                         <button type="button" class="btn btn-primary eye" data-toggle="modal" data-target="#preview-container" data-ext="${ext}" data-src="${src}" style="background: transparent; color: #007bff; border: none;">
                            <i class="fa fa-eye"></i>
                         </button>
                         <a class="font-size-18 text-gray hover-info" href=${src} download="${file.file_name}"><i class="fa fa-download"></i></a>
                     </div>
                    `)
                }
                else {
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
        fileContainer.on("click", ".eye", function() {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });
        mp4File.on("click", ".eye", function() {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });
        imgFile.on("click", ".eye", function() {
            const ext = $(this).data("ext");
            const src = $(this).data("src");
            viewFile(ext, src);
        });

        if(data.Assignment.type_assignment_id === 1){
            $("#student-container").css("display","none");
        }else if (data.Assignment.type_assignment_id === 2) {
            const fileAssignment = $("#file-student")
            fileAssignment.empty();
            loadTopic(fileAssignment);
            $('#openAddTopicModal').on('click', function () {
                $('#topic_id').empty();
                loadDropdowns();
            });

            $('#saveTopic').on('click', function () {
                const topicId = parseInt($('#topic_id').val());
                const requestData = {
                    assignment_id: id,
                    topic_id: topicId,
                };

                $.ajax({
                    url: '/api/v1/assignments/assign_topic',
                    method: 'POST',
                    contentType: 'application/json',
                    data: JSON.stringify(requestData),
                    success: function (response) {
                        if(response.code === 200){
                            showModalSuccess("Topic assign successfully");
                            window.location.href = "/assignments/" + id + "/courses/" + courseId;
                        }else{
                            showModalError(response.message);
                        }
                    },
                    error: function (xhr) {
                    },
                });
            });

        }
    })

})