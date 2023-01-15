<title>计划任务</title>

<div class="layui-fluid">
    <div class="layui-card">
        <div class="layui-card-header">添加计划任务</div>
        <div class="layui-card-body">
            <form class="layui-form" action="" lay-filter="cron-add-form">
                <div class="layui-form-item">
                    <label class="layui-form-label">任务名</label>
                    <div class="layui-input-inline">
                        <input type="text" name="name" lay-verify="required" placeholder="请输入任务名称"
                               autocomplete="off" class="layui-input">
                    </div>
                    <div class="layui-form-mid layui-word-aux">请填写任务名称</div>
                </div>
                <div class="layui-form-item">
                    <label class="layui-form-label">执行周期</label>
                    <div class="layui-input-inline">
                        <input type="text" name="time" id="cron-add-time"
                               lay-verify="required" placeholder="请选择或输入cron表达式" class="layui-input">
                    </div>
                    <div class="layui-form-mid layui-word-aux">请务必正确填写执行周期</div>
                </div>
                <div class="layui-form-item layui-form-text">
                    <label class="layui-form-label">脚本内容</label>
                    <div class="layui-input-block">
                        <div id="cron-add-script-editor"
                             style="height: 250px;"># 在此输入你要执行的脚本内容</div>
                    </div>
                </div>
                <div class="layui-form-item">
                    <div class="layui-input-block">
                        <button class="layui-btn" lay-submit="" lay-filter="cron-add-submit">立即提交</button>
                        <button type="reset" class="layui-btn layui-btn-primary">重置</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <div class="layui-card">
        <div class="layui-card-header">计划任务列表</div>
        <div class="layui-card-body">
            <table id="panel-cron" lay-filter="panel-cron"></table>
            <!-- 操作按钮模板 -->
            <script type="text/html" id="cron-table-edit">
                <a class="layui-btn layui-btn-xs" lay-event="log">日志</a>
                <a class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>
                <a class="layui-btn layui-btn-warm layui-btn-xs" lay-event="del">删除</a>
            </script>
            <!-- 运行开关 -->
            <script type="text/html" id="cron-table-status">
                <input type="checkbox" name="cron-status" lay-skin="switch" lay-text="ON|OFF"
                       lay-filter="cron-status"
                       value="@{{ d.status }}" data-id="@{{ d.id }}"
                       @{{ d.status==
                       1 ? 'checked' : '' }}>
            </script>
        </div>
    </div>
</div>

