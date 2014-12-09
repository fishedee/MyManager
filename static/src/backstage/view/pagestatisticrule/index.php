<?php require(__DIR__."/../common/header.php"); ?>
<!-- 查询窗口 -->
<form id="user-form" class="form-inline">
    名称：
    <input type="text" name="name" class="input-small"/>
	状态：
    <select name="state">
		<option value="">请选择状态</option>
		<option value="0">可用</option>
		<option value="1">不可用</option>
    </select>
    <button type="button" class="btn query">查询</button>
    <button type="reset" class="btn">重置</button>
</form>
<div class="m10">
    <button class="btn add">添加页面统计规则</button>
    <div id="table"></div>
</div>
<div>
</div>

<script>
function view(statisticRuleId){
	_href('/pagestatisticrule/view.html?statisticRuleId=' + statisticRuleId);
}
function add(){
	_href('/pagestatisticrule/view.html');
}
function get(params){
	var fields = {
		statisticRuleId:"统计规则ID",
        name: '名称',
		rule: '规则',
		state: {
			name:'状态',
			format:function(data){
				var m = {0:"可用",1:"不可用"};
				return m[data];
			}
		},
        createTime: '创建时间',
        modifyTime: '修改时间',
        oper:{
			name:'操作',
			format:function(data){
				return "<a href='#' class='_edit'>编辑</a>";
			}
		}
    };
    _table('/pagestatisticrule/search', { key_index: 'statisticRuleId', data: 'data', fields: fields,params:params}, function(){
        $("._edit").unbind("click").click(function(){
            var statisticRuleId = $.trim($(this).parent().parent().find(".statisticRuleId").text());
            view(statisticRuleId);
        });
    });
}
$(function(){
	$(".query").unbind("click").click(function(){
       get({
			name:$("input[name=name]").val(),
			state:$("select[name=state]").val()
	   });
    });
    $(".add").unbind("click").click(function(){
       add();
    });
	get();
});
</script>
<?php require(__DIR__."/../common/footer.php"); ?>
