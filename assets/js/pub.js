const dirOpenKey = "openDir-"
const viewDoc = "viewDoc-"
const nowViewDoc = "nowViewDoc-"

function viewDocKey(pid) {
    return viewDoc+pid
}

function viewDocAdd(pid, docId) {
    var listStr = localStorage.getItem(viewDocKey(pid))
    var list = [];
    if (listStr !== null) {
        list = listStr.split(",")
    } else {
        list.push(docId)
        localStorage.setItem(viewDocKey(pid), list.join(","))
        return
    }
    var has = false
    for(var i=0; i<list.length; i++) {
        if (list[i] === docId) {
            has = true
        }
    }
    if (!has) {
        list.push(docId)
        localStorage.setItem(viewDocKey(pid), list.join(","))
    }
}

function viewDocRefresh(pid, list) {
    localStorage.setItem(viewDocKey(pid), list.join(","))
}

function viewDocGet(pid)  {
    var listStr = localStorage.getItem(viewDocKey(pid))
    var list = [];
    if (listStr !== null) {
        list = listStr.split(",")
    }
    return list
}

function viewDocDel(pid, docId) {
    var listStr = localStorage.getItem(viewDocKey(pid))
    var list = [];
    if (listStr !== null) {
        list = listStr.split(",")
    }
    var newList = list.filter(function (value) {
        return value !== docId;
    });
    localStorage.setItem(viewDocKey(pid), newList.join(","))
}

async function copyContent (content) {
    let copyResult = true
    const text = content || '复制内容为空哦~';
    if (!!window.navigator.clipboard) {
        await window.navigator.clipboard.writeText(text).then((res) => {
            console.log('复制成功');
        }).catch((err) => {
            copyResult =  copyContent2(text)
        })
    } else {
        copyResult =  copyContent2(text)
    }
    return copyResult;
}

// 拖拽
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
                    //'box-shadow': '10px 10px 20px 0 #eee',
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


document.addEventListener('keydown', function(event) {
    if (event.ctrlKey && event.key === 's') {
        event.preventDefault();
    }
});

function isObject(obj) {
    return obj === Object(obj);
}

function isArrayOrObject(data) {
    if (Array.isArray(data)) {
        return 'array';
    } else if (isObject(data)) {
        return 'object';
    }
    return 'other';
}

function removeHtmlTags(htmlString) {
    return htmlString.replace(/<[^>]*>?/gm, '');
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

function AjaxPostNotAsync(url, param, func) {
    $.ajax({
        type: "post",
        url: url,
        data: JSON.stringify(param),
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

function ToastShow(msg) {
    const toastLiveExample = $('#liveToast')
    const toast = new bootstrap.Toast(toastLiveExample)
    var liveToastMsg = $("#liveToastMsg");
    liveToastMsg.empty(msg);
    liveToastMsg.append(msg);
    toast.show()
}

function newAPIDoc(id) {
    clearDocEditor(0)
    if (id !== "") {
        dirId = id
        $('#addDocDir').val(dirId);
    }
    $('#apiDocAddModal').modal('show');
}

function projectValueClear() {
    $("#projectName").val("");
    $("#projectPrivate").val(1);
    $("#projectDescription").val("");
}

function addProject() {
    $('#addProjectModalLabel').empty();
    $('#addProjectModalLabel').text("创建项目")
    $('#projectCreateSubmit').show();
    $('#projectModifySubmit').hide();
    projectValueClear()
    $('#addProjectModal').modal('show');
}

function ProjectCreate() {
    const url = "/project/create"
    var param = {
        "name": $("#projectName").val(),
        "description": $("#projectDescription").val(),
        "private": Number($("#projectPrivate").val()),
    }
    AjaxPost(url, param, function (data){
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                location.reload();
            }, 1000);
        }
    })
}

function modifyProject(pid, name, private, description) {
    $('#addProjectModalLabel').empty();
    $('#addProjectModalLabel').text("编辑项目");
    $('#projectCreateSubmit').hide();
    $('#projectModifySubmit').show();
    projectValueClear()
    $('#modifyProjectId').val(pid);
    $("#projectName").val(name);
    $("#projectPrivate").val(private);
    $("#projectDescription").val(description);
    $('#addProjectModal').modal('show');
}

function ProjectModify() {
    var param = {
        "projectId": $('#modifyProjectId').val(),
        "name": $("#projectName").val(),
        "description": $("#projectDescription").val(),
        "private": Number($("#projectPrivate").val()),
    }
    AjaxPost("/project/modify", param, function (data){
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
    AjaxPost("/mange/create/user", param, function (data){
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
    var param = {
        "name": $("#userModifyName").val(),
        "account": $('#userModifyAcc').val(),
    }

    AjaxPost("/user/modify", param, function (data){
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
    var param = {
        "password": $("#resetPassword").val(),
        "password2": $("#resetPassword2").val(),
        "account": $('#userResetPasswordAcc').val(),
    }
    AjaxPost("/user/reset/password", param, function (data){
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
    if (select.val() === null) {
        return
    }
    var param = {
        "pid": pid,
        "accounts": select.val().join(","),
    }
    AjaxPost("/project/adduser", param, function (data){
        if (data.code === 1) {
            ToastShow(data.msg);
        } else {
            location.reload();
        }

    })
}

function teamWorkerProjectDel(account, pid) {
    var param = {
        "pid": pid,
        "account": account,
    }
    AjaxPost("/project/deluser", param, function (data){
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
    var param = {
        "projectId": id,
        "name": name,
    }
    AjaxPost("/project/delete", param, function (data){
        console.log(data)
        ToastShow(data.msg);
        if (data.code === 0) {
            setTimeout(function() {
                window.location.href = "/home";
            }, 1000);
        }
    })
}

function openPConf(projectId) {
    window.location.href='/project/index/' + projectId + '?doc=' + projectId
}