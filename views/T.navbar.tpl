{{define "header-bar"}}
   <div class="navbar navbar-default navbar-static-top" role="banner">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".header-navbar-collapse">
            <span class="sr-only">切换菜单</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/" rel="home" itemprop="headline">「一个 · 轻博客」</a>          
        </div>

        <nav id="navbar" class="collapse navbar-collapse header-navbar-collapse" role="navigation" >
        <ul class="nav navbar-nav">
            <!-- <li><a href="/" itemprop="url"><i class="_mi _before dashicons dashicons-admin-home" aria-hidden="true"></i><span>主页</span></a></li> -->

            <li {{if .IsHome}}class="active"{{end}}><a href="/"><span>首页</span></a></li>
            <li {{if .IsTholeaf}}class="active"{{end}}><a href="/tholeaf"><span>轻博</span></a></li>
            <li {{if .IsCategory}}class="active"{{end}}><a href="/category"><span>分类</span></a></li>
            <li {{if .IsTopic}}class="active"{{end}}><a href="/topic"><span>文章</span></a></li>

        </ul>

        <ul class="nav navbar-nav navbar-right">
                
            <li class="dropdown"><a href="/admin" data-toggle="dropdown" class="dropdown-toggle" itemprop="url">管理 <span class="caret"></span></a></a>         
                <ul class="dropdown-menu" role="menu">
                    {{if .IsLogin}}
                    <li><a href="/admin"><i aria-hidden="true"></i><span>后台管理</span></a></li>
                    <li><a href="/login?exit=true"><span>退出登录</span></a></li>
                    {{else}}
                    <li><a href="/login"><span>登录</span></a></li>
                    {{end}}
                </ul>
            </li>
        </ul>            

        </nav><!-- #navbar -->
    </div>
</div>

{{end}}

<li>
<a href="http://www.codebeta.cn/?p=31" title="历史" itemprop="url">
<i class="_mi _before dashicons dashicons-info" aria-hidden="true"></i>
<span>关于本站</span>
</a>
</li>