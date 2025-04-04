$(document).ready(function () {
    const apiUrl = '/api/v1/skills';

    $('#skillTable').DataTable({
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
                        <button class="btn btn-primary btn-sm edit-btn" data-id="${row.skill_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${row.skill_id}"><i class="fa fa-trash" aria-hidden="true"></i> </button>`;
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

    $('#openAddSkillModal').on('click', function () {
        $('#name').val('');
        $('.modal-add-skill').modal('show');
    });

    $('#saveSkill').on('click', function () {
           const  skillName=  $('#name').val();

        if (!skillName) {
            showModalError("Please enter a skill name");
            return;
        }
        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                name: skillName,
            }),
            success: function (response) {
                showModalSuccess("Skill created successfully");
                $('#skillTable').DataTable().ajax.reload();
                $('.modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Failed to create skill");
            }
        });
    });

    let skillIdToDelete = null;
    $('#skillTable').on('click', '.delete-btn', function () {
        skillIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');
    });
    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${skillIdToDelete}`,
            method: 'DELETE',
            success: function (response) {
                if (response.code === 200) {
                    showModalSuccess("Skill deleted successfully");
                    $('#skillTable').DataTable().ajax.reload();
                    $('.delete-confirm-modal').modal('hide');
                }
            },
            error: function (xhr) {
                showModalError("Fail to delete skill");
                // $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#skillTable').on('click', '.edit-btn', function () {
        const skillId = $(this).data('id');
        $('.modal-edit-skill').data('skill-id', skillId);

        $.ajax({
            url: `${apiUrl}/${skillId}`,
            method: 'GET',
            success: function (response) {
                if (response.code === 200) {
                    const skill = response.data;
                    $('#editName').val(skill.name);
                    $('.modal-edit-skill').modal('show');
                } else {
                    showModalError(response.message || "Failed to fetch skill data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching skill details");
            }
        });
    });


    $('#saveEditSkill').on('click', function () {
        const  skillId = $('.modal-edit-skill').data('skill-id');
        const formData = {
            name: $('#editName').val(),
        };
        $.ajax({
            url: `${apiUrl}/${skillId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {

                    showModalSuccess("Skill updated successfully");
                    $('#skillTable').DataTable().ajax.reload();
                    $('.modal-edit-skill').modal('hide');
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