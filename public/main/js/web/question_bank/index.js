$(document).ready(function () {
    const apiUrl = '/api/v1/question_bank';
    const programApi = '/api/v1/programs';
    const levelApi = '/api/v1/levels';
    const skillApi = '/api/v1/skills';
    const challengeApi = '/api/v1/challenges';
    const typeQuestionApi = '/api/v1/type-questions';

    $('#questionTable').DataTable({
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
            {
                data: 'content',
                render: function (data) {
                    return `<div style="max-width: 300px; white-space: normal;">${data}</div>`;
                },
                orderable: false
            },
            { data: 'TypeQuestion.name',orderable: false},
            { data: 'score',orderable: false},
            { data: 'Program.name',orderable: false},
            { data: 'Level.name',orderable: false},
            { data: 'Challenge.name',orderable: false},
            { data: 'Skill.name',orderable: false},
            {
                render: function (_, __, row) {
                    return `
                   <button class="btn btn-primary btn-sm edit-btn" title="Edit" data-id="${row.question_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>
                   </button>
                    <button class="btn btn-danger btn-sm delete-btn" title="Delete" data-id="${row.question_id}"><i class="fa fa-trash" aria-hidden="true"></i>
                   </button>
                   <a href="/question_bank/${row.question_id}" title="Answer" class="btn btn-secondary btn-sm answer-link"  data-id="${row.question_id}" data-content="${row.content}"><i class="fa fa-link"></i>
                   </a>
                `;
                },
                orderable: false
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


    function loadDropdowns(selectedProgramId = null, selectedLevelId = null, selectedSkillId = null, selectedChallengeId = null, selectedTypeQuestionId = null) {

        $.get(programApi, function (response) {
            const programSelect = $('#editProgramId, #program_id');
            programSelect.empty();
            programSelect.append('<option value="">Select Program</option>');
            response.data.forEach(program => {
                programSelect.append(
                    `<option value="${program.program_id}" ${program.program_id == selectedProgramId ? 'selected' : ''}>
                        ${program.name}
                    </option>`
                );
            });
        });

        $.get(levelApi, function (response) {
            const levelSelect = $('#editLevelId, #level_id');
            levelSelect.empty();
            levelSelect.append('<option value="">Select Level</option>');
            response.data.forEach(level => {
                levelSelect.append(
                    `<option value="${level.level_id}" ${level.level_id == selectedLevelId ? 'selected' : ''}>
                        ${level.name}
                    </option>`
                );
            });
        });
        $.get(skillApi, function (response) {
            const skillSelect = $('#editSkillId, #skill_id');
            skillSelect.empty();
            skillSelect.append('<option value="">Select skill</option>');
            response.data.forEach(skill => {
                skillSelect.append(
                    `<option value="${skill.skill_id}" ${skill.skill_id == selectedSkillId ? 'selected' : ''}>
                        ${skill.name}
                    </option>`
                );
            });
        });
        $.get(challengeApi, function (response) {
            const challengeSelect = $('#editChallengeId, #challenge_id');
            challengeSelect.empty();
            challengeSelect.append('<option value="">Select Challenge</option>');
            response.data.forEach(challenge => {
                challengeSelect.append(
                    `<option value="${challenge.challenge_id}" ${challenge.challenge_id == selectedChallengeId ? 'selected' : ''}>
                        ${challenge.name}
                    </option>`
                );
            });
        });

        $.get(typeQuestionApi, function (response) {
            const typeQuestionSelect = $('#editTypeQuestionId, #type_question_id');
            typeQuestionSelect.empty();
            typeQuestionSelect.append('<option value="">Select Type Question </option>');
            response.data.forEach(typeQuestion => {
                typeQuestionSelect.append(
                    `<option value="${typeQuestion.type_question_id}" ${typeQuestion.type_question_id == selectedTypeQuestionId ? 'selected' : ''}>
                        ${typeQuestion.name}
                    </option>`
                );
            });
        });



    }

    $('#openAddQuestionModal').on('click', function () {
        $('#content, #score').val('');
        $('#program_id, #level_id, #challenge_id, #type_question_id, #skill_id').empty();
        loadDropdowns();
    });

    $('#saveQuestion').on('click',function () {
        const questionData = {
            content: $('#content').val(),
            score: parseInt($('#score').val()),
            type_question_id: parseInt($('#type_question_id').val()),
            program_id: parseInt($('#program_id').val()),
            level_id: parseInt($('#level_id').val()),
            skill_id: parseInt($('#skill_id').val()),
            challenge_id: parseInt($('#challenge_id').val()),
        };

        $.ajax({
            url: `${apiUrl}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(questionData),
            success: function () {
                showModalSuccess("Create Successfully");
                $('#questionTable').DataTable().ajax.reload();
                $('.modal-add-question').modal('hide');
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });

    });

    let questionIdToDelete = null;
    $('#questionTable').on('click', '.delete-btn', function (){
        questionIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');

    })

    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${questionIdToDelete}`,
            method: 'DELETE',
            success: function (response) {
                if (response.success) {
                    showModalSuccess("Delete Successfully");
                    $('#questionTable').DataTable().ajax.reload();
                } else {
                    showModalError("Fail to delete");
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError("Fail to delete question");
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });


    $('#questionTable').on('click', '.edit-btn', function (){
       const questionId = $(this).data('id');
        $.get(`${apiUrl}/${questionId}`, function (response){
            const question = response.data || response;
            $('#editContent').val(question.content);
            $('#editScore').val(question.score);
            loadDropdowns(question.program_id,question.level_id,question.skill_id,question.challenge_id, question.type_question_id);
            $('#editQuestion').data('id', questionId);
            $('.modal-edit-question').modal('show');
        });
    });

    $('#editQuestion').on('click', function (){
       const questionId = $(this).data('id');

       const updateData = {
           content: $('#editContent').val(),
           score: parseInt($('#editScore').val()),
           program_id: parseInt($('#editProgramId').val()),
           level_id: parseInt($('#editLevelId').val()),
           skill_id: parseInt($('#editSkillId').val()),
           challenge_id: parseInt($('#editChallengeId').val()),
           type_question_id: parseInt($('#editTypeQuestionId').val()),

       };

       $.ajax({
           url: `${apiUrl}/${questionId}`,
           method: 'PUT',
           contentType: 'application/json',
           data: JSON.stringify(updateData),
           success: function (response){
               showModalSuccess("Question update successfully");
               $('#questionTable').DataTable().ajax.reload();
               $('.modal-edit-question').modal('hide');
           },
           error: function (xhr) {
               showModalError("Failed to update course");
           }
       })

    });

    $('#questionTable').on('click', '.answer-link', function () {
        const questionId = $(this).data('id');
        const questionContent = $(this).data('content')
        localStorage.setItem('questionId', questionId);
        localStorage.setItem('questionContent', questionContent)
    });

});

