$(document).ready(function () {
    const apiUrl = '/api/v1/programs';
    
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
            {data:'program_code'},
            // { data: 'name' },
            {
                render: function (data, type, row) {
                    return `
                        <a href="/programs/${row.program_id}/details">${row.name}</a>`
                }
            },
            {
                render: function (data, type, row) {
                    return `
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.program_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.program_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddProgramModal').on('click', function () {
        $('#name').val('');
        $('.modal-add-program').modal('show');
    });

    $('#saveProgram').on('click', function () {
        const formData = {
            name: $('#name').val(),

        };
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("Program created successfully");
                $('#userTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to create program");
            }
        });
    });

    let programIdToDelete = null;
    $('#userTable').on('click', '.delete-btn', function () {
        programIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${programIdToDelete}`, // URL API
            method: 'DELETE',
            beforeSend: function () {
            },
            success: function (response) {

                if (response.success) {
                    showModalSuccess("Program deleted successfully");
                    $('#userTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete program");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Fail to delete role");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#userTable').on('click', '.edit-btn', function () {
        const programId = $(this).data('id');
        $('.modal-edit-program').data('program-id', programId);

        $.ajax({
            url: `${apiUrl}/${programId}`,
            method: 'GET',
            success: function (response) {
                if (response.success) {
                    const program = response.user;
                    $('#editName').val(program.name);
                    $('.modal-edit-program').modal('show');
                } else {
                    showModalError("Failed to fetch program data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching program details");
            }
        });
    });


    $('#saveEditProgram').on('click', function () {
        const programId = $('.modal-edit-program').data('program-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${programId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Program updated successfully");
                    $('#userTable').DataTable().ajax.reload();
                    $('.modal-edit-program').modal('hide');
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