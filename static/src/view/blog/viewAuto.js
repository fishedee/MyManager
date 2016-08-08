var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
var blogSyncAutoId = $.location.getQueryArgv('blogSyncAutoId');
function go(value){
	input.verticalInput({
		id:'container',
		field:[
			{id:'gitUrl',type:'text',name:'git地址'},
			{id:'accessToken',type:'text',name:'csdn访问授权'},
		],
		value:value,
		submit:function(data){
			if( blogSyncAutoId != null ){
				data = $.extend({blogSyncAutoId:blogSyncAutoId},data);
				$.post('/blog/modAuto',data,function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			}else{
				$.post('/blog/addAuto',data,function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			}
		},
		cancel:function(){
			history.back();
		}
	});
}
if( blogSyncAutoId != null ){
	$.get('/blog/getAuto',{blogSyncAutoId:blogSyncAutoId},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 ){
			dialog.message(data.msg);
			return;
		}
		go(data.data);
	});
}else{
	go({});
}
