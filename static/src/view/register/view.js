var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
var registerId = $.location.getQueryArgv('registerId');
var needDealType = {1:'不需要',2:'需要'};
var haveDealType = {1:'未挂号',2:'已挂号'};
function go(value){
	input.verticalInput({
		id:'container',
		field:[
			{id:'registerId',type:'text',name:'挂号ID'},
			{id:'name',type:'text',name:'姓名'},
			{id:'beginTime',type:'text',name:'开始时间'},
			{id:'endTime',type:'text',name:'结束时间'},
			{id:'mail',type:'text',name:'提醒邮箱'},
			{id:'needDealType',type:'enum',name:'是否需要自动挂号',map:needDealType},
			{id:'haveDealType',type:'enum',name:'是否已经自动挂号',map:haveDealType},
			{id:'haveDealTime',type:'text',name:'已经自动挂号时间'},
		],
		value:value,
		submit:function(data){
			if( registerId != null ){
				data = $.extend({registerId:registerId},data);
				$.post('/register/mod',data,function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			}else{
				$.post('/register/add',data,function(data){
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
if( registerId != null ){
	$.get('/register/get',{registerId:registerId},function(data){
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
