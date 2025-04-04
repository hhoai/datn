$(document).ready(function () {
    const clearErrorMessage = () => {
        $(".error-text").text("");
    };

    $(".form").on("submit", function (event) {
        event.preventDefault();


        var formData = $(this).serialize();
        $.ajax({
            url: "/api/v1/signup",
            method: "POST",
            data: formData,
        })
            .done(function (response) {
                if(response.code === 200){
                    window.location.href = "/";
                }

            })
            .fail(function (xhr) {
                // This function runs if the request fails
                showModalError(xhr.responseJSON.message);
            });
    });
});

function showModalError(message){
    $.toast({
        heading:  'Warning!',
        text: message,
        position: 'top-right',
        loaderBg: '#ff6849',
        icon: 'error',
        hideAfter: 3000,
        stack: 6
    });
}