webpackJsonp([16],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	var cardId = $.location.getQueryArgv('cardId');
	function go(value) {
		input.verticalInput({
			id: 'container',
			field: [{ id: 'name', type: 'text', name: '名称' }, { id: 'bank', type: 'text', name: '银行' }, { id: 'card', type: 'text', name: '卡号' }, { id: 'money', type: 'text', name: '初始金额' }, { id: 'remark', type: 'text', name: '备注' }],
			value: value,
			submit: function (data) {
				if (cardId != null) {
					data = $.extend({ cardId: cardId }, data);
					$.post('/card/mod', data, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						history.back();
					});
				} else {
					$.post('/card/add', data, function (data) {
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
	if (cardId != null) {
		$.get('/card/get', { cardId: cardId }, function (data) {
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