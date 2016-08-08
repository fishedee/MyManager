var $ = require('fishfront/ui/global');
var loginPage = require('fishfront/ui/loginPage');
var dialog = require('fishfront/ui/dialog');
loginPage.use({
	title:'Fish个人管理系统',
	init:function(){
		$.get('/login/islogin?t='+Math.random(),{},function(data){
			data = $.JSON.parse(data);
			if( data.code == 0 ){
				location.href = 'index.html';
				return;
			}
		});
	},
	login:function(data){
		$.post('/login/checkin?t='+Math.random(),data,function(data){
			data = $.JSON.parse(data);
			if( data.code != 0 ){
				dialog.message(data.msg);
				return;
			}
			location.href = 'index.html';
		});
	}
});
