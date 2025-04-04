$(document).ready(function () {
    $("#exportExcelBtn").click(function () {
        let url = "/api/v1/question_bank/export";

        fetch(url)
            .then(response => response.blob())
            .then(blob => {
                let link = document.createElement('a');
                link.href = URL.createObjectURL(blob);
                link.download = 'sample_file.xlsx';
                link.click();
            });
    });
});
