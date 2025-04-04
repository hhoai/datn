$(document).ready(function () {
    $("#infoForm").on("submit", function (event){
        event.preventDefault();

        var formData = {
            "name": $('#inputName').val(),
            "email": $('#inputEmail').val(),
            "odd_password": $('#inputPassword').val(),
            // "new_password": $('#inputNewPassword').val(),
            // "confirm_password": $('#inputConfirmPassword').val(),
        }

        $(".update-confirm-modal").modal("show")

        $('#confirmUpdate').on('click', function () {
            $.ajax({
                url: "/api/v1/information",
                method: "PUT",
                contentType: "application/json",
                data: JSON.stringify(formData),
            })
                .done(function (response) {
                    if (response.code === 200) {
                        showModalSuccess("Send email required change email success!")
                        $(".update-confirm-modal").modal("hide")
                    }
                })
                .fail(function (xhr) {
                    showModalError(xhr.responseJSON.message);
                });
        })

    })

    $("#updatePassword").on('click', function (event){
        event.preventDefault();

        var formData = {
            "odd_password": $('#u-password').val(),
            "new_password": $('#u-newPassword').val(),
            "confirm_password": $('#u-confirmPassword').val(),
        }

        $(".update-confirm-modal").modal("show")

        $('#confirmUpdate').on('click', function () {
            $.ajax({
                url: "/api/v1/information/change-password",
                method: "PUT",
                contentType: "application/json",
                data: JSON.stringify(formData),
            })
                .done(function (response) {
                    if (response.code === 200) {
                        $(".update-confirm-modal").modal("hide")
                        $("#myModal").modal("hide")
                        showModalSuccess("Change password successfully!")
                    }
                })
                .fail(function (xhr) {
                    showModalError(xhr.responseJSON.message);
                });
        })
    })

})