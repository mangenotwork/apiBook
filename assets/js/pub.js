;(function( $ ){
    /**
     * Author: https://github.com/Barrior
     *
     * DDSort: drag and drop sorting.
     * @param {Object} options
     *        target[string]: 		可选，jQuery事件委托选择器字符串，默认'li'
     *        cloneStyle[object]: 	可选，设置占位符元素的样式
     *        floatStyle[object]: 	可选，设置拖动元素的样式
     *        down[function]: 		可选，鼠标按下时执行的函数
     *        move[function]: 		可选，鼠标移动时执行的函数
     *        up[function]: 		可选，鼠标抬起时执行的函数
     */
    $.fn.DDSort = function( options ){
        var $doc = $( document ),
            fnEmpty = function(){},

            settings = $.extend( true, {
                down: fnEmpty,
                move: fnEmpty,
                up: fnEmpty,
                target: 'li',
                cloneStyle: {
                    'background-color': '#eee'
                },
                floatStyle: {
                    //用固定定位可以防止定位父级不是Body的情况的兼容处理，表示不兼容IE6，无妨
                    'position': 'fixed',
                    'box-shadow': '10px 10px 20px 0 #eee',
                    // 'webkitTransform': 'rotate(4deg)',
                    // 'mozTransform': 'rotate(4deg)',
                    // 'msTransform': 'rotate(4deg)',
                    // 'transform': 'rotate(4deg)'
                }

            }, options );

        return this.each(function(){

            var that = $( this ),
                height = 'height',
                width = 'width';

            if( that.css( 'box-sizing' ) == 'border-box' ){
                height = 'outerHeight';
                width = 'outerWidth';
            }

            that.on( 'mousedown.DDSort', settings.target, function( e ){
                //只允许鼠标左键拖动
                if( e.which != 1 ){
                    return;
                }

                //防止表单元素失效
                var tagName = e.target.tagName.toLowerCase();
                if( tagName == 'input' || tagName == 'textarea' || tagName == 'select' ){
                    return;
                }

                var THIS = this,
                    $this = $( THIS ),
                    offset = $this.offset(),
                    disX = e.pageX - offset.left,
                    disY = e.pageY - offset.top,

                    clone = $this.clone()
                        .css( settings.cloneStyle )
                        .css( 'height', $this[ height ]() )
                        .empty(),

                    hasClone = 1,

                    //缓存计算
                    thisOuterHeight = $this.outerHeight(),
                    thatOuterHeight = that.outerHeight(),

                    //滚动速度
                    upSpeed = thisOuterHeight,
                    downSpeed = thisOuterHeight,
                    maxSpeed = thisOuterHeight * 3;

                settings.down.call( THIS );

                $doc.on( 'mousemove.DDSort', function( e ){
                    if( hasClone ){
                        $this.before( clone )
                            .css( 'width', $this[ width ]() )
                            .css( settings.floatStyle )
                            .appendTo( $this.parent() );

                        hasClone = 0;
                    }

                    var left = e.pageX - disX,
                        top = e.pageY - disY,

                        prev = clone.prev(),
                        next = clone.next().not( $this );

                    $this.css({
                        left: left,
                        top: top
                    });

                    //向上排序
                    if( prev.length && top < prev.offset().top + prev.outerHeight()/2 ){

                        clone.after( prev );

                        //向下排序
                    }else if( next.length && top + thisOuterHeight > next.offset().top + next.outerHeight()/2 ){

                        clone.before( next );

                    }

                    /**
                     * 处理滚动条
                     * that是带着滚动条的元素，这里默认以为that元素是这样的元素（正常情况就是这样），如果使用者事件委托的元素不是这样的元素，那么需要提供接口出来
                     */
                    var thatScrollTop = that.scrollTop(),
                        thatOffsetTop = that.offset().top,
                        scrollVal;

                    //向上滚动
                    if( top < thatOffsetTop ){

                        downSpeed = thisOuterHeight;
                        upSpeed = ++upSpeed > maxSpeed ? maxSpeed : upSpeed;
                        scrollVal = thatScrollTop - upSpeed;

                        //向下滚动
                    }else if( top + thisOuterHeight - thatOffsetTop > thatOuterHeight ){

                        upSpeed = thisOuterHeight;
                        downSpeed = ++downSpeed > maxSpeed ? maxSpeed : downSpeed;
                        scrollVal = thatScrollTop + downSpeed;
                    }

                    that.scrollTop( scrollVal );

                    settings.move.call( THIS );

                })
                    .on( 'mouseup.DDSort', function(){

                        $doc.off( 'mousemove.DDSort mouseup.DDSort' );

                        //click的时候也会触发mouseup事件，加上判断阻止这种情况
                        if( !hasClone ){
                            clone.before( $this.removeAttr( 'style' ) ).remove();
                            settings.up.call( THIS );
                        }
                    });

                return false;
            });
        });
    };

})( jQuery );


