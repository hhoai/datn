$(document).ready(function () {
    var lessonID = $('input[name="lesson_id"]').val()
    const apiUrl = '/api/v1/posts/' + lessonID;

    var table = $('#postTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: (json) => {
                json.data.forEach((element, index) => {
                    element.RowNumber = `<p>${index + 1}</p>`
                    element.Detail = `<a href="/posts/${element.post_id}">${element.post_title}</a>`;
                    element.Action = `<button class="btn btn-primary btn-sm edit-btn" data-id="${element.post_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></i></button>
                        <button class="btn btn-danger btn-sm delete-btn" data-id="${element.post_id}"><i class="fa fa-trash" aria-hidden="true"></i></button>`;
                });
                return json.data;
            },
        },
        columns: [
            { data: 'RowNumber'},
            { data: 'Detail'},
            { data: 'post_body'},
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
            $(row).on("click", function () {
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
});
