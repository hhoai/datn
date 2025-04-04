
$(document).ready(function () {
    const clearErrorMessage = () => {
        $(".error-text").text("");
    };

    $(".form").on("submit", function (event) {
        event.preventDefault();
        var formData = $(this).serialize();
        $.ajax({
            url: "/api/v1/login",
            method: "POST",
            data: formData,
        })
            .done(function (response) {
                if (response.code === 200) {
                    window.location.href = "/dashboard";
                }

            })
            .fail(function (xhr) {
                
                showModalError(xhr.responseJSON.message);
            });
    });

    $("#fg-pass").on("click", function () {
        var formData = $(".form").serialize();
        $.ajax({
            url: "/api/v1/forget-password",
            method: "POST",
            data: formData,
        })
            .done(function (response) {
                if (response.code === 200) {
                    showModalSuccess(response.message);
                }

            })
            .fail(function (xhr) {
                
                showModalError(xhr.responseJSON.message);
            });
    });
});

function showModalSuccess(message){
    $.toast({
        heading:  'Successfully!',
        text: message,
        position: 'top-right',
        loaderBg: '#ff6849',
        icon: 'success',
        hideAfter: 5000,
        stack: 6
    });
}

function showModalError(message){
    $.toast({
        heading:  'Warning!',
        text: message,
        position: 'top-right',
        loaderBg: '#ff6849',
        icon: 'error',
        hideAfter: 5000,
        stack: 6
    });
}