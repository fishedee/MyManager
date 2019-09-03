webpackJsonp([24],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var query = __webpack_require__(36);
	var dialog = __webpack_require__(32);
	query.simpleQuery({
		id: 'container',
		url: '/user/search',
		column: [{ id: 'userId', type: 'text', name: '用户ID' }, { id: 'name', type: 'text', name: '姓名' }, { id: 'type', type: 'enum', name: '类型', map: { '1': '管理员', '2': '普通会员' } }, { id: 'createTime', type: 'text', name: '创建时间' }, { id: 'modifyTime', type: 'text', name: '修改时间' }],
		queryColumn: ['name', 'type'],
		operate: [{
			name: '修改类型',
			click: function (data) {
				location.href = 'modType.html?userId=' + data.userId;
			}
		}, {
			name: '修改密码',
			click: function (data) {
				location.href = 'modPassword.html?userId=' + data.userId;
			}
		}, {
			name: '删除',
			click: function (data) {
				dialog.confirm('确认删除该用户，不可回退操作？!', function () {
					$.post('/user/del', { userId: data.userId }, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						location.href = 'index.html';
					});
				});
			}
		}],
		button: [{
			name: '添加用户',
			click: function () {
				location.href = 'add.html';
			}
		}]
	});

/***/ }
]);