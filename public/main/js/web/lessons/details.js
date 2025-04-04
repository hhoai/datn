$(document).ready(function () {
    var lessonID = $('input[name="lesson_id"]').val()
    const apiUrl = '/api/v1/assignments/' + lessonID;

    var table = $('#assignmentTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: (json) => {
                if (json.data && json.data.assignments) {
                    return json.data.assignments.map((element, index) => {
                        element.RowNumber = `<p>${index + 1}</p>`;
                        element.Detail = `<a href="/assignments/${element.AssignmentID}/courses/${json.data.lessons.course_id}">${element.title}</a>`;
                        element.Action = `
                       <button class="btn btn-success btn-sm scoring-btn" data-id="${element.AssignmentID}">
                       <i class="fa fa-magic" aria-hidden="true"></i> </button>
                      <button class="btn btn-primary btn-sm edit-btn" data-id="${element.AssignmentID}">
                      <i class="fa fa-pencil-square-o" aria-hidden="true"></i> </button>
                      <button class="btn btn-danger btn-sm delete-btn" data-id="${element.AssignmentID}">
                       <i class="fa fa-trash" aria-hidden="true"></i></button>
                     `;
                        element.Checkbox = `
                      <input type="checkbox" id="checkbox-assignment-${element.AssignmentID}" class="item-account filled-in" value="${element.AssignmentID}" name="assignment_id[]">
                      <label for="checkbox-assignment-${element.AssignmentID}"></label>
                       `;
                        return element;
                    });
                }
                return [];
            },

        },
        columns: [
            { data: 'RowNumber'},
            { data: 'Detail'},
            { data: 'assignment_body'},
            { data: 'TypeAssignment.name'},
            { data: 'Action'},
            // { data: 'Checkbox' },
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },
        rowCallback: function (row, data) {
            $(row).on("click", function (e) {
                var checkbox = $(this).find(".item-account");
                checkbox.prop("checked", !checkbox.prop("checked"));
                //add class for selected row

                if (!row.classList.contains("selectedRow")) {
                    row.classList.add("selectedRow");
                } else {
                    row.classList.remove("selectedRow");
                }
            });
        },
    });

    table.on('click', '.scoring-btn', function () {
        var assignmentID = $(this).data('id');
        window.location.href = "/lessons/assignment/" + assignmentID + "/scoring"
    })
});