<script>
    var cronAddScriptEditor = ace.edit("cron-add-script-editor", {
        mode: "ace/mode/sh",
        selectionStyle: "text"
    });
    var cronEditScriptEditor;
    layui.use(['admin', 'table', 'jquery', 'cron'], function () {
        var $ = layui.$
            , form = layui.form
            , table = layui.table
            , admin = layui.admin
            , cron = layui.cron;

        cron.render({
            elem: "#cron-add-time",
            btns: ['confirm'],
            show: false,
            done: function (value) {
                $('#cron-add-time').val(value);
            }
        });

        form.render();

        table.render({
            elem: '#panel-cron'
            , url: '/api/panel/cron/getList'
            , cols: [[
                {field: 'id', hide: true, title: 'ID'}
                , {field: 'name', width: 150, title: '任务名', sort: true}
                , {field: 'type', width: 150, title: '任务类型', sort: true}
                , {field: 'status', title: '启用', width: 100, templet: '#cron-table-status', unresize: true}
                , {field: 'time', width: 200, title: '任务周期（cron表达式）'}
                , {field: 'updated_at', title: '上次运行时间'}
                , {
                    field: 'edit',
                    width: 180,
                    title: '操作',
                    templet: '#cron-table-edit',
                    fixed: 'right',
                    align: 'left'
                }
            ]]
            , page: true
            , limit: 10
            , limits: [10, 100, 200, 500, 1000]
            , text: {
                none: '暂无数据'
            }
            , done: function () {
                //element.render('progress');
            }
        });

        // 工具条
        table.on('tool(panel-cron)', function (obj) {
            let data = obj.data;
            if (obj.event === 'log') {
                // 打开日志弹窗
                admin.popup({
                    title: '日志'
                    ,
                    area: ['80%', '80%']
                    ,
                    id: 'cron-log'
                    ,
                    content: '<pre id="cron-log-view" style="overflow: auto; height: 95%;border: 0 none; line-height:23px; padding: 15px; margin: 0; white-space: pre-wrap; background-color: rgb(51,51,51); color:#f1f1f1; border-radius:0;"></pre>'
                    ,
                    success: function (layero, index) {
                        admin.req({
                            url: '/api/panel/cron/getLog?id=' + data.id
                            , type: 'GET'
                            , success: function (res) {
                                if (res.code === 0) {
                                    $('#cron-log-view').html(res.data);
                                } else {
                                    layer.msg(res.msg, {icon: 2, time: 1000});
                                }
                            }
                            , error: function (xhr, status, error) {
                                console.log('耗子Linux面板：ajax请求出错，错误' + error);
                            }
                        });
                    }
                });

            } else if (obj.event === 'edit') {
                // 打开编辑弹窗
                admin.popup({
                    title: '编辑'
                    ,
                    area: ['80%', '80%']
                    ,
                    id: 'cron-log'
                    ,
                    content: '任务名&nbsp;&nbsp;&nbsp;&nbsp;<div class="layui-input-inline" style="width: 190px;"><input type="text" name="cron-edit-name" placeholder="请输入任务名称" autocomplete="off" class="layui-input" value="' + data.name + '"></div>&nbsp;&nbsp;&nbsp;&nbsp;执行周期&nbsp;&nbsp;&nbsp;&nbsp;<div class="layui-input-inline" style="width: 190px;"><input id="cron-edit-time-' + data.id + '" type="text" name="cron-edit-time" placeholder="请输入执行周期" autocomplete="off" class="layui-input" value="' + data.time + '"/></div><hr><div id="cron-edit-script-editor" style="height: 80%;">' + data.script + '</div><br><button id="cron-edit-' + data.id + '" class="layui-btn">保存</button>'
                    ,
                    success: function (layero, index) {
                        cronEditScriptEditor = ace.edit("cron-edit-script-editor", {
                            mode: "ace/mode/sh",
                            selectionStyle: "text"
                        });
                        cron.render({
                            elem: "#cron-edit-time-" + data.id,
                            btns: ['confirm'],
                            show: false,
                            done: function (value) {
                                $('#cron-add-time').val(value);
                            }
                        });
                        $('#cron-edit-' + data.id).click(function () {
                            admin.req({
                                url: '/api/panel/cron/edit'
                                , type: 'POST'
                                , data: {
                                    id: data.id,
                                    name: $('input[name="cron-edit-name"]').val(),
                                    time: $('input[name="cron-edit-time"]').val(),
                                    script: cronEditScriptEditor.getValue()
                                }
                                , success: function (res) {
                                    if (res.code === 0) {
                                        layer.msg('保存成功', {icon: 1, time: 1000});
                                        table.reload('panel-cron');
                                        layer.close(index);
                                    } else {
                                        layer.msg(res.msg, {icon: 2, time: 1000});
                                    }
                                }
                                , error: function (xhr, status, error) {
                                    console.log('耗子Linux面板：ajax请求出错，错误' + error);
                                }
                            });
                        });
                    }
                });
            } else if (obj.event === 'del') {
                layer.confirm('确定删除计划任务' + data.name + '吗？', function (index) {
                    layer.close(index);
                    admin.req({
                        url: '/api/panel/cron/delete',
                        type: 'POST',
                        data: {
                            id: data.id
                        }
                        , success: function (res) {
                            if (res.code === 0) {
                                table.reload('panel-cron');
                                layer.msg('计划任务：' + data.name + ' 已删除', {
                                    icon: 1,
                                    time: 1000
                                });
                            } else {
                                layer.msg(res.msg, {icon: 2, time: 1000});
                            }
                        }
                        , error: function (xhr, status, error) {
                            console.log('耗子Linux面板：ajax请求出错，错误' + error);
                        }
                    });
                });
            }
        });

        form.on('switch(cron-status)', function (obj) {
            let $ = layui.$;
            let id = $(this).data('id');
            let status = obj.elem.checked ? 1 : 0;

            admin.req({
                url: '/api/panel/cron/setStatus',
                type: 'POST',
                data: {
                    id: id,
                    status: status
                }
                , success: function (res) {
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1, time: 1000});
                    } else {
                        layer.msg(res.msg, {icon: 2, time: 1000});
                    }
                }
                , error: function (xhr, status, error) {
                    console.log('耗子Linux面板：ajax请求出错，错误' + error);
                }
            });
        });

        form.on('submit(cron-add-submit)', function (data) {
            data.field.script = cronAddScriptEditor.getValue();
            admin.req({
                url: "/api/panel/cron/add"
                , method: 'post'
                , data: data.field
                , success: function (result) {
                    if (result.code !== 0) {
                        console.log('耗子Linux面板：计划任务添加失败，接口返回' + result);
                        layer.msg('计划任务添加失败！')
                        return false;
                    }
                    table.reload('panel-cron');
                    layer.alert('计划任务添加成功！', {
                        icon: 1
                        , title: '提示'
                        , btn: ['确定']
                        , yes: function (index) {
                            layer.closeAll();
                            //location.reload();
                        }
                    });
                }
                , error: function (xhr, status, error) {
                    console.log('耗子Linux面板：ajax请求出错，错误' + error);
                }
            });
            return false;
        });
    });
</script>
