CKEDITOR.plugins.add('pdf', {
    icons: 'pdf',
    init: function (editor) {

        editor.addCommand('pdf', new CKEDITOR.dialogCommand('pdfDialog'));

        editor.ui.addButton('Pdf', {
            label: 'Chèn pdf',
            command: 'pdf',
            toolbar: 'insert'
        });        

        CKEDITOR.dialog.add('pdfDialog', this.path + 'dialogs/pdf.js');
    }
});