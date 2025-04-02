
function getCT() {
    var URL = `${ospath}PMIS_TBI/getCT`;
    $("#__ct").select2({
        multiple: true,
        allowClear: true,
        placeholder: "Chọn",
        ajax: {
            url: URL,
            dataType: 'json',
            type: "GET",
            quietMillis: 50,
            data: function (term) {
                return {
                    cda: $('#__cda').val(),
                    dvi: $('#__dvi').val(),
                    kw: term
                };
            },
            results: function (data) {
                return {
                    results: $.map(data, function (item) {
                        return {
                            text: item.Title,
                            slug: item.Index,
                            id: item.Index
                        }
                    })
                };
            }
        }
    });
}
function getDz() {
    var URL = `${ospath}PMIS_TBI/getDz`;
    $("#__dz").select2({
        multiple: false,
        allowClear: true,
        placeholder: "Chọn",
        ajax: {
            url: URL,
            dataType: 'json',
            type: "GET",
            quietMillis: 50,
            data: function (term) {
                return {
                    cda: $('#__cda').val(),
                    dvi: $('#__dvi').val(),
                    kw: term
                };
            },
            results: function (data) {
                return {
                    results: $.map(data, function (item) {
                        return {
                            text: item.Title,
                            slug: item.Index,
                            id: item.Index
                        }
                    })
                };
            }
        }
    });
}
function getTBA() {
    var URL = `${ospath}PMIS_TBI/getTBA`;
    $("#__tba").select2({
        multiple: false,
        allowClear: true,
        placeholder: "Chọn",
        ajax: {
            url: URL,
            dataType: 'json',
            type: "GET",
            quietMillis: 50,
            data: function (term) {
                return {
                    cda: $('#__cda').val(),
                    dvi: $('#__dvi').val(),
                    kw: term
                };
            },
            results: function (data) {
                return {
                    results: $.map(data, function (item) {
                        return {
                            text: item.Title,
                            slug: item.Index,
                            id: item.Index
                        }
                    })
                };
            }
        }
    });
}
function getNLO() {
    var URL = `${ospath}PMIS_TBI/getNLO`;
    $("#__nlo").select2({
        multiple: false,
        allowClear: true,
        placeholder: "Chọn",
        ajax: {
            url: URL,
            dataType: 'json',
            type: "GET",
            quietMillis: 50,
            data: function (term) {
                return {
                    tba: $('#__tba').val(),
                    cda: $('#__cda').val(),
                    dvi: $('#__dvi').val(),
                    kw: term
                };
            },
            results: function (data) {
                return {
                    results: $.map(data, function (item) {
                        return {
                            text: item.Title,
                            slug: item.Index,
                            id: item.Index
                        }
                    })
                };
            }
        }
    });
}
function getVTRI() {
    var URL = `${ospath}PMIS_TBI/getVTRI`;
    $("#__vtri").select2({
        multiple: true,
        allowClear: true,
        placeholder: "Chọn",
        ajax: {
            url: URL,
            dataType: 'json',
            type: "GET",
            quietMillis: 50,
            data: function (term) {
                if ($('#__dvi').val() === "" && $('#__ct').val() === "") {
                    showErrorMessage("Vui lòng chọn Đơn vị, Công trình")
                    return false;
                }
                return {
                    dvi: $('#__dvi').val(),
                    ct: $('#__ct').val(),
                    kw: term
                };
            },
            cache: true,
            results: function (data) {
                return {
                    results: $.map(data, function (item) {
                        return {
                            text: item.Title,
                            slug: item.Index,
                            id: item.Index
                        }
                    })
                };
            }
        }
    });
}