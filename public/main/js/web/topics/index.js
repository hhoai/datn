$(document).ready(function () {
    const apiUrl = '/api/v1/topics';
    $('#topicTable').DataTable({
        ajax: {
            url: apiUrl,
            method: 'GET',
            dataSrc: 'data'
        },
        columns: [
            {
                data: null,
                render: function (data, type, row, meta) {
                    return meta.row + 1;
                },
                orderable: false,
                searchable: false,
            },
            {data: 'name'},
            { data: 'description' },
            {
                render: function (_, __, row) {
                    return `
                    <a href="/topics/${row.topic_id}" class="btn btn-success btn-sm answer-link" title="Overview" data-id="${row.topic_id}"><i class="fa fa-eye" aria-hidden="true"></i></i>
                   </a>
                    <a href="/topics/update/${row.topic_id}" class="btn btn-primary btn-sm answer-link" title="Edit answer" data-id="${row.topic_id}"><i class="fa fa-pencil-square-o" aria-hidden="true"></i></i>
                   </a>
                   <button class="btn btn-danger btn-sm delete-btn" data-id="${row.topic_id}" title="Delete"><i class="fa fa-trash" aria-hidden="true"></i>
                   </button>
                   
                        
                    `;
                }
            }
        ],
        responsive: true,
        paging: true,
        searching: true,
        ordering: true,
        language: {
            emptyTable: "No data available"
        },
    });

    $('#create').click(function (e) {
        e.preventDefault();
        window.location.href = '/topics/create';
    });

    let topicIdToDelete = null;
    $('#topicTable').on('click', '.delete-btn', function (){
        topicIdToDelete = $(this).data('id');
        $('.delete-confirm-modal').modal('show');

    })

    $('#confirmDelete').on('click', function () {
        $.ajax({
            url: `${apiUrl}/${topicIdToDelete}`,
            method: 'DELETE',
            success: function (response) {
                if (response.success) {
                    showModalSuccess("Delete Successfully");
                    $('#topicTable').DataTable().ajax.reload();
                } else {
                    showModalError(response.message);
                }
                $('.delete-confirm-modal').modal('hide');
            },
            error: function (xhr) {
                showModalError(response.message);
                $('.delete-confirm-modal').modal('hide');
            }
        });
    });






});
