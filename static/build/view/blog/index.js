webpackJsonp([8],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var query = __webpack_require__(36);
	var state = { 1: '未开始', 2: '进行中', 3: '失败', 4: '成功' };
	var syncType = { 1: '增量转换', 2: '全量转换' };
	function go() {
		query.simpleQuery({
			id: 'container',
			url: '/blog/searchTask',
			column: [{ id: 'blogSyncId', type: 'text', name: '同步ID' }, { id: 'gitUrl', type: 'text', name: 'git地址' }, { id: 'accessToken', type: 'text', name: 'csdn访问授权' }, { id: 'syncType', type: 'enum', name: '同步类型', map: syncType }, { id: 'state', type: 'enum', name: '状态', map: state }, { id: 'stateMessage', type: 'text', name: '状态描述' }, { id: 'createTime', type: 'text', name: '创建时间' }, { id: 'modifyTime', type: 'text', name: '修改时间' }],
			queryColumn: ['gitUrl', 'syncType', 'state'],
			operate: [{
				name: '重新开始',
				click: function (data) {
					$.post('/blog/restartTask', { blogSyncId: data.blogSyncId }, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						location.href = 'index.html';
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