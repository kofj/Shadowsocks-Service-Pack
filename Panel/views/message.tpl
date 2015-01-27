<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="HandheldFriendly" content="True" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" />
  <meta name="_xsrf" content="{{.xsrf_token}}" />
  <title>页面跳转 - 西直门加速网络</title>
  <!--link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon"-->
  <link rel="stylesheet" type="text/css" href="/static/css/semantic.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
</head>
<body>
  <div class="ui page grid body">

    <dic class="column">
      <div style="min-height: 8rem;" class="ui segment">
        <div class="ui top attached label">{{ .msg.title }}</div>
        <h3>{{.msg.msg}}</h3>
        <div id="myurl" style="display:none;">{{.msg.url}}</div>
        <div class="ui bottom attached label">页面自动 <a id="href" href="{{.msg.url}}">跳转</a> 等待时间： <b id="wait">{{.msg.wait}}</b></div>
      </div>
    </dic>

  </div>

  <script type="text/javascript">
  (function(){
    document.getElementById('href').href = document.getElementById('myurl').innerHTML;
    var wait = document.getElementById('wait'),href = document.getElementById('href').href;
    totaltime=parseInt(wait.innerHTML);
    var interval = setInterval(function(){
      var time = --totaltime;
      wait.innerHTML=""+time;
      if(time === 0) {
        location.href = href;
        clearInterval(interval);
      };
    }, 1000);
  })(); 
  </script>

</body>
</html>