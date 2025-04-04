$(document).ready(function () {
    const apiUrl = '/api/v1/levels';

    $('#levelTable').DataTable({
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
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.level_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.level_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddLevelModal').on('click', function () {
        $('#name').val('');
        $('.modal-add-level').modal('show');
    });

    $('#saveLevel').on('click', function () {
        const formData = {
            name: $('#name').val(),
        };
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Level created successfully");
                $('#levelTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
                $('#name').val("");
            },
            error: function (xhr) {
                showModalError("Failed to create program");
            }
        });
    });

    let levelIdToDelete = null;
    $('#levelTable').on('click', '.delete-btn', function () {
        levelIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${levelIdToDelete}`, // URL API
            method: 'DELETE',
            success: function (response) {

                if (response.success) {
                    showModalSuccess("Level deleted successfully");
                    $('#levelTable').DataTable().ajax.reload();
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


    $('#levelTable').on('click', '.edit-btn', function () {
        const levelId = $(this).data('id');
        $('.modal-edit-level').data('level-id', levelId);

        $.ajax({
            url: `${apiUrl}/${levelId}`,
            method: 'GET',
            success: function (response) {
                if (response.success) {
                    const level = response.level;
                    $('#editName').val(level.name);
                    $('.modal-edit-level').modal('show');
                } else {
                    showModalError("Failed to fetch level data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching level details");
            }
        });
    });


    $('#saveEditLevel').on('click', function () {
        const  levelId = $('.modal-edit-level').data('level-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${levelId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Program updated successfully");
                    $('#levelTable').DataTable().ajax.reload();
                    $('.modal-edit-level').modal('hide');
                } else {
                    showModalError("Failed to update program");
                }
            },
            error: function (xhr) {
                showModalError("Failed to update program");
            }
        });
    });




});