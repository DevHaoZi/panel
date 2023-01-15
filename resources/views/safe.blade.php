<title>系统安全</title>

<div class="layui-fluid">
    <div class="layui-card">
        <div class="layui-form layui-card-header layuiadmin-card-header-auto">
            <div class="layui-inline">
                <span style="margin-right: 10px;">防火墙</span>
                <input type="checkbox" id="safe_firewall" lay-filter="safe_firewall" lay-skin="switch"
                       lay-text="ON|OFF">
                <span style="margin: 0px 10px;">启用SSH</span>
                <input type="checkbox" id="safe_ssh" lay-filter="safe_ssh" lay-skin="switch" lay-text="ON|OFF">
                <span style="margin: 0px 10px 0px 20px;">SSH端口</span>
                <div class="layui-input-inline" style="width: 80px;">
                    <input type="number" id="safe_ssh_port" class="layui-input" style="height: 30px; margin-top: 5px;"
                           min=1
                           max=65535 disabled>
                </div>
                <div class="layui-input-inline">
                    <button id="safe_ssh_port_save" class="layui-btn layui-btn-sm layui-btn-primary">确定
                    </button>
                </div>
                <span style="margin: 0px 10px 0px 20px;">允许Ping</span>
                <input type="checkbox" id="switch_ping" lay-filter="safe_ping" lay-skin="switch" lay-text="ON|OFF">
            </div>
            <div class="layui-inline" style="float: right;">
                {{--<button class="layui-btn layui-btn-sm layui-btn-danger">清空 OpenResty 日志
                </button>--}}
            </div>
        </div>
    </div>

    <div id="vm_security">
        <div class="layui-card">
            <div class="layui-form layui-card-header layuiadmin-card-header-auto">
                <div class="layui-inline">
                    <span style="margin-right: 10px;">端口控制</span>
                    <div class="layui-input-inline">
                        <input id="safe_add_firewall_rule_port" type="text" name="safe_add_firewall_rule_port" class="layui-input"
                               placeholder="例如：3306、1000-2000">
                    </div>
                    <div class="layui-input-inline">
                        <select id="safe_add_firewall_rule_protocol" lay-filter="safe_add_firewall_rule_protocol"
                                style="height: 30px; margin-top: 5px;">
                            <option value="tcp">TCP</option>
                            <option value="udp">UDP</option>
                        </select>
                    </div>

                    <div class="layui-input-inline">
                        <button id="safe_add_firewall_rule" class="layui-btn layui-btn-sm" style="margin-top: -4px;">放行
                        </button>
                    </div>
                </div>
            </div>
            <div class="layui-card-body">
                <table class="layui-hide" id="safe-port" lay-filter="safe-port"></table>
                <!-- 右侧删除端口 -->
                <script type="text/html" id="safe-port-setting">
                    <a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
                </script>
            </div>
        </div>

    </div>
</div>

