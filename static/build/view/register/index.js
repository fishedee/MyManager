webpackJsonp([21],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var query = __webpack_require__(36);
	var needDealType = { 1: '不需要', 2: '需要' };
	var haveDealType = { 1: '未挂号', 2: '已挂号' };
	function go() {
		query.simpleQuery({
			id: 'container',
			url: '/register/search',
			column: [{ id: 'registerId', type: 'text', name: '挂号ID' }, { id: 'name', type: 'text', name: '姓名' }, { id: 'beginTime', type: 'text', name: '开始时间' }, { id: 'endTime', type: 'text', name: '结束时间' }, { id: 'mail', type: 'text', name: '提醒邮箱' }, { id: 'needDealType', type: 'enum', name: '是否需要自动挂号', map: needDealType }, { id: 'haveDealType', type: 'enum', name: '是否已经自动挂号', map: haveDealType }, { id: 'haveDealResult', type: 'text', name: '自动挂号结果' }, { id: 'createTime', type: 'text', name: '创建时间' }, { id: 'modifyTime', type: 'text', name: '修改时间' }],
			queryColumn: ['name'],
			operate: [{
				name: '编辑',
				click: function (data) {
					location.href = 'view.html?registerId=' + data.registerId;
				}
			}, {
				name: '删除',
				click: function (data) {
					$.post('/register/del', { registerId: data.registerId }, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						location.reload();
					});
				}
			}],
			button: [{
				name: '添加任务',
				click: function () {
					location.href = 'view.html';
				}
			}]
		});
	}
	go();

/***/ }
]);