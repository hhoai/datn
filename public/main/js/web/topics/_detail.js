$(document).ready(function () {
    const apiUrl = '/api/v1/topics';
    var url = window.location.href;
    var parts = url.split('/');
    var topicID = parseInt(parts[parts.length - 1], 10);
    $.ajax({
        url: apiUrl + "/" + topicID + "/questions",
        type: 'GET',
        success: function (response) {
            const container = document.getElementById('questions-container');
            let data = response.data;
            data.forEach((item, questionIndex) => {
                const questionDiv = document.createElement('div');
                questionDiv.classList.add('form-group');

                // Create question label
                const questionLabel = document.createElement('label');
                questionLabel.classList.add('lb-question');
                questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point)`;
                questionDiv.appendChild(questionLabel);

                // Create options
                if (item.options && item.options.length > 0) {
                    switch (item.type_question_id){
                        case 1:{
                            item.options.forEach(option => {
                                const optionDiv = document.createElement('div');
                                optionDiv.classList.add('checkbox', 'lb-answer');

                                const optionInput = document.createElement('input');
                                optionInput.type = 'checkbox';
                                optionInput.id = `Checkbox_${item.topic_question_id}_${option.option_id}`;
                                optionInput.name = `check${item.topic_question_id}[]`
                                if (option.is_correct) optionInput.checked = true;

                                const optionLabel = document.createElement('label');
                                optionLabel.setAttribute('for', `Checkbox_${item.topic_question_id}_${option.option_id}`);
                                optionLabel.textContent = option.content;

                                optionDiv.appendChild(optionInput);
                                optionDiv.appendChild(optionLabel);
                                questionDiv.appendChild(optionDiv);
                            });
                            break;
                        }
                        //FILL BLANK
                        case 3: {
                            const selectBox = document.createElement('select');
                            selectBox.classList.add('form-control', 'lb-answer-inline');
                            selectBox.name = `${item.topic_question_id}`;
                            selectBox.id = `${item.question_id}`;

                            item.options.forEach(option => {
                                const optionElement = document.createElement('option');
                                optionElement.value = option.option_id;
                                optionElement.textContent = option.content;

                                if (option.is_correct) optionElement.selected = true;

                                selectBox.appendChild(optionElement);
                            });

                            questionLabel.textContent = `${questionIndex + 1}.  `;

                            const questionParts = item.content.split("...");
                            const container = document.createElement('div');
                            container.classList.add('question-container');

                            const beforeText = document.createTextNode(questionParts[0]);
                            container.appendChild(beforeText);

                            container.appendChild(selectBox);

                            const afterText = document.createTextNode(questionParts[1]);
                            container.appendChild(afterText);

                            questionLabel.appendChild(container);

                            const scoreText = document.createElement('span');
                            scoreText.textContent = ` (${item.score} point)`;
                            container.appendChild(scoreText);
                            break;
                        }



                        default:{
                            item.options.forEach(option => {
                                const optionDiv = document.createElement('div');
                                optionDiv.classList.add('radio', 'lb-answer');

                                const optionInput = document.createElement('input');
                                optionInput.type = 'radio';
                                optionInput.id = `Option_${item.topic_question_id}_${option.option_id}`;
                                optionInput.name = `group_${item.topic_question_id}`
                                if (option.is_correct) optionInput.checked = true;

                                const optionLabel = document.createElement('label');
                                optionLabel.setAttribute('for', `Option_${item.topic_question_id}_${option.option_id}`);
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