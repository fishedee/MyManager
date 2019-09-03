webpackJsonp([11],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var input = __webpack_require__(39);
	var dialog = __webpack_require__(32);
	var blogSyncAutoId = $.location.getQueryArgv('blogSyncAutoId');
	function go(value) {
		input.verticalInput({
			id: 'container',
			field: [{ id: 'gitUrl', type: 'text', name: 'git地址' }, { id: 'accessToken', type: 'text', name: 'csdn访问授权' }],
			value: value,
			submit: function (data) {
				if (blogSyncAutoId != null) {
					data = $.extend({ blogSyncAutoId: blogSyncAutoId }, data);
					$.post('/blog/modAuto', data, function (data) {
						data = $.JSON.parse(data);
						if (data.code != 0) {
							dialog.message(data.msg);
							return;
						}
						history.back();
					});
				} else {
					$.post('/blog/addAuto', data, function (data) {
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
	if (blogSyncAutoId != null) {
		$.get('/blog/getAuto', { blogSyncAutoId: blogSyncAutoId }, function (data) {
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