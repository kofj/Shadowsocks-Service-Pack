<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="HandheldFriendly" content="True" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
  <meta name="_xsrf" content="{{.xsrf_token}}" />
  <title>面板 - 西直门加速网络</title>
  <!--link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon"-->
  <link rel="stylesheet" type="text/css" href="/static/css/semantic.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
  <header class="following bar light fixed">
    <div class="ui page grid">
      <div class="column">
        <div class="ui logo shape">
          <div class="sides">
            <div class="active ui side">
              <img class="ui image" src="/static/img/logo/small.png">
            </div>
          </div>
        </div>
        <div class="ui right floated secondary menu">
          <div class="item"></div>
          <div class="ui compact menu">
            <div class="ui simple dropdown item">
              <i class="dropdown icon"></i>
              FrankKung
              <div class="menu">
                <div class="item">修改资料</div>
                <div class="item">修改密码</div>
                <div class="item"><a href="{{.logout}}">退出登录</a></div>
              </div>
            </div>
          </div>
        </div>
        <div class="ui large secondary network menu">
          <a class="view-ui item">
            <i class="sidebar icon"></i> 概览
          </a>
          <a class="view-ui item">充值</a>
        </div>
      </div>
    </div>
  </header>

  <div class="ui page grid body">

    <div class="three column row">
          <div class="column">
            <img src="http://semantic-ui.com/images/wireframe/image.png">
          </div>
          <div class="column">
            <img src="http://semantic-ui.com/images/wireframe/image.png">
          </div>
          <!--div class="column">
            <img src="http://semantic-ui.com/images/wireframe/image.png">
          </div-->
        </div>
    <div class="column">


    </div>
  </div>

  {{template "public/js.tpl" .}}

</body>
</html>