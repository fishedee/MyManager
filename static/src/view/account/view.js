var $ = require('fishfront/ui/global');
var dialog = require('fishfront/ui/dialog');
var input = require('fishfront/ui/input');
var chart = require('fishfront/ui/chart');
var accountId = $.location.getQueryArgv('accountId');
var categorys = {0:'未分类'};
var cards = {0:'无银行卡'};
var types = {1:'收入',2:'支出',3:'转账收入',4:'转账支出',5:'借还账收入',6:'借还账支出'};
var account = {};
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
function getAccount(nextStep){
	$.get('/account/get',{accountId:accountId},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 ){
			dialog.message(data.msg);
			return;
		}
		account = data.data;
		nextStep();
	});
}
function go(){
	input.verticalInput({
		id:'container',
		field:[
			{id:'name',type:'text',name:'名称'},
			{id:'money',type:'text',name:'金额'},
			{id:'type',type:'enum',name:'类型',map:types},
			{id:'categoryId',type:'enum',name:'分类',map:categorys},
			{id:'cardId',type:'enum',name:'银行卡',map:cards},
			{id:'remark',type:'text',name:'备注'},
		],
		value:account,
		submit:function(data){
			if( accountId != null ){
				data = $.extend({accountId:accountId},data);
				$.post('/account/mod',data,function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			}else{
				$.post('/account/add',data,function(data){
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
getCategory(function(){
	getCard(function(){
		if( accountId != null ){
			getAccount(function(){
				go();
			});
		}else{
			go();
		}
	});
});
