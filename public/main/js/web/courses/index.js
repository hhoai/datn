$(document).ready(function () {
    const apiUrl = 'api/v1/courses';
    const programApi = '/api/v1/programs';
    const categoryApi = '/api/v1/course_categories';
    const levelApi = '/api/v1/levels'
    const prerequisiteCourseApi = '/api/v1/programs/prerequisite-course'

    
    function loadDropdowns(selectedProgramId = null, selectedLevelId = null, selectedCategoryId = null, selectedPrequisiteCourseID = null) {
        const programSelect = $('#editProgramId, #program_id');
        $.get(programApi, function (response) {
            programSelect.empty();
            programSelect.append('<option value="">Select Program</option>');
            response.data.forEach(program => {
                programSelect.append(`<option value="${program.program_id}" ${program.program_id == selectedProgramId ? 'selected' : ''}>${program.name}</option>`);
            });
        });

        // Load Levels
        $.get(levelApi, function (response) {
            const levelSelect = $('#editLevelId, #level_id');
            levelSelect.empty();
            levelSelect.append('<option value="">Select Level</option>');
            response.data.forEach(level => {
                levelSelect.append(`<option value="${level.level_id}" ${level.level_id == selectedLevelId ? 'selected' : ''}>${level.name}</option>`);
            });
        });

        // Load Categories
        $.get(categoryApi, function (response) {
            const categorySelect = $('#editCategoryId, #category_id');
            categorySelect.empty();
            categorySelect.append('<option value="">Select Category</option>');


            response.data.forEach(category => {

                const isSelected = selectedCategoryId && category.course_categories_id == selectedCategoryId;
                categorySelect.append(`<option value="${category.course_categories_id}" ${isSelected ? 'selected' : ''}>${category.name}</option>`);
            });
        });

        const prerequisiteCourseSelect = $('#edit-prerequisite_course_id, #prerequisite_course_id');
        prerequisiteCourseSelect.empty();
        prerequisiteCourseSelect.append('<option value="">Select Prerequisite Course</option>');

        programSelect.on("change", function (){
            const programID = $(this).val()
            const api = prerequisiteCourseApi + "/" + programID
            // Load prerequisite course
            $.get(api, function (response) {
                prerequisiteCourseSelect.empty();
                prerequisiteCourseSelect.append('<option value="">Select Prerequisite Course</option>');
                response.data.forEach(c => {
                    prerequisiteCourseSelect.append(`<option value="${c.course_id}">${c.code} - ${c.title}</option>`);
                })
            });
        })
    }


    $('#openAddCourseModal').on('click', function () {
        $('#title, #amount, #startTime, #endTime, #description').val('');
        $('#program_id, #level_id, #category_id, #prerequisite_course_id').empty();
        loadDropdowns();
        $('.modal-add-course').modal('show');
    });
    $( '#saveCourse').on('click',function () {
        const courseData = {
            title: $('#title').val(),
            amount: parseFloat($('#amount').val()),
            start_time: $('#startTime').val(),
            end_time: $('#endTime').val(),
            description: $('#description').val(),
            course_categories_id: parseInt($('#category_id').val()),
            program_id: parseInt($('#program_id').val()),
            level_id: parseInt($('#level_id').val()),
            prerequisite_course_id: parseInt($('#prerequisite_course_id').val())
        };

        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(courseData),
            success: function (response) {
                if(response.code === 200){
                    showModalSuccess("Course updated successfully");
                    $('#userTable').DataTable().ajax.reload();
                    $('.modal-add-course').modal('hide');
                }else{
                    showModalError(response.message);
                }
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });
    });

    let courseIdToDelete = null;
    $('#userTable').on('click', '.delete-btn', function () {

        courseIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });

    $('#confirmDelete').on('click', function () {
        if (!courseIdToDelete) {
            showModalError("Failed to delete course");
            return;
        }

        $.ajax({
            url: `${apiUrl}/${courseIdToDelete}`,
            method: 'DELETE',
            success: function (response) {
                if (response.success) {
                    showModalSuccess("Course deleted successfully");
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete course");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to delete course");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });
});