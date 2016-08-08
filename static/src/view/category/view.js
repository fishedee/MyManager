var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
var categoryId = $.location.getQueryArgv('categoryId');
function go(value){
	input.verticalInput({
		id:'container',
		field:[
			{id:'name',type:'text',name:'名字'},
			{id:'remark',type:'text',name:'备注'},
		],
		value:value,
		submit:function(data){
			if( categoryId != null ){
				data = $.extend({categoryId:categoryId},data);
				$.post('/category/mod',data,function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			}else{
				$.post('/category/add',data,function(data){
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
if( categoryId != null ){
	$.get('/category/get',{categoryId:categoryId},function(data){
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
