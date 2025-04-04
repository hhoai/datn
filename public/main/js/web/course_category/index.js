$(document).ready(function () {
    const apiUrl = '/api/v1/course_categories';

    $('#categoryTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: 'data'
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
            { data: 'name' },
            { data: 'description'},

            {
                render: function (data, type, row) {
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.course_categories_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.course_categories_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
                }
            }
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        }
    });

    $('#openAddCategoryCourseModal').on('click', function () {
        $('#name').val('');
        $('#description').val('');

    });

    $('#saveCourseCategory').on('click', function () {
        const formData = {
            name: $('#name').val(),
            description: $('#description').val(),

        };
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Category created successfully");
                $('#categoryTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to create program");
            }
        });
    });

    let courseCategoryIdToDelete = null;
    $('#categoryTable').on('click', '.delete-btn', function () {
        courseCategoryIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${courseCategoryIdToDelete}`,
            method: 'DELETE',
            success: function (response) {

                if (response.code === 200) {
                    showModalSuccess("Category deleted successfully");
                    $('#categoryTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete program");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Fail to delete level");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#categoryTable').on('click', '.edit-btn', function () {
        const courseCategoryId = $(this).data('id');
        $('.modal-edit-course-category').data('course-category-id', courseCategoryId);

        $.ajax({
            url: `${apiUrl}/${courseCategoryId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const courseCategory = response.data;
                    $('#editName').val(courseCategory.name);
                    $('#editDescription').val(courseCategory.description);
                    $('.modal-edit-course-category').modal('show');
                } else {
                    showModalError(response.message || "Failed to fetch skill data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching skill details");
            }
        });
    });


    $('#saveEditCourseCategory').on('click', function () {
        const  courseCategoryId = $('.modal-edit-course-category').data('course-category-id');
        const formData = {
            name: $('#editName').val(),
            description: $('#editDescription').val(),
        };
        $.ajax({
            url: `${apiUrl}/${courseCategoryId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Category updated successfully");
                    $('#categoryTable').DataTable().ajax.reload();
                    $('.modal-edit-course-category').modal('hide');
                } else {
                    showModalError("Failed to update skill");
                }
            },
            error: function (xhr) {
                showModalError("Failed to update skill");
            }
        });
    });

});