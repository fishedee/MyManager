var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
input.verticalInput({
	id:'container',
	field:[
		{id:'url',type:'text',name:'刷榜链接'},
		{id:'type',type:'enum',name:'类型',map:{'1':'直接抓取','2':'代理抓取'}},
		{id:'totalNum',type:'text',name:'刷榜数量'},
	],
	submit:function(data){
		$.post('/brush/addTask',data,function(data){
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

