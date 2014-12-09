<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class User extends CI_Controller {

	public function __construct()
    {
        parent::__construct();
		$this->load->model('user/userAo','userAo');
		$this->load->model('user/loginAo','loginAo');
		$this->load->library('argv','argv');
    }
	
	private function _checkAdmin()
	{
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 )
			return $result;
		
		$userId = $result["data"];
		$result = $this->userAo->get($userId);
		if( $result["code"] != 0 )
			return $result;
			
		if( $result["data"]['type'] != $this->userAo->TYPE_ADMIN )
			return array(
				"code"=>1,
				"msg"=>"你没有权限执行此操作",
				"data"=>""
			);
		
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>$userId
		);
	}
	
	public function search()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数		
		$result = $this->argv->getOptionInput(array('name','type'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return;
		}
		$dataWhere = $result["data"];
		
		$result = $this->argv->getRequireInput(array('pageIndex','pageSize'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return;
		}
		$dataLimit = $result["data"];
			
		//执行业务逻辑
		$data = $this->userAo->search($dataWhere,$dataLimit);
		if( $data["code"] != 0 ){
			$this->load->view('json',$data);
			return;
		}
		$this->load->view('json',$data);
	}
	
	public function get()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数
		$result = $this->argv->getRequireInput(array('userId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result["data"]["userId"];
		
		//执行业务逻辑
		$data = $this->userAo->get(
			$userId
		);
		$this->load->view('json',$data);
	}
	
	public function add()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('name','type','password'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$data = $result["data"];
		
		//执行业务逻辑
		$data = $this->userAo->add(
			$data
		);
		$this->load->view('json',$data);
	}
	
	public function del()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('userId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result["data"]['userId'];
		
		//执行业务逻辑
		$data = $this->userAo->del(
			$userId
		);
		$this->load->view('json',$data);
	}
	
	public function modType()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('userId','type'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result["data"]["userId"];
		$type = $result["data"]["type"];
		
		//执行业务逻辑
		$data = $this->userAo->modType(
			$userId,
			$type
		);
		$this->load->view('json',$data);
	}
	
	public function modPassword()
	{
		//检查权限
		$result = $this->_checkAdmin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('userId','password'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result["data"]["userId"];
		$password = $result["data"]["password"];
		
		//执行业务逻辑
		$data = $this->userAo->modPassword(
			$userId,
			$password
		);
		$this->load->view('json',$data);
	}
	
	public function modMyPassword()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('oldPassword','newPassword'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$oldPassword = $result["data"]["oldPassword"];
		$newPassword = $result["data"]["newPassword"];
		
		//执行业务逻辑
		$data = $this->userAo->modPasswordByOld(
			$userId,
			$oldPassword,
			$newPassword
		);
		$this->load->view('json',$data);
	}
}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
