$(document).ready(function () {
    const apiUrl = '/api/v1/challenges';

    $('#challengeTable').DataTable({
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
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.challenge_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.challenge_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddChallengeModal').on('click', function () {
        $('#name').val('');

    });

    $('#saveChallenge').on('click', function () {
        const formData = {
            name: $('#name').val(),
        };
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                showModalSuccess("challenge created successfully");
                $('#challengeTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
                $('#name').val("");
            },
            error: function (xhr) {
                showModalError("Failed to create program");
            }
        });
    });

    let challengeIdToDelete = null;
    $('#challengeTable').on('click', '.delete-btn', function () {
        challengeIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${challengeIdToDelete}`,
            method: 'DELETE',
            success: function (response) {

                if (response.code === 200) {
                    showModalSuccess("Challenge deleted successfully");
                    $('#challengeTable').DataTable().ajax.reload();
                } else {
                    showModalError("Failed to delete challenge");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Fail to delete challenge");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#challengeTable').on('click', '.edit-btn', function () {
        const challengeId = $(this).data('id');
        $('.modal-edit-challenge').data('challenge-id', challengeId);

        $.ajax({
            url: `${apiUrl}/${challengeId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const challenge = response.data;
                    $('#editName').val(challenge.name);
                    $('.modal-edit-challenge').modal('show');
                } else {
                    showModalError(response.message || "Failed to fetch challenge data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching challenge details");
            }
        });
    });


    $('#saveEditChallenge').on('click', function () {
        const  challenge = $('.modal-edit-challenge').data('challenge-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${challenge}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("challenge updated successfully");
                    $('#challengeTable').DataTable().ajax.reload();
                    $('.modal-edit-challenge').modal('hide');
                } else {
                    showModalError("Failed to update challenge");
                }
            },
            error: function (xhr) {
                showModalError("Failed to update challenge");
            }
        });
    });

});