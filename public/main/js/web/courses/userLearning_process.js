
$(document).ready(function () {
    var url = window.location.href;
    var parts = url.split('/');
    var user_id = parts[parts.indexOf("users") + 1];

    function groupProcessLearningByDate(data) {
        return data.reduce((acc, record) => {
            const date = record.timeline_date;

            if (!acc[date]) {
                acc[date] = [];
            }
            acc[date].push({
                course_id: record.course_id,
                course_name : record.course_name,
                lesson_id: record.lesson_id,
                lesson_name: record.lesson_name,
                assignment_id: record.assignment_id,
                assignment_name: record.assignment_name,
                submit_at: record.submit_at,
            });

            return acc;
        }, {});
    }
    function formatTime (timestamp) {
        const date = new Date(timestamp);
        const formatter = new Intl.DateTimeFormat('en-GB', {
            day: 'numeric',
            month: 'numeric',
            year: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            hour12: false
        });
        return formatter.format(date);
    }

    $.get(`/api/v1/users/${user_id}/learning_process`, function (res){
        let logs = res.data
        const groupedLogs = groupProcessLearningByDate(logs);
        renderTimeline(groupedLogs);
    })


    function renderTimeline(groupedLogs) {
        const timelineElement = document.querySelector(".timeline");
        Object.keys(groupedLogs).forEach(date => {
            const timelineDateHtml = `
                <span class="timeline-label">
                    <span class="badge badge-pill badge-primary badge-lg">${date}</span>
                </span>
            `;

            timelineElement.innerHTML += timelineDateHtml;
            groupedLogs[date].forEach(log => {
                if (log.assignment_id !== 0) {
                    const timelineItemHtml = `
                        <div class="timeline-item">
                            <div class="timeline-point timeline-point-success">
                                <i class="fa fa-money"></i>
                            </div>
                            <div class="timeline-event">
                                <div class="timeline-heading">
                                  <a href="/users/${user_id}/courses/${log.course_id}/lessons">  
                                      <h4 class="timeline-title">
                                          <span>Course</span>: ${log.course_name}
                                      </h4>
                                  </a>
                                </div>
                                <div class="timeline-body">
                                    <span>Assignment Submitted:</span>
                                    <a href="/student-courses/assignments/${log.assignment_id}/details">"${log.assignment_name}"</a>- 
                                    <span>Lesson</span>
                                    <a href="/users/${user_id}/lessons/${log.lesson_id}/assignments">"${log.lesson_name}"</a>  
                                </div>
                                <div class="timeline-footer">
                                    <p class="text-right">${formatTime(log.submit_at)}</p>
                                </div>
                            </div>
                        </div>
                        `;
                    timelineElement.innerHTML += timelineItemHtml;
                }
                else if (log.lesson_id !== 0) {
                    const timelineItemHtml = `
                        <div class="timeline-item">
                            <div class="timeline-point timeline-point-success">
                                <i class="fa fa-money"></i>
                            </div>
                            <div class="timeline-event">
                                <div class="timeline-heading">
                                  <a href="/users/${user_id}/courses/${log.course_id}/lessons">  
                                      <h4 class="timeline-title">
                                          <span>Course</span>: ${log.course_name}
                                      </h4>
                                  </a>
                                </div>
                                <div class="timeline-body">
                                    <span>Lesson completed</span>:
                                    <a href="/users/${user_id}/lessons/${log.lesson_id}/assignments">"${log.lesson_name}"</a>  
                                </div>
                                <div class="timeline-footer">
                                    <p class="text-right">${formatTime(log.submit_at)}</p>
                                </div>
                            </div>
                        </div>
                        `;
                    // Chèn nội dung từng bản ghi vào timeline
                    timelineElement.innerHTML += timelineItemHtml;
                }
                else {
                    const timelineItemHtml = `
                        <div class="timeline-item">
                            <div class="timeline-point timeline-point-success">
                                <i class="fa fa-money"></i>
                            </div>
                            <div class="timeline-event">
                                <div class="timeline-heading">
                                  <a href="/users/${user_id}/courses/${log.course_id}/lessons">  
                                      <h4 class="timeline-title">
                                          <span>Course</span>: ${log.course_name}
                                      </h4>
                                  </a>
                                </div>
                                <div class="timeline-body">
                                    <span>Course completed</span>
                                </div>
                                <div class="timeline-footer">
                                    <p class="text-right">${formatTime(log.submit_at)}</p>
                                </div>
                            </div>
                        </div>
                        `;
                    // Chèn nội dung từng bản ghi vào timeline
                    timelineElement.innerHTML += timelineItemHtml;
                }
            });
        });
    }


})