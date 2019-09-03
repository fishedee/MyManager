webpackJsonp([4],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var table = __webpack_require__(37);
	var chart = __webpack_require__(53);
	var statistic = {};
	$('body').prepend('<div id="chart" style="height:400px"></div>');
	function getData(nextStep) {
		$.get('/account/getWeekCardStatistic', {}, function (data) {
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
			column: [{ id: 'name', type: 'text', name: '时间' }, { id: 'cardName', type: 'text', name: '银行卡' }, { id: 'money', type: 'text', name: '结余金额' }, { id: 'cardId', type: 'hidden', name: '银行卡ID' }, { id: 'year', type: 'hidden', name: '年份' }, { id: 'week', type: 'hidden', name: '周份' }],
			operate: [{
				name: '详细信息',
				click: function (data) {
					location.href = 'viewTotalStatistic.html?year=' + data.year + '&week=' + data.week + '&cardId=' + data.cardId;
				}
			}]
		});
		chart.simpleBrokeLine({
			id: 'chart',
			data: statistic,
			xAxis: 'name',
			yAxis: 'money',
			category: 'cardName'
		});
	}
	getData(function () {
		go();
	});

/***/ }
]);