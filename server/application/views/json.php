<?php
	header("Expires: Mon, 26 Jul 1970 05:00:00 GMT");
	header("Last-Modified:".gmdate("D, d M Y H:i:s ")."GMT");
	header("Cache-control:no-cache,no-store,must-revalidate"); 
	header("Pragma:no-cache");
	$result = array(
		"code"=>$code,	
		"msg"=>$msg,
		"data"=>$data
	);
	$output = json_encode($result);
	if( $output == null )
		$output = json_encode(array(
			'code'=>1,
			'msg'=>'输出中含有非UTF8编码',
			'data'=>''
		));
	echo $output;
?>
