var $ = require('fishfront/ui/global');
var query = require('fishfront/ui/query');
var dialog = require('fishfront/ui/dialog');
query.simpleQuery({
	id:'container',
	url:'/brush/searchTask',
	column:[
		{id:'brushTaskId',type:'text',name:'任务ID'},
		{id:'url',type:'text',name:'链接'},
		{id:'type',type:'enum',name:'类型',map:{'1':'直接抓取','2':'xici代理抓取'}},
		{id:'retryNum',type:'text',name:'最大重试次数'},
		{id:'successNum',type:'text',name:'成功数'},
		{id:'failNum',type:'text',name:'失败数'},
		{id:'totalNum',type:'text',name:'总数'},
		{id:'state',type:'enum',name:'状态',map:{'1':'未开始','2':'进行中','3':'失败','4':'成功'}},
		{id:'stateMessage',type:'text',name:'状态描述'},
		{id:'createTime',type:'text',name:'创建时间'},
		{id:'modifyTime',type:'text',name:'修改时间'},
	],
	queryColumn:['type','state'],
	operate:[
	{
		name:'查看详细爬取进度',
		click:function(data){
			location.href = 'indexCrawl.html?brushTaskId='+data.brushTaskId;
		}
	}],
	button:[
	{
		name:'添加刷榜',
		click:function(){
			location.href = 'add.html';
		}
	}
	],
});
