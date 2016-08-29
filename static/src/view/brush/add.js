var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
$('body').prepend('<div class="alert alert-danger" role="alert">注意！<br>刷榜链接的填写格式为：http://www.baidu.com<br/>类型可以选择为不走代理，或走xici免费代理(http://xicidaili.com/nn/1)<br/>最大重试次数为单个刷榜请求的最大重试次数，失败后会在1s后自动重试</div>')
input.verticalInput({
	id:'container',
	field:[
		{id:'url',type:'text',name:'刷榜链接'},
		{id:'type',type:'enum',name:'类型',map:{'1':'直接抓取','2':'xici代理抓取'}},
		{id:'totalNum',type:'text',name:'刷榜数量'},
		{id:'retryNum',type:'text',name:'最大重试次数'},
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

