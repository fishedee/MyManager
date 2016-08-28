var $ = require('fishfront/ui/global');
var query = require('fishfront/ui/query');
var dialog = require('fishfront/ui/dialog');
var brushTaskId = $.location.getQueryArgv('brushTaskId');
query.simpleQuery({
	id:'container',
	url:'/brush/searchCrawl?brushTaskId='+brushTaskId,
	column:[
		{id:'brushCrawlId',type:'text',name:'爬取ID'},
		{id:'brushTaskId',type:'text',name:'任务ID'},
		{id:'proxy',type:'text',name:'代理地址'},
		{id:'retryNum',type:'text',name:'重试次数'},
		{id:'state',type:'enum',name:'状态',map:{'1':'未开始','2':'进行中','3':'失败重试中','4':'失败','5':'成功'}},
		{id:'stateMessage',type:'text',name:'状态描述'},
		{id:'createTime',type:'text',name:'创建时间'},
		{id:'modifyTime',type:'text',name:'修改时间'},
	],
	queryColumn:['state'],
	operate:[],
	button:[
	{
		name:'返回',
		click:function(){
			history.back();
		}
	}
	],
});