function tooltipInit() {
    var tooltipTriggerList = Array.prototype.slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl)
    })
}

function getJsonPaths(json, currentPath = '') {
    let paths = [];
    function traverse(obj, path) {
        Object.keys(obj).forEach(key => {
            const newPath = path + (path === '' ? '' : '.') + key;
            paths.push(newPath);
            if (typeof obj[key] === 'object' && obj[key] !== null && !Array.isArray(obj[key])) {
                traverse(obj[key], newPath);
            }
      });
    }
    traverse(json, currentPath);
    return paths;
}

function AjaxGet(url, func) {
    $.ajax({
        type: "get",
        url: url,
        data: "",
        dataType: 'json',
        async: true,
        success: function(data){
            func(data);
        },
        error: function(xhr,textStatus) {
            console.log(xhr, textStatus);
        }
    });
}

function AjaxGetNotAsync(url, func) {
    $.ajax({
        type: "get",
        url: url,
        data: "",
        dataType: 'json',
        async: false,
        success: function(data){
            func(data);
        },
        error: function(xhr,textStatus) {
            console.log(xhr, textStatus);
        }
    });
}

function AjaxPost(url, param, func) {
    $.ajax({
        type: "post",
        url: url,
        data: JSON.stringify(param),
        dataType: 'json',
        async: true,
        success: function(data){
            func(data);
        },
        error: function(xhr,textStatus) {
            console.log(xhr, textStatus);
        }
    });
}

function ToastShow(msg) {
    const toastLiveExample = $('#liveToast')
    const toast = new bootstrap.Toast(toastLiveExample)
    $("#liveToastMsg").empty(msg);
    $("#liveToastMsg").append(msg);
    toast.show()
}

function newAPIDoc() {
    $('#apiDocAddModal').modal('show');
}

function addProject() {
    $('#addProjectModal').modal('show');
}

