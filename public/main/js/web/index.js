$(document).ready(function () {
    const isActive = $('input[name="isActive"]').val()

    if (isActive === "false") {
        $('.active-account-modal').modal('show');
    }

    $('#confirmActive').on("click", function () {
        const button = $(this)
        button.prop("disabled", true).text("Sending...")

        $.ajax({
            url: '/api/v1/resend-activation',
            method: 'POST',
            success: function (response) {
                setTimeout(() => {
                    $('.active-account-modal').modal('hide');
                }, 1000)
                showModalSuccess("Resend Activation Successful. Check Your Mail!");
                button.hide();
            },
            error: function (xhr) {
                showModalError("Failed to send activation email.");
                button.prop('disabled', false).text('Resend Activation Email');
            },
            complete: function () {

                setTimeout(() => {
                    button.prop('disabled', false).text('Resend Activation Email').show();
                }, 10000);
            }
        });
    })
})
