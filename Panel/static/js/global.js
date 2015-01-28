var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});


// ajax remote
$.fn.form.settings.rules["ajax"] = function(value,url) {
  $.ajax({
    async : false,
    url : url,
    type : "POST",
    dataType: "json",
    data : {
        value : value
    },
    success: function(data){
        console.log(data,url);
        return data["result"];
    }
  });
};


// ajax check invite code
$.fn.form.settings.rules["invite"] = function(value) {
    var result;
    $.ajax({
        async : false,
        url : "/register/checkinvite/"+ $("#invite").val(),
        type : "GET",
        dataType: "json",
        success: function(data){
            result = data.result;
        }
    });
    return result
};


// ajax check username
$.fn.form.settings.rules["username"] = function(value) {
    var result;
    $.ajax({
        async : false,
        url : "/register/checkusername/"+ $("#username").val(),
        type : "GET",
        dataType: "json",
        success: function(data){
            result = data.result;
        }
    });
    return result
};

// ajax check email
$.fn.form.settings.rules["mail"] = function(value) {
    var result;
    $.ajax({
        async : false,
        url : "/register/checkemail/"+ $("#email").val(),
        type : "GET",
        dataType: "json",
        success: function(data){
            result = data.result;
        }
    });
    return result
};

// not contains
$.fn.form.settings.rules["notcontains"] = function(value, field) {
    text = text.replace(/[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]/g, "\\$&");
    console.log("[debug]v: "+ value +", t:"+text);
    return !(value.search(text) !== -1);
};

