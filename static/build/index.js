webpackJsonp([0],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var indexPage = __webpack_require__(1);
	var dialog = __webpack_require__(32);
	var $ = __webpack_require__(6);
	indexPage.use({
		title: 'Fish个人管理系统',
		init: function () {
			$.get('/login/islogin?t=' + Math.random(), {}, function (data) {
				data = $.JSON.parse(data);
				if (data.code != 0) {
					location.href = 'login.html';
					return;
				}
			});
		},
		logout: function () {
			$.get('/login/checkout?t=' + Math.random(), {}, function (data) {
				data = $.JSON.parse(data);
				if (data.code != 0) {
					dialog.message(data.msg);
					return;
				}
				location.href = 'login.html';
			});
		},
		menu: {
			'系统管理': {
				'系统管理': [{ name: '帐号管理', url: 'view/user/index.html' }, { name: '密码管理', url: 'view/password/index.html' }]
			},
			'财务管理': {
				'财务管理': [{ name: '分类管理', url: 'view/category/index.html' }, { name: '银卡管理', url: 'view/card/index.html' }, { name: '账务管理', url: 'view/account/index.html' }, { name: '收支统计', url: 'view/account/inoutStatistic.html' }, { name: '结余统计', url: 'view/account/totalStatistic.html' }]
			},
			'便捷工具': {
				'博客同步': [{ name: '自动同步', url: 'view/blog/indexAuto.html' }, { name: '一键同步', url: 'view/blog/index.html' }],
				'刷榜工具': [{ name: '自动刷榜', url: 'view/brush/index.html' }],
				'市一产科挂号': [{ name: '自动挂号', url: 'view/register/index.html' }],
				'便捷工具': [{ name: '优惠卷工具', url: 'view/coupon/view.html' }]
			}
		}
	});

/***/ },
/* 1 */
/***/ function(module, exports, __webpack_require__) {

	__webpack_require__(2);
	var $ = __webpack_require__(6);
	module.exports = {
		use:function(option){
			//处理Option
			option = option || {};
			var defaultOption = {
				title:'后台管理系统',
				menu:[],
				logout:function(){
				},
				init:function(){
				}
			};
			for( var i in option )
				defaultOption[i] = option[i];
			//添加基本框架
			var menu = "";
			for( var name in defaultOption.menu ){
				menu = '<div class="comm_rightcol item button" name="'+name+'">'+name+'</div>' + menu;
			}
			var div = 
				'<div class="comm_abtop" id="topbar">'+
				'	<div class="comm_leftcol item" id="title">'+defaultOption.title+'</div>'+
				'	<div class="comm_rightcol item button" name="logout">[退出]</div>'+
					menu+
				'</div>'+
				'<div class="comm_ableft" id="leftbar">'+
				'</div>'+
				'<div class="comm_abright" id="rightbar">'+
				'	<div id="pageTitle">用户管理</div>'+
				'	<div id="pageLine"></div>'+
				'</div>'+
				'<div class="comm_abfull" id="centerbar">'+
				'	<iframe name="myframe" src="" frameborder="no"/>'+
				'</div>';
			div = $(div);
			$('body').append(div);
			//设置样式
			$.addCssToHead(
				'#topbar{'+
				'	height:50px;'+
				'	background:#222;'+
				'	background-image:linear-gradient(to bottom,#3c3c3c 0,#222 100%);'+
				'	background-repeat:repeat-y;'+
				'	border-color:#080808;'+
				'}'+
				'#topbar .item{'+
				'	font-size:18px;'+
				'	line-height:50px;'+
				'	color:#777;'+
				'	text-shadow:0 -1px 0 rgba(0,0,0,.25);'+
				'	transition:color 0.5s;'+
				'}'+
				'#topbar .item:hover{'+
				'	cursor:pointer;'+
				'	color:#fff;'+
				'}'+
				'#topbar #title{'+
				'	margin-left:15px;'+
				'}'+
				'#topbar .button{'+
				'	font-size:14px;'+
				'	margin-right:15px;'+
				'}'+
				'#leftbar{'+
				'	top:50px;'+
				'	width:200px;'+
				'	background:#f5f5f5;'+
				'	border-right:1px solid #eee;'+
				'}'+
				'#leftbar .category:first-child{'+
				'	margin-top:15px;'+
				'}'+
				'#leftbar .category .head{'+
				'	line-height: 25px;'+
				'	height: 25px;'+
				'	font-size: 12px;'+
				'	margin-left: 15px;'+
				'	color: black;'+
				'	font-weight:bold;'+
				'}'+
				'#leftbar .category .title{'+
				'	height:40px;'+
				'	line-height:40px;'+
				'	padding-left:30px;'+
				'	background:transparent;'+
				'	transition:all 0.5s;'+
				'	font-size:14px;'+
				'	color:#2a6496;'+
				'	text-decoration:none;'+
				'}'+
				'#leftbar .category .title:hover{'+
				'	background:#eee;'+
				'	cursor:pointer;'+
				'}'+
				'#leftbar .category .activetitle{'+
				'	background:rgb(66, 139, 202);'+
				' 	color:white;'+
				'}'+
				'#leftbar .category .activetitle:hover{'+
				'	background:rgb(66, 139, 202);'+
				' 	color:white;'+
				'}'+
				'#centerbar{'+
				'	top:100px;'+
				'	left:200px;'+
				'	bottom:10px;'+
				'	padding-top:15px;'+
				'	padding-left:10px;'+
				'	padding-right:10px;'+
				'	width:auto;'+
				'	height:auto;'+
				'}'+
				'#centerbar iframe{'+
				'	width:100%;'+
				'	height:100%;'+
				'	border:0px;'+
				'}'+
				'#rightbar{'+
				'	display:none;'+
				'	top:50px;'+
				'	left:200px;'+
				'	height:50px;'+
				'	padding-top:15px;'+
				'	padding-left:10px;'+
				'	padding-right:10px;'+
				'	width:auto;'+
				'	height:auto;'+
				'}'+
				'#rightbar #pageTitle{'+
				'	font-size:24px;'+
				'	color:rgb(180,180,180);'+
				'}'+
				'#rightbar #pageLine{'+
				'	border-bottom:1px solid rgb(220,220,220);'+
				'	margin-top:20px;'+
				'}'
			);
			function getRedirectUrl(url){
				if(url.indexOf('?')!=-1)
					return url + '&t='+new Date().getTime();
				else
					return url + '?t='+new Date().getTime();
			}
			function chooseTopMenu(topMenuItemName){
				if( topMenuItemName == 'logout'){
					defaultOption.logout();
				}else{
					var menu = "";
					for( var name in defaultOption.menu[topMenuItemName] ){
						var items = defaultOption.menu[topMenuItemName][name];
						menu +=
						'	<div class="category">'+
						'		<div class="head">'+name+'</div>';
						for( var i in items ){
							var item = items[i];
							menu += '<a href="'+getRedirectUrl(item.url)+'" target="myframe" data-href="'+item.url+'">'+
							'<div class="title">'+item.name+'</div>'+
							'</a>';
						}
						menu += '</div>';
					}
					$('#leftbar').html(menu);
					$.location.setHashArgv({
						'menu':topMenuItemName
					});
					var firstMenuClick = _.keys(defaultOption.menu[topMenuItemName])[0];
					chooseLeftBarAndClick(defaultOption.menu[topMenuItemName][firstMenuClick][0].url);
				}
			}
			function chooseLeftBar(leftMenuHref){
				var leftMenu = null;
				div.filter('#leftbar').find('a').each(function(){
					if( $(this).attr('data-href') != leftMenuHref )
						return;
					leftMenu = $(this);
				});
				if( leftMenu == null )
					return;
				$('#rightbar').show();
				$('#rightbar #pageTitle').text(leftMenu.find('.title').text());
				$('#leftbar .category .title').removeClass('activetitle');
				leftMenu.find('.title').addClass('activetitle');
				leftMenu.attr('href',getRedirectUrl(leftMenu.attr('data-href')));
				$.location.setHashArgv({
					'menu':$.location.getHashArgv('menu'),
					'location':leftMenuHref
				});
			}
			function chooseLeftBarAndClick(leftMenuHref){
				var leftMenu = null;
				div.filter('#leftbar').find('a').each(function(){
					if( $(this).attr('data-href') != leftMenuHref )
						return;
					leftMenu = $(this);
				});
				leftMenu.find('div').click();
			}
			//设置事件
			div.on('click','a',function(){
				chooseLeftBar($(this).attr('data-href'));
			});
			div.filter('#topbar').on('click','.button',function(){
				chooseTopMenu($(this).attr('name'));
			});
			//启动
			var menu = $.location.getHashArgv('menu');
			if( menu != null ){
				var location = $.location.getHashArgv('location');
				if( location != null ){
					chooseTopMenu(menu);
					chooseLeftBarAndClick(location);
				}else{
					chooseTopMenu(menu);
				}
			}else{
				var firstMenuKey = _.keys(defaultOption.menu)[0];
				chooseTopMenu(firstMenuKey);
			}
			defaultOption.init();
		}
	};

/***/ },
/* 2 */
/***/ function(module, exports, __webpack_require__) {

	// style-loader: Adds some css to the DOM by adding a <style> tag

	// load the styles
	var content = __webpack_require__(3);
	if(typeof content === 'string') content = [[module.id, content, '']];
	// add the styles to the DOM
	var update = __webpack_require__(5)(content, {});
	if(content.locals) module.exports = content.locals;
	// Hot Module Replacement
	if(false) {
		// When the styles change, update the <style> tags
		if(!content.locals) {
			module.hot.accept("!!./../../css-loader/index.js!./indexPage.css", function() {
				var newContent = require("!!./../../css-loader/index.js!./indexPage.css");
				if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
				update(newContent);
			});
		}
		// When the module is disposed, remove the <style> tags
		module.hot.dispose(function() { update(); });
	}

/***/ },
/* 3 */
/***/ function(module, exports, __webpack_require__) {

	exports = module.exports = __webpack_require__(4)();
	// imports


	// module
	exports.push([module.id, "/*全局css*/\n*{\n\tmargin:0px;\n\tpadding:0px;\n\tborder:0px;\n\tfont-size:0px;\n\ttext-decoration:none;\n\tfont-family:\"Helvetica Neue\",Helvetica,STHeiTi,sans-serif;\n\t/*IE浏览器的字体*/\n\tfont-family:\"Microsoft Yahei\",Helvetica,Arial\\9;\n\t-webkit-touch-callout: none; \n\t-webkit-highlight: none; \n\t-webkit-tap-highlight-color: rgba(0,0,0,0);\n}\n.comm_body{\n\tmargin:0 auto;\n\tmax-width:750px;\n}\n/*float布局的常用css*/\n.comm_row{\n\toverflow:auto; \n\tzoom:1;\n}\n.comm_leftcol{\n\tfloat:left;\n}\n.comm_rightcol{\n\tfloat:right;\n}\n.comm_centercol{\n\toverflow:auto; \n\tzoom:1;\n\tmargin:0 auto;\n}\n/*absolue布局的常用css*/\n.comm_abfull{\n\tposition:absolute;\n\ttop:0px;\n\tbottom:0px;\n\tleft:0px;\n\tright:0px;\n\twidth:100%;\n\theight:100%;\n}\n.comm_abcenter{\n\tposition:absolute;\n\ttop:0px;\n\tbottom:0px;\n\tleft:0px;\n\tright:0px;\n\tmargin:auto auto;\n}\n.comm_abtop{\n\tposition:absolute;\n\ttop:0px;\t\n\tleft:0px;\n\tright:0px;\n\tmargin:0 auto;\n}\n.comm_ableft{\n\tposition:absolute;\n\ttop:0px;\n\tbottom:0px;\n\tleft:0px;\n\tmargin:auto 0;\n}\n.comm_abright{\n\tposition:absolute;\n\ttop:0px;\n\tbottom:0px;\n\tright:0px;\n\tmargin:auto 0;\n}\n.comm_abbottomcenter{\n\tposition:absolute;\n\tbottom:0px;\n\tleft:0px;\n\tright:0px;\n\tmargin:0 auto;\n}\n/*fix布局的常用css*/\n/*控件常用css*/\n.comm_img{\n\twidth:100%;\n\theight:100%;\n}\n.comm_rowimg{\n\twidth:100%;\n}\n.comm_input{\n\twidth:100%;\n\theight:30px;\n\tvertical-align:middle;\n\tfont-size:16px;\n\tborder-bottom:1px dotted rgb(170,170,170);\n}\n.comm_show{\n\twidth:100%;\n\tline-height:30px;\n\tfont-size:16px;\n\tborder-bottom:1px dotted rgb(170,170,170);\n\tcolor:rgb(170,170,170);\n}\n", ""]);

	// exports


/***/ }
]);