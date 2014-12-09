define('ui/dialog',['core/global','ui/gri/gri'], function(require, exports, module) {
	//载入JQuery，以及bootstrap插件
	var $ = require('core/global');
	var GRI = require('ui/gri/gri');
	return {
		message:function(msg,callback){
			var title = '提示信息';
			var itype = 3;
			var desc ='';
			var dialog = new GRI.Dialog({ 
				title : title, 
				type : itype, 
				btnType : 3, 
				content : msg,
				winSize : 2,
				desc:desc,
				extra : {
					zIndex : 99999,
					winSize : 2
				}
			}, callback ); 
		},
		confirm:function(msg,callback){
			var title = '提示信息';
			var itype = 3;
			var desc ='';
			var dialog = new GRI.Dialog({ 
				title : title, 
				type : itype, 
				btnType : 1, 
				content : msg,
				winSize : 2,
				desc:desc,
				extra : {
					zIndex : 99999,
					winSize : 2
				}
			}, callback );
		}
	};
});