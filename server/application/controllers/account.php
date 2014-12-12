<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class Account extends CI_Controller {
	var $TYPE_PAY = 1;
	var $TYPE_INCOME = 2;
	var $TYPE_TRANSFER_IN = 3;
	var $TYPE_TRANSFER_OUT = 4;
	public function __construct()
    {
        parent::__construct();
		$this->load->model('account/accountAo','accountAo');
		$this->load->model('account/accountStatistic','accountStatistic');
		$this->load->model('user/loginAo','loginAo');
		$this->load->library('argv','argv');
    }
	
	public function search()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数		
		$result = $this->argv->getOptionInput(array('name','remark','categoryId','cardId','type'));
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
		$data = $this->accountAo->search($userId,$dataWhere,$dataLimit);
		if( $data["code"] != 0 ){
			$this->load->view('json',$data);
			return;
		}
		$this->load->view('json',$data);
	}
	
	public function get()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->getRequireInput(array('accountId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$accountId = $result["data"]["accountId"];
		
		//执行业务逻辑
		$data = $this->accountAo->get(
			$userId,
			$accountId
		);
		$this->load->view('json',$data);
	}
	
	public function add()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('name','money','remark','categoryId','cardId','type'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$data = $result["data"];
		
		//执行业务逻辑
		$data = $this->accountAo->add(
			$userId,
			$data
		);
		$this->load->view('json',$data);
	}
	
	public function del()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('accountId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$accountId = $result["data"]['accountId'];
		
		//执行业务逻辑
		$data = $this->accountAo->del(
			$userId,
			$accountId
		);
		$this->load->view('json',$data);
	}
	
	public function mod()
	{
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->postRequireInput(array('accountId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$accountId = $result["data"]["accountId"];
		
		$result = $this->argv->postRequireInput(array('name','money','remark','categoryId','cardId','type'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$data = $result["data"];
		
		//执行业务逻辑
		$data = $this->accountAo->mod(
			$userId,
			$accountId,
			$data
		);
		$this->load->view('json',$data);
	}
	
	public function getMonthTypeStatistic(){
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//执行业务逻辑
		$data = $this->accountStatistic->getMonthTypeStatistic(
			$userId
		);
		$this->load->view('json',$data);
	}
	
	public function getDetailTypeStatistic(){
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->getRequireInput(array('year','month','type'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$year = $result["data"]["year"];
		$month = $result["data"]["month"];
		$type = $result["data"]["type"];
		
		//执行业务逻辑
		$data = $this->accountStatistic->getDetailTypeStatistic(
			$userId,
			$year,
			$month,
			$type
		);
		$this->load->view('json',$data);
	}
	
	public function getMonthCardStatistic(){
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//执行业务逻辑
		$data = $this->accountStatistic->getMonthCardStatistic(
			$userId
		);
		$this->load->view('json',$data);
	}
	
	public function getDetailCardStatistic(){
		//检查权限
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//检查输入参数
		$result = $this->argv->getRequireInput(array('year','month','cardId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$year = $result["data"]["year"];
		$month = $result["data"]["month"];
		$cardId = $result["data"]["cardId"];
		
		//执行业务逻辑
		$data = $this->accountStatistic->getDetailCardStatistic(
			$userId,
			$year,
			$month,
			$cardId
		);
		$this->load->view('json',$data);
	}

}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
