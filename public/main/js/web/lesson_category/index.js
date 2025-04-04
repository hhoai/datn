
$(document).ready(function () {
    const apiUrl = '/api/v1/lesson-categories';

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


            {
                render: function (data, type, row) {
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.lesson_category_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.lesson_category_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddCategoryLessonModal').on('click', function () {
        $('#name').val('');
    });

    $('#saveLessonCategory').on('click', function () {
        const formData = {
            name: $('#name').val(),
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

    let lessonCategoryIdToDelete = null;
    $('#categoryTable').on('click', '.delete-btn', function () {
        lessonCategoryIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${lessonCategoryIdToDelete}`,
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
        const lessonCategoryId = $(this).data('id');
        $('.modal-edit-lesson-category').data('lesson-category-id', lessonCategoryId);

        $.ajax({
            url: `${apiUrl}/${lessonCategoryId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const courseCategory = response.data;
                    $('#editName').val(courseCategory.name);
                    $('.modal-edit-lesson-category').modal('show');
                } else {
                    showModalError(response.message || "Failed to fetch  data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching  details");
            }
        });
    });


    $('#saveEditLessonCategory').on('click', function () {
        const  lessonCategoryId = $('.modal-edit-lesson-category').data('lesson-category-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${ lessonCategoryId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Category updated successfully");
                    $('#categoryTable').DataTable().ajax.reload();
                    $('.modal-edit-lesson-category').modal('hide');
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