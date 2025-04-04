$(document).ready(function () {
    const apiUrl = '/api/v1/topics';
    var url = window.location.href;
    var parts = url.split('/');
    var topicID = parseInt(parts[parts.length - 1], 10);
    var assignmentID = parseInt(parts[parts.length - 3], 10);

    let currentIndex = 0;
    let questions = [];
    let answersState = {};
    let status = false;

    function renderQuestion(index) {
        const container = $('#question-container');
        container.empty();

        const question = questions[index];
        const questionDiv = $('<div>').addClass('form-group');

        if(question.type_question_id !== 3){
            const questionLabel = $('<label>').addClass('lb-question')
                .text(`${index + 1}. ${question.content} (${question.score} point)`);

            if (status) {
                if (question.is_correct) {

                    questionLabel.css({
                        color: '#365cef',
                        'font-style': 'oblique',
                        'padding-left': '10px',
                    }).text(`${index + 1}. ${question.content} (${question.score} point) Correct`);
                } else {
                    questionLabel.css({
                        color: '#e7573e',
                        'font-style': 'oblique',
                        'padding-left': '10px',
                    }).text(`${index + 1}. ${question.content} (${question.score} point) Incorrect`);
                }
            }

            questionDiv.append(questionLabel);
        }else {
        }

        if (question.options && question.options.length > 0) {
            const savedState = answersState[question.question_id] || {};

            switch (question.type_question_id) {
                case 1: // Checkbox
                    question.options.forEach(option => {
                        const optionDiv = $('<div>').addClass('checkbox lb-answer');
                        const optionInput = $('<input>')
                            .attr({
                                type: 'checkbox',
                                id: option.option_id,
                                name: question.topic_question_id,
                                'data-id': question.question_id,
                            })
                            .prop('checked', savedState[option.option_id] || false || (status && option.is_correct))
                            .on('change', function (){
                                saveCurrentState();
                            })


                        const optionLabel = $('<label>').attr('for', option.option_id).text(option.content);

                        optionDiv.append(optionInput).append(optionLabel);
                        questionDiv.append(optionDiv);
                    });
                    break;

                case 3: // Dropdown
                    const questionParts = question.content.split('...');
                    const questionContainer = $('<label>').addClass('lb-question d-flex align-items-center');

                    if (questionParts[0]) {
                        questionContainer.append($('<span>').text(`${index + 1}. ${questionParts[0]}`));
                    }

                    const selectBox = $('<select>').addClass('form-control lb-answer-inline mx-2')
                        .attr({ name: question.topic_question_id, id: question.question_id })
                        .append($('<option>').attr('value', '').text(''));

                    question.options.forEach(option => {
                        const optionElement = $('<option>')
                            .attr('value', option.option_id)
                            .text(option.content)
                            .prop('selected', savedState[option.option_id] || false || (status && option.is_correct));
                        selectBox.append(optionElement);
                    });

                    questionContainer.append(selectBox);
                    if (questionParts[1]) {
                        questionContainer.append($('<span>').text(questionParts[1]));
                    }
                    questionContainer.append($('<span>').text(`(${question.score} point)`));
                    questionDiv.append(questionContainer);

                    if (status) {
                        const selectedOptionId = selectBox.val();
                        const selectedOption = question.options.find(option => option.option_id === Number(selectedOptionId));

                        const isCorrect = selectedOption ? selectedOption.is_correct : false;

                        if (!isCorrect || !question.is_correct) {
                            questionContainer.css('color', '#e90909');
                        } else {
                        }

                        const resultLabel = $('<span>')
                            .addClass('ml-2')
                            .css({
                                color: isCorrect && question.is_correct ? '#365cef' : '#e90909',
                                'font-style': 'oblique',
                            })
                            .text(isCorrect && question.is_correct ? 'Correct' : 'Incorrect');

                        questionContainer.append(resultLabel);
                    }
                    break;


                default: // Radio
                    question.options.forEach(option => {
                        const optionDiv = $('<div>').addClass('radio lb-answer');
                        const optionInput = $('<input>')
                            .attr({
                                type: 'radio',
                                id: option.option_id,
                                name: question.topic_question_id,
                                'data-id': question.question_id,
                            })
                            .prop('checked', savedState[option.option_id] || false || (status && option.is_correct))
                            .on('change', function (){
                                saveCurrentState();
                            })


                        const optionLabel = $('<label>').attr('for', option.option_id).text(option.content);


                        optionDiv.append(optionInput).append(optionLabel);
                        questionDiv.append(optionDiv);
                    });
                    break;
            }
        } else {
            questionDiv.append('<p>No options available</p>');
        }

        container.append(questionDiv);
        updateNavigation();
    }



    function updateNavigation() {
        $('#btn-previous').toggle(currentIndex > 0);
        $('#btn-next').toggle(currentIndex < questions.length - 1);
    }

    function saveCurrentState() {
        const container = $('#question-container');
        const inputs = container.find('input');
        const selects = container.find('select');
        const questionId = questions[currentIndex].question_id;

        const state = {};
        let isAnswered = false;

        inputs.each(function () {
            const input = $(this);
            state[input.attr('id')] = input.prop('checked');
            if (input.prop('checked')) isAnswered = true;
        });

        selects.each(function () {
            const select = $(this);
            const selectedValue = select.val();
            if (selectedValue !== '') {
                state[selectedValue] = true;
                isAnswered = true;
            }
        });

        answersState[questionId] = state;
        questions[currentIndex].isAnswered = isAnswered;
        renderQuestionList();
    }

    function renderQuestionList() {
        const navContainer = $('#question-navigation');
        navContainer.empty();

        questions.forEach((question, index) => {
            const navItem = $('<div>')
                .addClass('question-item col-4')
                .text(`${index + 1}`)
                .on('click', function () {
                    saveCurrentState();
                    currentIndex = index;
                    renderQuestion(currentIndex);
                });

            if (status && question.is_correct !== undefined) {
                navItem.css('background-color', question.is_correct ? '#42b729' : '#e90909');
            } else if (question.isAnswered) {
                navItem.css({
                    'background-color': '#fbff00',
                    'color': '#080000',
                    'border-color': '#FFC107'
                });
            }

            navContainer.append(navItem);
        });
    }

    $.ajax({
        url: `${apiUrl}/${topicID}/student/questions/${assignmentID}`,
        type: 'GET',
        success: function (response) {
            questions = response.data.questions;
            status = response.data.status;


            $("#total-score").text(response.data.total_score);
            $("#minimum-score").text(response.data.minimum_score);

            if(status){
                $('#btn-submit').hide();
                $("#score").text(response.data.score);
            }

            renderQuestionList();
            renderQuestion(currentIndex);

        },
        error: function (xhr) {
            showModalError(xhr.responseJSON.message);
        }
    });

    $('#btn-previous').on('click', function () {
        if (currentIndex > 0) {
            saveCurrentState();
            currentIndex--;
            renderQuestion(currentIndex);
        }
    });

    $('#btn-next').on('click', function () {
        if (currentIndex < questions.length - 1) {
            saveCurrentState();
            currentIndex++;
            renderQuestion(currentIndex);
        }
    });

    $('#btn-submit').on('click', function () {
        $('.submit-confirm-modal').modal('show');
    });

    $('#confirmSubmit').on('click', function (){
        saveCurrentState();
        const groupedData = [];
        questions.forEach(question => {
            const savedState = answersState[question.question_id] || {};
            const options = [];

            question.options.forEach(option => {
                const isSelected = savedState[option.option_id] || false;
                options.push({
                    option_id: option.option_id,
                    is_correct: isSelected
                });
            });

            if (options.length > 0) {
                groupedData.push({
                    topic_question_id: question.topic_question_id,
                    topic_id: topicID,
                    question_id: question.question_id,
                    options: options
                });
            }
        });

        $.ajax({
            url: `${apiUrl}/${topicID}/student/questions/${assignmentID}`,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(groupedData),
            success: function (response) {
                if (response.code === 200) {
                    status = true;
                    renderQuestionList();
                    renderQuestion(currentIndex);
                    $('btn-submit').hide();

                    let title = `Điểm của bạn là: ${response.data.score}`;
                    let content = "Chúc mừng bạn đã vượt qua bài kiểm tra trắc nghiệm";
                    swal({
                        title: title,
                        text: content,
                        type: "success",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "OK",
                        closeOnConfirm: true
                    }, function () {
                        $('.submit-confirm-modal').modal('hide');
                        window.location.href = url;

                    });
                } else {
                    let title = `Điểm của bạn là: ${response.data.score}`;
                    let content = "Rất tiếc, bạn không đủ điểm để vượt qua bài kiểm tra trắc nghiệm";
                    swal(title, content, "error");
                    $('.submit-confirm-modal').modal('hide');
                }
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });
     })
});
