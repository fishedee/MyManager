var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
var userId = $.location.getQueryArgv('userId');
$.get('/user/get',{userId:userId},function(data){
	data = $.JSON.parse(data);
	if( data.code != 0 ){
		dialog.message(data.msg);
		return;
	}
	data = data.data;
	input.verticalInput({
		id:'container',
		field:[
			{id:'name',type:'read',name:'姓名'},
			{id:'type',type:'enum',name:'类型',map:{'1':'管理员','2':'普通会员'}},
		],
		value:data,
		submit:function(data){
			data = {userId:userId,type:data.type};
			$.post('/user/modType',data,function(data){
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
});
