const url = window.location.href.split('/');
const userId = url[url.indexOf('users') + 1];
const lessonId = url[url.indexOf('lessons') + 1];
$(document).ready(function () {

    const apiUrl = `/api/v1/users/${userId}/lessons/${lessonId}/assignments`;

    $.ajax({
        url: apiUrl,
        method: 'GET',
        dataType: 'json',
        success: function (response) {
            if (response && Array.isArray(response.data)) {

                response.data.forEach(assignment => {
                    const assignmentHtml = `
                <div data-assignment-id="${assignment.Assignment.AssignmentID}">
                <div style="padding: 10px 25px; margin-bottom: 20px; cursor: pointer;" class="assignment_content" onclick="toggleDetails(this)">
                 <div class="align-items-center lesson-title">
                <span class="badge badge-dot badge-danger opacity-0"></span>
                <div class="flexbox flex-grow gap-items text-truncate">
                    <div class="mr-15 bg-primary-light l-h-60 rounded text-center" style="height: 50px; width: 50px;">
                        <span class="icon-Library font-size-24"><span class="path1"></span><span class="path2"></span></span>
                    </div>
                    <div class="media-body text-truncate d-flex" style="margin-bottom: 0px; justify-content: space-between; align-items: center;">
                        <div style="display: flex; justify-content: space-between; align-items: center; width: 100%;">
                            <b>${assignment.Assignment.title}</b>
                            <span class="status" style="margin-left: auto;"></span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="detail" style="display: none">
                <div style="display: flex; justify-content: space-between;padding: 15px 0; border-top: 1px solid #e1eaf9;">
                    <p>${assignment.Assignment.assignment_body}</p>
                    <div style="display: flex; flex-direction: column; border-left: 1px solid #e1eaf9; padding-left: 10px;">
                        <p>Score</p>
                        <h4 style="color: #c63838">${assignment.Assignment.score === 0 ? 'Not Graded Yet' : assignment.Assignment.score}</h4>
                    </div>
                </div>
                <div style="display: flex; justify-content: space-between; padding: 15px 0; border-top: 1px solid #e1eaf9;">
               <p style="color: ${assignment.status === true ? '#28a670' : '#d30909'}; margin-bottom: 0">
                  ${assignment.status === true ? 'Assignment Submitted' : 'Late Submission'}
                </p>
               </div>

            </div>
        </div>
    </div>
`;
       $("#assignments").append(assignmentHtml);
});
         setTimeout(function () {
                    processCompletedAssignments();
                }, 100);
            }
        },
        error: function (xhr, status, error) {
            console.error("Error fetching lessons:", error);
        }
    });

    function processCompletedAssignments() {
        $.ajax({
            url: `/api/v1/users/${userId}/assignments_process`,
            method: 'GET',
            dataType: 'json',
            success: function (processResponse) {
                if (processResponse.code === 200 && Array.isArray(processResponse.data)) {
                    const completedAssignments = processResponse.data;

                    $(".assignment_content").each(function () {
                        const assignmentID = $(this).parent().data("assignment-id");
                        const matchingAssignment = completedAssignments.find(assignment => {
                            return parseInt(assignment.assignment_id) === parseInt(assignmentID);
                        });

                        if (matchingAssignment) {
                            $(this).find(".status")
                                .text("Completed")
                                .css("color", "green");
                        } else {
                            $(this).find(".status")
                                .text("Incomplete")
                                .css("color", "red");
                        }
                    });

                } else {
                    console.error("Failed to fetch completed assignments data");
                }
            },
            error: function (xhr, status, error) {
                console.error("Error occurred while fetching process data:", error);
            }
        });
    }
});
function toggleDetails(element) {
    const details = element.querySelector('.detail');
    const assignmentId = element.parentElement.getAttribute('data-assignment-id');
    const fileListContainer = details.querySelector('.file-list');

    if (details.style.display === 'none') {
        details.style.display = 'block';

        if (!fileListContainer || fileListContainer.innerHTML.trim() === '') {
            $.ajax({
                url: `/api/v1/file-assignments/${assignmentId}`,
                method: 'GET',
                dataType: 'json',
                success: function (response) {
                    if (response.code === 200 && Array.isArray(response.data)) {
                        const fileListHtml = response.data.map(file => `
                            <li>
                                 <li>
                                <a href="/student/${userId}/${assignmentId}/${file.file_name}" target="_blank">${file.file_name}</a>
                            </li>
                            </li>
                        `).join('');

                        if (!fileListContainer) {
                            const fileListElement = document.createElement('ul');
                            fileListElement.classList.add('file-list');
                            fileListElement.innerHTML = fileListHtml;
                            details.appendChild(fileListElement);
                        } else {
                            fileListContainer.innerHTML = fileListHtml;
                        }
                    } else {
                        console.error('Failed to fetch file assignments:', response.message);
                    }
                },
                error: function (xhr, status, error) {
                    console.error('Error fetching file assignments:', error);
                }
            });
        }
    } else {
        details.style.display = 'none';
    }
}
