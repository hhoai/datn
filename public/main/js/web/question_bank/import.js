$(document).ready(function () {
    $('#importQuestions').click(function () {
        var formData = new FormData();
        var fileInput = $('#importFile')[0];

        formData.append('file', fileInput.files[0]);

        $.ajax({
            url: '/api/v1/question_bank/import',
            type: 'POST',
            data: formData,
            processData: false,
            contentType: false,
            success: function (response) {
                if (response.code === 200) {
                    location.reload();
                    $('.modal-import-questions').modal('hide');
                }

            },
            error: function () {
                showModalError(response.message);
                $('.modal-import-questions').modal('hide');
            }
        });
    });


});
