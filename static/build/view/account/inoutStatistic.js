webpackJsonp([3],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var table = __webpack_require__(37);
	var chart = __webpack_require__(53);
	var statistic = {};
	$('body').prepend('<div id="chart" style="height:400px"></div>');
	function getData(nextStep) {
		$.get('/account/getWeekTypeStatistic', {}, function (data) {
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
		table.staticSimpleTable({
			id: 'container',
			data: statistic,
			column: [{ id: 'name', type: 'text', name: '时间' }, { id: 'typeName', type: 'text', name: '类型' }, { id: 'money', type: 'text', name: '金额' }, { id: 'year', type: 'hidden', name: '年份' }, { id: 'week', type: 'hidden', name: '周份' }, { id: 'type', type: 'hidden', name: '类型' }],
			operate: [{
				name: '详细信息',
				click: function (data) {
					location.href = 'viewInoutStatistic.html?year=' + data.year + '&week=' + data.week + '&type=' + data.type;
				}
			}]
		});
		chart.simpleBrokeLine({
			id: 'chart',
			data: statistic,
			xAxis: 'name',
			yAxis: 'money',
			category: 'typeName'
		});
	}
	getData(function () {
		go();
	});

/***/ }
]);