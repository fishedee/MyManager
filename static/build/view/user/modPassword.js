webpackJsonp([25],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	var userId = $.location.getQueryArgv('userId');
	$.get('/user/get', { userId: userId }, function (data) {
		data = $.JSON.parse(data);
		if (data.code != 0) {
			dialog.message(data.msg);
			return;
		}
		data = data.data;
		input.verticalInput({
			id: 'container',
			field: [{ id: 'name', type: 'read', name: '姓名' }, { id: 'password', type: 'text', name: '新密码' }],
			value: {
				name: data.name
			},
			submit: function (data) {
				data = $.extend({ userId: userId }, { password: data.password });
				$.post('/user/modPassword', data, function (data) {
					data = $.JSON.parse(data);
					if (data.code != 0) {
						dialog.message(data.msg);
						return;
					}
					history.back();
				});
			},
			cancel: function () {
				history.back();
			}
		});
	});

/***/ }
]);