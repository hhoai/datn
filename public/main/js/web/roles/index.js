$(document).ready(function () {
    const apiUrl = '/api/v1/roles';
    const permissionApiUrl = "/api/v1/permissions"

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

            {
                render: function (data, type, row) {
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.role_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.role_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddRoleModal').on('click', function () {
        $('#name').val('');
        $('#permissionsContainer').empty();

        $.ajax({
            url: permissionApiUrl,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const permissions = response.data;

                    permissions.forEach(permission => {
                        $('#permissionsContainer').append(`
                            <div class="form-check">
                                <input class="form-check-input" type="checkbox" value="${permission.permission_id}" id="permission-${permission.permission_id}">
                                <label class="form-check-label" for="permission-${permission.permission_id}">
                                    ${permission.name}
                                </label>
                            </div>
                        `);
                    });


                    $('.modal-add-role').modal('show');
                } else {
                    showModalError("Failed to load permissions");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching permissions");
            }
        });
    });

    $('#saveRole').on('click', function () {
        const roleName = $('#name').val();
        const selectedPermissions = [];


        $('#permissionsContainer .form-check-input:checked').each(function () {
            selectedPermissions.push(parseInt($(this).val(), 10));
        });

        if (!roleName) {
            showModalError("Please enter a role name");
            return;
        }

        $.ajax({
            url: '/api/v1/roles/create',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                name: roleName,
                permissions: selectedPermissions
            }),
            success: function (response) {
                if (response.code === 200) {
                    showModalSuccess("Role created successfully");
                    $('.modal-add-role').modal('hide');
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to create role");
                }
            },
            error: function (xhr) {
                showModalError("Error while saving role");
            }
        });
    });


    let roleIdToDelete = null;


    $('#userTable').on('click', '.delete-btn', function () {
        roleIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });

    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${roleIdToDelete}`, // URL API
            method: 'DELETE',
            success: function (response) {
                if (response.success) {
                    showModalSuccess("Role deleted successfully");
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete role");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to delete role");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });



    $('#userTable').on('click', '.edit-btn', function () {
        const roleId = $(this).data('id');
        $('.modal-edit-role').data('role-id', roleId);

        $.ajax({
            url: `${apiUrl}/${roleId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const role = response.data.role;
                    const assignedPermissions = response.data.permissions || [];
                    const allPermissions = response.data.allPermissions || [];

                    $('#editName').val(role.name);
                    $('#editPermissions').empty();

                    const assignedPermissionIds = assignedPermissions.map(p => p.Permission.permission_id);

                    if (allPermissions.length > 0) {
                        allPermissions.forEach(permission => {
                            const isChecked = assignedPermissionIds.includes(permission.permission_id) ? 'checked' : '';
                            $('#editPermissions').append(`
                            <div class="form-check">
                                <input class="form-check-input" type="checkbox" value="${permission.permission_id}" id="edit-permission-${permission.permission_id}" ${isChecked}>
                                <label class="form-check-label" for="edit-permission-${permission.permission_id}">
                                    ${permission.name}
                                </label>
                            </div>
                        `);
                        });
                    } else {
                        $('#editPermissions').append('<p>No permissions available</p>');
                    }

                    $('.modal-edit-role').modal('show');
                } else {
                    showModalError("Failed to fetch role data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching role details");
            }
        });
    });


    $('#saveEditRole').on('click', function () {
        const roleId = $('.modal-edit-role').data('role-id');
        const roleName = $('#editName').val();
        const selectedPermissions = [];

        $('#editPermissions .form-check-input:checked').each(function () {
            selectedPermissions.push(parseInt($(this).val(), 10));
        });

        if (!roleName) {
            showModalError("Please enter a role name");
            return;
        }
        const url = `${apiUrl}/${roleId}`;

        $.ajax({
            url: url,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify({
                name: roleName,
                permissions: selectedPermissions
            }),
            success: function (response) {
                if (response.code === 200) {
                    showModalSuccess("Role updated successfully");
                    $('.modal-edit-role').modal('hide');
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to update user");
                }
            },
            error: function (xhr) {
                showModalError("Failed to update user");
            }
        });
    });

});
