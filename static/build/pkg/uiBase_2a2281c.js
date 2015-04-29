define('fishstrap/ui/query.js', function(require, exports, module){ var $ = require('fishstrap/core/global.js');
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
		$.console.log($.JSON.stringify(option.column));
		//生成staticTable框架
		var tableOperation = table.staticTable({
			id:tableId,
			params:defaultOption.params,
			column:defaultOption.column,
			operate:defaultOption.operate,
			checkAll:defaultOption.checkAll,
			url:defaultOption.url
		});
		//生成flowInput框架
		var field = [];
		for( var i = 0 ; i != defaultOption.queryColumn.length; ++i ){
			var param = defaultOption.queryColumn[i];
			var columnInfo;
			for( var j = 0 ; j != defaultOption.column.length ; ++j ){
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
		for( var i = 0 ; i != defaultOption.button.length ; ++i ){
			var button = defaultOption.button[i];
			(function(button){
				var div = $('<button class="btn">'+
					button.name+'</button>');
				div.click(button.click);
				$('#'+buttonListId).append(div);
			})(button);
		}
		return tableOperation;
	},
}; });
;define('fishstrap/ui/input.js', function(require, exports, module){ var $ = require('fishstrap/core/global.js');
var dialog = require('fishstrap/ui/dialog.js');
var editor = require('fishstrap/ui/editor.js');
var table = require('fishstrap/ui/table.js');
var upload = require('fishstrap/util/upload.js');
var subPage = require('fishstrap/page/subPage.js');
require('fishstrap/util/jqueryDatetimePicker.js');
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
		for( var i = 0 ; i != defaultOption.field.length; ++i ){
			var field = defaultOption.field[i];
			if( field.type == 'text'){
				div += '<span>&nbsp;'+field.name+'：</span>' + '<input type="text" name="'+field.id+'" class="input-small"/>';
			}else if(field.type == 'time'){
				div += '<span>&nbsp;'+field.name+'：</span>' + '<input type="text" name="'+field.id+'" class="time input-small"/>';
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
		//挂载控件事件
		for( var i = 0 ; i != defaultOption.field.length ; ++i ){
			var field = defaultOption.field[i];
			if( field.type == 'time'){
				$('#'+defaultOption.id).find('input[name='+field.id+']').datetimepicker({
					lang:'ch',
					timepicker:false,
					format: 'Y-m-d',
					closeOnDateSelect:true
				});
			}
		}
		//挂载事件
		$('#'+defaultOption.id).find('.query').click(function(){
			var data = {};
			for( var i = 0 ; i != defaultOption.field.length; ++i ){
				var field = defaultOption.field[i];
				if( field.type == 'text'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('input[name='+field.id+']').val());
				}else if( field.type == 'time'){
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
		$.console.log($.JSON.stringify(defaultOption.field));
		for( var i = 0; i != defaultOption.field.length; ++i ){
			var field = defaultOption.field[i];
			contentDiv += 
				'<tr>'+
					'<td class="tableleft" style="width:20%;">'+field.name+'</td>';
			if( typeof field.targetId != 'undefined'){
				contentDiv += '<td id="'+field.targetId+'">';
			}else{
				contentDiv += '<td>';
			}
			if( field.type == 'read'){
				contentDiv += '<div name="'+field.id+'"/>';
			}else if(field.type == 'time'){
				contentDiv += '<input type="text" name="'+field.id+'" class="time input-small"/>';
			}else if( field.type == 'link'){
				contentDiv += '<a name="'+field.id+'" target="_blank"><div name="'+field.id+'"/></a>';
			}else if( field.type == 'fullEditor'){
				field.editorTargetId = $.uniqueNum();
				contentDiv += '<div class="alert alert-danger" role="alert">注意！为了保证微信内观看视频的兼容性，强烈建议你只插入腾讯视频网址，不要插入其它视频网站，或直接上传视频。</div><div name="'+field.id+'" id="'+field.editorTargetId+'"/>';
			}else if( field.type == 'simpleEditor'){
				field.editorTargetId = $.uniqueNum();
				contentDiv += '<div name="'+field.id+'" id="'+field.editorTargetId+'"/>';
			}else if( field.type == 'image'){
				field.imageTargetId = $.uniqueNum();
				field.imageProgressTargetId = $.uniqueNum();
				contentDiv += '<div><img name="'+field.id+'" src=""/></div>'+
					'<div class="progress"><div class="bar" id="'+field.imageProgressTargetId+'"></div></div>'+
					'<div class="btn" id="'+field.imageTargetId+'"><span>点击这里上传图片</span></div>';
			}else if( field.type == 'compressFile'){
				field.tableId = $.uniqueNum();
				field.fileTargetId = $.uniqueNum();
				field.fileProgressTargetId = $.uniqueNum();
				contentDiv += 
					'<div id="'+field.tableId+'"></div>'+
					'<div class="progress"><div class="bar" id="'+field.fileProgressTargetId+'"></div></div>'+
					'<div class="btn" id="'+field.fileTargetId+'"><span>点击这里上传压缩文件</span></div>';
			}else if( field.type == 'area'){
				var disabledDiv = '';
				if( _.isUndefined(field.disabled) == false &&  field.disabled === true )
					disabledDiv = 'disabled="true"';
				contentDiv += '<textarea name="'+field.id+'" style="width:90%;height:300px;" '+disabledDiv+'></textarea>';
			}else if( field.type == 'text'){
				contentDiv += '<input type="text" name="'+field.id+'"/>';
			}else if( field.type == 'password'){
				contentDiv += '<input type="password" name="'+field.id+'"/>';
			}else if( field.type == 'enum'){
				var disabledDiv = '';
				if( _.isUndefined(field.disabled) == false &&  field.disabled === true )
					disabledDiv = 'disabled="true"';
				var option = "";
				for( var j in field.map ){
					option += '<option value="'+j+'">'+field.map[j]+'</option>';
				}
				contentDiv += '<select '+disabledDiv+' name="'+field.id+'">'+option+'</select>';
			}else if( field.type == 'check'){
				var disabledDiv = '';
				if( _.isUndefined(field.disabled) == false &&  field.disabled === true )
					disabledDiv = 'disabled="true"';
				var option = "";
				for( var j in field.map ){
					option += '<span><input type="checkbox" '+disabledDiv+' name="'+field.id+'" value="'+j+'">'+field.map[j]+'</span>&nbsp;&nbsp;';
				}
				contentDiv += option;
			}else if( field.type == 'table'){
				field.tableId = $.uniqueNum();
				contentDiv += '<div>';
				for( var j = 0 ; j != field.option.button.length ; ++j ){
					var singleButton = field.option.button[j];
					singleButton.buttonId = $.uniqueNum();
					contentDiv += '<button type="button" class="btn" id="'+singleButton.buttonId+'">'+singleButton.name+'</button>';
					field.option.button[j] = singleButton;
				}
				contentDiv += '</div>';
				contentDiv += '<div id="'+field.tableId+'"></div>';
			}
			contentDiv +=
				'</td>'+
			'</tr>';
		}
		var buttonDiv = '';
		if( _.isUndefined(defaultOption.submit) == false ||
			_.isUndefined(defaultOption.cancel) == false ){
			buttonDiv += 
			'<tr>'+
				'<td class="tableleft"></td>'+
				'<td>';
			if( _.isUndefined(defaultOption.submit) == false)
				buttonDiv += '<button type="button" class="btn btn-primary submit" >提交</button>';
			if( _.isUndefined(defaultOption.cancel) == false)
				buttonDiv += '<button type="button" class="btn btn-success cancel">返回列表</button>';
			buttonDiv += 
				'</td>'+
			'</tr>';
		}
		div += '<table class="table table-bordered table-hover definewidth m10">'+
			contentDiv+
			buttonDiv+
		'</table>';
		div = $(div);
		//插入到页面中
		$('#'+defaultOption.id).append(div);
		//挂载控件事件
		for( var i = 0 ; i != defaultOption.field.length; ++i ){
			var field = defaultOption.field[i];
			(function(field){
				if( field.type == 'image'){
					upload.image({
						url:field.option.url,
						urlToken:field.option.urlToken,
						urlType:field.option.urlType,
						target:field.imageTargetId,
						field:'data',
						width:field.option.width,
						height:field.option.height,
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
				}else if( field.type == 'compressFile'){
					field.tableOperation = table.staticSimpleTable({
						id:field.tableId,
						data:[],
						column:[
							{id:'name',type:'html',name:'文件'}
						],
						operate:[]
					});
					upload.file({
						url:field.option.url,
						target:field.fileTargetId,
						field:'data',
						type:field.option.type,
						maxSize:field.option.maxSize,
						onProgress:function(progress){
							$.console.log(progress);
							$('#'+field.fileProgressTargetId).text(progress+'%');
							$('#'+field.fileProgressTargetId).css('width',progress+'%');
						},
						onSuccess:function(data){
							data = $.JSON.parse(data);
							if( data.code != 0 ){
								dialog.message(data.msg);
								return;
							}
							var fileAddress = _.map(data.data,function(single){
								return {name:'<a href="'+single+'" target="_blank">'+single+'</a>'};
							});
							field.tableOperation.clear();
							field.tableOperation.add(fileAddress);
						},
						onFail:function(msg){
							dialog.message(msg);
						}
					});
				}else if( field.type == 'simpleEditor'){
					field._editor = editor.simpleEditor({
						id:field.editorTargetId,
						url:field.option.url
					});
				}else if( field.type == 'fullEditor'){
					field._editor = editor.fullEditor({
						id:field.editorTargetId,
						url:field.option.url
					});
				}else if( field.type == 'time'){
					$('#'+defaultOption.id).find('input[name='+field.id+']').datetimepicker({
						lang:'ch',
						timepicker:false,
						format: 'Y-m-d',
						closeOnDateSelect:true
					});
				}else if( field.type == 'table'){
					field.tableOperation = table.staticSimpleTable({
						id:field.tableId,
						data:[],
						column:field.option.column,
						operate:field.option.operate
					});
					for( var i = 0 ; i != field.option.button.length ; ++i ){
						var singleButton = field.option.button[i];
						$('#'+singleButton.buttonId).click(function(){
							singleButton.click(field.tableOperation);
						});
					}
				}
			})(field);
		}
		//设置value
		function setAllData(dataValue){
			for( var i = 0 ; i != defaultOption.field.length ; ++i ){
				var field = defaultOption.field[i];
				if( typeof dataValue[field.id] == 'undefined' )
					continue;
				if( field.type == 'read'){
					div.find('div[name='+field.id+']').text(dataValue[field.id]);
				}else if( field.type == 'link'){
					div.find('div[name='+field.id+']').text(dataValue[field.id]);
					div.find('a[name='+field.id+']').attr('href',dataValue[field.id]);
				}else if( field.type == 'fullEditor'){
					field._editor.setContent(dataValue[field.id]);
				}else if( field.type == 'simpleEditor'){
					field._editor.setFormatData(dataValue[field.id]);
				}else if( field.type == 'image'){
					div.find('img[name='+field.id+']').attr("src",dataValue[field.id]);
				}else if( field.type == 'compressFile'){
					var fileAddress = _.map(dataValue[field.id],function(single){
						return {name:'<a href="'+single+'" target="_blank">'+single+'</a>'};
					});
					field.tableOperation.clear();
					field.tableOperation.add(fileAddress);
				}else if( field.type == 'area'){
					div.find('textarea[name='+field.id+']').val(dataValue[field.id]);
				}else if( field.type == 'text' || field.type == 'password' || field.type == 'time'){
					div.find('input[name='+field.id+']').val(dataValue[field.id]);
				}else if( field.type == 'enum'){
					div.find('select[name='+field.id+']').val(dataValue[field.id]);
				}else if( field.type == 'check'){
					div.find('input[name='+field.id+']').each(function(){
						for( var i = 0 ; i != dataValue[field.id].length; ++i ){
							var value = dataValue[field.id][i];
							if( $(this).val() == value ){
								$(this).attr('checked',true);
								return;
							}
						}
						$(this).attr('checked',false);
					});
				}else if( field.type == 'table'){
					field.tableOperation.add(dataValue[field.id]);
				}
			}
		}
		//挂载事件
		function getAllData(){
			var data = {};
			for( var i = 0 ; i != defaultOption.field.length ; ++i ){
				var field = defaultOption.field[i];
				if( field.type == 'read'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('div[name='+field.id+']').text());
				}else if( field.type == 'link'){
					data[field.id] = div.find('a[name='+field.id+']').attr('href');
				}else if( field.type == 'simpleEditor'){
					data[field.id] = field._editor.getFormatData();
				}else if( field.type == 'fullEditor'){
					data[field.id] = field._editor.getContent();
				}else if( field.type == 'image'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('img[name='+field.id+']').attr("src"));
				}else if( field.type == 'compressFile'){
					data[field.id] = _.map(field.tableOperation.get(),function(single){
						return single.name;
					});
				}else if( field.type == 'area'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('textarea[name='+field.id+']').val());
				}else if( field.type == 'text' || field.type == 'password' || field.type == 'time'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('input[name='+field.id+']').val());
				}else if( field.type == 'enum'){
					data[field.id] = $.trim($('#'+defaultOption.id).find('select[name='+field.id+']').val());
				}else if( field.type == 'check'){
					data[field.id] = [];
					$('#'+defaultOption.id).find('input[name='+field.id+']:checked').each(function(){
						data[field.id].push($(this).val());
					});
				}else if( field.type == 'table'){
					data[field.id] = field.tableOperation.get();
				}
			}
			return data;
		}
		setAllData(defaultOption.value);
		div.find('.submit').click(function(){
			defaultOption.submit(getAllData());
		});
		div.find('.cancel').click(defaultOption.cancel);
		return {
			get:getAllData,
			set:setAllData
		};
	}
}; });
;define('fishstrap/ui/table.js', function(require, exports, module){ /*
*加载griTable
*@require fishstrap/lib/gri/griTable.js
*/
var $ = require('fishstrap/core/global.js');
var dialog = require('fishstrap/ui/dialog.js');
module.exports = {
	staticSimpleTable:function(option){
		//$.console.log(option.column);
		//执行Option
		var defaultOption = {
			id:'',
			data:'',
			column:[],
			operate:[]
		};
		for( var i in option )
			defaultOption[i] = option[i];
		//清除数据
		function clearAllData(){
			$('#'+defaultOption.id).find('tbody').empty();
		}
		//设置单行数据
		function setSingleData(tr,singleData){
			tr.find('td').each(function(){
				var singleColumn =  null;
				for( var j = 0 ; j != defaultOption.column.length ; ++j ){
					if( defaultOption.column[j].id == $(this).attr('class') ){
						singleColumn = defaultOption.column[j];
						break;
					}
				}
				if( singleColumn == null )
					return;
				if( _.isUndefined(singleData[singleColumn.id]) == true )
					return;
				if( singleColumn.type == 'hidden')
					$(this).text(singleData[singleColumn.id]);
				else if( singleColumn.type == 'image')
					$(this).find('img').attr('src',singleData[singleColumn.id]);
				else if( singleColumn.type == 'textInput')
					$(this).find('input').val(singleData[singleColumn.id]);
				else 
					$(this).text(singleData[singleColumn.id]);
			});
		}
		//获取单行数据
		function getSingleData(tr){
			var singleData = {};
			tr.find('td').each(function(){
				var singleColumn =  null;
				for( var j = 0 ; j != defaultOption.column.length ; ++j ){
					if( defaultOption.column[j].id == $(this).attr('class') ){
						singleColumn = defaultOption.column[j];
						break;
					}
				}
				if( singleColumn == null )
					return;
				if( singleColumn.type == 'hidden')
					singleData[singleColumn.id] = $(this).text();
				else if( singleColumn.type == 'image')
					singleData[singleColumn.id] = $(this).find('img').attr('src');
				else if( singleColumn.type == 'textInput')
					singleData[singleColumn.id] = $(this).find('input').val();
				else 
					singleData[singleColumn.id] = $(this).text();
			});
			return singleData;
		}
		//获取数据
		function getAllData(){
			var data = [];
			$('#'+defaultOption.id).find('tbody tr').each(function(){
				var tr = $(this);
				while( tr.is('tr') == false )
					tr = tr.parent();
				data.push(getSingleData(tr));
			});
			return data;
		}
		//挂在事件
		function addEvent(){
			//设置触发方法
			var triggerEvent = function(tr,next){
				while( tr.is('tr') == false )
					tr = tr.parent();
				var data = getSingleData(tr);
				var operation = {
					remove:function(){
						tr.remove();
					},
					mod:function(data){
						setSingleData(tr,data);
					},
					moveUp:function(){
						var prev = tr.prev();
						tr.insertBefore(prev);
					},
					moveDown:function(){
						var next = tr.next();
						next.insertBefore(tr);
					}
				};
				next(data,operation);
			}
			//挂载operate事件
			for( var i  = 0 ; i != defaultOption.operate.length ; ++i ){
				(function(i){
					$('#'+defaultOption.id).find('.operate_'+defaultOption.operate[i].id).unbind("click").click(function(){
						triggerEvent($(this),defaultOption.operate[i].click);
					});
				})(i);
			}
			//挂载数据事件
			for( var i = 0 ; i != defaultOption.column.length; ++i ){
				(function(i){
					var column = defaultOption.column[i];
					if( column.type == 'textInput' && _.isUndefined(column.change) == false  ){
						$('#'+defaultOption.id).find('.'+column.id+' input').unbind("input").on('input',function(){
							triggerEvent($(this),column.change);
						});
					}
				})(i);
			}
		}
		//添加数据
		function addData(data){
			//构造操作
			var operateDiv = '';
			for( var i = 0 ; i != defaultOption.operate.length ; ++i ){
				defaultOption.operate[i].id = $.uniqueNum();
				operateDiv += "<a href='javascript: void(0)' class=operate_"+defaultOption.operate[i].id
					+">"+defaultOption.operate[i].name+"</a>&nbsp;";
			}
			//构造添加数据
			var div = '';
			for( var i = 0 ; i != data.length; ++i ){
				var item = data[i];
				if( defaultOption.operate.length == 0 )
					var width = 'width:'+(1/defaultOption.column.length*100)+'%;';
				else
					var width = 'width:'+(1/(defaultOption.column.length+1)*100)+'%;';
				div += '<tr>';
				for( var j = 0 ; j != defaultOption.column.length ; ++j ){
					var column = defaultOption.column[j];
					var style = '';
					if( column.type == 'hidden'){
						div += '<td style="display:none;'+width+'" class="'+column.id+'">'+item[column.id]+'</td>';
					}else if( column.type == 'image'){
						div += '<td style="'+width+'" class="'+column.id+'"><img src="'+item[column.id]+'"/></td>';
					}else if( column.type == 'textInput'){
						div += '<td style="'+width+'" class="'+column.id+'"><input style="width:100%" type="text" value="'+item[column.id]+'"/></td>';
					}else {
						div += '<td style="'+width+'" class="'+column.id+'">'+item[column.id]+'</td>';
					}
					
				}
				if( defaultOption.operate.length != 0 ){
					div += '<td style="'+width+'" >'+operateDiv+'</td>';
				}
				div += '</tr>';
			}
			return div;
		}
		//显示数据
		function showData(data){
			var div = '';
			div += '<div class="mod-basic">';
			div += '<table class="mod_table" style="table-layout: auto;">';
			//显示列表头数据
			div += '<thead><tr>';
			if( defaultOption.operate.length == 0 )
				var width = 'width:'+(1/defaultOption.column.length*100)+'%;';
			else
				var width = 'width:'+(1/(defaultOption.column.length+1)*100)+'%;';
			for( var i = 0; i != defaultOption.column.length ;++i ){
				var column = defaultOption.column[i];
				var style = '';
				if( column.type == 'hidden'){
					div += '<th style="display:none;'+width+'"><span class="label">'+column.name+'</span></th>';
				}else {
					div += '<th style="'+width+'"><span class="label">'+column.name+'</span></th>';
				}
			}
			if( defaultOption.operate.length != 0 ){
				div += '<th><span class="label" style="'+width+'">操作</span></th>';
			}
			div += '</tr></thead>';
			//显示列表身数据
			div += '<tbody>';
			div += addData(data);
			div += '</tbody>';
			div += '</table>';
			div += '</div>';
			div = $(div);
			$('#'+defaultOption.id).empty();
			$('#'+defaultOption.id).append(div);
			//挂载事件
			addEvent();
		}
		function addDataAndRefreshEvent(data){
			//添加数据
			$('#'+defaultOption.id).find('tbody').append(addData(data));
			//挂载事件
			addEvent();
		}
		function preaddDataAndRefreshEvent(data){
			//添加数据
			$('#'+defaultOption.id).find('tbody').prepend(addData(data));
			//挂载事件
			addEvent();
		}
		showData(defaultOption.data);
		return {
			preadd:preaddDataAndRefreshEvent,
			add:addDataAndRefreshEvent,
			get:getAllData,
			clear:clearAllData,
		};
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
		_option.checkAll    = defaultOption.checkAll;
		_option.params      = {};

		//拼接url
		var paramStr = "";
		for(var i in defaultOption.params) {
			if( $.trim(defaultOption.params[i]) != "")
				paramStr += (i+"="+encodeURIComponent($.trim(defaultOption.params[i])) + "&");
		}
		
		paramStr += ("t=" + new Date().getTime());
		if(defaultOption.url.indexOf('?') == -1 )
			sendUrl = defaultOption.url + "?" + paramStr;
		else
			sendUrl = defaultOption.url + "&" + paramStr;

		//拼接列
		_option.fields = {};
		for(var i = 0 ; i != defaultOption.column.length; ++i ){
			var column = defaultOption.column[i];
			if( column.hidden ){
				continue;
			}
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
				}else if( column.type == 'image'){		
					single = {
						thText:column.name,		
						format:function(data){		
							return '<img src="'+data+'" style="width:100%;max-width:128px;">';		
						}		
					};		
 				}
			})(column);
			_option.fields[column.id] = single;
		}
		var info = "";
		for( var i = 0 ; i != defaultOption.operate.length; ++i ){
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
					for( var i = 0 ; i != defaultOption.operate.length; ++i ){
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
		function getSingleRowData(parent){
			var singleData = {};
			parent.find('td').each(function(){
				var id = $(this).attr('class');
				var column = null;
				for(var i = 0 ; i != defaultOption.column.length; ++i ){
					var curcolumn = defaultOption.column[i];
					if( curcolumn.id != id )
						continue;
					column = curcolumn;
					break;
				}
				if( column == null )
					return;
				var columnValue;
				if( column.type == 'image'){
					columnValue = $(this).find('img').attr('src');
				}else{
					columnValue = $(this).text();
				}
				singleData[id] = columnValue;
			});
			return singleData;
		}
		return {
			getCheckData:function(){
				var target = $('#'+defaultOption.id+' .gri_td_checkbox:checked');
				var data = [];
				target.each(function(){
					var parent = $(this).parent().parent();
					data.push(getSingleRowData(parent));
				});
				return data;
			}
		};
	},
	
}; });
;define('fishstrap/ui/editor.js', function(require, exports, module){ /*
*加载依赖的ueditor
*@require fishstrap/lib/ueditor/ueditor.config.js
*@require fishstrap/lib/ueditor/ueditor.all.min.js
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
}; });