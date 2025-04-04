$(document).ready(function () {
    var assignmentID = $('input[name="assignmentId"]').val()

    const apiUrl = '/api/v1/file-assignments/' + assignmentID;

    // data table
    var table = $('#fileTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: (json) => {
                json.data.forEach((element, index) => {
                    element.RowNumber = `<p>${index + 1}</p>`;
                    element.Checkbox = `<input type="checkbox" id="checkbox_course" class="item-account filled-in" value="${element.FileAssignmnetID}" name="file_id[]">
                                <label for="checkbox_course"></label>`;
                });
                return json.data;
            },
        },
        columns: [
            {
                data: 'RowNumber',
                orderable: false,
            },
            {data: 'file_name'},
            {data: 'Checkbox'},
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },
        bPaginate: false, // Tắt phân trang
        bFilter: false,   // Tắt ô tìm kiếm
        bInfo: false,     // Tắt thông tin tổng số bản ghi
        rowCallback: function (row, data) {
            $(row).on("click", function (e) {
                //toogle checkbox
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
})