$(document).ready(function () {
    const apiUrl = '/api/v1/question_bank';
    const questionId = getId();
   const questionContent = localStorage.getItem('questionContent')
    if (questionId) {
        $.ajax({
            url: `${apiUrl}/${questionId}`,
            method: 'GET',
            success: function (response) {
                if (response.success) {
                    const question = response.data;
                    handleQuestionType(question.type_question_id,questionId);
                    $('#questionContent').text(questionContent);

                } else {
                    showModalError("Failed to fetch question data");
                }
            },
            error: function (xhr) {
                showModalError("Error fetching question data");
            }
        });
    } else {
        showModalError("No Question ID found in LocalStorage");
    }
});

function getHtmlForCase1() {
    return `
       <div class="form-group col-md-6">
       <h4 style="margin-top: 25px; color: red"><strong style="color: #0b0b0b">Question:</strong> <span id="questionContent"></span></h4>
       </div>
        <div class="form-group col-md-6">
            <label for="newAnswer">Answer:</label>
            <input type="text" id="newAnswer" class="form-control">
        </div>
        <div class="form-group col-md-6">
            <button id="addAnswerBtn" class="btn btn-primary mt-3">Add Answer</button>
            <h5 class="mt-4">List of answer:</h5>
             <ul id="answersContainer" class="list-group"></ul>
        </div>
        <div class="form-group col-md-6">
          <button id="save13" class="btn btn-primary mt-3">Save Answer</button>
       </div>
    `;
}
function getHtmlForCase2() {
    return `
        <div class="form-group col-md-6">
       <h4 style="margin-top: 25px; color: red"><strong style="color: #302828">Question:</strong> <span id="questionContent"></span></h4>
       </div>
       <div class="form-group col-md-6">
           <ul id="answersContainer" class="list-group"></ul>
       </div>
       <div class="form-group col-md-6">
          <button id="save2" class="btn btn-primary mt-3">Save Answer</button>
       </div>
    `;
}

