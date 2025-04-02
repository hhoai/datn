CKEDITOR.dialog.add('pdfDialog', function (editor) {
    var beginUrl = "https://docs.google.com/gview?url=";
    var endUrl = "&embedded=true";
    return {
        title: 'Chèn pdf',
        minWidth: 400,
        minHeight: 125,

        contents: [
            {
                id: 'tab-main',
                label: '',
                elements: [
                    {
                        type: 'text',
                        id: 'url',
                        label: 'Đường dẫn',
                        validate: CKEDITOR.dialog.validate.notEmpty("Đường dẫn không được để trống."),

                        setup: function (element) {
                            //var src = element.getAttribute("src");                            
                            //var url = src.substring(beginUrl.lenght + 1, src.lenght - endUrl.lenght);
                            //this.setValue(url);
                        },

                        commit: function (element) {
                            if (this.getValue()) {
                                if (this.getValue().substring(1, 1) == '/')
                                    element.setAttribute("src", beginUrl + location.protocol + "//" + document.domain + ":" + location.port + this.getValue() + endUrl);
                                else element.setAttribute("src", this.getValue());
                            }
                        }
                    },
                    { type: "button", id: "browse", filebrowser: { action: "Browse", target: "tab-main:url" }, style: "float:right", hidden: !0, label: editor.lang.common.browseServer },
                    //{ type: "file", id: "upload", label: editor.lang.image.btnUpload, style: "height:40px", size: 38 },
                    //{
                    //    type: "fileButton", id: "uploadButton", filebrowser: "tab-main:url", label: editor.lang.image.btnUpload,
                    //    "for": ["tab-main", "upload"]
                    //},                   
                ]
            }
        ],

        onShow: function () {
            var selection = editor.getSelection();
            var element = selection.getStartElement();

            if (element)
                element = element.getAscendant('iframe', true);

            if (!element || element.getName() != 'iframe') {
                element = editor.document.createElement('iframe');
                element.setAttribute('frameborder', 0);
                element.setAttribute("style", "width:100%; height:1000px;");
                this.insertMode = true;
            }
            else
                this.insertMode = false;

            this.element = element;
            if (!this.insertMode)
                this.setupContent(this.element);
        },

        onOk: function () {
            var dialog = this;
            var pdf = this.element;
            this.commitContent(pdf);

            if (this.insertMode)
                editor.insertElement(pdf);
        }
    };
});