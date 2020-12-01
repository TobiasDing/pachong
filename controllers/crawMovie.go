package controllers

import (
	beego "github.com/astaxie/beego/server/web"
	"pachong/models"
)

type CrawMovieController struct {
	beego.Controller
}

func (c CrawMovieController) CrawMovie() {
	sMovieHtml := `

<!DOCTYPE html>
<html lang="zh-CN" class="ua-mac ua-webkit">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="renderer" content="webkit">
    <meta name="referrer" content="always">
    <meta name="google-site-verification" content="ok0wCgT20tBBgo9_zat2iAcimtN4Ftf5ccsh092Xeyw" />
    <title>
        一秒钟 (豆瓣)
</title>
    
    <meta name="baidu-site-verification" content="cZdR4xxR7RxmM4zE" />
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">
    
    <link rel="apple-touch-icon" href="https://img3.doubanio.com/f/movie/d59b2715fdea4968a450ee5f6c95c7d7a2030065/pics/movie/apple-touch-icon.png">
    <link href="https://img3.doubanio.com/f/shire/859dba5cddc7ed1435808cf5a8ddde5792cd6e0c/css/douban.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/shire/db02bd3a4c78de56425ddeedd748a6804af60ee9/css/separation/_all.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/movie/252bef058b97005c6a41e8f1b9f7b06b84bc71b3/css/movie/base/init.css" rel="stylesheet">
    <script type="text/javascript">var _head_start = new Date();</script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/movie/0495cb173e298c28593766009c7b0a953246c5b5/js/movie/lib/jquery.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/5ecaf46d6954d5a30bc7d99be86ae34031646e00/js/douban.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/0efdc63b77f895eaf85281fb0e44d435c6239a3f/js/separation/_all.js"></script>
    
    <meta name="keywords" content="一秒钟,一秒钟,一秒钟影评,剧情介绍,电影图片,预告片,影讯,在线购票,论坛">
    <meta name="description" content="一秒钟电影简介和剧情介绍,一秒钟影评、图片、预告片、影讯、论坛、在线购票">
    <meta name="mobile-agent" content="format=html5; url=https://m.douban.com/movie/subject/30257787/"/>
    <link rel="alternate" href="android-app://com.douban.movie/doubanmovie/subject/30257787/" />
    <link rel="stylesheet" href="https://img3.doubanio.com/dae/cdnlib/libs/LikeButton/1.0.5/style.min.css">
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript">
      Do.add('dialog', {path: 'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js', type: 'js'});
      Do.add('dialog-css', {path: 'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css', type: 'css'});
      Do.add('handlebarsjs', {path: 'https://img3.doubanio.com/f/movie/3d4f8e4a8918718256450eb6e57ec8e1f7a2e14b/js/movie/lib/handlebars.current.js', type: 'js'});
    </script>
    
  <script type='text/javascript'>
  var _vwo_code = (function() {
    var account_id = 249272,
      settings_tolerance = 0,
      library_tolerance = 2500,
      use_existing_jquery = false,
      // DO NOT EDIT BELOW THIS LINE
      f=false,d=document;return{use_existing_jquery:function(){return use_existing_jquery;},library_tolerance:function(){return library_tolerance;},finish:function(){if(!f){f=true;var a=d.getElementById('_vis_opt_path_hides');if(a)a.parentNode.removeChild(a);}},finished:function(){return f;},load:function(a){var b=d.createElement('script');b.src=a;b.type='text/javascript';b.innerText;b.onerror=function(){_vwo_code.finish();};d.getElementsByTagName('head')[0].appendChild(b);},init:function(){settings_timer=setTimeout('_vwo_code.finish()',settings_tolerance);var a=d.createElement('style'),b='body{opacity:0 !important;filter:alpha(opacity=0) !important;background:none !important;}',h=d.getElementsByTagName('head')[0];a.setAttribute('id','_vis_opt_path_hides');a.setAttribute('type','text/css');if(a.styleSheet)a.styleSheet.cssText=b;else a.appendChild(d.createTextNode(b));h.appendChild(a);this.load('//dev.visualwebsiteoptimizer.com/j.php?a='+account_id+'&u='+encodeURIComponent(d.URL)+'&r='+Math.random());return settings_timer;}};}());

  +function () {
    var bindEvent = function (el, type, handler) {
        var $ = window.jQuery || window.Zepto || window.$
       if ($ && $.fn && $.fn.on) {
           $(el).on(type, handler)
       } else if($ && $.fn && $.fn.bind) {
           $(el).bind(type, handler)
       } else if (el.addEventListener){
         el.addEventListener(type, handler, false);
       } else if (el.attachEvent){
         el.attachEvent("on" + type, handler);
       } else {
         el["on" + type] = handler;
       }
     }

    var _origin_load = _vwo_code.load
    _vwo_code.load = function () {
      var args = [].slice.call(arguments)
      bindEvent(window, 'load', function () {
        _origin_load.apply(_vwo_code, args)
      })
    }
  }()

  _vwo_settings_timer = _vwo_code.init();
  </script>


    


<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "name": "一秒钟",
  "url": "/subject/30257787/",
  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2626693662.webp",
  "director": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1054398/",
      "name": "张艺谋 Yimou Zhang"
    }
  ]
,
  "author": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1054398/",
      "name": "张艺谋 Yimou Zhang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1276029/",
      "name": "邹静之 Jingzhi Zou"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1410271/",
      "name": "周晓枫 Xiaofeng Zhou"
    }
  ]
,
  "actor": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1274761/",
      "name": "张译 Yi Zhang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1410267/",
      "name": "刘浩存 Haocun Liu"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1051526/",
      "name": "范伟 Wei Fan"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1313809/",
      "name": "余皑磊 Ailei Yu"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1410678/",
      "name": "张邵勃  Shaobo Zhang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1353889/",
      "name": "李延 Yan Li"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1338461/",
      "name": "于洋 Yang Yu"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1407244/",
      "name": "刘云龙 Yunlong Liu"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1424139/",
      "name": "常海军 Haijun Chang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1364324/",
      "name": "唐孜悦 Ziyue Tang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1367528/",
      "name": "曹瑞 Rui Cao"
    }
  ]
,
  "datePublished": "2020-11-27",
  "genre": ["\u5267\u60c5"],
  "duration": "PT1H44M",
  "description": "影片讲述了没赶上电影场次的张九声与刘闺女因一场电影结下了不解之缘的故事。故事灵感来源于张艺谋导演早期经历，是其一贯对文化展开追忆和寻根的风格。",
  "@type": "Movie",
  "aggregateRating": {
    "@type": "AggregateRating",
    "ratingCount": "72784",
    "bestRating": "10",
    "worstRating": "2",
    "ratingValue": "7.9"
  }
}
</script>


    <style type="text/css">
  
</style>
    <style type="text/css">img { max-width: 100%; }</style>
    <script type="text/javascript"></script>
    <link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/6b909164236e53c6.css">

    <link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
</head>

<body>
  
    <script type="text/javascript">var _body_start = new Date();</script>

    
    



    <link href="//img3.doubanio.com/dae/accounts/resources/19870c3/shire/bundle.css" rel="stylesheet" type="text/css">



<div id="db-global-nav" class="global-nav">
  <div class="bd">
    
<div class="top-nav-info">
  <a href="https://accounts.douban.com/passport/login?source=movie" class="nav-login" rel="nofollow">登录/注册</a>
</div>


    <div class="top-nav-doubanapp">
  <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
  <div id="doubanapp-tip">
    <a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 <span class="version">6.0</span> 全新发布</a>
    <a href="javascript: void 0;" class="tip-close">×</a>
  </div>
  <div id="top-nav-appintro" class="more-items">
    <p class="appintro-title">豆瓣</p>
    <p class="qrcode">扫码直接下载</p>
    <div class="download">
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=iOS">iPhone</a>
      <span>·</span>
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=Android" class="download-android">Android</a>
    </div>
  </div>
</div>

    


<div class="global-nav-items">
  <ul>
    <li class="">
      <a href="https://www.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-main&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣</a>
    </li>
    <li class="">
      <a href="https://book.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-book&quot;,&quot;uid&quot;:&quot;0&quot;}">读书</a>
    </li>
    <li class="on">
      <a href="https://movie.douban.com"  data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-movie&quot;,&quot;uid&quot;:&quot;0&quot;}">电影</a>
    </li>
    <li class="">
      <a href="https://music.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-music&quot;,&quot;uid&quot;:&quot;0&quot;}">音乐</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/location" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-location&quot;,&quot;uid&quot;:&quot;0&quot;}">同城</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/group" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-group&quot;,&quot;uid&quot;:&quot;0&quot;}">小组</a>
    </li>
    <li class="">
      <a href="https://read.douban.com&#47;?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-read&quot;,&quot;uid&quot;:&quot;0&quot;}">阅读</a>
    </li>
    <li class="">
      <a href="https://douban.fm&#47;?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-fm&quot;,&quot;uid&quot;:&quot;0&quot;}">FM</a>
    </li>
    <li class="">
      <a href="https://time.douban.com&#47;?dt_time_source=douban-web_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-time&quot;,&quot;uid&quot;:&quot;0&quot;}">时间</a>
    </li>
    <li class="">
      <a href="https://market.douban.com&#47;?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-market&quot;,&quot;uid&quot;:&quot;0&quot;}">豆品</a>
    </li>
  </ul>
</div>

  </div>
</div>
<script>
  ;window._GLOBAL_NAV = {
    DOUBAN_URL: "https://www.douban.com",
    N_NEW_NOTIS: 0,
    N_NEW_DOUMAIL: 0
  };
</script>



    <script src="//img3.doubanio.com/dae/accounts/resources/19870c3/shire/bundle.js" defer="defer"></script>




    



    <link href="//img3.doubanio.com/dae/accounts/resources/19870c3/movie/bundle.css" rel="stylesheet" type="text/css">




<div id="db-nav-movie" class="nav">
  <div class="nav-wrap">
  <div class="nav-primary">
    <div class="nav-logo">
      <a href="https:&#47;&#47;movie.douban.com">豆瓣电影</a>
    </div>
    <div class="nav-search">
      <form action="https:&#47;&#47;search.douban.com&#47;movie/subject_search" method="get">
        <fieldset>
          <legend>搜索：</legend>
          <label for="inp-query">
          </label>
          <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" placeholder="搜索电影、电视剧、综艺、影人" value=""></div>
          <div class="inp-btn"><input type="submit" value="搜索"></div>
          <input type="hidden" name="cat" value="1002" />
        </fieldset>
      </form>
    </div>
  </div>
  </div>
  <div class="nav-secondary">
    

<div class="nav-items">
  <ul>
    <li    ><a href="https://movie.douban.com/cinema/nowplaying/"
     >影讯&购票</a>
    </li>
    <li    ><a href="https://movie.douban.com/explore"
     >选电影</a>
    </li>
    <li    ><a href="https://movie.douban.com/tv/"
     >电视剧</a>
    </li>
    <li    ><a href="https://movie.douban.com/chart"
     >排行榜</a>
    </li>
    <li    ><a href="https://movie.douban.com/tag/"
     >分类</a>
    </li>
    <li    ><a href="https://movie.douban.com/review/best/"
     >影评</a>
    </li>
    <li    ><a href="https://movie.douban.com/annual/2019?source=navigation"
     >2019年度榜单</a>
    </li>
    <li    ><a href="https://m.douban.com/standbyme/annual2019?source=navigation"
            target="_blank"
     >2019书影音报告</a>
    </li>
  </ul>
</div>

    <a href="https://movie.douban.com/annual/2019?source=movie_navigation" class="movieannual"></a>
  </div>
</div>

<script id="suggResult" type="text/x-jquery-tmpl">
  <li data-link="{{= url}}">
            <a href="{{= url}}" onclick="moreurl(this, {from:'movie_search_sugg', query:'{{= keyword }}', subject_id:'{{= id}}', i: '{{= index}}', type: '{{= type}}'})">
            <img src="{{= img}}" width="40" />
            <p>
                <em>{{= title}}</em>
                {{if year}}
                    <span>{{= year}}</span>
                {{/if}}
                {{if sub_title}}
                    <br /><span>{{= sub_title}}</span>
                {{/if}}
                {{if address}}
                    <br /><span>{{= address}}</span>
                {{/if}}
                {{if episode}}
                    {{if episode=="unknow"}}
                        <br /><span>集数未知</span>
                    {{else}}
                        <br /><span>共{{= episode}}集</span>
                    {{/if}}
                {{/if}}
            </p>
        </a>
        </li>
  </script>




    <script src="//img3.doubanio.com/dae/accounts/resources/19870c3/movie/bundle.js" defer="defer"></script>





    
    <div id="wrapper">
        

        
    <div id="content">
        

    <div id="dale_movie_subject_top_icon"></div>
    <h1>
        <span property="v:itemreviewed">一秒钟</span>
            <span class="year">(2020)</span>
    </h1>

        <div class="grid-16-8 clearfix">
            

            
            <div class="article">
                
    

    





        <div class="indent clearfix">
            <div class="subjectwrap clearfix">
                <div class="subject clearfix">
                    



<div id="mainpic" class="">
    <a class="nbgnbg" href="https://movie.douban.com/subject/30257787/photos?type=R" title="点击看更多海报">
        <img src="https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2626693662.webp" title="点击看更多海报" alt="一秒钟" rel="v:image" />
   </a>
</div>

                    


<div id="info">
        <span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/1054398/" rel="v:directedBy">张艺谋</a></span></span><br/>
        <span ><span class='pl'>编剧</span>: <span class='attrs'><a href="/celebrity/1054398/">张艺谋</a> / <a href="/celebrity/1276029/">邹静之</a></span></span><br/>
        <span class="actor"><span class='pl'>主演</span>: <span class='attrs'><a href="/celebrity/1274761/" rel="v:starring">张译</a> / <a href="/celebrity/1410267/" rel="v:starring">刘浩存</a> / <a href="/celebrity/1051526/" rel="v:starring">范伟</a> / <a href="/celebrity/1313809/" rel="v:starring">余皑磊</a> / <a href="/celebrity/1410678/" rel="v:starring">张邵勃 </a> / <a href="/celebrity/1353889/" rel="v:starring">李延</a> / <a href="/celebrity/1338461/" rel="v:starring">于洋</a> / <a href="/celebrity/1407244/" rel="v:starring">刘云龙</a> / <a href="/celebrity/1424139/" rel="v:starring">常海军</a></span></span><br/>
        <span class="pl">类型:</span> <span property="v:genre">剧情</span><br/>
        
        <span class="pl">制片国家/地区:</span> 中国大陆<br/>
        <span class="pl">语言:</span> 汉语普通话<br/>
        <span class="pl">上映日期:</span> <span property="v:initialReleaseDate" content="2020-11-27(中国大陆)">2020-11-27(中国大陆)</span><br/>
        <span class="pl">片长:</span> <span property="v:runtime" content="104">104分钟</span><br/>
        <span class="pl">又名:</span> One Second<br/>
        <span class="pl">IMDb链接:</span> <a href="https://www.imdb.com/title/tt8959680" target="_blank" rel="nofollow">tt8959680</a><br>

</div>




                </div>
                    


<div id="interest_sectl">
    <div class="rating_wrap clearbox" rel="v:rating">
        <div class="clearfix">
          <div class="rating_logo ll">豆瓣评分</div>
          <div class="output-btn-wrap rr" style="display:none">
            <img src="https://img3.doubanio.com/f/movie/692e86756648f29457847c5cc5e161d6f6b8aaac/pics/movie/reference.png" />
            <a class="download-output-image" href="#">引用</a>
          </div>
          
          
        </div>
        



<div class="rating_self clearfix" typeof="v:Rating">
    <strong class="ll rating_num" property="v:average">7.9</strong>
    <span property="v:best" content="10.0"></span>
    <div class="rating_right ">
        <div class="ll bigstar bigstar40"></div>
        <div class="rating_sum">
                <a href="comments" class="rating_people">
                    <span property="v:votes">72860</span>人评价
                </a>
        </div>
    </div>
</div>
<div class="ratings-on-weight">
    
        <div class="item">
        
        <span class="stars5 starstop" title="力荐">
            5星
        </span>
        <div class="power" style="width:22px"></div>
        <span class="rating_per">19.9%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars4 starstop" title="推荐">
            4星
        </span>
        <div class="power" style="width:64px"></div>
        <span class="rating_per">56.0%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars3 starstop" title="还行">
            3星
        </span>
        <div class="power" style="width:25px"></div>
        <span class="rating_per">22.0%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars2 starstop" title="较差">
            2星
        </span>
        <div class="power" style="width:1px"></div>
        <span class="rating_per">1.7%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars1 starstop" title="很差">
            1星
        </span>
        <div class="power" style="width:0px"></div>
        <span class="rating_per">0.4%</span>
        <br />
        </div>
</div>

    </div>
        <div class="rating_betterthan">
            好于 <a href="/typerank?type_name=剧情&type=11&interval_id=75:65&action=">70% 剧情片</a><br/>
        </div>
</div>


                
            </div>
                




<div id="interest_sect_level" class="clearfix">
        
            <a href="https://www.douban.com/reason=collectwish&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-30257787-wish">
                <span>想看</span>
            </a>
            <a href="https://www.douban.com/reason=collectcollect&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-30257787-collect">
                <span>看过</span>
            </a>
        <div class="ll j a_stars">
            
    
    评价:
    <span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">

            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30257787-1">
            <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star1" width="16" height="16"/>
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30257787-2">
            <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star2" width="16" height="16"/>
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30257787-3">
            <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star3" width="16" height="16"/>
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30257787-4">
            <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star4" width="16" height="16"/>
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30257787-5">
            <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star5" width="16" height="16"/>
        </a>
    </span><span id="rateword" class="pl"></span>
    <input id="n_rating" type="hidden" value=""  />
    </span>

        </div>
    

</div>


            

















<div class="gtleft">
    <ul class="ul_subject_menu bicelink color_gray pt6 clearfix">
        
    
        
                <li> 
    <img src="https://img3.doubanio.com/f/shire/cc03d0fcf32b7ce3af7b160a0b85e5e66b47cc42/pics/short-comment.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_cmnt_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写短评</a>
 </li>
                    <li> 
    
    <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_rv_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写影评</a>
 </li>
                <li> 
    



 </li>
                <li> 
   

   
    
    <span class="rec" id="电影-30257787">
    <a href= "#"
        data-type="电影"
        data-url="https://movie.douban.com/subject/30257787/"
        data-desc="电影《一秒钟》 (来自豆瓣) "
        data-title="电影《一秒钟》 (来自豆瓣) "
        data-pic="https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2626693662.jpeg"
        class="bn-sharing ">
        分享到
    </a> &nbsp;&nbsp;
    </span>

    <script>
        if (!window.DoubanShareMenuList) {
            window.DoubanShareMenuList = [];
        }
        var __cache_url = __cache_url || {};

        (function(u){
            if(__cache_url[u]) return;
            __cache_url[u] = true;
            window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';

            var initShareButton = function() {
                $.ajax({url:u,dataType:'script',cache:true});
            };

            if (typeof Do == 'function' && 'ready' in Do) {
                Do(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js',
                    initShareButton
                );
            } else if(typeof Douban == 'object' && 'loader' in Douban) {
                Douban.loader.batch(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js'
                ).done(initShareButton);
            }

        })('https://img3.doubanio.com/f/movie/32be6727ed3ad8f6c4a417d8a086355c3e7d1d27/js/movie/lib/sharebutton.js');
    </script>


  </li>
            

    </ul>

    <script type="text/javascript">
        $(function(){
            $(".ul_subject_menu li.rec .bn-sharing").bind("click", function(){
                $.get("/blank?sbj_page_click=bn_sharing");
            });
            $(".ul_subject_menu .create_from_menu").bind("click", function(e){
                e.preventDefault();
                var $el = $(this);
                var glRoot = document.getElementById('gallery-topics-selection');
                if (window.has_gallery_topics && glRoot) {
                    // 判断是否有话题
                    glRoot.style.display = 'block';
                } else {
                    location.href = $el.attr('href');
                }
            });
        });
    </script>
</div>




                




    <style type="text/css">
        
.suggestions-list li { position: relative; left: 0; top: 0; margin-bottom: 7px; height: 35px }
.suggestions-list li .user-thumb { display: inline-block; *display: inline; float: left; margin: 2px 5px 0 0; vertical-align: top }
.suggestions-list li .user-thumb img { width: 25px; height: 25px }
.suggestions-list li .user-name-info { display: inline-block; *display: inline; line-height: 1.4em }
.suggestions-list li .user-name-info .user-profile-link { color: #333; font-weight: 800 }
.suggestions-list li .user-name-info .user-profile-link:hover { color: #4b8dc5 }
.suggestions-list li .user-name-info p { color: #999 }
.suggestions-list li .user-name-info b { color: #4b8dc5; font-weight: normal; cursor: pointer }
.suggestions-list li .user-name-info b:hover { text-decoration: underline }
.suggestions-list li .dismiss { position: absolute }
.suggestions-list li .dismiss { color: #aaa; margin: 0 0 0 10px; top: 0; right: 0 }
.suggestions-list li .dismiss:hover { color: #333; text-decoration: none }


.suggest-overlay { position: absolute; z-index: 99; width: auto; background: #fff; border: 1px solid #c5c7d2;
    -moz-border-radius : 3px;
    -webkit-border-radius : 3px;
    border-radius: 3px
}
.suggest-overlay .bd { min-width: 220px; line-height: 1; background: #fafafa; color: #b3b3b3; padding: 5px;
    -moz-border-radius : 3px;
    -webkit-border-radius : 3px;
    border-radius: 3px
}
.suggest-overlay ul { color: #999; padding: 3px 0; min-width: 214px }
.suggest-overlay li { cursor: pointer; padding: 3px 7px }
.suggest-overlay li b { font-weight: bold }
.suggest-overlay li .username { color: #333 }
.suggest-overlay img { margin-right: 5px; width: 20px; height: 20px; vertical-align: middle }
.suggest-overlay .on { background: #e9f0f8 }

.mentioned-highlighter { font: 14px/20px "Helvetica Neue",Helvetica,Arial,sans-serif; position: absolute; left: 4px; top: 4px; font-size: 14px; height: 60px; width: 98.5%; overflow: hidden; background: #fff; white-space: pre-wrap; word-wrap: break-word; color: transparent }
.mentioned-highlighter b { font-weight: normal; background-color: #d2e1f3; color: transparent;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px
}

        .movie-share-dialog .bn-flat input {
    font-size: 14px;
}
.movie-share-dialog {
    z-index: 100;
}
.movie-share-dialog .form-ft-inner{
    text-align: right;
}
.movie-share-dialog div.bd {
    padding: 0;
}

.movie-share .form-bd .input-area {
    position: relative;
    margin: 15px;
    height: 91px;
    zoom: 1;
}

.movie-share-no-media .form-bd {
    height: 140px;
}

.movie-share-dialog .share-text {
    height: 85px;
    position: absolute;
    z-index: 9;
    background: transparent;
    font: 14px/18px "Helvetica Neue",Helvetica,Arial,sans-serif;
    width: 98%;
    -webkit-border-radius: 4px 4px 4px 4px;
    border-radius: 4px 4px 4px 4px;
}

.movie-share-dialog .mentioned-highlighter {
    width: 483px;
    padding: 3px 4px 4px;
    color: white;
    position: absolute;
    top:0;
    left:0;
    z-index: 0;
}

.movie-share-dialog .mentioned-highlighter code {
    color: #D2E1F3;
    background: #D2E1F3;
    border-radius: 2px;
    padding-right: 1px;
    display: inline-block;
    font: 14px/18px "Helvetica Neue",Helvetica,Arial,sans-serif;
}


.movie-share .form-ft {
    background: #e9eef2;
    height: 25px;
    padding-top: 10px;
    padding-bottom: 10px;
}

.movie-share .form-ft-inner {
    height: 25px;
    padding-left: 15px;
    padding-right: 15px;
}

.movie-share-dialog .dialog-only-text {
    text-align: center;
    font-size: 14px;
    line-height: 1.5;
    padding-top: 30px;
    padding-bottom: 30px;
    color: #0c7823;
}

.movie-share-dialog .ll {
    float: left;
    display: inline;
}
.movie-share-dialog .share-label {
    width: auto;
    display: inline;
    float: none;
}

.movie-share-dialog .leading-label {
    _vertical-align: -2px;
}
.movie-share-dialog .media {
    float: left;
    margin-right: 10px;
    max-width: 100px;
    max-height: 100px;
    *width: 100px;
}
.movie-share-dialog .info-area{
    overflow: hidden;
    zoom: 1;
    margin: 0 15px 15px;
    height: 100px;
}
.movie-share-dialog .info-area strong{
    font-weight: bold;
}
.movie-share-dialog .info-area p{
    margin: 3px 0;
    color: #999;
}

.movie-share-dialog #sync-setting {
    _vertical-align: -5px;
    margin-left: 10px;
}

.movie-share-dialog .info-area .server-error {
    position: absolute;
    bottom: 45px;
    right: 15px;
    color: red;
}

.movie-share-dialog .avail-num-indicator {
    color: #aaa;
    font-weight: 800;
    padding-right: 3px;
}

.movie-share-dialog .bottom-setting {
    width: 432px;
}
.movie-share-dialog .input-checkbox {
    vertical-align: -2px;
    _vertical-align: -1px;
}

.movie-share-dialog #sync-setting img {
    _vertical-align: 2px;
}



.suggest-overlay {
    z-index: 2000;
}

.movie-bar {
    position: relative;
    margin-top: 10px;
}

.movie-bar-fav {
    position: absolute;
    top: 0;
    right: 0;
}

    </style>
    <script src="https://img3.doubanio.com/f/shire/a40c5220b3f40ce737b366c0030ecf810b37bfea/js/lib/mustache.js" type="text/javascript"></script>
    <script src="https://img3.doubanio.com/f/shire/1d985568f3cc434b145983919d9954e2ca627e9c/js/lib/textarea-mention.js" type="text/javascript"></script>
    <script src="https://img3.doubanio.com/f/movie/230795bbf6a9a7700cc6f395067493d5dcc572ad/js/movie/share.js" type="text/javascript"></script>

<div class="rec-sec">
<span class="rec">
    <script id="movie-share" type="text/x-html-snippet">
        
    <form class="movie-share" action="/j/share" method="POST">
        <div class="clearfix form-bd">
            <div class="input-area">
                <textarea name="text" class="share-text" cols="72" data-mention-api="https://api.douban.com/shuo/in/complete?alt=xd&amp;callback=?"></textarea>
                <input type="hidden" name="target-id" value="30257787">
                <input type="hidden" name="target-type" value="0">
                <input type="hidden" name="title" value="一秒钟‎ (2020)">
                <input type="hidden" name="desc" value="导演 张艺谋 主演 张译 / 刘浩存 / 中国大陆 / 7.9分(72860评价)">
                <input type="hidden" name="redir" value=""/>
                <div class="mentioned-highlighter"></div>
            </div>

            <div class="info-area">
                    <img class="media" src="https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2626693662.webp" />
                <strong>一秒钟‎ (2020)</strong>
                <p>导演 张艺谋 主演 张译 / 刘浩存 / 中国大陆 / 7.9分(72860评价)</p>
                <p class="error server-error">&nbsp;</p>
            </div>
        </div>
        <div class="form-ft">
            <div class="form-ft-inner">
                



                <span class="avail-num-indicator">140</span>
                <span class="bn-flat">
                    <input type="submit" value="推荐" />
                </span>
            </div>
        </div>
    </form>
    
    <div id="suggest-mention-tmpl" style="display:none;">
        <ul>
            {{#users}}
            <li id="{{uid}}">
              <img src="{{avatar}}">{{{username}}}&nbsp;<span>({{{uid}}})</span>
            </li>
            {{/users}}
        </ul>
    </div>


    </script>

        
        <a href="/accounts/register?reason=recommend"  class="j a_show_login lnk-sharing" 
            share-id="30257787" 
            data-mode="plain" data-name="一秒钟‎ (2020)" 
            data-type="movie" data-desc="导演 张艺谋 主演 张译 / 刘浩存 / 中国大陆 / 7.9分(72860评价)" 
            data-href="https://movie.douban.com/subject/30257787/" data-image="https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2626693662.webp" 
            data-properties="{}" 
            data-redir="" data-text="" 
            data-apikey="" data-curl="" 
            data-count="10" data-object_kind="1002" 
            data-object_id="30257787" data-target_type="rec" 
            data-target_action="0" 
            data-action_props="{&#34;subject_url&#34;:&#34;https:\/\/movie.douban.com\/subject\/30257787\/&#34;,&#34;subject_title&#34;:&#34;一秒钟‎ (2020)&#34;}">推荐</a>
</span>


</div>






            <script type="text/javascript">
                $(function() {
                    $('.collect_btn', '#interest_sect_level').each(function() {
                        Douban.init_collect_btn(this);
                    });
                    $('html').delegate(".indent .rec-sec .lnk-sharing", "click", function() {
                        moreurl(this, {
                            from : 'mv_sbj_db_share'
                        });
                    });
                });
            </script>
        </div>
            


    <div id="collect_form_30257787"></div>


        



<div class="related-info" style="margin-bottom:-10px;">
    <a name="intro"></a>
    
        
            
            
    <h2>
        <i class="">一秒钟的剧情简介</i>
              · · · · · ·
    </h2>

            <div class="indent" id="link-report">
                    
                        <span property="v:summary" class="">
                                　　影片讲述了没赶上电影场次的张九声与刘闺女因一场电影结下了不解之缘的故事。故事灵感来源于张艺谋导演早期经历，是其一贯对文化展开追忆和寻根的风格。
                        </span>
                        

            </div>
</div>


    <div id="dale_movie_subject_banner_after_intro"></div>

    








<div id="celebrities" class="celebrities related-celebrities">

  
    <h2>
        <i class="">一秒钟的演职员</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="/subject/30257787/celebrities">全部 25</a>
            )
            </span>
    </h2>


  <ul class="celebrities-list from-subject __oneline">
        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1054398/" title="张艺谋 Yimou Zhang" class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p568.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1054398/" title="张艺谋 Yimou Zhang" class="name">张艺谋</a></span>

      <span class="role" title="导演">导演</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1274761/" title="张译 Yi Zhang" class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1489386626.47.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1274761/" title="张译 Yi Zhang" class="name">张译</a></span>

      <span class="role" title="饰 张九声">饰 张九声</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1410267/" title="刘浩存 Haocun Liu" class="">
      <div class="avatar" style="background-image: url(https://img2.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1596685488.53.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1410267/" title="刘浩存 Haocun Liu" class="name">刘浩存</a></span>

      <span class="role" title="饰 刘闺女">饰 刘闺女</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1051526/" title="范伟 Wei Fan" class="">
      <div class="avatar" style="background-image: url(https://img9.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1511362485.44.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1051526/" title="范伟 Wei Fan" class="name">范伟</a></span>

      <span class="role" title="饰 范电影">饰 范电影</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1313809/" title="余皑磊 Ailei Yu" class="">
      <div class="avatar" style="background-image: url(https://img2.doubanio.com/view/celebrity/s_ratio_celebrity/public/p51592.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1313809/" title="余皑磊 Ailei Yu" class="name">余皑磊</a></span>

      <span class="role" title="饰 崔干事">饰 崔干事</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1410678/" title="张邵勃  Shaobo Zhang" class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1577343418.59.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1410678/" title="张邵勃  Shaobo Zhang" class="name">张邵勃 </a></span>

      <span class="role" title="饰 刘弟弟">饰 刘弟弟</span>

    </div>
  </li>


  </ul>
</div>


    


<link rel="stylesheet" href="https://img3.doubanio.com/f/verify/16c7e943aee3b1dc6d65f600fcc0f6d62db7dfb4/entry_creator/dist/author_subject/style.css">
<div id="author_subject" class="author-wrapper">
    <div class="loading"></div>
</div>
<script type="text/javascript">
    var answerObj = {
      ISALL: 'False',
      TYPE: 'movie',
      SUBJECT_ID: '30257787',
      USER_ID: 'None'
    }
</script>
<script type="text/javascript" src="https://img3.doubanio.com/f/movie/61252f2f9b35f08b37f69d17dfe48310dd295347/js/movie/lib/react/15.4/bundle.js"></script>
<script type="text/javascript" src="https://img3.doubanio.com/f/verify/ac140ef86262b845d2be7b859e352d8196f3f6d4/entry_creator/dist/author_subject/index.js"></script>


    









    
    <div id="related-pic" class="related-pic">
        
    
    
    <h2>
        <i class="">一秒钟的视频和图片</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30257787/trailer#trailer">预告片10</a>&nbsp;|&nbsp;<a href="/video/create?subject_id=30257787">添加视频评论</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/30257787/all_photos">图片178</a>&nbsp;·&nbsp;<a href="https://movie.douban.com/subject/30257787/mupload">添加</a>
            )
            </span>
    </h2>


        <ul class="related-pic-bd  ">
                <li class="label-trailer">
                    <a class="related-pic-video" href="https://movie.douban.com/trailer/268884/#content" title="预告片" style="background-image:url(https://img1.doubanio.com/img/trailer/medium/2626803168.jpg?1606212031)">
                    </a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2546939638/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2546939638.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2546778385/"><img src="https://img9.doubanio.com/view/photo/sqxs/public/p2546778385.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2545678267/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2545678267.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2624599628/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2624599628.webp" alt="图片" /></a>
                </li>
        </ul>
    </div>




    



<style type="text/css">
.award li { display: inline; margin-right: 5px }
.awards { margin-bottom: 20px }
.awards h2 { background: none; color: #000; font-size: 14px; padding-bottom: 5px; margin-bottom: 8px; border-bottom: 1px dashed #dddddd }
.awards .year { color: #666666; margin-left: -5px }
.mod { margin-bottom: 25px }
.mod .hd { margin-bottom: 10px }
.mod .hd h2 {margin:24px 0 3px 0}
</style>


<div class="mod">
    <div class="hd">
        
    <h2>
        <i class="">一秒钟的获奖情况</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30257787/awards/">全部</a>
            )
            </span>
    </h2>

    </div>
        
        <ul class="award">
            <li>
                <a href="https://movie.douban.com/awards/berlinale/69/">第69届柏林国际电影节</a>
            </li>
            <li>金熊奖(提名)</li>
            <li><a href='https://movie.douban.com/celebrity/1054398/' target='_blank'>张艺谋</a></li>
        </ul>
</div>

    








    <div id="recommendations" class="">
        
        
    <h2>
        <i class="">喜欢这部电影的人也喜欢</i>
              · · · · · ·
    </h2>

        
    
    <div class="recommendations-bd">
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/30204413/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2623499790.webp" alt="气球" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/30204413/?from=subject-page" class="" >气球</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/10604895/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2514078368.webp" alt="英格力士" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/10604895/?from=subject-page" class="" >英格力士</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/30355901/?from=subject-page" >
                    <img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2564136204.webp" alt="鸟鸣嘤嘤" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/30355901/?from=subject-page" class="" >鸟鸣嘤嘤</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26954859/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2568264011.webp" alt="兰心大剧院" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26954859/?from=subject-page" class="" >兰心大剧院</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/27172891/?from=subject-page" >
                    <img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2511811355.webp" alt="大象席地而坐" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/27172891/?from=subject-page" class="" >大象席地而坐</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1950330/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2555541430.webp" alt="罗马" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1950330/?from=subject-page" class="" >罗马</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1292365/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2513253791.webp" alt="活着" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1292365/?from=subject-page" class="" >活着</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1292434/?from=subject-page" >
                    <img src="https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2567845803.webp" alt="一一" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1292434/?from=subject-page" class="" >一一</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/2609258/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2620161520.webp" alt="芝加哥七君子审判" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/2609258/?from=subject-page" class="" >芝加哥七君子审判</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26752088/?from=subject-page" >
                    <img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2561305376.webp" alt="我不是药神" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26752088/?from=subject-page" class="" >我不是药神</a>
            </dd>
        </dl>
    </div>

    </div>



        


<script type="text/x-handlebar-tmpl" id="comment-tmpl">
    <div class="dummy-fold">
        {{#each comments}}
        <div class="comment-item" data-cid="id">
            <div class="comment">
                <h3>
                    <span class="comment-vote">
                            <span class="votes">{{votes}}</span>
                        <input value="{{id}}" type="hidden"/>
                        <a href="javascript:;" class="j {{#if ../if_logined}}vote-comment{{else}}a_show_login{{/if}}">有用</a>
                    </span>
                    <span class="comment-info">
                        <a href="{{user.path}}" class="">{{user.name}}</a>
                        {{#if rating}}
                        <span class="allstar{{rating}}0 rating" title="{{rating_word}}"></span>
                        {{/if}}
                        <span>
                            {{time}}
                        </span>
                        <p> {{content_tmpl content}} </p>
                    </span>
            </div>
        </div>
        {{/each}}
    </div>
</script>












    

    <div id="comments-section">
        <div class="mod-hd">
            
            
        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">
            <span>我要写短评</span>
        </a>

            
    <h2>
        <i class="">一秒钟的短评</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30257787/comments?status=P">全部 29037 条</a>
            )
            </span>
    </h2>

            
        </div>
        <div class="mod-bd">
                

    <div class="tab-hd">
        <a id="hot-comments-tab" href="comments" data-id="hot" class="on">热门</a>&nbsp;/&nbsp;
            <a id="new-comments-tab" href="comments?sort=time" data-id="new">最新</a>&nbsp;/&nbsp;
        <a id="following-comments-tab" href="comments?sort=follows" data-id="following"  class="j a_show_login">好友</a>
    </div>

    <div class="tab-bd">
        <div id="hot-comments" class="tab">
            
    
        
        <div class="comment-item " data-cid="2619065710">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                    <span class="votes vote-count">3282</span>

                    <input value="2619065710" type="hidden"/>
                    <a href="javascript:;" data-id="2619065710"
                        class="j a_show_login" 
                        onclick="">有用</a>
                    
                <!-- 删除短评 -->
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/148490548/" class="">3号厅检票小哥</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2020-11-27 15:06:22">
                    2020-11-27
                </span>
            </span>
        </h3>
        <p class=" comment-content">
            
                <span class="short">这根本就不是什么给电影的情书，整个电影没有人是爱电影的。

范伟放电影，但在编剧写法上其实是丑角。他不爱电影，他只是迷恋放电影给他带来的“虚假权力”，所以他会举bao张译，他又是受害者，所以会给张译胶片。

张译这个角色全程“追电影”，但也只是因为里面有他死去的女儿，他是被劳改的知识分子，他一定是和《英雄儿女》的氛围完全背离的，他也不认同简报里所谓的争先进。

ge命群众呢？他们是爱电影吗？不是的...</span>
                <span class="full">这根本就不是什么给电影的情书，整个电影没有人是爱电影的。

范伟放电影，但在编剧写法上其实是丑角。他不爱电影，他只是迷恋放电影给他带来的“虚假权力”，所以他会举bao张译，他又是受害者，所以会给张译胶片。

张译这个角色全程“追电影”，但也只是因为里面有他死去的女儿，他是被劳改的知识分子，他一定是和《英雄儿女》的氛围完全背离的，他也不认同简报里所谓的争先进。

ge命群众呢？他们是爱电影吗？不是的，他们只是贫瘠，他们负责倒映那个十年导致的精神贫瘠和集体狂热（合唱赞歌）。

所以整个电影里没有人爱电影，都是受害者。这不是什么狗屁情书，这是恐怖片。

整部电影也被删成了牛车后面的胶卷，都是窟窿。</span>
                <span class="expand">(<a href="javascript:;">展开</a>)</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item " data-cid="2618863725">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                    <span class="votes vote-count">2602</span>

                    <input value="2618863725" type="hidden"/>
                    <a href="javascript:;" data-id="2618863725"
                        class="j a_show_login" 
                        onclick="">有用</a>
                    
                <!-- 删除短评 -->
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/179461893/" class="">Le_Petit_P</a>
                    <span>看过</span>
                    <span class="allstar50 rating" title="力荐"></span>
                <span class="comment-time " title="2020-11-27 11:55:37">
                    2020-11-27
                </span>
            </span>
        </h3>
        <p class=" comment-content">
            
                <span class="short">若批评不自由，则赞美无意义。这么多年过去了，何必呢……</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item " data-cid="1614923462">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                    <span class="votes vote-count">1301</span>

                    <input value="1614923462" type="hidden"/>
                    <a href="javascript:;" data-id="1614923462"
                        class="j a_show_login" 
                        onclick="">有用</a>
                    
                <!-- 删除短评 -->
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/48369193/" class="">二十二岛主</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2020-11-23 20:47:55">
                    2020-11-23
                </span>
            </span>
        </h3>
        <p class=" comment-content">
            
                <span class="short">这部电影我等了2年，而张艺谋为了拍它等了50年。之所以这么想要把这段故事带给观众，除了张艺谋本人对于露天电影和胶片的情结之外，可能还在于，现在确实不会有导演再拍这样的电影了。这在一些人看来未免有些古板和过时，毕竟很多年轻人别说胶片电影了，就连露天电影都没有看过，短视频、流媒体和移动设备的崛起已经大大改变了年轻观众的观影习惯。当电影院都变得可有可无的时候，又有几个人会在意电影放映的过去与沿革呢？但没...</span>
                <span class="full">这部电影我等了2年，而张艺谋为了拍它等了50年。之所以这么想要把这段故事带给观众，除了张艺谋本人对于露天电影和胶片的情结之外，可能还在于，现在确实不会有导演再拍这样的电影了。这在一些人看来未免有些古板和过时，毕竟很多年轻人别说胶片电影了，就连露天电影都没有看过，短视频、流媒体和移动设备的崛起已经大大改变了年轻观众的观影习惯。当电影院都变得可有可无的时候，又有几个人会在意电影放映的过去与沿革呢？但没人拍不代表不该拍，事实上在我看来，尤其是在今年这样一个特殊的节点，《一秒钟》的出现是更具价值和意义的。现在看到70岁的张艺谋，在“从心所欲不逾矩”的年龄回归初心，拍了这样一部关于光影记忆的电影，我突然也有了勇气，希望在我70岁那年的时候，也能圆了自己少年时的梦，拍一部关于“电影”的电影。</span>
                <span class="expand">(<a href="javascript:;">展开</a>)</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item " data-cid="1461768640">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                    <span class="votes vote-count">779</span>

                    <input value="1461768640" type="hidden"/>
                    <a href="javascript:;" data-id="1461768640"
                        class="j a_show_login" 
                        onclick="">有用</a>
                    
                <!-- 删除短评 -->
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/yalindongdong/" class="">方枪枪</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2020-11-23 19:28:10">
                    2020-11-23
                </span>
            </span>
        </h3>
        <p class=" comment-content">
            
                <span class="short">《一秒钟》面子上讲的是对胶片电影往事的缅怀，但骨子里却是更直击观众内心的时代悲剧与情感表达——张译饰演的张九声与刘闺女两个陌生人因一盘胶片引发各种误会再到最后和解。故事里有太多让人五味杂陈的情节：张九声对见女儿的渴求，虽电影里没表明，但可猜测：他女儿早已去世，所以为了女儿《新闻简报》里出现的那一秒钟镜头他泪流满面，重复观看乐此不疲，甚至为那截被人丢弃的胶片歇斯底里；范伟饰演的范电影这一角色更复杂多...</span>
                <span class="full">《一秒钟》面子上讲的是对胶片电影往事的缅怀，但骨子里却是更直击观众内心的时代悲剧与情感表达——张译饰演的张九声与刘闺女两个陌生人因一盘胶片引发各种误会再到最后和解。故事里有太多让人五味杂陈的情节：张九声对见女儿的渴求，虽电影里没表明，但可猜测：他女儿早已去世，所以为了女儿《新闻简报》里出现的那一秒钟镜头他泪流满面，重复观看乐此不疲，甚至为那截被人丢弃的胶片歇斯底里；范伟饰演的范电影这一角色更复杂多面，一方面为保住自己工作举报张九声，同时也冒很大风险将张九声女儿那段胶片偷偷剪下给张九声.....电影在演员精湛演技下更触动人心，这就是大导演的功底，返璞归真，化繁为简。时代的浪潮下，你我无非只是一根野草，一粒沙子，被时代裹挟，被风沙掩埋！</span>
                <span class="expand">(<a href="javascript:;">展开</a>)</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item " data-cid="1460748471">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                    <span class="votes vote-count">625</span>

                    <input value="1460748471" type="hidden"/>
                    <a href="javascript:;" data-id="1460748471"
                        class="j a_show_login" 
                        onclick="">有用</a>
                    
                <!-- 删除短评 -->
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/42805786/" class="">⠀</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2020-11-26 22:16:29">
                    2020-11-26
                </span>
            </span>
        </h3>
        <p class=" comment-content">
            
                <span class="short">绝对是2020年到目前为止最值得看的国产片，没有之一，国师拍这种题材太驾轻就熟了，仿佛看到了九十年代的张艺谋。范伟的部分很搞笑，除此之外，从头到尾都是很干净很纯粹的叙事与展现，只剩下个别镜头和台词暗示着观众。结合阿德映后提到那个暗黑的彩蛋(女儿抗面袋时被车撞死了)，估计仅存的一点点反思与批判的成分都被提前咔嚓掉了，但还好没有影响到整片的气质。ps.超前观影还见到张导、张译、刘浩存了，张译好亲和一人...</span>
                <span class="full">绝对是2020年到目前为止最值得看的国产片，没有之一，国师拍这种题材太驾轻就熟了，仿佛看到了九十年代的张艺谋。范伟的部分很搞笑，除此之外，从头到尾都是很干净很纯粹的叙事与展现，只剩下个别镜头和台词暗示着观众。结合阿德映后提到那个暗黑的彩蛋(女儿抗面袋时被车撞死了)，估计仅存的一点点反思与批判的成分都被提前咔嚓掉了，但还好没有影响到整片的气质。ps.超前观影还见到张导、张译、刘浩存了，张译好亲和一人哈哈哈。</span>
                <span class="expand">(<a href="javascript:;">展开</a>)</span>
        </p>
    </div>

        </div>



    


                
                &gt; <a href="comments?sort=new_score&status=P" >
                    更多短评
                        29037条
                </a>
        </div>
        <div id="new-comments" class="tab">
            <div id="normal">
            </div>
            <div class="fold-hd hide">
                <a class="qa" href="/help/opinion#t2-q0" target="_blank">为什么被折叠？</a>
                <a class="btn-unfold" href="#">有一些短评被折叠了</a>
                <div class="qa-tip">
                    评论被折叠，是因为发布这条评论的帐号行为异常。评论仍可以被展开阅读，对发布人的账号不造成其他影响。如果认为有问题，可以<a href="https://help.douban.com/help/ask?category=movie">联系</a>豆瓣电影。
                </div>
            </div>
            <div class="fold-bd">
            </div>
            <span id="total-num"></span>
        </div>
        <div id="following-comments" class="tab">
            
    


        <div class="comment-item">
            你关注的人还没写过短评
        </div>

    


        </div>
    </div>


            
            
        </div>
    </div>



        

<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/1b80cf3bee515ce2.css">

<section class="topics mod">
    <header>
        <h2>
            一秒钟的话题 · · · · · ·
            <span class="pl">( <span class="gallery_topics">全部 <span id="topic-count"></span> 条</span> )</span>
        </h2>
    </header>

    




<section class="subject-topics">
    <div class="topic-guide" id="topic-guide">
        <img class="ic_question" src="//img3.doubanio.com/f/ithildin/b1a3edea3d04805f899e9d77c0bfc0d158df10d5/pics/export/icon_question.png">
        <div class="tip_content">
            <div class="tip_title">什么是话题</div>
            <div class="tip_desc">
                <div>无论是一部作品、一个人，还是一件事，都往往可以衍生出许多不同的话题。将这些话题细分出来，分别进行讨论，会有更多收获。</div>
            </div>
        </div>
        <img class="ic_guide" src="//img3.doubanio.com/f/ithildin/529f46d86bc08f55cd0b1843d0492242ebbd22de/pics/export/icon_guide_arrow.png">
        <img class="ic_close" id="topic-guide-close" src="//img3.doubanio.com/f/ithildin/2eb4ad488cb0854644b23f20b6fa312404429589/pics/export/close@3x.png">
    </div>

    <div id="topic-items"></div>

    <script>
        window.subject_id = 30257787;
        window.join_label_text = '写影评参与';

        window.topic_display_count = 4;
        window.topic_item_display_count = 1;
        window.no_content_fun_call_name = "no_topic";

        window.guideNode = document.getElementById('topic-guide');
        window.guideNodeClose = document.getElementById('topic-guide-close');
    </script>
    
        <link rel="stylesheet" href="https://img3.doubanio.com/f/ithildin/f731c9ea474da58c516290b3a6b1dd1237c07c5e/css/export/subject_topics.css">
        <script src="https://img3.doubanio.com/f/ithildin/d3590fc6ac47b33c804037a1aa7eec49075428c8/js/export/moment-with-locales-only-zh.js"></script>
        <script src="https://img3.doubanio.com/f/ithildin/c600fdbe69e3ffa5a3919c81ae8c8b4140e99a3e/js/export/subject_topics.js"></script>

</section>

    <script>
        function no_topic(){
            $('#content .topics').remove();
        }
    </script>
</section>
    <section id="reviews-wrapper" class="reviews mod movie-content">
        <header>
            
                <a href="new_review" rel="nofollow" class="create-review comment_btn "
                    data-isverify="False"
                    data-verify-url="https://www.douban.com/accounts/phone/verify?redir=http://movie.douban.com/subject/30257787/new_review">
                    <span>我要写影评</span>
                </a>
            <h2>
                一秒钟的影评 · · · · · ·
                <span class="pl">( <a href="reviews">全部 912 条</a> )</span>
            </h2>
        </header>

        

<style>
#gallery-topics-selection {
  position: fixed;
  width: 595px;
  padding: 40px 40px 33px 40px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
  top: 50%;
  left: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  z-index: 9999;
}
#gallery-topics-selection h1 {
  font-size: 18px;
  color: #007722;
  margin-bottom: 36px;
  padding: 0;
  line-height: 28px;
  font-weight: normal;
}
#gallery-topics-selection .gl_topics {
  border-bottom: 1px solid #dfdfdf;
  max-height: 298px;
  overflow-y: scroll;
}
#gallery-topics-selection .topic {
  margin-bottom: 24px;
}
#gallery-topics-selection .topic_name {
  font-size: 15px;
  color: #333;
  margin: 0;
  line-height: inherit;
}
#gallery-topics-selection .topic_meta {
  font-size: 13px;
  color: #999;
}
#gallery-topics-selection .topics_skip {
  display: block;
  cursor: pointer;
  font-size: 16px;
  color: #3377AA;
  text-align: center;
  margin-top: 33px;
}
#gallery-topics-selection .topics_skip:hover {
  background: transparent;
}
#gallery-topics-selection .close_selection {
  position: absolute;
  width: 30px;
  height: 20px;
  top: 46px;
  right: 40px;
  background: #fff;
  color: #999;
  text-align: right;
}
#gallery-topics-selection .close_selection:hover{
  background: #fff;
  color: #999;
}
</style>




            <div class="review_filter">
                <a href="javascript:;;" class="cur" data-sort="">热门</a> &#047;
                <a href="javascript:;;" data-sort="time">最新</a> &#047;
                <a href="javascript:;;" data-sort="follow">好友</a>
            </div>
            <script>
                var cur_sort = '';
                $('#reviews-wrapper .review_filter a').on('click', function () {
                    var sort = $(this).data('sort');
                    if(sort === cur_sort) return;

                    if(sort === 'follow' && true){
                        window.location.href = '//www.douban.com/accounts/login?source=movie';
                        return;
                    }

                    if($('#reviews-wrapper .review_filter').data('doing')) return;
                    $('#reviews-wrapper .review_filter').data('doing', true);

                    cur_sort = sort;

                    $('#reviews-wrapper .review_filter a').removeClass('cur');
                    $(this).addClass('cur');

                    $.getJSON('reviews', { sort: sort }, function(res){
                        $('#reviews-wrapper .review-list').remove();
                        $('#reviews-wrapper [href="reviews?sort=follow"]').parent().remove();
                        $('#reviews-wrapper .review_filter').after(res.html);
                        $('#reviews-wrapper .review_filter').data('doing', false);
                        $('#reviews-wrapper .review_filter').removeData('doing');

                        if(res.count === 0){
                            $('#reviews-wrapper .review-list').html('<span class="no-review">你关注的人还没写过长评</span>');
                        }
                    });
                });
            </script>


            



<div class="review-list  ">
        
    

        
    
    <div data-cid="13013938">
        <div class="main review-item" id="13013938">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/m19900714/" class="avator">
            <img width="24" height="24" src="https://img9.doubanio.com/icon/u19615070-16.jpg">
        </a>

        <a href="https://www.douban.com/people/m19900714/" class="name">闵思嘉</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 21:35:42</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13013938/">《一秒钟》，究竟删了几秒钟？</a></h2>

                <div id="review_13013938_short" class="review-short" data-rid="13013938">
                    <div class="short-content">

                        导语： 因为工作原因，我在去年冬天看过《一秒钟》，大概一月份，影片入围柏林前后。看片地点在朝阳区的一个小酒吧里，地方挺小，挤了几十个人。 我觉得“那个”版本的《一秒钟》很好，甚至可以说是张艺谋在进入新千年以后，最好的作品。 “这个”版本，长度差不多，气质太不同...

                        &nbsp;(<a href="javascript:;" id="toggle-13013938-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13013938_full" class="hidden">
                    <div id="review_13013938_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13013938" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13013938">
                                3315
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13013938" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13013938">
                                20
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13013938/#comments" class="reply ">435回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13014612">
        <div class="main review-item" id="13014612">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/mr_tree/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u60689518-41.jpg">
        </a>

        <a href="https://www.douban.com/people/mr_tree/" class="name">凹凸</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-28" class="main-meta">2020-11-28 00:37:56</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13014612/">献给电影的情书，还是控诉历史的血书？</a></h2>

                <div id="review_13014612_short" class="review-short" data-rid="13014612">
                    <div class="short-content">

                        未经授权，严禁转载！！！ 因为众所周知的“技术原因”，《一秒钟》几经波折才“顺利”上映，而片中的情节和影片本身的遭遇，也在某种程度上“不谋而合”。 乍看之下，《一秒钟》是个温情脉脉的故事：劳改犯张九声逃狱，只为了在电影放映前的新闻简报中看“一秒钟”的女儿；孤...

                        &nbsp;(<a href="javascript:;" id="toggle-13014612-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13014612_full" class="hidden">
                    <div id="review_13014612_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13014612" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13014612">
                                402
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13014612" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13014612">
                                19
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13014612/#comments" class="reply ">174回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13014054">
        <div class="main review-item" id="13014054">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/1441177/" class="avator">
            <img width="24" height="24" src="https://img9.doubanio.com/icon/u1441177-25.jpg">
        </a>

        <a href="https://www.douban.com/people/1441177/" class="name">临素光</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 22:07:31</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13014054/">电影狂欢背后的贫瘠：《一秒钟》的删改情节与真实背景</a></h2>

                <div id="review_13014054_short" class="review-short" data-rid="13014054">
                    <div class="short-content">

                        当我在电影院看到《一秒钟》的结局时，几乎想冷笑。 准确来说，这是原版没有、后来补拍的结局——在主线故事结束的两年后，张九声受益于1978年的平反政策，从劳改农场释放，排着队领取崭新的中山装和脸盆铺盖，高高兴兴开始新生活。 相较于那些在20年间蒙冤入狱、累死饿死的劳...

                        &nbsp;(<a href="javascript:;" id="toggle-13014054-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13014054_full" class="hidden">
                    <div id="review_13014054_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13014054" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13014054">
                                391
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13014054" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13014054">
                                9
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13014054/#comments" class="reply ">63回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13023899">
        <div class="main review-item" id="13023899">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/MovieL/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1128221-98.jpg">
        </a>

        <a href="https://www.douban.com/people/MovieL/" class="name">木卫二</a>


        <span content="2020-12-01" class="main-meta">2020-12-01 12:14:06</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13023899/">千审万查一秒钟</a></h2>

                <div id="review_13023899_short" class="review-short" data-rid="13023899">
                    <div class="short-content">

                        看完《一秒钟》，我头一次，冒出给张艺谋作品打五星的冲动。其实，我比任何人都清楚，《一秒钟》有一串问题，是会被扣分的，与完美无瑕、回味无穷的五星电影无法划等号。 我也同样说过：只要一部中国电影出现了“技术原因”，那么，我出于同情分，随手加上一颗星，也是很常见的...

                        &nbsp;(<a href="javascript:;" id="toggle-13023899-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13023899_full" class="hidden">
                    <div id="review_13023899_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13023899" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13023899">
                                119
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13023899" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13023899">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13023899/#comments" class="reply ">7回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13014136">
        <div class="main review-item" id="13014136">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/ruo1996/" class="avator">
            <img width="24" height="24" src="https://img9.doubanio.com/icon/u52623934-36.jpg">
        </a>

        <a href="https://www.douban.com/people/ruo1996/" class="name">徐若风</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 22:29:32</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13014136/">《一秒钟》的“无言之言”与表演的身体感</a></h2>

                <div id="review_13014136_short" class="review-short" data-rid="13014136">
                    <div class="short-content">

                        想必许多观众会同我一样，对最终呈现的《一秒钟》感到些许意外：它竟然是一部如此“简单”的电影，主线就聚焦于为了看逝世女儿的一秒钟影像而选择逃亡的张九声（张译饰）和被迫偷胶片做灯罩的刘闺女（刘浩存饰）两个人身上。简单的三天两夜的故事情节、简单的萍水相逢的人物关...

                        &nbsp;(<a href="javascript:;" id="toggle-13014136-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13014136_full" class="hidden">
                    <div id="review_13014136_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13014136" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13014136">
                                519
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13014136" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13014136">
                                11
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13014136/#comments" class="reply ">22回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13012153">
        <div class="main review-item" id="13012153">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/48369193/" class="avator">
            <img width="24" height="24" src="https://img9.doubanio.com/icon/u48369193-16.jpg">
        </a>

        <a href="https://www.douban.com/people/48369193/" class="name">二十二岛主</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 09:54:51</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13012153/">不要害怕电影！那只是一个人和一群人的执念与记忆</a></h2>

                <div id="review_13012153_short" class="review-short" data-rid="13012153">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        本文首发于公众号：电影岛赏（j_movie），欢迎关注。 先聊聊我关于《一秒钟》的一些记忆吧。2019年我去德国报道柏林电影节，那一年我最期待三部要放映的华语片：《一秒钟》《地久天长》和《少年的你》，没想到临时撤了两部。尤其是《一秒钟》的退出，就在放映的前一两天。当时...

                        &nbsp;(<a href="javascript:;" id="toggle-13012153-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13012153_full" class="hidden">
                    <div id="review_13012153_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13012153" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13012153">
                                444
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13012153" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13012153">
                                22
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13012153/#comments" class="reply ">49回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13017915">
        <div class="main review-item" id="13017915">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/philaro2009/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u26995231-39.jpg">
        </a>

        <a href="https://www.douban.com/people/philaro2009/" class="name">子戈</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-29" class="main-meta">2020-11-29 10:25:22</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13017915/">被删掉的和不能说的</a></h2>

                <div id="review_13017915_short" class="review-short" data-rid="13017915">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        01 看过《一秒钟》的遭遇，你就明白了当下的电影环境。 以前，只要拿到龙标，就意味着影片可以到国际上参赛，也可以顺利上映。 但如今还是这样吗？ 不是了。 尽管早在去年初，《一秒钟》就拿到了龙标。 可在随后的柏林电影节上，却因“技术原因”被紧急召回，也由此掀起了一阵...

                        &nbsp;(<a href="javascript:;" id="toggle-13017915-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13017915_full" class="hidden">
                    <div id="review_13017915_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13017915" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13017915">
                                121
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13017915" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13017915">
                                2
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13017915/#comments" class="reply ">18回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13011950">
        <div class="main review-item" id="13011950">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/AngelinaAnne/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u2342914-17.jpg">
        </a>

        <a href="https://www.douban.com/people/AngelinaAnne/" class="name">小A</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 07:42:29</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13011950/">历史群像里的一粒尘埃</a></h2>

                <div id="review_13011950_short" class="review-short" data-rid="13011950">
                    <div class="short-content">

                        《一秒钟》几乎可以确定是张艺谋职业生涯最好的剧本之一。影片从片名到情节安排了一系列微言大义的比喻，这些场景太耀眼，在脑海中挥之不去。“一秒钟”既是女儿留在人世最后影像的时长，也概括了动荡年代的一瞬间在漫长的历史时空中微不足道的分量。同样微不足道的还有那一瞬...

                        &nbsp;(<a href="javascript:;" id="toggle-13011950-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13011950_full" class="hidden">
                    <div id="review_13011950_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13011950" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13011950">
                                79
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13011950" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13011950">
                                10
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13011950/#comments" class="reply ">17回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13015375">
        <div class="main review-item" id="13015375">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/nothingkitsch/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u51931810-20.jpg">
        </a>

        <a href="https://www.douban.com/people/nothingkitsch/" class="name">nothing</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-28" class="main-meta">2020-11-28 12:57:10</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13015375/">电影已散场？谁又知道呢</a></h2>

                <div id="review_13015375_short" class="review-short" data-rid="13015375">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        相比于之前张艺谋那些动辄上亿大制作，追求拍摄规模与视觉突破的电影，《一秒钟》基调质朴，不太像人们习以为常那些个观影体验。很多东西不是靠视觉呈现的，而是以一种余韵的方式留给观众去品味和思考，以至于电影大幕落下，影院灯光亮起，自己一时还会有点回不过神来：“这就...

                        &nbsp;(<a href="javascript:;" id="toggle-13015375-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13015375_full" class="hidden">
                    <div id="review_13015375_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13015375" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13015375">
                                46
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13015375" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13015375">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13015375/#comments" class="reply ">2回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="13012684">
        <div class="main review-item" id="13012684">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/CanaanYoung/" class="avator">
            <img width="24" height="24" src="https://img2.doubanio.com/icon/u73496903-22.jpg">
        </a>

        <a href="https://www.douban.com/people/CanaanYoung/" class="name">迦南Canaan</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2020-11-27" class="main-meta">2020-11-27 13:31:52</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/13012684/">无名时代里的一秒钟</a></h2>

                <div id="review_13012684_short" class="review-short" data-rid="13012684">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        长久以来的叙述中，第五代导演都被认为是背负了历史重担的。无论是戴锦华将第五代称之为某时代之子，还是张艺谋自己说如果有机会要拍十部关于某时代的电影，都印证了这一点。如今，在《归来》六年之后，张艺谋终于又就这一题材挥洒才情。 一秒钟的双重维度 正如片名所昭示的，...

                        &nbsp;(<a href="javascript:;" id="toggle-13012684-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_13012684_full" class="hidden">
                    <div id="review_13012684_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="13012684" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-13012684">
                                39
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="13012684" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-13012684">
                                8
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/13012684/#comments" class="reply ">3回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>




    

    

    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/268f202f5305a01.js"></script>
    <!-- COLLECTED CSS -->
</div>








                <p class="pl">
                    &gt;
                    <a href="reviews">
                        更多影评
                            912篇
                    </a>
                </p>
    </section>
<!-- COLLECTED JS -->

    <br/>

            <div class="section-discussion">
                    
                    <div class="mod-hd">
                            <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow"><span>添加新讨论</span></a>
                        
    <h2>
        讨论区
         &nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;
    </h2>

                    </div>
                    
  <table class="olt"><tr><td><td><td><td></tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30257787/discussion/616818612/" title="一秒钟还可以再狠一点">一秒钟还可以再狠一点</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/139468945/">戴眼镜的怪蜀黍</a></td>
          <td class="pl"><span>34 回应</span></td>
          <td class="pl"><span>2020-12-01</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30257787/discussion/616820013/" title="电影是好电影但是有些不合时宜">电影是好电影但是有些不合时宜</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/153814967/">阿荣</a></td>
          <td class="pl"><span>44 回应</span></td>
          <td class="pl"><span>2020-12-01</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30257787/discussion/616821896/" title="时代车轮滚滚向前，碾死一批伤痕酸臭老古董。">时代车轮滚滚向前，碾死一批伤痕酸臭老古董。</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/157425845/">鱼怪</a></td>
          <td class="pl"><span>5 回应</span></td>
          <td class="pl"><span>2020-12-01</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30257787/discussion/616821416/" title="未来可以看到原版的吗？">未来可以看到原版的吗？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/116272612/">老叮</a></td>
          <td class="pl"><span>5 回应</span></td>
          <td class="pl"><span>2020-12-01</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30257787/discussion/616820630/" title="刘闺女的声音太尖了，尖的让人不适">刘闺女的声音太尖了，尖的让人不适</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/156377387/">三哥</a></td>
          <td class="pl"><span>14 回应</span></td>
          <td class="pl"><span>2020-12-01</span></td>
        </tr>
  </table>

                    <p class="pl" align="right">
                        <a href="/subject/30257787/discussion/" rel="nofollow">
                            &gt; 去这部影片的讨论区（全部443条）
                        </a>
                    </p>
            </div>


    <script type="text/javascript">
        $(function(){if($.browser.msie && $.browser.version == 6.0){
            var $info = $('#info'),
                maxWidth = parseInt($info.css('max-width'));
            if($info.width() > maxWidth) {
                $info.width(maxWidth);
            }
        }});
    </script>


            </div>
            <div class="aside">
                


    







            






    <div class="ticket">
        <a class="ticket-btn" href="https://movie.douban.com/ticket/redirect/?movie_id=30257787">购票</a>
    </div>



    <!-- douban ad begin -->
    <div id="dale_movie_subject_top_right"></div>
    <div id="dale_movie_subject_top_middle"></div>
    <!-- douban ad end -->

    



<style type="text/css">
    .m4 {margin-bottom:8px; padding-bottom:8px;}
    .movieOnline {background:#FFF6ED; padding:10px; margin-bottom:20px;}
    .movieOnline h2 {margin:0 0 5px;}
    .movieOnline .sitename {line-height:2em; width:160px;}
    .movieOnline td,.movieOnline td a:link,.movieOnline td a:visited{color:#666;}
    .movieOnline td a:hover {color:#fff;}
    .link-bt:link,
    .link-bt:visited,
    .link-bt:hover,
    .link-bt:active {margin:5px 0 0; padding:2px 8px; background:#a8c598; color:#fff; -moz-border-radius: 3px; -webkit-border-radius: 3px; border-radius: 3px; display:inline-block;}
</style>



    







    
    <div class="tags">
        
        
    <h2>
        <i class="">豆瓣成员常用的标签</i>
              · · · · · ·
    </h2>

        <div class="tags-body">
                <a href="/tag/文艺" class="">文艺</a>
                <a href="/tag/文革" class="">文革</a>
                <a href="/tag/中国大陆" class="">中国大陆</a>
                <a href="/tag/人性" class="">人性</a>
                <a href="/tag/剧情" class="">剧情</a>
                <a href="/tag/历史" class="">历史</a>
                <a href="/tag/社会" class="">社会</a>
                <a href="/tag/2020" class="">2020</a>
        </div>
    </div>


    <div id="dale_subject_right_guess_you_like"></div>
    <div id="dale_movie_subject_inner_middle"></div>
    <div id="dale_movie_subject_download_middle"></div>
        








<div id="subject-doulist">
    
    
    <h2>
        <i class="">以下豆列推荐</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30257787/doulists">全部</a>
            )
            </span>
    </h2>


    
    <ul>
            <li>
                <a href="https://www.douban.com/doulist/30299/" target="_blank">豆瓣电影【口碑榜】2020-07-10 更新</a>
                <span>(影志)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1768918/" target="_blank">在这里，看懂中国</a>
                <span>(时间之葬)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/10534/" target="_blank">『我是你爸爸~』</a>
                <span>(非魚)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1504454/" target="_blank">ღ♩♪生活有这些期待很有动力♫♬ღ</a>
                <span>(freedom♪)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1434921/" target="_blank">2020—2022值得关注的华语电影</a>
                <span>(closer)</span>
            </li>
    </ul>

</div>

            








<div id="subject-others-interests">
    
    
    <h2>
        <i class="">谁在看这部电影</i>
              · · · · · ·
    </h2>

    
    <ul class="">
            
            <li class="">
                <a href="https://www.douban.com/people/120578850/" class="others-interest-avatar">
                    <img src="https://img9.doubanio.com/icon/u120578850-5.jpg" class="pil" alt="可爱酒🐶">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/120578850/" class="">可爱酒🐶</a>
                    <div class="">
                        刚刚
                        看过
                        <span class="allstar50" title="力荐"></span>
                    </div>
                </div>
            </li>
            
            <li class="">
                <a href="https://www.douban.com/people/48600640/" class="others-interest-avatar">
                    <img src="https://img1.doubanio.com/icon/u48600640-9.jpg" class="pil" alt="迈克尔~壮壮">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/48600640/" class="">迈克尔~壮壮</a>
                    <div class="">
                        刚刚
                        想看
                        
                    </div>
                </div>
            </li>
            
            <li class="">
                <a href="https://www.douban.com/people/4273555/" class="others-interest-avatar">
                    <img src="https://img9.doubanio.com/icon/u4273555-6.jpg" class="pil" alt="Lindsay">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/4273555/" class="">Lindsay</a>
                    <div class="">
                        刚刚
                        看过
                        
                    </div>
                </div>
            </li>
    </ul>

    
    <div class="subject-others-interests-ft">
        
            <a href="https://movie.douban.com/subject/30257787/comments?status=P">77978人看过</a>
                &nbsp;/&nbsp;
            <a href="https://movie.douban.com/subject/30257787/comments?status=F">79041人想看</a>
    </div>

</div>



    

<!-- douban ad begin -->
<div id="dale_movie_subject_middle_right"></div>
<script type="text/javascript">
    (function (global) {
        if(!document.getElementsByClassName) {
            document.getElementsByClassName = function(className) {
                return this.querySelectorAll("." + className);
            };
            Element.prototype.getElementsByClassName = document.getElementsByClassName;

        }
        var articles = global.document.getElementsByClassName('article'),
            asides = global.document.getElementsByClassName('aside');

        if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_middle_right');
        }
    })(this);
</script>
<!-- douban ad end -->



    <br/>

    
<p class="pl">订阅一秒钟的评论: <br/><span class="feed">
    <a href="https://movie.douban.com/feed/subject/30257787/reviews"> feed: rss 2.0</a></span></p>


            </div>
            <div class="extra">
                
    
<!-- douban ad begin -->
<div id="dale_movie_subject_bottom_super_banner"></div>
<script type="text/javascript">
    (function (global) {
        var body = global.document.body,
            html = global.document.documentElement;

        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        if (height >= 2000) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_bottom_super_banner');
        }
    })(this);
</script>
<!-- douban ad end -->


            </div>
        </div>
    </div>

        
    <div id="footer">
            <div class="footer-extra"></div>
        
<span id="icp" class="fleft gray-link">
    &copy; 2005－2020 douban.com, all rights reserved 北京豆网科技有限公司
</span>

<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>

<span class="fright">
    <a href="https://www.douban.com/about">关于豆瓣</a>
    · <a href="https://www.douban.com/jobs">在豆瓣工作</a>
    · <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
    · <a href="https://www.douban.com/about/legal">法律声明</a>
    
    · <a href="https://help.douban.com/?app=movie" target="_blank">帮助中心</a>
    · <a href="https://www.douban.com/doubanapp/">移动应用</a>
    · <a href="https://www.douban.com/partner/">豆瓣广告</a>
</span>

    </div>

    </div>
    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/11f0fa39775e6347.js"></script>
        
        
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css" />
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/movie/4aca95d66d37ec0712b3d19973b5d8feb75f2f05/css/movie/mod/reg_login_pop.css" />
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js"></script>
    <script type="text/javascript">
        var HTTPS_DB='https://www.douban.com';
var account_pop={open:function(o,e){e?referrer="?referrer="+encodeURIComponent(e):referrer="?referrer="+window.location.href;var n="",i="",t=448;n="用户登录",i="https://accounts.douban.com/passport/login_popup?source=movie";var r=document.location.protocol+"//"+document.location.hostname,a=dui.Dialog({width:340,title:n,height:t,cls:"account_pop",isHideTitle:!0,modal:!0,content:"<iframe scrolling='no' frameborder='0' width='340' height='"+t+"' src='"+i+"' name='"+r+"'></iframe>"},!0),c=a.node;if(c.undelegate(),c.delegate(".dui-dialog-close","click",function(){var o=$("body");o.find("#login_msk").hide(),o.find(".account_pop").remove()}),$(window).width()<478){var d="";"reg"===o?d=HTTPS_DB+"/accounts/register"+referrer:"login"===o&&(d=HTTPS_DB+"/accounts/login"+referrer),window.location.href=d}else a.open();$(window).bind("message",function(o){"https://accounts.douban.com"===o.originalEvent.origin&&(c.find("iframe").css("height",o.originalEvent.data),c.height(o.originalEvent.data),a.update())})}};Douban&&Douban.init_show_login&&(Douban.init_show_login=function(o){var e=$(o);e.click(function(){var o=e.data("ref")||"";return account_pop.open("login",o),!1})}),Do(function(){$("body").delegate(".pop_register","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("reg",e),!1}),$("body").delegate(".pop_login","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("login",e),!1})});
    </script>

    
    
    
    




    
<script type="text/javascript">
    (function (global) {
        var newNode = global.document.createElement('script'),
            existingNode = global.document.getElementsByTagName('script')[0],
            adSource = '//erebor.douban.com/',
            userId = '',
            browserId = 'giqphxEOs64',
            criteria = '7:文革|7:社会|7:剧情|7:李延|7:刘浩存|7:中国大陆|7:常海军|7:于洋|7:范伟|7:柏林电影节|7:余皑磊|7:经典|7:刘云龙|7:张艺谋|7:张译|7:历史|7:唐孜悦|7:曹瑞|7:人性|7:文艺|7:2020|7:张邵勃 |3:/subject/30257787/?from=showing',
            preview = '',
            debug = false,
            adSlots = ['dale_movie_subject_top_icon', 'dale_movie_subject_top_right', 'dale_movie_subject_top_middle', 'dale_movie_subject_inner_middle', 'dale_movie_subject_download_middle', 'dale_movie_subject_banner_after_intro', 'dale_subject_right_guess_you_like'];

        global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
        global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);

        newNode.setAttribute('type', 'text/javascript');
        newNode.setAttribute('src', '//img1.doubanio.com/czF5ODV4Ni9mL2FkanMvMjBiYWY2MDg2ZWQ0MDE1YTNmMDJhNDYxMzhmNmM0MjQxYjExYWYwMC9hZC5yZWxlYXNlLmpz');
        newNode.setAttribute('async', true);
        existingNode.parentNode.insertBefore(newNode, existingNode);
    })(this);
</script>











    
  









<script type="text/javascript">
var _paq = _paq || [];
_paq.push(['trackPageView']);
_paq.push(['enableLinkTracking']);
(function() {
    var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
    _paq.push(['setTrackerUrl', u+'piwik']);
    _paq.push(['setSiteId', '100001']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
    g.type='text/javascript';
    g.defer=true;
    g.async=true;
    g.src=p+'://img3.doubanio.com/dae/fundin/piwik.js';
    s.parentNode.insertBefore(g,s);
})();
</script>

<script type="text/javascript">
var setMethodWithNs = function(namespace) {
  var ns = namespace ? namespace + '.' : ''
    , fn = function(string) {
        if(!ns) {return string}
        return ns + string
      }
  return fn
}

var gaWithNamespace = function(fn, namespace) {
  var method = setMethodWithNs(namespace)
  fn.call(this, method)
}

var _gaq = _gaq || []
  , accounts = [
      { id: 'UA-7019765-1', namespace: 'douban' }
    , { id: 'UA-7019765-19', namespace: '' }
    ]
  , gaInit = function(account) {
      gaWithNamespace(function(method) {
        gaInitFn.call(this, method, account)
      }, account.namespace)
    }
  , gaInitFn = function(method, account) {
      _gaq.push([method('_setAccount'), account.id]);
      _gaq.push([method('_setSampleRate'), '5']);

      
  _gaq.push([method('_addOrganic'), 'google', 'q'])
  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
  _gaq.push([method('_addOrganic'), 'soso', 'w'])
  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
  if (account.namespace) {
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
    _gaq.push([method('_addIgnoredOrganic'), 'douban'])
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
    _gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
  }

      if (account.namespace === 'douban') {
        _gaq.push([method('_setDomainName'), '.douban.com'])
      }

        _gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])

        _gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);

      _gaq.push([method('_trackPageview')])
    }

for(var i = 0, l = accounts.length; i < l; i++) {
  var account = accounts[i]
  gaInit(account)
}


;(function() {
    var ga = document.createElement('script');
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    ga.setAttribute('async', 'true');
    document.documentElement.firstChild.appendChild(ga);
})()
</script>








      

    <!-- dae-web-movie--default-fbb9d4b9f-dzszm-->

  <script>_SPLITTEST=''</script>
</body>

</html>



`

	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "cnt"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "lang"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "mc"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "type"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "time"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "span"))
	c.Ctx.WriteString(models.GetMovieDetail(sMovieHtml, "grd"))

}