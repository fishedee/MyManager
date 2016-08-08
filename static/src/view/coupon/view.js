var $ = require('fishfront/ui/global');
var input = require('fishfront/ui/input');
var dialog = require('fishfront/ui/dialog');
$('body').append('<div id="show"></div>');
function go(){
	inputOperation = input.verticalInput({
		id:'container',
		field:[
			{id:'couponId',type:'text',name:'优惠卷id'},
			{id:'userId',type:'area',name:'用户列表id'},
			{id:'url',type:'read',name:'批量优惠卷请求'},
		],
		value:{},
		submit:function(data){
			var couponId = $.trim(data.couponId);
			var userIds = data.userId.match(/\d+/g);
			if( couponId == ''){
				dialog.message('请输入优惠卷ID');
				return;
			}
			if( userIds == null || userIds.length == 0){
				dialog.message('请输入用户列表ID');
				return;
			}
			var result = "$.post('/index.php?m=Admin&c=Coupon&a=send',{coupon_id:"+
				couponId+",user_id:["+
				userIds.join(',')+"]},function(data){alert(JSON.stringify(data))});";
			inputOperation.set({url:result});
		},
		cancel:undefined
	});
}
go();
