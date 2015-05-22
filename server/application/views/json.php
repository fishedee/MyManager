<?php
	$result = array(
		"code"=>$code,	
		"msg"=>$msg,
		"data"=>$data
	);
	$output = json_encode($result,JSON_UNESCAPED_UNICODE);
	if( $output == null )
		$output = json_encode(array(
			'code'=>1,
			'msg'=>'输出中含有非UTF8编码',
			'data'=>''
		));
	echo $output;
?>
