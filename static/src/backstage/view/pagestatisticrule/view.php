<?php require(__DIR__."/../common/header.php"); ?>
<form id="user-form" class="definewidth m20" >
<input type="hidden" name="_method" value="POST"/>
<table class="table table-bordered table-hover definewidth m10">
    <tr>
        <td class="tableleft">名称</td>
        <td><input type="text" name="name"/></td>
    </tr>
	<tr>
        <td class="tableleft">规则</td>
        <td><input type="text" name="rule"/></td>
    </tr>
	<tr>
        <td class="tableleft">状态</td>
        <td>
        <select name="state">
            <option value="0">可用</option>
            <option value="1">不可用</option>
        </select>
        </td>
    </tr>
    <tr>
        <td class="tableleft"></td>
        <td>
            <button type="button" class="btn btn-primary submit" >提交</button>
			<button type="button" class="btn btn-success" onclick="location.href='index.html'">返回列表</button>
        </td>
    </tr>
</table>
</form>
<script>

function get( statisticRuleId ){
	$.get("/pagestatisticrule/get?t="+Math.random(),{statisticRuleId:statisticRuleId},function(data){
		data = $.JSON.parse(data);
		if( data.code != 0 )
			return;
		var data = data.data;
		$("input[name=name]").val(data.name);
		$("input[name=rule]").val(data.rule);
		$("select[name=state]").val(data.state);
	});
}
function mod( statisticRuleId ){
	var data = getAndValidate([
		{
			name:'name',
			input:'input[name=name]',
			desc:'名称',
			validate:'text'
		},
		{
			name:'rule',
			input:'input[name=rule]',
			desc:'规则',
			validate:'text'
		},
		{
			name:'state',
			input:'select[name=state]',
			desc:'状态',
			validate:'digit'
		}
	]);
	if( data == false )
		return;
	data.statisticRuleId = statisticRuleId;
	$.post("/pagestatisticrule/mod?t="+Math.random(),data,function(data){
		data = $.JSON.parse(data);
		if( data.code == 0)
			_href("/pagestatisticrule/index.html");
	});
}
function add(){
	var data = getAndValidate([
		{
			name:'name',
			input:'input[name=name]',
			desc:'名称',
			validate:'text'
		},
		{
			name:'rule',
			input:'input[name=rule]',
			desc:'规则',
			validate:'text'
		},
		{
			name:'state',
			input:'select[name=state]',
			desc:'状态',
			validate:'digit'
		}
	]);
	if( data == false )
		return;
	$.post("/pagestatisticrule/add?t="+Math.random(),data,function(data){
		data = $.JSON.parse(data);
		if( data.code == 0){
			_href("/pagestatisticrule/index.html");
		}
	});
}
$(function(){
    if( _get("statisticRuleId") && _get("statisticRuleId") != "null" ){
		get(_get("statisticRuleId"));
		$(".submit").click(function(){
			mod(_get("statisticRuleId"));		
		});
	}else{
		$(".submit").click(function(){
			add();
		});
	}
});

</script>
<?php require(__DIR__."/../common/footer.php"); ?>
