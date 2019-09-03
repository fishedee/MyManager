webpackJsonp([7],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var table = __webpack_require__(37);
	var chart = __webpack_require__(53);
	var year = $.location.getQueryArgv('year');
	var week = $.location.getQueryArgv('week');
	var cardId = $.location.getQueryArgv('cardId');
	var statistic = {};
	$('body').prepend('<div id="chart" style="height:400px"></div>');
	function getData(nextStep) {
		$.get('/account/getWeekDetailCardStatistic', { year: year, week: week, cardId: cardId }, function (data) {
			data = $.JSON.parse(data);
			if (data.code != 0) {
				dialog.message(data.msg);
				return;
			}
			statistic = data.data;
			nextStep();
		});
	}
	function go() {
		//扇形图
		chart.simpleSector({
			id: 'chart',
			data: statistic,
			xAxis: 'typeName',
			yAxis: 'money'
		});

		//表格
		var total = {};
		total.precent = '';
		total.typeName = '合计挣钱';
		total.money = 0;
		for (var i in statistic) {
			if (statistic[i].type == 1 || statistic[i].type == 3) total.money = total.money + parseInt(statistic[i].money);else total.money = total.money - parseInt(statistic[i].money);
		}
		statistic.push(total);
		table.staticSimpleTable({
			id: 'container',
			data: statistic,
			column: [{ id: 'typeName', type: 'text', name: '类型' }, { id: 'money', type: 'text', name: '金额' }, { id: 'precent', type: 'text', name: '占比' }],
			operate: []
		});
	}
	getData(function () {
		go();
	});

/***/ }
]);