webpackJsonp([18],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	var categoryId = $.location.getQueryArgv('categoryId');
	function go(value) {
		input.verticalInput({
			id: 'container',
			field: [{ id: 'name', type: 'text', name: '名字' }, { id: 'remark', type: 'text', name: '备注' }],
			value: value,
			submit: function (data) {
				if (categoryId != null) {
					data = $.extend({ categoryId: categoryId }, data);
					$.post('/category/mod', data, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						history.back();
					});
				} else {
					$.post('/category/add', data, function (data) {
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
	if (categoryId != null) {
		$.get('/category/get', { categoryId: categoryId }, function (data) {
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