$(document).ready(function () {
    function groupProcessLearningByDate(data) {
        return data.reduce((acc, record) => {
            const date = record.timeline_date;

            // Kiểm tra nếu ngày đã tồn tại trong nhóm, nếu chưa thì tạo mới
            if (!acc[date]) {
                acc[date] = [];
            }

            // Thêm bản ghi vào nhóm theo timeline_date
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
        // Convert the timestamp to a Date object
        const date = new Date(timestamp);

        // Format the date to "Jan 2, 2006"
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

    $.get("/api/v1/learning-process", function (res){
        let logs = res.data
        // Dữ liệu mẫu
        const groupedLogs = groupProcessLearningByDate(logs);
        // Gọi hàm để render dữ liệu
        renderTimeline(groupedLogs);
    })

    // Gọi hàm
    function renderTimeline(groupedLogs) {
        const timelineElement = document.querySelector(".timeline");

        // Lặp qua các khóa (timeline_date)
        Object.keys(groupedLogs).forEach(date => {
            // Tạo nội dung HTML cho từng timeline_date
            const timelineDateHtml = `
                <span class="timeline-label">
                    <span class="badge badge-pill badge-primary badge-lg">${date}</span>
                </span>
            `;

            // Chèn ngày vào timeline
            timelineElement.innerHTML += timelineDateHtml;

            // Duyệt qua từng log và thêm các timeline-item
            groupedLogs[date].forEach(log => {
                // check hoan thanh su kien nao (assignment / lesson / course)
                if (log.assignment_id !== 0) {
                    const timelineItemHtml = `
                        <div class="timeline-item">
                            <div class="timeline-point timeline-point-success">
                                <i class="fa fa-money"></i>
                            </div>
                            <div class="timeline-event">
                                <div class="timeline-heading">
                                      <a href="/student-courses/${log.course_id}">  
                                          <h4 class="timeline-title">
                                              <span>Course</span>: ${log.course_name}
                                          </h4>
                                      </a>
                                </div>
                                <div class="timeline-body">
                                    <span>Assignment Submitted:</span>
                                    <a href="/student-courses/assignments/${log.assignment_id}/details">"${log.assignment_name}"</a>- 
                                    <span>Lesson</span> 
                                    <a href="/student-courses/lessons/${log.lesson_id}/assignments">"${log.lesson_name}"</a>  
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
                else if (log.lesson_id !== 0) {
                    const timelineItemHtml = `
                        <div class="timeline-item">
                            <div class="timeline-point timeline-point-success">
                                <i class="fa fa-money"></i>
                            </div>
                            <div class="timeline-event">
                                <div class="timeline-heading">
                                  <a href="/student-courses/${log.course_id}">  
                                      <h4 class="timeline-title">
                                          <span>Course</span>: ${log.course_name}
                                      </h4>
                                  </a>
                                </div>
                                <div class="timeline-body">
                                    <span>Lesson completed</span>:  
                                    <a href="/student-courses/lessons/${log.lesson_id}/assignments">"${log.lesson_name}"</a>  
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
                                  <a href="/student-courses/${log.course_id}/certification">  
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

                    timelineElement.innerHTML += timelineItemHtml;
                }
            });
        });
    }


})