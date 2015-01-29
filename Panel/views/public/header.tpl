<!-- header start -->
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
              设置
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
<!-- header end -->