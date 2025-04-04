$(document).ready(function () {
    const apiUrl = '/api/v1/student_programs';

    $.get(apiUrl, function (response) {
        var data = response.data
        $(".username").append(`${data.Username}!`)
        $("#list-programs").empty()
        data.UserProgram.forEach(up => {
            $("#list-programs").append(`
            <tr>
                <td>
                    <div class="bg-primary h-50 w-50 l-h-50 rounded text-center">
                        <p class="mb-0 font-size-20 font-weight-600">T1</p>
                    </div>
                </td>
                <td class="font-weight-600"><a href="/student-programs/${up.ProgramID}/details">${up.ProgramName}</a></td>
                <td class="text-fade">${up.ProgramCode}</td>
                <td class="font-weight-500"><span class="badge badge-sm badge-dot badge-primary mr-10"></span>${up.Status}</td>
                <td class="text-fade">${up.Course} <span data-i18n="courses"></span></td>
            </tr>
        `)
        })
    })
})
