
$(document).ready(function () {
    const apiUrl = '/api/v1/topics';
    var url = window.location.href;
    var parts = url.split('/');
    var topicID = parseInt(parts[parts.length - 1], 10);
    var assignmentID = parseInt(parts[parts.length - 3], 10);
    $.ajax({
        url: apiUrl + "/" + topicID + "/student/questions/"+assignmentID,
        type: 'GET',
        success: function (response) {
            const container = document.getElementById('questions-container');
            let data = response.data.questions;
            let status = response.data.status;
            if(status){
                $("#btn-submit").css("display","none");
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
                if(status === false){
                    questionLabel.classList.add('lb-question');
                    questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point)`;
                }else if(status){
                    if(item.is_correct){
                        questionLabel.classList.add('lb-question-true');
                        questionLabel.textContent = `${questionIndex + 1}. ${item.content} (${item.score} point) Correct`;
                    }else{
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
                            container.classList.add('question_content');

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

    function getInputValues(topicID) {
        const container = document.getElementById("questions-container");
        const inputs = container.querySelectorAll("input");
        const selects = container.querySelectorAll("select");

        const groupedData = {};

        inputs.forEach((input) => {
            const topicQuestionId = parseInt(input.name);
            const questionId = parseInt(input.getAttribute('data-id'));

            if (!groupedData[topicQuestionId]) {
                groupedData[topicQuestionId] = {
                    topic_question_id: topicQuestionId,
                    topic_id: topicID,
                    question_id: questionId,
                    options: [],
                };
            }

            const option = {
                option_id: parseInt(input.id),
                is_correct: input.checked,
            };

            groupedData[topicQuestionId].options.push(option);

        });

        selects.forEach((select) => {
            const topicQuestionId = parseInt(select.name);
            const questionId = parseInt(select.getAttribute('id'));

            if (!groupedData[topicQuestionId]) {
                groupedData[topicQuestionId] = {
                    topic_question_id: topicQuestionId,
                    topic_id: topicID,
                    question_id: questionId,
                    options: [],
                };
            }

            const options = Array.from(select.options);
            options.forEach((option) => {
                groupedData[topicQuestionId].options.push({
                    option_id: parseInt(option.value),
                    is_correct: option.selected,
                });
            });
        });
        return Object.values(groupedData);


    }

    $("#btn-submit").on("click", function () {

        let data = getInputValues(topicID);
        $.ajax({
            url: apiUrl + "/" + topicID + "/student/questions/"+assignmentID,
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (response) {
                if(response.code === 200){
                    let title = "Điểm của bạn là: "+response.data.score;
                    let content = "Chúc mừng bạn đã vượt qua bài kiểm tra trắc nghiệm";

                    swal({
                        title: title,
                        text: content,
                        type: "success",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "OK",
                        closeOnConfirm: true
                    }, function(){
                        window.location.href = url;
                    });
                }else{
                    let title = "Điểm của bạn là: "+response.data.score;
                    let content = "Rất tiếc, bạn không đủ điểm để vượt qua bài kiểm tra trắc nghiệm";
                    swal(title, content, "success");
                }
            },
            error: function (xhr) {
                showModalError(xhr.responseJSON.message);
            }
        });
    });
});
