var $ = require('fishfront/ui/global');
var dialog = require('fishfront/ui/dialog');
var query = require('fishfront/ui/query');
var categorys = {0:'未分类'};
var cards = {0:'无银行卡'};
var types = {1:'收入',2:'支出',3:'转账收入',4:'转账支出',5:'借还账收入',6:'借还账支出'};
function getCategory(nextStep){
	$.get('/category/search',{pageIndex:-1,pageSize:-1},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 ){
			dialog.message( data.msg );
			return;
		}
		_.each(data.data.data,function(single){
			categorys[single.categoryId] = single.name;
		});
		nextStep();
	});
}
function getCard(nextStep){
	$.get('/card/search',{pageIndex:-1,pageSize:-1},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 ){
			dialog.message( data.msg );
			return;
		}
		_.each(data.data.data,function(single){
			cards[single.cardId] = single.name;
		});
		nextStep();
	});
}
function go(){
	query.simpleQuery({
		id:'container',
		url:'/account/search',
		column:[
			{id:'accountId',type:'text',name:'账务ID'},
			{id:'name',type:'text',name:'名称'},
			{id:'money',type:'text',name:'金额'},
			{id:'type',type:'enum',name:'类型',map:types},
			{id:'categoryId',type:'enum',name:'分类',map:categorys},
			{id:'cardId',type:'enum',name:'银行卡',map:cards},
			{id:'remark',type:'text',name:'备注'},
			{id:'createTime',type:'text',name:'创建时间'},
			{id:'modifyTime',type:'text',name:'修改时间'},
		],
		queryColumn:['name','remark','categoryId','cardId','type'],
		operate:[
		{
			name:'编辑',
			click:function(data){
				location.href = 'view.html?accountId='+data.accountId;
			}
		},
		{
			name:'删除',
			click:function(data){
				dialog.confirm('确认删除该账务，不可回退操作？!',function(){
					$.post('/account/del',{accountId:data.accountId},function(data){
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
			name:'添加账务',
			click:function(){
				location.href = 'view.html';
			}
		}
		],
	});
}
getCategory(function(){
	getCard(function(){
		go();
	});
});
