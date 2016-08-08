var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
input.verticalInput({
	id:'container',
	field:[
		{id:'name',type:'text',name:'姓名'},
		{id:'password',type:'text',name:'密码'},
		{id:'type',type:'enum',name:'类型',map:{'1':'管理员','2':'普通会员'}},
	],
	submit:function(data){
		$.post('/user/add',data,function(data){
			data = $.JSON.parse(data);
			if( data.code != 0 ){
				dialog.message(data.msg);
				return;
			}
			history.back();
		});
	},
	cancel:function(){
		history.back();
	}
});

