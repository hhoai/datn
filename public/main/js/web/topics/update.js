$(document).ready(function () {
    const programApi = '/api/v1/programs';
    const levelApi = '/api/v1/levels';
    const skillApi = '/api/v1/skills';
    const challengeApi = '/api/v1/challenges';
    const typeQuestionApi = '/api/v1/type-questions';
    const questionBankApi = '/api/v1/question_bank';
    const apiUrl = '/api/v1/topics';
    let lastFilters = {};
    var existingQuestionIDs = [];

    var totalScore = 0;
    function updateTotalScore(scoreChange){
        totalScore += scoreChange
        $('#totalScore').val(totalScore)
    }

    const topicId = getId();

    $.ajax({
        url: apiUrl +`/`+topicId,
        method: 'GET',
        success: function (response) {
            if (response.code === 200) {
                const topic = response.data.topic;
                const questions = response.data.questions;

                $('#name').val(topic.name);
                $('#description').val(topic.description);
                $('#totalScore').val(topic.total_score);
                $('#minimumScore').val(topic.minimum_score);


                const areaImprimir = $('#areaImprimir');
                areaImprimir.empty();
                let index = 1;
                questions.forEach(question => {

                    const questionContent = question.Question.content;
                    const questionScore = question.Question.score;
                    if (!existingQuestionIDs.includes(question.question_id)) {
                        existingQuestionIDs.push(question.question_id);

                    }

                    areaImprimir.append(`
                        <div class="uk-margin" data-question_id="${question.question_id}" data-score="${questionScore}" style="user-select: none;">
                            <div class="uk-card uk-card-hover uk-card-default uk-card-body uk-card-small">
                                <span class="uk-flex uk-flex-left uk-flex-middle">${questionContent}</span>
                                <span class="uk-flex uk-flex-right send-list">
                                    <a href="#" class="uk-icon-button uk-button-primary uk-icon" uk-icon="reply"></a>
                                </span>
                            </div>
                        </div>`);
                    index++;

                });
                questions.forEach(function(question) {
                    totalScore += question.Question.score ;
                    $('#totalScore').val(totalScore);
                });

            } else {
                showModalError("Failed to fetch topic data");
            }
        },
        error: function (xhr) {
            showModalError("Failed to fetch topic data");
        }
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

    $(document).on('click', 'a[uk-icon]', function () {
        const icon = $(this).attr("uk-icon");
        const questionID = $(this).closest('.uk-margin').data('question_id');
        const questionScore = $(this).closest('.uk-margin').data('score');

        if (icon === "forward") {
            $(this).closest(".uk-margin").appendTo("#areaImprimir");
            $(this).attr("uk-icon", "reply");
            if (!existingQuestionIDs.includes(questionID)) {
                existingQuestionIDs.push(questionID);
                updateTotalScore(questionScore);

            }
        } else if (icon === "reply") {
            $(this).closest(".uk-margin").appendTo("#areaAgenda");
            $(this).attr("uk-icon", "forward");
            const index = existingQuestionIDs.indexOf(questionID);
            if (index !== -1 ) {
                existingQuestionIDs.splice(index, 1);
                updateTotalScore(-questionScore)

            }

        }
    });




    function searchQuestionBank(filters = {}) {
        $.ajax({
            url: questionBankApi + '/search',
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(filters),
            success: function (response) {
                const areaAgenda = $('#areaAgenda');
                areaAgenda.empty();

                if (!response.data || response.data.length === 0) {
                    areaAgenda.append('<p>No questions found.</p>');
                    return;
                }

                const filteredQuestions = response.data.filter(question =>
                    !existingQuestionIDs.includes(question.question_id)
                );

                if (filteredQuestions.length === 0) {
                    return;
                }


                filteredQuestions.forEach((question, index) => {
                    const questionHTML = `
                <div class="uk-margin" data-question_id="${question.question_id}" data-score="${question.score}">
                    <div class="uk-card uk-card-hover uk-card-default uk-card-body uk-card-small">
                        <span class="uk-flex uk-flex-left uk-flex-middle">${question.content}</span>
                        <span class="uk-flex uk-flex-right send-list">
                            <a href="#" class="uk-icon-button uk-button-primary" uk-icon="forward"></a>
                        </span>
                    </div>
                </div>`;
                    areaAgenda.append(questionHTML);
                });

                UIkit.update(document.getElementById('areaAgenda'));
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });
    }
    loadDropdowns();
    searchQuestionBank();
    $('#btnSearch').on('click', function () {
        const filters = {
            type_question_id: parseInt($('#type_question_id').val()),
            program_id: parseInt($('#program_id').val()),
            level_id: parseInt($('#level_id').val()),
            skill_id: parseInt($('#skill_id').val()),
            challenge_id: parseInt($('#challenge_id').val()),
        };

        Object.keys(filters).forEach(key => {
            if (!filters[key]) {
                delete filters[key];
            }
        });

        if (JSON.stringify(filters) === JSON.stringify(lastFilters)) {
            return;
        }
        lastFilters = { ...filters };

        searchQuestionBank(filters);
    });


    $('#btnUpdate').on('click', function () {
        const name = $('#name').val();
        const description = $('#description').val();

        let totalScore = parseInt($('#totalScore').val());
        let minimumScore = parseInt($('#minimumScore').val());


        // if (!name || !description || existingQuestionIDs.length === 0) {
        //     showModalError("Please fill all fields and select at least one question");
        //     return;
        // }

        if (!name || !description) {
            showModalError("Please fill all fields");
            return;
        }

        const updatedTopic = {
            name: name,
            description: description,
            total_score: totalScore,
            minimum_score: minimumScore,
            question_ids: existingQuestionIDs
        };


        $.ajax({
            url: `${apiUrl}/${topicId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(updatedTopic),
            success: function (response) {
                if (response.code === 200) {
                    showModalSuccess("Topic updated successfully");
                    setTimeout(()=> {
                        window.location.href = "/topics"
                    }, 2000)
                } else {
                    showModalError(response.message);
                }
            },
            error: function (response) {
                showModalError(response.message);
            }
        });
    });


});

function getId() {
    var url = window.location.href;
    var parts = url.split('/');
    return parts[parts.length - 1];
}