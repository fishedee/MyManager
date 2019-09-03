webpackJsonp([20],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	input.verticalInput({
		id: 'container',
		field: [{ id: 'oldPassword', type: 'password', name: '旧密码' }, { id: 'newPassword1', type: 'password', name: '新密码' }, { id: 'newPassword2', type: 'password', name: '再输入一次新密码' }],
		submit: function (data) {
			if (data.newPassword1 != data.newPassword2) {
				dialog.message('两次输入密码不一致');
				return;
			}
			data = {
				oldPassword: data.oldPassword,
				newPassword: data.newPassword1
			};
			$.post('/user/modMyPassword', data, function (data) {
				data = $.JSON.parse(data);
				if (data.code != 0) {
					dialog.message(data.msg);
					return;
				}
				dialog.message('修改密码成功', function () {
					location.href = 'index.html';
				});
			});
		},
		cancel: function () {
			location.href = 'index.html';
		}
	});

/***/ }
]);