webpackJsonp([23],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	input.verticalInput({
		id: 'container',
		field: [{ id: 'name', type: 'text', name: '姓名' }, { id: 'password', type: 'text', name: '密码' }, { id: 'type', type: 'enum', name: '类型', map: { '1': '管理员', '2': '普通会员' } }],
		submit: function (data) {
			$.post('/user/add', data, function (data) {
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

/***/ }
]);