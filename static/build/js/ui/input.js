define('ui/input',['core/global','ui/editor','ui/dialog','util/upload'], function(require, exports, module){
	var $ = require('core/global');
	var dialog = require('ui/dialog');
	var editor = require('ui/editor');
	var upload = require('util/upload');
	return {
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