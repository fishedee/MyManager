<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');
class CI_Argv{
	var $CI;
	
	public function __construct()
    {
		$this->CI = & get_instance();
	}
	
	public function isUrl($s){
		return preg_match('/^http[s]?:\/\/'.
			'(([0-9]{1,3}\.){3}[0-9]{1,3}'. // IP形式的URL- 199.194.52.184
			'|'. // 允许IP和DOMAIN（域名）
			'([0-9a-z_!~*\'()-]+\.)*'. // 域名- www.
			'([0-9a-z][0-9a-z-]{0,61})?[0-9a-z]\.'. // 二级域名
			'[a-z]{2,6})'.  // first level domain- .com or .museum
			'(:[0-9]{1,4})?'.  // 端口- :80
			'((\/\?)|'.  // a slash isn't required if there is no file name
			'(\/[0-9a-zA-Z_!~\'\(\)\[\]\.;\?:@&=\+\$,%#-\/^\*\|]*)?)$/',
			$s) == 1;
	}
	
	public function postRequireInput( $input ,$xssFilt = true )
	{
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->post($value,$xssFilt) === false ){
				return array(
					"code"=>1,
					"msg"=>"请输入post参数".$value,
					"data"=>""
				);
			}else{
				$result[$value] = $this->CI->input->post($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
	
	public function postOptionInput( $input )
	{
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->post($value,true) === false ){
				continue;
			}else{
				$result[$value] = $this->CI->input->post($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
	
	public function postDefaultInput( $input ,$default )
	{
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->post($value,true) === false ){
				$result[$value] = $default;
			}else{
				$result[$value] = $this->CI->input->post($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
	
	public function getRequireInput( $input )
	{
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->get($value,true) === false ){
				return array(
					"code"=>1,
					"msg"=>"请输入get参数".$value,
					"data"=>""
				);
			}else{
				$result[$value] = $this->CI->input->get($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
	
	public function getOptionInput( $input ){
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->get($value,true) === false ){
				continue;
			}else{
				$result[$value] = $this->CI->input->get($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
	
	public function getDefaultInput( $input ,$default )
	{
		$result = array();
		foreach( $input as $key=>$value ){
			if( $this->CI->input->get($value,true) === false ){
				$result[$value] = $default;
			}else{
				$result[$value] = $this->CI->input->post($value,true);
			}
		}
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$result
		);
	}
}
?>