function renderAnswers(questionType, answers) {
    const $container = $('#answersContainer');
    $container.empty();

    answers.forEach((answer, index) => {
        const answerMarkup =
            questionType === 1
                ? `
                     <li class="list-group-item answer-item d-flex justify-content-between align-items-center">
                    <span class="d-flex align-items-center">
                    ${answer.text}
                     <span class="badge bg-success ml-2" style="color: white; ${answer.correct ? '' : 'display: none;'}" data-badge="correct">
                     correct
                     </span>
                     </span>
                    <div class="d-flex align-items-center">
                    <input type="checkbox" class="form-check-input toggle-correct" ${answer.correct ? 'checked' : ''} data-index="${index}" style="position: static; opacity: 1; vertical-align: middle; width:25px; height: 30px; color: #09d70c">
                    <button class="btn btn-danger btn-sm ml-2 delete-answer" data-index="${index}">
                   <i class="fa fa-trash"></i>
                    </button>
                   </div>
                    </li>

                  `
                : questionType === 2
                    ? `
                   <li class="list-group-item answer-item d-flex justify-content-between align-items-center">
                     <div class="d-flex align-items-center">
                        <input type="radio" class="correct-true" ${answer.correct === true ? 'checked' : ''} data-index="${index}" id="true${index}" name="answer" style="width:20px; height:20px; margin-right:5px;">
                        <label for="true${index}" class="form-check-label">True</label>
                        <input type="radio" class="correct-false" ${answer.correct === false ? 'checked' : ''} data-index="${index}" id="false${index}" name="answer" style="width:20px; height:20px; margin-left:10px;">
                        <label for="false${index}" style="margin-bottom: 0">False</label>
                     </div>
                   </li>
                    `
                : questionType === 3
                        ?`
                      <li class="list-group-item answer-item d-flex justify-content-between align-items-center">
                    <span class="d-flex align-items-center">${answer.text}
                   <span class="badge bg-success ml-2" style="color: white; ${answer.correct ? '' : 'display: none;'}" data-badge="correct">
                    correct
                   </span>
                 </span>
                 <div class="d-flex align-items-center">
                <input type="radio" name="correctAnswer" class="form-check-input toggle-correct-radio" ${answer.correct ? 'checked' : ''} data-index="${index}" style="position: static; opacity: 1; vertical-align: middle; width:25px; height: 30px; color: #09d70c">
                <button class="btn btn-danger btn-sm ml-2 delete-answer" data-index="${index}">
                    <i class="fa fa-trash"></i>
                </button>
               </div>
                </li>

               </li>
               `
                        :
                        `
                  <li class="list-group-item answer-item d-flex justify-content-between align-items-center">
                    <span class="d-flex align-items-center">${answer.text}
                   <span class="badge bg-success ml-2" style="color: white; ${answer.correct ? '' : 'display: none;'}" data-badge="correct">
                    correct
                   </span>
                 </span>
                 <div class="d-flex align-items-center">
                <input type="radio" name="correctAnswer" class="form-check-input toggle-correct-radio" ${answer.correct ? 'checked' : ''} data-index="${index}" style="position: static; opacity: 1; vertical-align: middle; width:25px; height: 30px; color: #09d70c">
                <button class="btn btn-danger btn-sm ml-2 delete-answer" data-index="${index}">
                    <i class="fa fa-trash"></i>
                </button>
               </div>
                </li>

               </li>
               `;

        $container.append(answerMarkup);
    });

    $('.toggle-correct').on('change', function () {
        const index = $(this).data('index');
        const isChecked = $(this).is(':checked');
        answers[index].correct = isChecked;

        const badge = $(this).closest('li').find('[data-badge="correct"]');
        if (isChecked) {
            badge.show();
        } else {
            badge.hide();
        }
    });

    if (questionType === 3 || questionType === 1) {

        $('.answer-text').on('input', function () {
            const index = $(this).data('index');
            answers[index].text = $(this).val().trim();
        });


        $('.toggle-correct').on('change', function () {
            const index = $(this).data('index');
            answers[index].correct = $(this).is(':checked');
        });
    }

    $('.toggle-correct-radio').on('change', function () {
        const selectedIndex = $(this).data('index');

        answers.forEach((answer, index) => {
            answer.correct = index === selectedIndex;
        });

        renderAnswers(4, answers);
    });

    $('.delete-answer').on('click', function () {
        const index = $(this).data('index');
            answers.splice(index, 1);
            renderAnswers(questionType, answers);

    });
}


