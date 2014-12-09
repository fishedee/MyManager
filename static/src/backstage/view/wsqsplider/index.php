<?php require(__DIR__."/../common/header.php"); ?>
<div class="alert alert-danger" role="alert">
<strong>注意! </strong>使用前请确认已经用--disable-web-security模式启动chrome浏览器，并已经登录微社区网站
</div>
<form id="user-form" class="form-inline">
	帖子地址：
    <input type="text" name="name" class="input-big"/>
	评论日期
	<input id="time" type="text" name="time" class="time input-small"> 
    <button type="button" class="btn btn-primary query">抓取</button>
	<button type="button" class="btn btn-warn excel">导出到excel</button>
    <button type="reset" class="btn">重置</button>
</form>
<form id="excel-form" action="/excel/exportFromUser" method="post" style="display:none">
	标题
	<input type="text" name="title" class="input-small" value="微社区评论数据"/>
	数据
	<input type="text" name="data" class="input-small"/>
</form>
<div class="totalinfo">
	共有0条回复
</div>
<table class="table table-bordered table-hover definewidth m10 step">
	<tr>
		<th>楼层</th>
		<th>用户名</th>
		<th>用户ID</th>
		<th>评论内容</th>
		<th>时间</th>
	</tr>
</table>
<script>
var result = {};
function getSinglePage( userId , bbsId , page ,fun ){
	$(".totalinfo").text("正在抓取第"+page+"页数据");
	var iframe = $('<iframe></iframe>');
	iframe.attr("src",
		"http://mp.wsq.qq.com/"+userId+"/t/"+bbsId+"?page="+page+"&resType=json&isAjax=2");
	iframe.on("load",function(){
		var data = JSON.parse(iframe.contents().text());
		fun(data);
	});
	iframe.hide();
	$(document.body).append(iframe);
}
function getAllPage(option,fun){
	if( !option.page ){
		option.page = 1;
	}
	if( option.page == 1000 ){
		fun();
		return;
	}
	getSinglePage(option.userId,option.bbsId,option.page,function(data){
		if( data.errCode != 0 ){
			_dialog( data.message );
			fun();
			return;
		}
		var hasData = false;
		for( var i in data.data.dataList ){
			var item = data.data.dataList[i];
			if( item.floor ){
				hasData = true;
				if( option.time && item.fCreatedTime.substr(0,10) != option.time)
					continue;
				result[i] = {
					floor:item.floor,
					author:item.author,
					authorUid:item.authorUid,
					content:item.content,
					fCreatedTime:item.fCreatedTime,
					replyList:[]
				};
				if(item.floorList){
					for( var j in item.floorList ){
						var reply = item.floorList[j];
						result[i].replyList[j] = {
							floor:"",
							author:reply.author,
							authorUid:reply.authorUid,
							content:reply.content,
							fCreatedTime:reply.fCreatedTime
						};
					}
				}
			}
		}
		if( hasData == false ){
			fun();
			return;
		}
		option.page = option.page + 1;
		getAllPage(option,fun);
	});
}
function splider(){
	//校验数据
	var option = {};
	var url = $("input[name=name]").val();
	var time = $("input[name=time]").val();
	var urlmatch = /http:\/\/mp\.wsq\.qq\.com\/cp#!(\d+)\/t\/(\d+)/;
	var matchresult = url.match(urlmatch);
	console.log(matchresult);
	if( matchresult && matchresult["1"] && matchresult["2"] ){
		option.userId = matchresult["1"];
		option.bbsId = matchresult["2"];
	}else{
		_dialog("请输入帖子地址");
		return;
	}
	if( $.trim(time) != ""){
		option.time = $.trim(time);
	}
	//获取数据
	$(".table .data").remove();
	result = {};
	$('body').append(loadingDiv);
	getAllPage(option,function(){
		//页面展示
		$('#__loading').remove();
		//设置提示
		var floorCount = 0;
		for( var i in result ){
			if( result[i].floor != ""){
				floorCount++;
			}
		}
		$(".totalinfo").text("共有"+floorCount+"回复");
		//设置表格
		var divs = "";
		for( var i in result ){
			var item = result[i];
			divs += 
				'<tr class="data">'+
					'<td>'+item.floor+'</td>'+
					'<td>'+item.author+'</td>'+
					'<td>'+item.authorUid+'</td>'+
					'<td>'+item.content+'</td>'+
					'<td>'+item.fCreatedTime+'</td>'+
				'</tr>';
			for( var j in item.replyList ){
				var reply = item.replyList[j];
				divs += 
					'<tr class="data">'+
						'<td>'+reply.floor+'</td>'+
						'<td>'+reply.author+'</td>'+
						'<td>'+reply.authorUid+'</td>'+
						'<td>'+reply.content+'</td>'+
						'<td>'+reply.fCreatedTime+'</td>'+
					'</tr>';
			}
		}
		$(".table").append($(divs));
		
	});
	
}
function excel(){
	var data = [];
	data.push([
		$($(".table th")[0]).text(),
		$($(".table th")[1]).text(),
		$($(".table th")[2]).text(),
		$($(".table th")[3]).text(),
		$($(".table th")[4]).text(),
	]);
	console.log(data);
	$(".table .data").each(function(){
		var div = $(this);
		data.push([
			$(div.find("td")[0]).text(),
			$(div.find("td")[1]).text(),
			$(div.find("td")[2]).text(),
			$(div.find("td")[3]).text(),
			$(div.find("td")[4]).text(),
		]);
	});
	console.log(data);
	$("#excel-form input[name=data]").val(JSON.stringify(data));
	$("#excel-form").submit();
}
$(document).ready( function(){
	$('#time').datetimepicker({
		lang:'ch',
		timepicker:false,
		format: 'Y-m-d',
		closeOnDateSelect:true
	});
	$(".query").click( splider );
	$(".excel").click( excel );
});
</script>
<?php require(__DIR__."/../common/footer.php"); ?>
