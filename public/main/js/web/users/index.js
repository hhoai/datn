$(document).ready(function () {
    const apiUrl = '/api/v1/users';
    const rolesApi = '/api/v1/roles';
    const typeUsersApi = '/api/v1/type_users';


    $('#userTable').DataTable({
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
            { data: 'email' },
            { data: 'Role.name' },
            { data: 'TypeUser.name' },
            { data: 'TypeAccount'},
            {
                render: function (_, __, row) {
                    if (row.TypeAccount) {
                        return `
                            <button class="btn btn-primary btn-sm edit-btn" data-id="${row.user_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>
                            <button class="btn btn-danger btn-sm delete-btn" data-id="${row.user_id}"><i class="fa fa-trash" aria-hidden="true"></i></button>
                        `
                    }
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.user_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.user_id}"><i class="fa fa-trash" aria-hidden="true"></i></button>
                        <button class="btn btn-success btn-sm password-btn" data-id="${row.user_id}"><i class="fa fa-key" aria-hidden="true"></i></button>
                    `;
                }
            }
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },
    });


    function loadDropdowns(selectedRoleId = null, selectedTypeUserId = null) {
        // Load Roles
        $.get(rolesApi, function (response) {
            const roleSelect = $('#editRoleId, #role_id');
            roleSelect.empty();
            // roleSelect.append('<option value="">Select Role</option>');
            response.data.forEach(role => {

                roleSelect.append(`<option value="${role.role_id}" ${role.role_id == selectedRoleId ? 'selected' : ''}>${role.name}</option>`);
            });
        });


        // Load TypeUsers
        $.get(typeUsersApi, function (response) {
            const typeUserSelect = $('#editTypeUserId, #type_user_id');
            // typeUserSelect.empty().append('<option value="">Select Type User</option>');
            typeUserSelect.empty()
            response.data.forEach(typeUser => {
                typeUserSelect.append(
                    `<option value="${typeUser.type_user_id}" ${typeUser.type_user_id == selectedTypeUserId ? 'selected' : ''}>${typeUser.name}</option>`
                );
            });
        });
    }


    $('#openAddUserModal').on('click', function () {
        $('#name, #email, #password').val('');
        $('#role_id, #type_user_id').empty();
        loadDropdowns();
        $('.modal-add-user').modal('show');
    });


    $('#saveUser').on('click', function () {
        const formData = {
            name: $('#name').val(),
            email: $('#email').val(),
            password: $('#password').val(),
            role_id: parseInt($('#role_id').val()),
            type_user_id: parseInt($('#type_user_id').val())
        };
        $.ajax({
            url: '/api/v1/users/create',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function () {
                showModalSuccess("User created successfully");
                $('#userTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
            },
            error: function (res) {
                showModalError(res.responseJSON.message);
            }
        });
    });


    $('#userTable').on('click', '.edit-btn', function () {
        const userId = $(this).data('id');
        $.get(`${apiUrl}/${userId}`, function (response) {
            const user = response.user || response;
            $('#editName').val(user.name);
            $('#editEmail').val(user.email);
            $('#editUserId').val(user.user_id);
            loadDropdowns(user.Role?.role_id, user.TypeUser?.type_user_id);
            $('.modal-edit-user').modal('show');
        });
    });


    $('#updateUserBtn').on('click', function () {
        const userId = $('#editUserId').val();
        const formData = {
            name: $('#editName').val(),
            email: $('#editEmail').val(),
            role_id: $('#editRoleId').val(),
            type_user_id: $('#editTypeUserId').val()
        };

        $.ajax({
            url: `${apiUrl}/${userId}`,
            method: 'PUT',
            data: JSON.stringify(formData),
            contentType: 'application/json',
            success: function () {
                showModalSuccess("User updated successfully");
                $('.modal-edit-user').modal('hide');
                $('#userTable').DataTable().ajax.reload();
            },
            error: function (res) {
                showModalError(res.responseJSON.message);
            }
        });
    });


    let userIdToDelete = null;


    $('#userTable').on('click', '.delete-btn', function () {
        userIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });

    $('#confirmDelete').on('click', function () {
        if (!userIdToDelete) {
            showModalError("Failed to find user to Delete");
            return;
        }

        $.ajax({
            url: `${apiUrl}/${userIdToDelete}`,
            method: 'DELETE',
            success: function (response) {
                if (response.success) {
                    showModalSuccess("User deleted successfully");
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete user");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (res) {
                showModalError(res.responseJSON.message);
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });

    $('#userTable').on('click', '.password-btn', function () {
        const userId = $(this).data('id');
        $('#editUserId2').val(userId);
        $.get(`${apiUrl}/${userId}`, function (response) {
            const user = response.user || response;
            $('#newPassword').val('');
            $('#confirmPassword').val('');
            $('.modal-edit-password').modal('show');
        });
    });

    $('#updatePasswordBtn').on('click', function () {
        const userId = $('#editUserId2').val();

        const formData = {
            new_password: $('#newPassword').val(),
            confirm_password: $('#confirmPassword').val()
        };
        $.ajax({
            url: `${apiUrl}/${userId}/change_password`,
            method: 'PUT',
            data: JSON.stringify(formData),
            contentType: 'application/json',
            success: function () {
                $.toast({
                    heading:  'Successfully!',
                    text: 'Change password successfully',
                    position: 'top-right',
                    loaderBg: '#ff6849',
                    icon: 'success',
                    hideAfter: 3500,
                    stack: 6
                });
                $('.modal-edit-password').modal('hide');
                $('#userTable').DataTable().ajax.reload();
            },
            error: function (res) {
                $.toast({
                    heading: 'Warning',
                    text: res.responseJSON.message,
                    position: 'top-right',
                    loaderBg: '#ff6849',
                    icon: 'error',
                    hideAfter: 3500

                });
            }
        });
    });

});