function handleQuestionType(type_question_id, questionId) {
    const container = $('.contentDetail');
    const answers = [];
    switch (type_question_id) {
        case 1:
            container.empty();
            container.append(getHtmlForCase1());
            fetchAnswers(questionId, answers, type_question_id);
            $('#addAnswerBtn').on('click', function () {
                const newAnswer = $('#newAnswer').val().trim();
                if (!newAnswer) {
                    showModalError("Please fill the answer");
                    return;
                }
                answers.push({ text: newAnswer, correct: false });
                renderAnswers(1, answers);
                $('#newAnswer').val('');
            });
            container.on('click', '#save13', function () {
                const updatedAnswers = answers.map(answer => ({
                    content: answer.text.trim(),
                    is_correct: answer.correct,
                }));

                $.ajax({
                    url: `/api/v1/question_bank/${questionId}/options`,
                    method: 'PUT',
                    contentType: 'application/json',
                    data: JSON.stringify(updatedAnswers),
                    success: function (response) {
                        if (response.success) {
                            showModalSuccess("Answer updated successfully");
                        } else {
                            showModalError("Failed to save answers");
                        }
                    },
                    error: function (xhr) {
                        showModalError("Error saving answers");
                    },
                });
            });
            break;

        case 2:
            container.empty();
            container.append(getHtmlForCase2());
            fetchAnswers(questionId, answers, type_question_id);
            container.on('click', '#save2', function () {
                const updatedAnswers = [];
                $('#answersContainer .answer-item').each(function () {
                    const $item = $(this);
                    const isCorrect = $item.find('.correct-true').is(':checked');
                    const optionContent = isCorrect ? 'True' : 'False';
                    updatedAnswers.push({
                        content: optionContent,
                        is_correct: isCorrect,
                    });
                });

                $.ajax({
                    url: `/api/v1/question_bank/${questionId}/options`,
                    method: 'PUT',
                    contentType: 'application/json',
                    data: JSON.stringify(updatedAnswers),
                    success: function (response) {
                        if (response.success) {
                            showModalSuccess("Answer updated successfully");
                        } else {
                            showModalError("Failed to save answers");
                        }
                    },
                    error: function (xhr) {
                        showModalError("Error saving answers");
                    },
                });
            });
            break;

        case 3:
            container.empty();
            container.append(getHtmlForCase1());
            fetchAnswers(questionId, answers, type_question_id);
            $('#addAnswerBtn').on('click', function () {
                const newAnswer = $('#newAnswer').val().trim();
                if (!newAnswer) {
                    showModalError("Please fill the answer");
                    return;
                }
                answers.push({ text: newAnswer, correct: false });
                renderAnswers(3, answers);
                $('#newAnswer').val('');
            });
            container.on('click', '#save13', function () {
                const updatedAnswers = answers.map(answer => ({
                    content: answer.text.trim(),
                    is_correct: answer.correct,
                }));

                $.ajax({
                    url: `/api/v1/question_bank/${questionId}/options`,
                    method: 'PUT',
                    contentType: 'application/json',
                    data: JSON.stringify(updatedAnswers),
                    success: function (response) {
                        if (response.success) {
                            showModalSuccess("Answer updated successfully");
                        } else {
                            showModalError("Failed to save answers");
                        }
                    },
                    error: function (xhr) {
                        showModalError("Error saving answers");
                    },
                });
            });
            break;


        case 4:
            container.empty();
            container.append(getHtmlForCase1());
            fetchAnswers(questionId, answers, type_question_id);
            $('#addAnswerBtn').on('click', function () {
                const newAnswer = $('#newAnswer').val().trim();
                if (!newAnswer) {
                    showModalError("Please fill the answer");
                    return;
                }
                answers.push({ text: newAnswer, correct: false });
                renderAnswers(4, answers);
                $('#newAnswer').val('');
            });
            container.on('click', '#save13', function () {
                const updatedAnswers = answers.map(answer => ({
                    content: answer.text.trim(),
                    is_correct: answer.correct,
                }));

                $.ajax({
                    url: `/api/v1/question_bank/${questionId}/options`,
                    method: 'PUT',
                    contentType: 'application/json',
                    data: JSON.stringify(updatedAnswers),
                    success: function (response) {
                        if (response.success) {
                            $.toast({
                                heading: 'Successfully!',
                                text: 'Answer updated successfully',
                                position: 'top-right',
                                loaderBg: '#ff6849',
                                icon: 'success',
                                hideAfter: 3500,
                                stack: 6,
                            });
                        } else {
                            showModalError("Failed to save answers: " + response.message);
                        }
                    },
                    error: function (xhr) {
                        showModalError("Error saving answers");
                    },
                });
            });
            break;

        default:
            showModalError('Unknown question type!');
    }
}

function fetchAnswers(questionId, answers, questionType) {
    $.ajax({
        url: `/api/v1/question_bank/${questionId}/options`,
        method: 'GET',
        success: function (response) {
            if (response.success) {
                if (Array.isArray(response.data)) {
                    const fetchedAnswers = response.data.map(option => {
                        return {
                            text: option.content || '',
                            correct: option.is_correct === true,
                        };
                    });
                    answers.push(...fetchedAnswers);

                    if (answers.length === 0 && questionType === 2) {
                        answers.push({ text: 'True False', correct: true });

                    }

                    renderAnswers(questionType, answers);
                } else {
                    showModalError("Options data format is invalid");
                }
            } else {
                showModalError("Failed to fetch answers");
            }
        },
        error: function (xhr) {
            showModalError("Error fetching options");
        },
    });
}

function getId(){
    var url = window.location.href;
    var parts = url.split('/');
    return parts[parts.length - 1];
}



