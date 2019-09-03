webpackJsonp([10],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var dialog = __webpack_require__(32);
	var input = __webpack_require__(39);
	var subPage = __webpack_require__(49);
	var syncType = { 1: '增量转换', 2: '全量转换' };
	function go() {
		input.verticalInput({
			id: 'container',
			field: [{ id: 'gitUrl', type: 'text', name: 'git地址(https)' }, { id: 'accessToken', type: 'text', name: 'csdn访问授权(例如：user,pass)' }, { id: 'syncType', type: 'enum', name: '同步类型', map: syncType }],
			submit: function (data) {
				$.post('/blog/addTask', data, function (data) {
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
	}
	go();

/***/ }
]);