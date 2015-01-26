<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="HandheldFriendly" content="True" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
  <meta name="_xsrf" content="{{.xsrf_token}}" />
  <title>用户登录 - 西直门加速网络</title>
  <!--link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon"-->
  <link rel="stylesheet" type="text/css" href="/static/css/semantic.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
  <!-- {{template "public/header.tpl" .}} -->
  <div id="register" class="ui two column middle aligned relaxed grid basic segment">
    <!-- left column -->
    <div class="column">
      <div class="ui attached message">
        <div class="header">{{.sitename}}</div>
        <p>免费获取网络加速流量，畅享海量信息资源</p>
      </div>
      <div class="ui form attached segment">

        <div class="field">
          <label>用户名</label>
          <div class="ui left icon input">
            <input type="text" id="username" name="username" placeholder="用户名">
            <i class="user icon"></i>
            <div class="ui corner label red" data-content="用户名以字母开头,且不区分大小写">
              <i class="asterisk icon"></i>
            </div>
          </div>
        </div>

        <div class="field">
          <label>密码</label>
          <div class="ui left icon input">
            <input type="password" id="password" name="password" placeholder="密码">
            <i class="lock icon"></i>
            <div class="ui corner label red" data-content="密码必须填写,将作为服务器登录密码,且不能为弱口令">
              <i class="asterisk icon"></i>
            </div>
          </div>
        </div>

        <div id="code" class="two fields">
          <div class="field">
            <label>验证码</label>
            <input id="captcha" type="text" placeholder="请输入验证码" name="captcha">
          </div>
          <div class="field">
            <div id="verifycode">{{create_captcha}}</div>
          </div>
        </div>

        <div class="inline field">
          <div class="ui checkbox">
            <input type="checkbox" id="remember" name="remember" checked="chekced">
            <label for="remember">一周免登录(请勿在公共计算机使用)</label>
          </div>
        </div>
        <div class="center aligned column submit">
          <div class=" green ui labeled icon button"><i class="plus icon"></i>登录控制台</div>
        </div>

      </div>
      <div id="notice" class="ui bottom attached info message"><i class="icon warning sign"></i>敬告：请不要使用西直门进行非法活动,防止封号</div>
    </div>
    <!--/ left column -->

    <div class="ui vertical divider">Or</div>

    <!-- right column -->
    <div class="center aligned column">
      <a href="/register">
        <div class="huge blue ui labeled icon button">
          <i class="sign in icon"></i>注册新用户
        </div>
      </a>
    </div>
    <!--/ right column -->
  </div>

  <!-- {{template "public/footer.tpl" .}} -->
  {{template "public/js.tpl" .}}
  <script type="text/javascript">
  $(document).ready(function(){
    $("#username").focus();
    // 提交
    $(".submit").click(function(){
      var url       = "/login";
      var username  = $("#username").val();
      var password  = $("#password").val();
      var captcha   = $("#captcha").val();
      var captcha_id = $("input[name=captcha_id]").val();
      $(this).submit();
      $.post(url, {
        "username": username, 
        "password": password, 
        "captcha_id": captcha_id,
        "captcha": captcha
      }, 
      function(data){
        if (data.result == "-1") { // 验证码错误
          $(".captcha img").click()
        };
        if (data.result == "true") {
          document.location.href = "/home"
        };
        console.log(data)
      });
    });
  });

  // 提示信息
    $("*").popup({
      position: 'right center'
    });
  // 表单验证
  $('.ui.form').form({
    username: {
      identifier  : 'username',
      rules: [{
        type   : 'empty',
        prompt : '请填写用户名'
      },{
        type  : 'length[4]',
        prompt  : '用户名至少4个字符'
      }]
    },
    password: {
      identifier  : 'password',
      rules: [{
        type   : 'empty',
        prompt : '请填写密码'
      },{
        type   : 'length[6]',
        prompt : '密码最少为六位'
      }]
    },
    captcha: {
      identifier : 'captcha',
      rules : [{
        type   : 'length[6]',
        prompt : '请完整填写验证码'
      }] 
    },
  }, {
    inline : true,
    on     : 'blur'
  });
  </script>
</body>
</html>