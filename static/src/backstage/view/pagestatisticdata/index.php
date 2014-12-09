<?php require(__DIR__."/../common/header.php"); ?>
<!-- 查询窗口 -->
<form id="user-form" class="form-inline">
    页面：
	<select id="statisticRuleId" name="statisticRuleId">
		<option value="">请选择页面</option>
    </select>
	开始日期
	<input id="beginTime" type="text" name="beginTime" class="time input-small"> 
	结束日期
	<input id="endTime" type="text" name="endTime" class="time input-small"> 
	选择时间：
    <select id="choosetime" name="choosetime">
	</select>
    <button type="button" class="btn query">查询</button>
    <button type="reset" class="btn">重置</button>
</form>
<div class="m10" id="showChart" style="height:400px">
</div>
<div class="m10">
    <table class="table table-bordered table-hover definewidth m10 step">
		<tr>
			<th>日期</th>
			<th>页面</th>
			<th>PV</th>
			<th>UV</th>
		</tr>
	</table>
</div>
<div>
</div>

<script>
function getAllPage(fun){
	$.get('/pagestatisticrule/getAllValid',{},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		var div = "";
		for( var i in data.data ){
			var item = data.data[i];
			div += '<option value="'+item.statisticRuleId+
				'">'+item.name+'</option>';
		}
		$('#statisticRuleId').append($(div));
		$('#statisticRuleId').val($($("#statisticRuleId option")[1]).val());
		fun();
	});
}
function get(params){
	var option = {};
	for( var i in params ){
		if( params[i] != "")
			option[i] = params[i];
	}
	$.get('/pagestatisticdata/search',option,function(data){
		//绘制表格
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		$(".table .data").remove();
		var div = "";
		for( var i in data.data.data ){
			var item = data.data.data[i];
			div += 
			'<tr class="data">'+
				'<td>'+item.day+'</td>'+
				'<td>'+item.name+'</td>'+
				'<td>'+item.pv+'</td>'+
				'<td>'+item.uv+'</td>'+
			'</tr>';
		}
		$(".table").append(div);
		
		//绘制折线图
		var allDay = {};
		var allPage = {};
		var allData = {};
		for( var i in data.data.data ){
			var item = data.data.data[i];
			allDay[ item.day ] = "";
			allPage[ item.name ] = "";
			allData[ item.day+"_"+item.name+"_PV"] = item.pv;
			allData[ item.day+"_"+item.name+"_UV"] = item.uv;
		}
		var xAxis = Object.keys(allDay);
		var yAxis = Object.keys(allPage);
		var series = [];
		xAxis.sort();
		for( var i in yAxis ){
			var pvs = [];
			var uvs = [];
			var name = yAxis[i];
			for( var j in xAxis ){
				var day = xAxis[j];
				if(allData[day+"_"+name+"_PV"]){
					pvs.push(allData[day+"_"+name+"_PV"]);
					uvs.push(allData[day+"_"+name+"_UV"]);
				}else{
					pvs.push(0);
					uvs.push(0);
				}
			}
			series.push({
				name:name+"PV",
				data:pvs
			});
			series.push({
				name:name+"UV",
				data:uvs
			});
		}
		for( var i in xAxis ){
			xAxis[i] = xAxis[i].substr(5);
		}
        _drawBrokeLine("#showChart",{
			xAxis:xAxis,
			series:series,
		});
	});
}
$(function(){
	_setTimePicker("#beginTime","#endTime","#choosetime");
	getAllPage( function(){
		$(".query").unbind("click").click(function(){
		   get({
				statisticRuleId:$("select[name=statisticRuleId]").val(),
				beginTime:$("input[name=beginTime]").val(),
				endTime:$("input[name=endTime]").val(),
				state:$("select[name=state]").val()
		   });
		});
		get({
			statisticRuleId:$("select[name=statisticRuleId]").val(),
			beginTime:$("input[name=beginTime]").val(),
			endTime:$("input[name=endTime]").val(),
			state:$("select[name=state]").val()
		});
	});
});
</script>
<?php require(__DIR__."/../common/footer.php"); ?>
