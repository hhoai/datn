$(document).ready(function () {
    const apiUrl = '/api/v1/topics';
    var url = window.location.href;
    var parts = url.split('/');
    var topicID = parseInt(parts[parts.length - 1], 10);
    var assignmentID = parseInt(parts[parts.length - 3], 10);
    var userID = parseInt(parts[parts.length - 5], 10);
    $.ajax({
        url: apiUrl + "/" + topicID + "/student/questions/" + assignmentID + "/" + userID,
        type: 'GET',
        success: function (response) {
            const container = document.getElementById('questions-container');
            let data = response.data.questions;
            let status = response.data.status;
            if (status) {
                $("#btn-submit").css("display", "none");
            }
            // let questionHTML = "";
            $("#score").text(response.data.score);
            $("#total-score").text(response.data.total_score);
            $("#minimum-score").text(response.data.minimum_score);
            data.forEach((item, questionIndex) => {
                const questionDiv = document.createElement('div');
                questionDiv.classList.add('form-group');

                // Create question label
                const questionLabel = document.createElement('label');
                if (status === false) {
                    questionLabel.classList.add('lb-question');
                    questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point)`;
                } else if (status) {
                    if (item.is_correct) {
                        questionLabel.classList.add('lb-question-true');
                        questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point) Correct`;
                    } else {
                        questionLabel.classList.add('lb-question-false');
                        questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point) Incorrect`;
                    }

                }

                questionDiv.appendChild(questionLabel);

                // Create options
                if (item.options && item.options.length > 0) {
                    switch (item.type_question_id) {
                        case 1: {
                            item.options.forEach(option => {
                                const optionDiv = document.createElement('div');
                                optionDiv.classList.add('checkbox', 'lb-answer');

                                const optionInput = document.createElement('input');
                                optionInput.type = 'checkbox';
                                optionInput.id = `${option.option_id}`;
                                optionInput.name = `${item.topic_question_id}`;
                                optionInput.setAttribute('data-id', `${item.question_id}`);
                                if (option.is_correct) optionInput.checked = true;

                                const optionLabel = document.createElement('label');
                                optionLabel.setAttribute('for', `${option.option_id}`);
                                optionLabel.textContent = option.content;

                                optionDiv.appendChild(optionInput);
                                optionDiv.appendChild(optionLabel);
                                questionDiv.appendChild(optionDiv);
                            });
                            break;
                        }
                        default: {
                            item.options.forEach(option => {
                                const optionDiv = document.createElement('div');
                                optionDiv.classList.add('radio', 'lb-answer');

                                const optionInput = document.createElement('input');
                                optionInput.type = 'radio';
                                optionInput.id = `${option.option_id}`;
                                optionInput.name = `${item.topic_question_id}`;
                                optionInput.setAttribute('data-id', `${item.question_id}`);
                                if (option.is_correct) optionInput.checked = true;

                                const optionLabel = document.createElement('label');
                                optionLabel.setAttribute('for', `${option.option_id}`);
                                optionLabel.textContent = option.content;

                                optionDiv.appendChild(optionInput);
                                optionDiv.appendChild(optionLabel);
                                questionDiv.appendChild(optionDiv);
                            });
                        }
                    }

                } else {
                    const noOptions = document.createElement('p');
                    noOptions.textContent = 'No options available';
                    questionDiv.appendChild(noOptions);
                }

                container.appendChild(questionDiv);
            });
        },
        error: function (xhr) {
            showModalError(xhr.responseJSON.message);
        }
    });
});