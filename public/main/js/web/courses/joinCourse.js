$(document).ready(function () {
    var input = $('#course-code')
    input.on('input', function () {
        // Check if the input field has a value
        const hasValue = $(this).val().trim().length >= 5;

        // Enable or disable the button based on the input value
        $('#join-new-course-btn').prop('disabled', !hasValue);
    });

    $("#join-new-course-btn").on("click", function () {
        var courseCode = input.val();

        var formData = {
            "course_code" : courseCode
        }

        $.ajax({
            url: "/api/v1/request-course",
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                if (response.code === 200) {
                    $("#modal-join-course").modal('hide')
                    $.toast({
                        heading: 'Success',
                        text: 'Successfully sent the course join request',
                        position: 'top-right',
                        loaderBg: '#0fa535',
                        icon: 'success',
                        hideAfter: 1000
                    });
                    setTimeout(function() {
                        location.reload();  // Reload lại trang
                    }, 1000);
                }
                else {
                    $.toast({
                        heading: 'Warning',
                        text: response.message,
                        position: 'top-right',
                        loaderBg: '#ff6849',
                        icon: 'error',
                        hideAfter: 3500

                    });
                }
            },
            error: function (res) {
                $.toast({
                    heading: 'Warning',
                    text: res.responseJSON.message,
                    position: 'top-right',
                    loaderBg: '#ff6849',
                    icon: 'error',
                    hideAfter: 3500

                });
            }
        });

    })

})