<script>
    layui.use(['layer', 'admin', 'form', 'laypage', 'table'], function () {
        var $ = layui.$;
        var admin = layui.admin;
        var table = layui.table;
        var form = layui.form;
        var layer = layui.layer;

        // 获取防火墙状态
        admin.req({
            url: '/api/panel/safe/getFirewallStatus'
            , type: 'get'
            , dataType: 'json'
            , success: function (res) {
                if (res.code === 0) {
                    // 防火墙
                    if (res.data) {
                        $('#safe_firewall').attr('checked', true);
                    } else {
                        $('#safe_firewall').attr('checked', false);
                    }
                    form.render();
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }
        });
        // 获取SSH状态
        admin.req({
            url: '/api/panel/safe/getSshStatus'
            , type: 'get'
            , dataType: 'json'
            , success: function (res) {
                if (res.code === 0) {
                    // SSH
                    if (res.data) {
                        $('#safe_ssh').attr('checked', true);
                    } else {
                        $('#safe_ssh').attr('checked', false);
                    }
                    form.render();
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }
        });
        // 获取SSH端口
        admin.req({
            url: '/api/panel/safe/getSshPort'
            , type: 'get'
            , dataType: 'json'
            , success: function (res) {
                if (res.code === 0) {
                    // SSH端口
                    $('#safe_ssh_port').val(res.data);
                    $('#safe_ssh_port').attr('disabled', false);
                    form.render();
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }
        });
        // 获取ping状态
        admin.req({
            url: '/api/panel/safe/getPingStatus'
            , type: 'get'
            , dataType: 'json'
            , success: function (res) {
                if (res.code === 0) {
                    // ping
                    if (res.data) {
                        $('#switch_ping').attr('checked', true);
                    } else {
                        $('#switch_ping').attr('checked', false);
                    }
                    form.render();
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }
        });

        // 设置防火墙开关
        form.on('switch(safe_firewall)', function (data) {
            admin.req({
                url: '/api/panel/safe/setFirewallStatus'
                , type: 'post'
                , dataType: 'json'
                , data: {
                    status: data.elem.checked === true ? 1 : 0
                }
                , success: function (res) {
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1});
                    } else {
                        layer.msg(res.msg, {icon: 2});
                    }
                }
            });
        });

        // 设置SSH开关
        form.on('switch(safe_ssh)', function (data) {
            admin.req({
                url: '/api/panel/safe/setSshStatus'
                , type: 'post'
                , dataType: 'json'
                , data: {
                    status: data.elem.checked === true ? 1 : 0
                }
                , success: function (res) {
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1});
                    } else {
                        layer.msg(res.msg, {icon: 2});
                    }
                }
            });
        });

        // 设置ping开关
        form.on('switch(safe_ping)', function (data) {
            admin.req({
                url: '/api/panel/safe/setPingStatus'
                , type: 'post'
                , dataType: 'json'
                , data: {
                    status: data.elem.checked === true ? 1 : 0
                }
                , success: function (res) {
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1});
                    } else {
                        layer.msg(res.msg, {icon: 2});
                    }
                }
            });
        });

        table.render({
            elem: '#safe-port'
            , url: '/api/panel/safe/getFirewallRules'
            , title: '防火墙'
            , cols: [[
                {field: 'port', title: '端口', width: 100, sort: true}
                , {field: 'protocol', title: '协议', sort: true}
                , {fixed: 'right', title: '操作', toolbar: '#safe-port-setting', width: 150}
            ]]
        });
        table.on('tool(safe-port)', function (obj) {
            let data = obj.data;
            if (obj.event === 'del') {
                layer.confirm('确定要删除 <b style="color: red;">' + data.protocol + '</b> 端口 <b style="color: red;">' + data.port + '</b> 吗？', function (index) {
                    admin.req({
                        url: "/api/panel/safe/deleteFirewallRule"
                        , method: 'post'
                        , data: data
                        , success: function (result) {
                            if (result.code !== 0) {
                                console.log('耗子Linux面板：端口删除失败，接口返回' + result);
                                layer.msg('网站删除失败，请刷新重试！')
                                return false;
                            }
                            obj.del();
                            layer.alert('<b style="color: red;">' + data.protocol + '</b> 端口 <b style="color: red;">' + data.port + '</b> 删除成功！');
                        }
                        , error: function (xhr, status, error) {
                            console.log('耗子Linux面板：ajax请求出错，错误' + error);
                        }
                    });
                    layer.close(index);
                });
            }
        });

        // 监听ssh端口保存
        $('#safe_ssh_port_save').click(function () {
            var port = Number($('#safe_ssh_port').val());
            // 判断端口是否合法
            if (isNaN(port) || port < 1 || port > 65535) {
                layer.msg('端口号不合法', {icon: 2});
                return false;
            }
            var index = layer.load();
            admin.req({
                url: '/api/panel/safe/setSshPort'
                , type: 'post'
                , dataType: 'json'
                , data: {
                    port: port
                }
                , success: function (res) {
                    layer.close(index);
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1});
                    } else {
                        layer.msg(res.msg, {icon: 2});
                    }
                }
            });
        });

        // 监听添加端口保存
        $('#safe_add_firewall_rule').click(function () {
            var port = $('#safe_add_firewall_rule_port').val();
            var protocol = $('#safe_add_firewall_rule_protocol').val();
            var index = layer.load();
            admin.req({
                url: '/api/panel/safe/addFirewallRule'
                , type: 'post'
                , dataType: 'json'
                , data: {
                    port: port,
                    protocol: protocol
                }
                , success: function (res) {
                    layer.close(index);
                    if (res.code === 0) {
                        layer.msg('设置成功', {icon: 1});
                        table.reload('safe-port');
                    } else {
                        layer.msg(res.msg, {icon: 2});
                    }
                }
            });
        });
    });
</script>