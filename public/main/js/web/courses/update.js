$(document).ready(function () {
    const apiUrl = 'api/v1/courses';
    const programApi = '/api/v1/programs';
    const categoryApi = '/api/v1/course_categories';
    const levelApi = '/api/v1/levels';
    const prerequisiteCourseApi = '/api/v1/programs/prerequisite-course'

    $('#userTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: (json) => {
                json.data.forEach((element) => {
                    element.Detail = `<a href="/courses/${element.course_id}">${element.title}</a>`;
                    element.Status = `<span class="badge badge-pill ${
                        element.status === true ? 'badge-success' : 'badge-danger'
                    }">${element.status === true ? 'Open' : 'Close'}</span>`;
                });
                return json.data;
            }
        },
        columns: [
            {
                data: null,
                render: function (data, type, row, meta) {
                    return meta.row + 1;
                },
                orderable: false,
                searchable: false,
            },
            { data: 'code'},
            { data: 'Detail' },
            { data: 'amount' },
            { data: 'Level.name' },
            { data: 'CourseCategory.name' },
            { data: 'Program.name' },
            { data: 'Status', orderable: false },
            {
                render: function (_, __, row) {
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.course_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.course_id}"><i class="fa fa-trash" aria-hidden="true"></i></button>
                    `;
                }
            }
        ],

        autoWidth: false,
        columnDefs: [
            { targets: -1, className: 'text-center', width: '80px' },
            { targets: '_all', defaultContent: '', width: 'auto' }
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },
    });

    function loadDropdowns(selectedProgramId = null, selectedLevelId = null, selectedCategoryId = null) {
        const programSelect = $('#editProgramId');
        $.get(programApi, function (response) {

            programSelect.empty();
            programSelect.append('<option value="">Select Program</option>');
            response.data.forEach(program => {
                programSelect.append(
                    `<option value="${program.program_id}" ${program.program_id == selectedProgramId ? 'selected' : ''}>
                        ${program.name}
                    </option>`
                );
            });


        });
        
        $.get(levelApi, function (response) {
            const levelSelect = $('#editLevelId, #level_id');
            levelSelect.empty();
            levelSelect.append('<option value="">Select Level</option>');
            response.data.forEach(level => {
                levelSelect.append(
                    `<option value="${level.level_id}" ${level.level_id == selectedLevelId ? 'selected' : ''}>
                        ${level.name}
                    </option>`
                );
            });
        });


        $.get(categoryApi, function (response) {
            const categorySelect = $('#editCategoryId, #category_id');
            categorySelect.empty();
            categorySelect.append('<option value="">Select Category</option>');
            response.data.forEach(category => {
                const isSelected = selectedCategoryId && category.course_categories_id == selectedCategoryId;
                categorySelect.append(
                    `<option value="${category.course_categories_id}" ${isSelected ? 'selected' : ''}>
                        ${category.name}
                    </option>`
                );
            });
        });

        if(selectedProgramId) {
            const api = prerequisiteCourseApi + "/" + selectedProgramId
            // Load prerequisite course
            $.get(api, function (response) {
                const prerequisiteCourseSelect = $('#edit-prerequisite_course_id');
                prerequisiteCourseSelect.empty();
                response.data.forEach(c => {
                    prerequisiteCourseSelect.append(`<option value="${c.course_id}">${c.code} - ${c.title}</option>`);
                })
            });
        }

        programSelect.on("change", function (){
            programID = $(this).val()
            const api = prerequisiteCourseApi + "/" + programID
            // Load prerequisite course
            $.get(api, function (response) {
                const prerequisiteCourseSelect = $('#edit-prerequisite_course_id');
                prerequisiteCourseSelect.empty();
                response.data.forEach(c => {
                    prerequisiteCourseSelect.append(`<option value="${c.course_id}">${c.code} - ${c.title}</option>`);
                })
            });
        })
    }


    $('#userTable').on('click', '.edit-btn', function () {
        const courseId = $(this).data('id');
        $.get(`${apiUrl}/${courseId}`, function (response) {
            const course = response.course || response;
            $('#editTitle').val(course.title);
            $('#editAmount').val(course.amount);
            $('#editStartTime').val(course.start_time.split('T')[0]);
            $('#editEndTime').val(course.end_time.split('T')[0]);
            $('#editDescription').val(course.description);
            $('#editStatus').prop('checked', course.status === true);
            loadDropdowns(course.program_id, course.level_id, course.course_categories_id);

            $('#editCourse').data('id', courseId);
            $('.modal-edit-course').modal('show');
        });
    });

    $('#editCourse').on('click', function () {
        const courseId = $(this).data('id');
        const updatedData = {
            title: $('#editTitle').val(),
            amount: parseFloat($('#editAmount').val()),
            start_time: $('#editStartTime').val(),
            end_time: $('#editEndTime').val(),
            description: $('#editDescription').val(),
            status: $('#editStatus').prop('checked'),
            course_categories_id: parseInt($('#editCategoryId').val()),
            program_id: parseInt($('#editProgramId').val()),
            level_id: parseInt($('#editLevelId').val()),
            prerequisite_course_id: parseInt($('#edit-prerequisite_course_id').val())
        };

        $.ajax({
            url: `${apiUrl}/${courseId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(updatedData),
            success: function (response) {
                if(response.code === 200){
                    showModalSuccess("Course updated successfully");
                    $('#userTable').DataTable().ajax.reload();
                    $('.modal-edit-course').modal('hide');
                }else{
                    showModalError(response.message);
                }

            },
            error: function (xhr) {
                showModalError("Failed to update course");
            }
        });
    });
});
