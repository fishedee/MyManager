var $ = require('fishfront/ui/global');
var dialog = require('fishfront/ui/dialog');
var query = require('fishfront/ui/query');
function go(){
	query.simpleQuery({
		id:'container',
		url:'/blog/searchAuto',
		column:[
			{id:'blogSyncAutoId',type:'text',name:'自动同步ID'},
			{id:'gitUrl',type:'text',name:'git地址'},
			{id:'accessToken',type:'text',name:'csdn访问授权'},
			{id:'createTime',type:'text',name:'创建时间'},
			{id:'modifyTime',type:'text',name:'修改时间'},
		],
		queryColumn:['gitUrl'],
		operate:[
		{
			name:'删除',
			click:function(data){
				$.post('/blog/delAuto',{blogSyncAutoId:data.blogSyncAutoId},function(data){
					data = $.JSON.parse(data);
					if( data.code != 0 ){
						dialog.message(data.msg);
						return;
					}
					location.href = 'indexAuto.html';
				});
			}
		},
		{
			name:'查看',
			click:function(data){
				location.href = 'viewAuto.html?blogSyncAutoId='+data.blogSyncAutoId;
			}
		}],
		button:[
		{
			name:'添加自动同步',
			click:function(){
				location.href = 'viewAuto.html';
			}
		}
		],
	});
}
go();
