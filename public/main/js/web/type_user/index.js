$(document).ready(function () {
    const apiUrl = '/api/v1/type_users';

    $('#typeUserTable').DataTable({
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
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.type_user_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.type_user_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddTypeUserModal').on('click', function () {
        $('#name').val('');

    });

    $('#saveTypeUser').on('click', function () {
        const formData = {
            name: $('#name').val(),
        };
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Type User created successfully");
                $('#typeUserTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to create program");
            }
        });
    });

    let typeUserIdToDelete = null;
    $('#typeUserTable').on('click', '.delete-btn', function () {
        typeUserIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${typeUserIdToDelete}`,
            method: 'DELETE',
            success: function (response) {

                if (response.code === 200) {
                    showModalSuccess("Type User deleted successfully");
                    $('#typeUserTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete type user");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Fail to delete level");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#typeUserTable').on('click', '.edit-btn', function () {
        const typeUserId = $(this).data('id');
        $('.modal-edit-type-user').data('type-user-id', typeUserId);

        $.ajax({
            url: `${apiUrl}/${typeUserId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const typeUser = response.data;
                    $('#editName').val(typeUser.name);
                    $('.modal-edit-type-user').modal('show');
                } else {
                    showModalError(response.message || "Failed to fetch data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching details");
            }
        });
    });


    $('#saveEditTypeUser').on('click', function () {
        const  typeUserId = $('.modal-edit-type-user').data('type-user-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${typeUserId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Type User updated successfully");
                    $('#typeUserTable').DataTable().ajax.reload();
                    $('.modal-edit-type-user').modal('hide');
                } else {
                    showModalError("Failed to update ");
                }
            },
            error: function (xhr) {
                showModalError("Failed to update type user");
            }
        });
    });

});