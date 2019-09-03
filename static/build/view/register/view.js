webpackJsonp([22],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	var registerId = $.location.getQueryArgv('registerId');
	var needDealType = { 1: '不需要', 2: '需要' };
	var haveDealType = { 1: '未挂号', 2: '已挂号' };
	function go(value) {
		input.verticalInput({
			id: 'container',
			field: [{ id: 'name', type: 'text', name: '姓名' }, { id: 'beginTime', type: 'time', name: '开始时间' }, { id: 'endTime', type: 'time', name: '结束时间' }, { id: 'mail', type: 'text', name: '提醒邮箱' }, { id: 'needDealType', type: 'enum', name: '是否需要自动挂号', map: needDealType }],
			value: value,
			submit: function (data) {
				if (registerId != null) {
					data = $.extend({ registerId: registerId }, data);
					$.post('/register/mod', data, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						history.back();
					});
				} else {
					$.post('/register/add', data, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						history.back();
					});
				}
			},
			cancel: function () {
				history.back();
			}
		});
	}
	if (registerId != null) {
		$.get('/register/get', { registerId: registerId }, function (data) {
			data = $.JSON.parse(data);
			if (data.code != 0) {
				dialog.message(data.msg);
				return;
			}
			go(data.data);
		});
	} else {
		go({});
	}

/***/ }
]);