// 创建项目
function ProjectCreate() {
    const url = "/project/create"
    var param = {
        "name": $("#projectName").val(),
        "description": $("#projectDescription").val(),
        "private": Number($("#projectPrivate").val()),
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function addUser() {
    $('#addUserModal').modal('show');
}

function UserCreate() {
    const url = "/mange/create/user"
    var isAdmin = 0
    if ($("#isAdmin").is(":checked")) {
        isAdmin = 1
    }
    var param = {
        "name": $("#name").val(),
        "account": $("#account").val(),
        "password": $("#password").val(),
        "password2": $("#password2").val(),
        "isAdmin": isAdmin,
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function modifyUser(acc) {
    $('#userModifyAcc').val(acc);
    $('#modifyUserModal').modal('show');
}

function UserModify() {
    const url = "/user/modify"

    var param = {
        "name": $("#userModifyName").val(),
        "account": $('#userModifyAcc').val(),
    }

    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function resetPasswordUser(acc) {
    $('#userResetPasswordAcc').val(acc);
    $('#resetPasswordUserModal').modal('show');
}

function UserResetPassword() {
    const url = "/user/reset/password"
    var param = {
        "password": $("#resetPassword").val(),
        "password2": $("#resetPassword2").val(),
        "account": $('#userResetPasswordAcc').val(),
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function disableUser(acc, isDisable) {
    const url = "/mange/disable/user"
    var param = {
        "account": acc,
        "isDisable": isDisable,
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function deleteUser(acc) {
    const url = "/mange/delete/user"
    var param = {
        "account": acc,
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}


function teamWorkerProjectAdd(pid) {
    var select = $('#userList')
    console.log(select.val())
    if (select.val() === null) {
        return
    }
    const url = "/project/adduser"
    var param = {
        "pid": pid,
        "accounts": select.val().join(","),
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        if (data.code === 1) {
            ToastShow(data.msg);
        } else {
            location.reload();
        }

    })

}

function teamWorkerProjectDel(account, pid) {
    const url = "/project/deluser"
    var param = {
        "pid": pid,
        "account": account,
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        if (data.code === 1) {
            ToastShow(data.msg);
        } else {
            location.reload();
        }
    })

}

function delProjectOpen() {
    $('#delProjectModal').modal('show');
}

function delProject(name, id) {
    if ($('#delProjectName').val() != name) {
        ToastShow("项目名字错误");
        return
    }
    const url = "/project/delete"
    var param = {
        "projectId": id,
        "name": name,
    }
    AjaxPost(url, param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                window.location.href = "/home";
            }, 1000);
        }
    })
}

function openBodyMain(type) {
    var bodyJson = $("#bodyJson");
    var bodyFromData = $("#bodyFromData");
    var bodyXwwwFrom = $("#bodyXwwwFrom");
    var bodyXml = $("#bodyXml");
    var bodyText = $("#bodyText");

    var openBodyBtnJson = $("#openBodyBtn-json");
    var openBodyBtnFromData = $("#openBodyBtn-from-data");
    var openBodyBtnXWwwFormUrlencoded = $("#openBodyBtn-x-www-form-urlencoded");
    var openBodyBtnXml = $("#openBodyBtn-xml")
    var openBodyBtnPlain = $("#openBodyBtn-plain")

    bodyJson.hide();
    bodyFromData.hide();
    bodyXwwwFrom.hide();
    bodyXml.hide();
    bodyText.hide();

    openBodyBtnJson.removeClass('btn-light');
    if (!openBodyBtnJson.hasClass('btn-dark')) {
        openBodyBtnJson.addClass('btn-dark')
    }

    openBodyBtnFromData.removeClass('btn-light');
    if (!openBodyBtnFromData.hasClass('btn-dark')) {
        openBodyBtnFromData.addClass('btn-dark')
    }

    openBodyBtnXWwwFormUrlencoded.removeClass('btn-light');
    if (!openBodyBtnXWwwFormUrlencoded.hasClass('btn-dark')) {
        openBodyBtnXWwwFormUrlencoded.addClass('btn-dark')
    }

    openBodyBtnXml.removeClass('btn-light');
    if (!openBodyBtnXml.hasClass('btn-dark')) {
        openBodyBtnXml.addClass('btn-dark')
    }

    openBodyBtnPlain.removeClass('btn-light');
    if (!openBodyBtnPlain.hasClass('btn-dark')) {
        openBodyBtnPlain.addClass('btn-dark')
    }

    switch (type) {
        case 'json':
            bodyJson.show();
            openBodyBtnJson.removeClass('btn-dark');
            openBodyBtnJson.addClass('btn-light')
            break
        case 'from-data':
            bodyFromData.show();
            openBodyBtnFromData.removeClass('btn-dark');
            openBodyBtnFromData.addClass('btn-light')
            break
        case 'x-www-form-urlencoded':
            bodyXwwwFrom.show();
            openBodyBtnXWwwFormUrlencoded.removeClass('btn-dark');
            openBodyBtnXWwwFormUrlencoded.addClass('btn-light')
            break
        case 'xml':
            bodyXml.show();
            openBodyBtnXml.removeClass('btn-dark');
            openBodyBtnXml.addClass('btn-light')
            break
        case 'plain':
            bodyText.show();
            openBodyBtnPlain.removeClass('btn-dark');
            openBodyBtnPlain.addClass('btn-light')
            break
    }
}

function openRespMain(type) {
    var respJson = $("#respJson")
    var respXml = $("#respXml")
    var respText = $("#respText")

    var openRespBtnJson = $("#openRespBtn-json")
    var openRespBtnXml = $("#openRespBtn-xml")
    var openRespBtnPlain = $("#openRespBtn-plain")

    respJson.hide();
    respXml.hide();
    respText.hide();

    openRespBtnJson.removeClass('btn-light');
    if (!openRespBtnJson.hasClass('btn-dark')) {
        openRespBtnJson.addClass('btn-dark')
    }

    openRespBtnXml.removeClass('btn-light');
    if (!openRespBtnXml.hasClass('btn-dark')) {
        openRespBtnXml.addClass('btn-dark')
    }

    openRespBtnPlain.removeClass('btn-light');
    if (!openRespBtnPlain.hasClass('btn-dark')) {
        openRespBtnPlain.addClass('btn-dark')
    }

    switch (type) {
        case 'json':
            respJson.show();
            openRespBtnJson.removeClass('btn-dark');
            openRespBtnJson.addClass('btn-light')
            break
        case 'xml':
            respXml.show();
            openRespBtnXml.removeClass('btn-dark');
            openRespBtnXml.addClass('btn-light')
            break
        case 'plain':
            respText.show();
            openRespBtnPlain.removeClass('btn-dark');
            openRespBtnPlain.addClass('btn-light')
            break
    }
}