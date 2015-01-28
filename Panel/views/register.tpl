<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="HandheldFriendly" content="True" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
  <meta name="_xsrf" content="{{.xsrf_token}}" />
  <title>用户注册 - 西直门加速网络</title>
  <!--link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon"-->
  <link rel="stylesheet" type="text/css" href="/static/css/semantic.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
  <div id="register" class="ui two column middle aligned relaxed grid basic segment">
    <!-- left column -->
    <div class="column">
      <div class="ui attached message">
        <div class="header">{{.sitename}}</div>
        <p>免费获取网络加速流量，畅享海量信息资源</p>
      </div>
      <div class="ui form attached segment">

        <div class="field">
          <label>邀请码</label>
          <div class="ui left icon input">
            <input type="text" id="invite" name="invite" value="{{.invite}}" placeholder="邀请码">
            <i class="gift icon"></i>
            <div class="ui corner label red" data-content="邀请码,必须填写">
              <i class="asterisk icon"></i>
            </div>
          </div>
        </div>

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
          <label>邮箱</label>
          <div class="ui left icon input">
            <input type="text" id="email" name="email" placeholder="邮箱">
            <i class="mail outline icon"></i>
            <div class="ui corner label red" data-content="请填写常用邮箱,我们将向此邮箱发送注册确认通知.">
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
            <input type="checkbox" id="terms" name="terms" checked="chekced">
            <label for="terms"><a href="/site/terms">同意用户协议</a></label>
          </div>
        </div>
        <div class="center aligned column">
          <div class=" blue ui labeled icon button"><i class="plus icon"></i>注册新用户</div>
        </div>

      </div>
      <div id="notice" class="ui bottom attached info message"><i class="icon warning sign"></i>敬告：请不要使用西直门进行非法活动,防止封号</div>
    </div>
    <!--/ left column -->

    <div class="ui vertical divider">Or</div>

    <!-- right column -->
    <div class="center aligned column">
      <a href="/login">
        <div class="huge green ui labeled icon button">
          <i class="sign in icon"></i>登录控制台
        </div>
      </a>
    </div>
    <!--/ right column -->
  </div>
  
  <script src="/static/js/sha1.js" type="text/javascript"></script>
  {{template "public/js.tpl" .}}
  <script type="text/javascript">
  $(document).ready(function(){
    {{if .invite }}
    $("#invite").focus();
    {{end}}
  });

  // 提示信息
    $("*").popup({
      position: 'right center'
    });
  // 表单验证
  $('.ui.form').form({
    invite: {
      identifier  : 'invite',
      rules: [{
        type   : 'empty',
        prompt : '请填写邀请码'
      }, {
        type   : 'length[9]',
        prompt : '邀请码无效',
      }, {
        type   : 'maxLength[9]',
        prompt : '邀请码无效',
      }, {
        type   : 'invite',
        prompt : '邀请码不可用'
      }]
    },
    username: {
      identifier  : 'username',
      rules: [{
        type   : 'empty',
        prompt : '请填写用户名'
      },{
		type	: 'length[4]',
		prompt  : '用户名至少4个字符'
		}, {
        type   : 'username',
        prompt   : '用户名已被注册,请重新填写'
      }]
    },
    email: {
      identifier  : 'email',
      rules: [{
        type   : 'empty',
        prompt : '请填写联系邮箱'
      }, {
        type   : 'email',
        prompt : '无效邮件地址'
      }, {
        type   : 'mail',
        prompt : '邮箱不可用'
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
      }, {
        type   : 'notcontains[username]',
        prompt : '密码中不能包含用户名'
      }]
    },
    captcha: {
      identifier : 'captcha',
      rules : [{
        type   : 'length[6]',
        prompt : '请完整填写验证码'
      }] 
    },
    terms: {
      identifier : 'terms',
      rules: [{
        type   : 'checked',
        prompt : '使用本站提供的服务需要同意服务协议'
      }]
    }
  }, {
    inline : true,
    on     : 'blur'
  }, {
    debug  : true
  });
  </script>
</body>
</html>