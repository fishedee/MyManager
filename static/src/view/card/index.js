var $ = require('fishfront/ui/global');
var query = require('fishfront/ui/query');
var dialog = require('fishfront/ui/dialog');
query.simpleQuery({
	id:'container',
	url:'/card/search',
	column:[
		{id:'cardId',type:'text',name:'银行卡ID'},
		{id:'name',type:'text',name:'名称'},
		{id:'bank',type:'text',name:'银行'},
		{id:'remark',type:'text',name:'备注'},
		{id:'createTime',type:'text',name:'创建时间'},
		{id:'modifyTime',type:'text',name:'修改时间'},
	],
	queryColumn:['name','remark'],
	operate:[
	{
		name:'编辑',
		click:function(data){
			location.href = 'view.html?cardId='+data.cardId;
		}
	},
	{
		name:'删除',
		click:function(data){
			dialog.confirm('确认删除该栏目，不可回退操作？!',function(){
				$.post('/card/del',{cardId:data.cardId},function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					location.href = 'index.html';
				});
			});
		}
	}],
	button:[
	{
		name:'添加银行卡',
		click:function(){
			location.href = 'view.html';
		}
	}
	],
});
