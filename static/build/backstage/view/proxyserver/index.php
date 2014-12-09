<?php require(__DIR__."/../common/header.php"); ?>
<div class="alert alert-danger" role="alert">
<strong>注意! </strong>使用完代理服务器后记得关掉，不然会影响现网网站的性能，而且长期使用单一IP的代理服务器容易造成IP被封，切记使用完关掉代理服务器
</div>
<div class="totalinfo">
	代理服务器地址：<span class="addressinfo"></span><br/>
	代理服务器端口：3128<br/>
	代理服务器状态：<span class="statusinfo"></span>
</div>
<div class="buttonlist">
	代理服务器操作：
	<button type="button" class="btn btn-primary start">开始</button>
	<button type="button" class="btn btn-warn stop">停止</button>
</div>
<div>
	<div>使用说明</div>
	<button type="button" class="btn btn-warn download">下载说明文档</button>
</div>
<script>
function updateTotalInfo(){
	$(".addressinfo").text(location.host);
	$.get('/proxy/status',{},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		if( data.data == true ){
			$(".statusinfo").text('正在运行中');
			$(".stop").removeAttr('disabled');
			$(".start").attr('disabled','disabled');
		}else{
			$(".statusinfo").text('停止中');
			$(".stop").attr('disabled','disabled');
			$(".start").removeAttr('disabled');
		}
	});
}
function start(){
	$.post('/proxy/start',{},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		updateTotalInfo();
	});
}
function stop(){
	$.post('/proxy/stop',{},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		updateTotalInfo();
	});
}
function download(){
	location.href = "../../assets/document/%E4%BB%A3%E7%90%86%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%BD%BF%E7%94%A8.docx";
}
$(document).ready( function(){
	$(".stop").click(stop);
	$(".start").click(start);
	$(".download").click(download);
	updateTotalInfo();
});
</script>
<?php require(__DIR__."/../common/footer.php"); ?>
