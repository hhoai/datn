var courseID = $('input[name="course_id"]').val();

function teacherList(){
    $.ajax({
        url: '/api/v1/courses/' + courseID + '/teachers',
        method: "GET",
        success: function (response) {
            if (response.code === 200) {
                const teachers = response.data.teachers || [];
                let htmlContent = teachers.map(teacher => `
                <div class="media media-custom">
                    <span class="badge badge-dot badge-danger badge-custom"></span>
                    <a class="avatar avatar-lg status-success" href="#">
                        <img src="/images/avatar/1.jpg" alt="${teacher.User.name}">
                    </a>
                    <div class="media-body">
                        <p><a href="#"><strong>${teacher.User.name}</strong></a></p>
                        <p>Email: ${teacher.User.email}</p>
                    </div>
                      <button class="btn btn-danger btn-sm" data-id="${teacher.User.user_id}" onclick="showDeleteModal(this)">
                      <i class="fa fa-trash" aria-hidden="true"></i></button>
                </div>
            `).join("");

                $("#teacher-list").html(htmlContent);
            } else {
                $("#teacher-list").html("<p>No teachers found.</p>");
            }
        },
        error: function (xhr) {
            showModalError(xhr.responseJSON.message);
        },
    });
}

function studentList(){
    $.ajax({
        url: '/api/v1/courses/' + courseID + '/students',
        method: "GET",
        success: function (response) {
            if (response.code === 200) {
                const students = response.data.students || [];
                let htmlContent = students.map(student => `
                <div class="media media-custom">
                    <span class="badge badge-dot badge-danger"></span>
                    <a class="avatar avatar-lg status-success" href="#">
                        <img src="/images/avatar/1.jpg" alt="${student.User.name}">
                    </a>
                    <div class="media-body">
                        <p><a href="#"><strong>${student.User.name}</strong></a></p>
                        <p>Email: ${student.User.email}</p>
                    </div>
                     <a href="/users/${student.User.user_id}/courses" 
                       class="btn btn-success btn-sm">
                      <i class="fa fa-book" aria-hidden="true"></i>
                    </a>
                     <button class="btn btn-danger btn-sm" data-id="${student.User.user_id}" onclick="showDeleteModal(this)">
                      <i class="fa fa-trash" aria-hidden="true"></i></button>    
                </div>
            `).join("");

                $("#student-list").html(htmlContent);
            } else {
                $("#student-list").html("<p>No teachers found.</p>");
            }
        },
        error: function (xhr) {
            showModalError(xhr.responseJSON.message);
        },
    });
}


$(document).ready(function () {

    const apiUrl = '/api/v1/lessons/' + courseID;
    const teacherApi = '/api/v1/courses/' + courseID + '/teachers_not_course';
    const studentApi = '/api/v1/courses/' + courseID + '/students_not_course';

    var table = $('#lessonTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: (json) => {
                json.data.forEach((element, index) => {
                    element.RowNumber = `<p>${index + 1}</p>`
                    element.Detail = `<a href="/lessons/${element.LessonID}/assignments">${element.title}</a>`;
                    element.Action = `<button class="btn btn-primary btn-sm edit-btn" data-id="${element.LessonID}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></i></button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${element.LessonID}"><i class="fa fa-trash" aria-hidden="true"></i></button>`;
                    element.Checkbox = `<input type="checkbox" id="checkbox-lesson" class="item-account filled-in" value="${element.LessonID}" name="lesson_id[]">
                                <label for="checkbox-lesson"></label>`;
                });
                return json.data;
            },
        },
        columns: [
            { data: 'RowNumber'},
            { data: 'Detail' },
            { data: 'Level.name'},
            { data: 'LessonCategory.name'},
            {data: 'Action'},
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },

    });
    teacherList();
    studentList();



    function loadTeacherDropdowns(selectedUserId = null) {

        $.get(teacherApi, function (response) {
            const userSelect = $('#user_id');
            userSelect.empty();
            response.data.teachers.forEach(user => {
                userSelect.append(`<option value="${user.user_id}" ${user.user_id == selectedUserId ? 'selected' : ''}>${user.name}</option>`);
            });
        });

    }
    function loadStudentDropdowns(selectedUserId = null) {

        $.get(studentApi, function (response) {
            const userSelect = $('#student_id');
            userSelect.empty();
            response.data.students.forEach(user => {
                userSelect.append(`<option value="${user.user_id}" ${user.user_id == selectedUserId ? 'selected' : ''}>${user.name}</option>`);
            });
        });

    }
    function loadStudentCheckboxes(selectedUserIds = []) {
        $.get(studentApi, function (response) {
            const userList = $('#student_list');
            userList.empty();
            response.data.students.forEach(user => {
                const isChecked = selectedUserIds.includes(user.user_id);
                userList.append(`
                <div>
                    <input type="checkbox" class="student-checkbox" id="student_${user.user_id}" value="${user.user_id}" ${isChecked ? 'checked' : ''}>
                    <label for="student_${user.user_id}">${user.name}</label>
                </div>
            `);
            });
        });
    }
    $('#openAddTeacherModal').on('click', function () {
        $('#user_id').empty();
        loadTeacherDropdowns();

    });

    $(document).ready(function () {
        function filterStudents(keyword) {
            if (!keyword.trim()) {
                $('#student_list > div').show(); // Hiển thị tất cả học viên nếu từ khóa rỗng
                return;
            }

            $('#student_list > div').each(function () {
                const studentName = $(this).find('label').text().toLowerCase();
                if (studentName.includes(keyword.toLowerCase())) {
                    $(this).show();
                } else {
                    $(this).hide();
                }
            });
        }

        $('#search_student').on('input', function () { 
            const keyword = $(this).val();
            filterStudents(keyword);
        });
    });


    $('#saveUser').on('click', function () {
        const formData = {
            user_id: parseInt($('#user_id').val())
        };

        $.ajax({
            url: '/api/v1/courses/' + courseID + '/teachers',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Add teacher successfully");
                $('.create-teacher-modal').modal('hide');
                teacherList();

            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });
    });

    $('#openAddStudentModal').on('click', function () {
        $('#student_id').empty();
        loadStudentCheckboxes()
    });

    $('#saveStudent').on('click', function () {
        const selectedUserIds = [];
        $('.student-checkbox:checked').each(function () {
            selectedUserIds.push(parseInt($(this).val()));
        });

        const formData = {
            user_ids: selectedUserIds
        };

        $.ajax({
            url: '/api/v1/courses/' + courseID + '/students',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Add Student successfully");
                $('#addStudentModal').modal('hide');
                studentList();

            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
               }
        });
    });



});
let selectedUserID = null;
function showDeleteModal(button) {
    selectedUserID = $(button).data('id');
    $('.delete-user-modal').modal('show');
}

$(document).ready(function () {
    $('#deleteUser').on('click', function () {
        if (selectedUserID) {
            $.ajax({
                url: '/api/v1/courses/' + courseID + '/remove_user',
                method: "DELETE",
                contentType: "application/json",
                data: JSON.stringify({ user_id: selectedUserID }),
                success: function (response) {
                    if (response.code === 200) {
                        $("#teacher-" + selectedUserID).remove();
                        showModalSuccess("Deleted successfully");
                        teacherList();
                        studentList();
                    }
                    $('.delete-user-modal').modal('hide');

                },
                error: function (xhr) {
                    showModalError(xhr.responseJSON.message);
                    $('.delete-user-modal').modal('hide');

                }
            });
        }
    });
});


