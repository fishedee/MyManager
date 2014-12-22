define('fishstrap/ui/query.js', function(require, exports, module){

var $ = require('fishstrap/core/global.js');
var dialog = require('fishstrap/ui/dialog.js');
var table = require('fishstrap/ui/table.js');
var input = require('fishstrap/ui/input.js');
module.exports = {
	simpleQuery:function( option ){
		//处理option
		var defaultOption = {
			id:'',
			column:[],
			queryColumn:[],
			params:{},
			operate:[],
			checkAll:false,
			url:'',
			button:[],
		};
		for( var i in option )
			defaultOption[i] = option[i];
		//生成基本框架
		var formId = $.uniqueNum();
		var tableId = $.uniqueNum();
		var buttonListId = $.uniqueNum();
		var target = $('#'+defaultOption.id);
		var div = $(
		'<form id="'+formId+'" class="form-inline m10"></form>'+
		'<div class="m10"><div id="'+buttonListId+'"></div><div id="'+tableId+'"></div></div>');
		target.empty();
		target.append(div);
		//生成staticTable框架
		table.staticTable({
			id:tableId,
			params:defaultOption.params,
			column:defaultOption.column,
			operate:defaultOption.operate,
			checkAll:defaultOption.checkAll,
			url:defaultOption.url
		});
		//生成flowInput框架
		var field = [];
		for( var i in defaultOption.queryColumn ){
			var param = defaultOption.queryColumn[i];
			var columnInfo;
			for( var j in defaultOption.column ){
				var column = defaultOption.column[j];
				if( column.id == param ){
					columnInfo = column;
					break;
				}
			}
			if( columnInfo.type == 'enum'){
				columnInfo.map = $.extend({
					"":"请选择"
				},columnInfo.map);
			}
			field.push(columnInfo);
		}
		input.flowInput({
			id:formId,
			field:field,
			submit:function(data){
				table.staticTable({
					id:tableId,
					params:$.extend(defaultOption.params,data),
					column:defaultOption.column,
					operate:defaultOption.operate,
					checkAll:defaultOption.checkAll,
					url:defaultOption.url
				});
			}
		});
		//生成按钮框架
		for( var i in defaultOption.button ){
			var button = defaultOption.button[i];
			(function(button){
				var div = $('<button class="btn">'+
					button.name+'</button>');
				div.click(button.click);
				$('#'+buttonListId).append(div);
			})(button);
		}
		
	},
};

});
;define('fishstrap/ui/input.js', function(require, exports, module){

var $ = require('fishstrap/core/global.js');
var dialog = require('fishstrap/ui/dialog.js');
var editor = require('fishstrap/ui/editor.js');
var upload = require('fishstrap/util/upload.js');
module.exports = {
	flowInput:function( option ){
		//处理option
		var defaultOption = {
			id:'',
			field:[],
			submit:function(){
			},
		};
		for( var i in option )
			defaultOption[i] = option[i];
		if( defaultOption.field.length == 0 )
			return;
		var div = "";
		//基本框架
		for( var i in defaultOption.field ){
			var field = defaultOption.field[i];
			if( field.type == 'text'){
				div += '<span>&nbsp;'+field.name+'：</span>' + '<input type="text" name="'+field.id+'" class="input-small"/>';
			}else if( field.type == 'enum'){
				var option = "";
				if( typeof field.map[""] != 'undefined')
					option += '<option value="">'+field.map[""]+'</option>';
				for( var j in field.map ){
					if( j != "")
						option += '<option value="'+j+'">'+field.map[j]+'</option>';
				}
				div += '<span>&nbsp;'+field.name+'：</span>' + '<select name="'+field.id+'">'+option+'</select>';
			}
		}
		div += '<button type="button" class="btn query">查询</button>'+
			'<button type="reset" class="btn">重置</button>';
		div = $(div);
		//加入页面
		$('#'+defaultOption.id).append(div);
		//挂载事件
		$('#'+defaultOption.id).find('.query').click(function(){
			var data = {};
			for( var i in defaultOption.field ){
				var field = defaultOption.field[i];
				if( field.type == 'text'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('input[name='+field.id+']').val());
				}else if( field.type == 'enum'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('select[name='+field.id+']').val());
				}
			}
			defaultOption.submit(data);
		});
	},
	verticalInput:function( option ){
		//处理option
		var defaultOption = {
			id:'',
			field:[],
			value:{},
			submit:function(){
			},
			cancel:function(){
			}
		};
		for( var i in option )
			defaultOption[i] = option[i];
		
		//执行业务逻辑
		var div = "";
		var contentDiv = "";
		for( var i in defaultOption.field ){
			var field = defaultOption.field[i];
			contentDiv += 
				'<tr>'+
					'<td class="tableleft">'+field.name+'</td>';
			if( typeof field.targetId != 'undefined'){
				contentDiv += '<td id="'+field.targetId+'">';
			}else{
				contentDiv += '<td>';
			}
			if( field.type == 'read'){
				contentDiv += '<div name="'+field.id+'"/>';
			}else if( field.type == 'fullEditor'){
				field.editorTargetId = $.uniqueNum();
				contentDiv += '<div name="'+field.id+'" id="'+field.editorTargetId+'"/>';
			}else if( field.type == 'simpleEditor'){
				field.editorTargetId = $.uniqueNum();
				contentDiv += '<div name="'+field.id+'" id="'+field.editorTargetId+'"/>';
			}else if( field.type == 'image'){
				field.imageTargetId = $.uniqueNum();
				field.imageProgressTargetId = $.uniqueNum();
				contentDiv += '<div><img name="'+field.id+'" src=""/></div>'+
					'<div class="progress"><div class="bar" id="'+field.imageProgressTargetId+'"></div></div>'+
					'<div class="btn" id="'+field.imageTargetId+'"><span>点击这里上传图片</span></div>';
			}else if( field.type == 'area'){
				contentDiv += '<textarea name="'+field.id+'" style="width:90%;height:300px;"></textarea>';
			}else if( field.type == 'text'){
				contentDiv += '<input type="text" name="'+field.id+'"/>';
			}else if( field.type == 'password'){
				contentDiv += '<input type="password" name="'+field.id+'"/>';
			}else if( field.type == 'enum'){
				var option = "";
				for( var j in field.map ){
					option += '<option value="'+j+'">'+field.map[j]+'</option>';
				}
				contentDiv += '<select name="'+field.id+'">'+option+'</select>';
			}
			contentDiv +=
				'</td>'+
			'</tr>';
		}
		div += '<table class="table table-bordered table-hover definewidth m10">'+
			contentDiv+
			'<tr>'+
				'<td class="tableleft"></td>'+
				'<td>'+
					'<button type="button" class="btn btn-primary submit" >提交</button>'+
					'<button type="button" class="btn btn-success cancel">返回列表</button>'+
				'</td>'+
			'</tr>'+
		'</table>';
		div = $(div);
		//插入到页面中
		$('#'+defaultOption.id).append(div);
		//挂载控件事件
		for( var i in defaultOption.field ){
			var field = defaultOption.field[i];
			(function(field){
				if( field.type == 'image'){
					upload.image({
						url:field.url,
						target:field.imageTargetId,
						field:'data',
						width:2048,
						quality:0.8,
						onProgress:function(progress){
							$.console.log(progress);
							$('#'+field.imageProgressTargetId).text(progress+'%');
							$('#'+field.imageProgressTargetId).css('width',progress+'%');
						},
						onSuccess:function(data){
							data = $.JSON.parse(data);
							if( data.code != 0 ){
								dialog.message(data.msg);
								return;
							}
							div.find('img[name='+field.id+']').attr('src',data.data);
						},
						onFail:function(msg){
							dialog.message(msg);
						}
					});
				}else if( field.type == 'simpleEditor'){
					field._editor = editor.simpleEditor({
						id:field.editorTargetId
					});
				}else if( field.type == 'fullEditor'){
					field._editor = editor.fullEditor({
						id:field.editorTargetId
					});
				}
			})(field);
		}
		//设置value
		for( var i in defaultOption.field ){
			var field = defaultOption.field[i];
			if( typeof defaultOption.value[field.id] == 'undefined' )
				continue;
			if( field.type == 'read')
				div.find('div[name='+field.id+']').text(defaultOption.value[field.id]);
			else if( field.type == 'fullEditor')
				field._editor.setContent(defaultOption.value[field.id]);
			else if( field.type == 'simpleEditor')
				field._editor.setFormatData(defaultOption.value[field.id]);
			else if( field.type == 'image')
				div.find('img[name='+field.id+']').attr("src",defaultOption.value[field.id]);
			else if( field.type == 'area')
				div.find('textarea[name='+field.id+']').val(defaultOption.value[field.id]);
			else if( field.type == 'text' || field.type == 'password')
				div.find('input[name='+field.id+']').val(defaultOption.value[field.id]);
			else if( field.type == 'enum')
				div.find('select[name='+field.id+']').val(defaultOption.value[field.id]);
		}
		//挂载事件
		div.find('.submit').click(function(){
			var data = {};
			for( var i in defaultOption.field ){
				var field = defaultOption.field[i];
				if( field.type == 'read'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('div[name='+field.id+']').text());
				}else if( field.type == 'simpleEditor'){
					data[field.id] = field._editor.getFormatData();
				}else if( field.type == 'fullEditor'){
					data[field.id] = field._editor.getContent();
				}else if( field.type == 'image'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('img[name='+field.id+']').attr("src"));
				}else if( field.type == 'area'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('textarea[name='+field.id+']').val());
				}else if( field.type == 'text' || field.type == 'password'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('input[name='+field.id+']').val());
				}else if( field.type == 'enum'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('select[name='+field.id+']').val());
				}
			}
			defaultOption.submit(data);
		});
		div.find('.cancel').click(defaultOption.cancel);
	}
};

});
;define('fishstrap/ui/table.js', function(require, exports, module){

/*
*加载griTable
*@require fishstrap/lib/gri/griTable.js
*/
var $ = require('fishstrap/core/global.js');
var dialog = require('fishstrap/ui/dialog.js');
module.exports = {
	staticSimpleTable:function(option){
		//执行Option
		var defaultOption = {
			id:'',
			data:'',
			column:[],
			operate:[]
		};
		for( var i in option )
			defaultOption[i] = option[i];
		//显示数据
		function showData(data){
			var div = '';
			var operateDiv = '';
			div += '<div class="mod-basic">';
			div += '<table class="mod_table" style="table-layout: auto;">';
			//构造操作
			for( var i in defaultOption.operate ){
				defaultOption.operate[i].id = $.uniqueNum();
				operateDiv += "<a href='#' class=operate_"+defaultOption.operate[i].id
					+">"+defaultOption.operate[i].name+"</a>&nbsp;";
			}
			//显示列表头数据
			div += '<thead><tr>';
			for( var i in defaultOption.column ){
				var column = defaultOption.column[i];
				var style = '';
				if( column.type == 'hidden')
					style = 'style="display:none;"';
				div += '<th '+style+' ><span class="label">'+column.name+'</span></th>';
			}
			if( defaultOption.operate.length != 0 ){
				div += '<th><span class="label">操作</span></th>';
			}
			div += '</tr></thead>';
			//显示列表身数据
			div += '<tbody>';
			for( var i in data ){
				var item = data[i];
				div += '<tr>';
				for( var j in defaultOption.column ){
					var column = defaultOption.column[j];
					var style = '';
					if( column.type == 'hidden')
						style = 'style="display:none;"';
					div += '<td '+style+' class="'+column.id+'">'+item[column.id]+'</td>';
				}
				if( defaultOption.operate.length != 0 ){
					div += '<td>'+operateDiv+'</td>';
				}
				div += '</tr>';
			}
			div += '</tbody>';
			div += '</table>';
			div += '</div>';
			div = $(div);
			$('#'+defaultOption.id).append(div);
			//挂载事件
			for( var i in defaultOption.operate ){
				(function(i){
					$('.operate_'+defaultOption.operate[i].id).unbind("click").click(function(){
						var tr = $(this);
						while( tr.is('tr') == false )
							tr = tr.parent();
						var data = {};
						tr.find('td').each(function(){
							data[$(this).attr('class')] = $(this).text();
						});
						defaultOption.operate[i].click(data);
					});
				})(i);
			}
		}
		showData(defaultOption.data);
	},
	staticTable:function(option){
		//执行Option
		var defaultOption = {
			id:'',
			params:{},
			column:[],
			operate:[],
			checkAll:false,
			url:null,
		};
		for( var i in option )
			defaultOption[i] = option[i];
			
		//执行_option
		var sendUrl = '';
		var _option = {};
		_option.table_id    = defaultOption.id;
		_option.key_index   = '_id';
		_option.order_field = '_id';
		_option.order_type  = 'asc';
		_option.page_size   = 10;
		_option.container_class = "mod-basic box-table";
		_option.table_class = "mod_table";
		_option.fields      = '';
		_option.summary     = '';
		_option.ifRealPage  = true;
		_option.data        = 'data';
		_option.checkAll    = false;
		_option.params      = {};

		//拼接url
		var paramStr = "";
		for(var i in defaultOption.params) {
			if( $.trim(defaultOption.params[i]) != "")
				paramStr += (i+"="+encodeURIComponent($.trim(defaultOption.params[i])) + "&");
		}
		
		paramStr += ("t=" + new Date().getTime());
		sendUrl = defaultOption.url + "?" + paramStr;

		//拼接列
		_option.fields = {};
		for(var i in defaultOption.column){
			var column = defaultOption.column[i];
			var single;
			(function(column){
				if( column.type == 'text'){
					single = {
						thText:column.name
					};
				}else if( column.type == 'enum'){
					single = {
						thText:column.name,
						format:function(data){
							return column.map[data];
						}
					};
				}
			})(column);
			_option.fields[column.id] = single;
		}
		var info = "";
		for( var i in defaultOption.operate ){
			defaultOption.operate[i].id = $.uniqueNum();
			info += "<a href='#' class=operate_"+defaultOption.operate[i].id
				+">"+defaultOption.operate[i].name+"</a>&nbsp;";
		}
		if( defaultOption.operate.length != 0 ){
			_option.fields['_oper'] = {
				thText:'操作',
				format:function(){
					return info;
				}
			};
		}
		//GRITABLE
		$.get(sendUrl+"&pageIndex=0&pageSize=10", {}, function (result) {
			result = $.JSON.parse(result);
			if( result.code != 0 ){
				dialog.message(result.msg);
				return;
			}
			dt = GRI.initDataTable({
				resultObj: result,
				name: _option.data,
				tableId: _option.table_id,
				data: result.data[_option.data],
				summary: _option.summary,
				allFields: _option.fields,
				layout: 'auto',
				checkAll: _option.checkAll,
				enableThClick: true,  
				stopToday: false,
				keyIndex: _option.key_index,
				page:{
					orderField: _option.order_field,
					orderType: _option.order_type,
					ifRealPage: _option.ifRealPage,
					size: _option.page_size,
					rowCount: result.data.count,
					index: 0,
					url:sendUrl
				},
				cssSetting:{
					containerClass: _option.container_class,
					tableClass: _option.table_class
				},
				callback: function(data){
					for( var i in defaultOption.operate ){
						(function(i){
							$('.operate_'+defaultOption.operate[i].id).unbind("click").click(function(){
								var tr = $(this);
								while( tr.is('tr') == false )
									tr = tr.parent();
								var data = {};
								tr.find('td').each(function(){
									data[$(this).attr('class')] = $(this).text();
								});
								defaultOption.operate[i].click(data);
							});
						})(i);
					}
				}
			});
		});
	},
	
};

});
;define('fishstrap/ui/editor.js', function(require, exports, module){

/*
*加载依赖的ueditor
*@require fishstrap/lib/uedit/ueditor.config.js
*@require fishstrap/lib/uedit/ueditor.all.min.js
*/
var $ = require('fishstrap/core/global.js');
module.exports = {
	fullEditor:function( option ){
		//处理option
		var defaultOption =	{
			id:'',
			url:'/uedit/control',
		};
		for( var i in option )
			defaultOption[i] = option[i];
		var editorId = $.uniqueNum();
		$('#'+defaultOption.id).html(
			'<script id="'+editorId+'" name="content" type="text/plain">'+
			'</script>'
		);
		//初始化editor
		var ue = UE.getEditor(editorId, {
			serverUrl:defaultOption.url,
			autoHeightEnabled: true,
			toolbars: [
			['fullscreen', 'source', '|', 'undo', 'redo'] ,
			['bold', 'italic', 'underline', 'fontborder', 'strikethrough', 'superscript', 'subscript', 'removeformat', 'formatmatch', 'autotypeset', 'blockquote', 'pasteplain', '|', 'forecolor', 'backcolor', 'insertorderedlist', 'insertunorderedlist', 'selectall', 'cleardoc','|','rowspacingtop', 'rowspacingbottom', 'lineheight', '|','customstyle', 'paragraph', 'fontfamily', 'fontsize', '|',
				'directionalityltr', 'directionalityrtl', 'indent', '|',
				'simpleupload', 'insertimage','pagebreak', 'template', 'background'],
			['justifyleft', 'justifycenter', 'justifyright', 'justifyjustify', '|', 'touppercase', 'tolowercase', '|',
				'link', 'unlink', 'anchor', '|', 'imagenone', 'imageleft', 'imageright', 'imagecenter'],
			['inserttable', 'deletetable', 'insertparagraphbeforetable', 'insertrow', 'deleterow', 'insertcol', 'deletecol', 'mergecells', 'mergeright', 'mergedown', 'splittocells', 'splittorows', 'splittocols', 'charts']],
		});
		//计算editor的函数
		return {
			getContent:function(){
				return ue.getContent();
			},
			setContent:function(content){
				ue.ready(function(){
					ue.setContent(content);
				});
			},
		};
	},
	simpleEditor:function(option){
		//处理option
		var defaultOption =	{
			id:'',
			url:'/uedit/control',
		};
		for( var i in option )
			defaultOption[i] = option[i];
		var editorId = $.uniqueNum();
		$('#'+defaultOption.id).html(
			'<script id="'+editorId+'" name="content" type="text/plain">'+
			'</script>'
		);
		//初始化editor
		var ue = UE.getEditor(editorId, {
			serverUrl:defaultOption.url,
			autoHeightEnabled: true,
			toolbars: [['fullscreen','simpleupload']],
		});
		//计算editor的函数
		return {
			getContent:function(){
				return ue.getContent();
			},
			getFormatData:function(){
				var content = ue.getContent();

				var div = $(content);
				var data = [];
				div.each(function(){
					var self = $(this);
					if( self.is('p') == true ){
						var hasImg = self.has('img');
						var hasText = self.text();
						if( hasImg && ! hasText ){
							self.find('img').each(function(){
								data.push({
									type:1,
									data:$(this).attr('src'),
								});
							});
						}else if( hasImg && hasText ){
							data.push({
								type:0,
								data:self.text(),
							});
							self.find('img').each(function(){
								data.push({
									type:1,
									data:$(this).attr('src'),
								});
							});
						}else{
							data.push({
								type:0,
								data:self.text(),
							});
						}
					}else if( self.is('img') == true ){
						data.push({
							type:1,
							data:self.attr('src'),
						});
					}
				});
				return data;
			},
			setContent:function(content){
				ue.ready(function(){
					ue.setContent(content);
				});
			},
			setFormatData:function(data){
				var content = "";
				for( var i in data ){
					var singleData = data[i];
					if( singleData.type == 0 )
						content += '<p>'+singleData.data+'</p>';
					else if( singleData.type == 1 )
						content += '<p><img src="'+singleData.data+'"/></p>';
				}
				ue.ready(function(){

					ue.setContent(content);
				});
			}
		};
	}
};

});