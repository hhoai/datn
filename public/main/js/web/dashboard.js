$(document).ready(function (){
    const apiGetPost = "/api/v1/dashboard"
    const apiGetNews = "/api/v1/news"

    function formatTime (timestamp) {
        // Convert the timestamp to a Date object
        const date = new Date(timestamp);

        // Format the date to "Jan 2, 2006"
        const formatter = new Intl.DateTimeFormat('en-GB', {
            day: 'numeric',
            month: 'numeric',
            year: 'numeric',
        });
        return formatter.format(date);
    }

    $.get(apiGetPost, function (response) {
        var data = response.data
        $(".count-courses").append(`${data.CountCourse}`)
        $(".count-users").append(`${data.CountUser}`)
        $(".count-topics").append(`${data.CountTopic}`)
        $(".count-questions").append(`${data.CountQuestion}`)

        const courses = $("#courses-admin")
        courses.empty()

        data.Courses.forEach(course => {
            courses.append(`
            <div class="col-xl-3 col-md-6 col-12">
                <a href="/courses/${course.course_id}" class="box bg-secondary-light pull-up" style="background-image: url(/images/svg-icon/color-svg/st-1.svg); background-position: right bottom; background-repeat: no-repeat;">
                    <div class="box-body">
                        <div class="flex-grow-1">
                            <div class="d-flex align-items-center pr-2 justify-content-between">
                                <div class="d-flex">
                                    <span class="badge badge-primary mr-5"><i class="fa fa-lock"></i></span>
                                    <span class="badge badge-primary"><i class="fa fa-clock-o"></i></span>
                                </div>
                            </div>
                            <h4 class="mt-25 mb-5">${course.title}</h4>
                            <p class="text-fade mb-0 font-size-12">45 Days Left</p>
                        </div>
                    </div>
                </a>
            </div>
            `)
        })

        const dates = data.CountNewUser.map(item => item.date);
        const counts = data.CountNewUser.map(item => item.count);

        const chart = Highcharts.chart('user-admin', {
            chart: {
                type: 'column'
            },
            title: {
                text: ''
            },
            xAxis: {
                categories: dates // ['Apples', 'Bananas', 'Oranges'] // day
            },
            series: [{
                name: 'New User',
                data: counts // [1, 0, 4]  // count new user
            }]
        });

    })

    $(".search-programs").on("click", function (){
        var search = $("#search-content").val()
        if (search) {
            $.get('/api/v1/search/' + search, function (response) {
                var data = response.data

                const searchContainer = $(".search-container")
                searchContainer.empty()

                if(data.length > 0) {
                    data.forEach(p => {
                        searchContainer.append(`
                        <div class="box mb-15">
                          <div class="box-body">
                            <div class="d-flex align-items-center justify-content-between">
                              <div class="d-flex align-items-center">
                                <div class="mr-15 bg-warning rounded text-center" style="width: 60px;height: 60px; line-height: 70px;">
                                  <span class="icon-Book-open font-size-24"><span class="path1"></span><span class="path2"></span></span>
                                </div>
                                <div class="d-flex flex-column font-weight-500">
                                  <a href="/student-programs/${p.program_id}" class="text-dark hover-primary mb-1 font-size-16">${p.name}</a>
                                  <span class="text-fade">${p.program_code}</span>
                                </div>
                              </div>
                              <a href="/student-programs/${p.program_id}">
                                <span class="icon-Arrow-right font-size-24"><span class="path1"></span><span class="path2"></span></span>
                              </a>
                            </div>
                          </div>
                        </div>
                    `)
                    })
                }
                else {
                    searchContainer.append(`<p class="text-danger" style="margin: 10px 0">* No learning programs found</p>`)
                }
            })
        }
    })

})
