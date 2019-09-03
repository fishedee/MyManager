webpackJsonp([17],[
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	var $ = __webpack_require__(6);
	var query = __webpack_require__(36);
	var dialog = __webpack_require__(32);
	query.simpleQuery({
		id: 'container',
		url: '/category/search',
		column: [{ id: 'categoryId', type: 'text', name: '分类ID' }, { id: 'name', type: 'text', name: '名字' }, { id: 'remark', type: 'text', name: '备注' }, { id: 'createTime', type: 'text', name: '创建时间' }, { id: 'modifyTime', type: 'text', name: '修改时间' }],
		queryColumn: ['name', 'remark'],
		operate: [{
			name: '编辑',
			click: function (data) {
				location.href = 'view.html?categoryId=' + data.categoryId;
			}
		}, {
			name: '删除',
			click: function (data) {
				dialog.confirm('确认删除该分类，不可回退操作？!', function () {
					$.post('/category/del', { categoryId: data.categoryId }, function (data) {
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
			name: '添加分类',
			click: function () {
				location.href = 'view.html';
			}
		}]
	});

/***/ }
]);