var $ = require('fishfront/ui/global');
var dialog = require('fishfront/ui/dialog');
var table = require('fishfront/ui/table');
var chart = require('fishfront/ui/chart');
var year = $.location.getQueryArgv('year');
var week = $.location.getQueryArgv('week');
var type = $.location.getQueryArgv('type');
var statistic = {};
$('body').prepend('<div id="chart" style="height:400px"></div>');
function getData( nextStep ){
	$.get('/account/getWeekDetailTypeStatistic',{year:year,week:week,type:type},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 ){
			dialog.message(data.msg);
			return;
		}
		statistic = data.data;
		nextStep();
	});
}
function go(){
	//扇形图
	chart.simpleSector({
		id:'chart',
		data:statistic,
		xAxis:'categoryName',
		yAxis:'money',
	});
	
	//表格
	var total = {};
	total.precent = '100%';
	total.categoryName = '合计';
	total.money = 0;
	for( var i in statistic )
		total.money = total.money + parseInt(statistic[i].money);
	statistic.push(total);
	table.staticSimpleTable({
		id:'container',
		data:statistic,
		column:[
			{id:'categoryName',type:'text',name:'分类'},
			{id:'money',type:'text',name:'金额'},
			{id:'precent',type:'text',name:'占比'},
		],
		operate:[],
	});
}
getData(function(){
	go();
});
