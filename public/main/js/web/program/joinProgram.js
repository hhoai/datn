$(document).ready(function () {
    var input = $('#course-code')
    input.on('input', function () {
        // Check if the input field has a value
        const hasValue = $(this).val().trim().length >= 5;

        // Enable or disable the button based on the input value
        $('#join-new-program-btn').prop('disabled', !hasValue);
    });

    $("#join-new-program-btn").on("click", function () {
        var courseCode = input.val();

        var formData = {
            "program_code" : courseCode
        }

        $.ajax({
            url: "/api/v1/request-program",
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
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
            },
            error: function (xhr) {
                $.toast({
                    heading: 'Warning',
                    text: 'Failed to Failed to join the program request.',
                    position: 'top-right',
                    loaderBg: '#ff6849',
                    icon: 'error',
                    hideAfter: 3500

                });
            }
        });

    })
    $(".join-program").on("click", function (){
        var url = window.location.pathname;
        var programID = url.split('/')[2];

        const formData = {
            "program_id": programID
        }

        $.ajax({
            url: "/api/v1/request-program",
            method: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
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
            },
            error: function (xhr) {
                $.toast({
                    heading: 'Warning',
                    text: 'Failed to Failed to join the program request.',
                    position: 'top-right',
                    loaderBg: '#ff6849',
                    icon: 'error',
                    hideAfter: 3500

                });
            }
        });

    })
})