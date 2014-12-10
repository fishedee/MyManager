<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class Card extends CI_Controller {

	public function __construct()
    {
        parent::__construct();
		$this->load->model('card/cardAo','cardAo');
		$this->load->model('user/loginAo','loginAo');
		$this->load->library('argv','argv');
    }
	
	public function search()
	{
		//���Ȩ��
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//����������		
		$result = $this->argv->getOptionInput(array('name','remark'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return;
		}
		$dataWhere = $result["data"];
		
		$result = $this->argv->getOptionInput(array('pageIndex','pageSize'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return;
		}
		$dataLimit = $result["data"];
			
		//ִ��ҵ���߼�
		$data = $this->cardAo->search($userId,$dataWhere,$dataLimit);
		if( $data["code"] != 0 ){
			$this->load->view('json',$data);
			return;
		}
		$this->load->view('json',$data);
	}
	
	public function get()
	{
		//���Ȩ��
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//����������
		$result = $this->argv->getRequireInput(array('cardId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$cardId = $result["data"]["cardId"];
		
		//ִ��ҵ���߼�
		$data = $this->cardAo->get(
			$userId,
			$cardId
		);
		$this->load->view('json',$data);
	}
	
	public function add()
	{
		//���Ȩ��
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//����������
		$result = $this->argv->postRequireInput(array('name','remark','card','bank','money'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$data = $result["data"];
		
		//ִ��ҵ���߼�
		$data = $this->cardAo->add(
			$userId,
			$data
		);
		$this->load->view('json',$data);
	}
	
	public function del()
	{
		//���Ȩ��
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//����������
		$result = $this->argv->postRequireInput(array('cardId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$cardId = $result["data"]['cardId'];
		
		//ִ��ҵ���߼�
		$data = $this->cardAo->del(
			$userId,
			$cardId
		);
		$this->load->view('json',$data);
	}
	
	public function mod()
	{
		//���Ȩ��
		$result = $this->loginAo->islogin();
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$userId = $result['data'];
		
		//����������
		$result = $this->argv->postRequireInput(array('cardId'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$cardId = $result["data"]["cardId"];
		
		$result = $this->argv->postRequireInput(array('name','remark','card','bank','money'));
		if( $result["code"] != 0 ){
			$this->load->view('json',$result);
			return $result;
		}
		$data = $result["data"];
		
		//ִ��ҵ���߼�
		$data = $this->cardAo->mod(
			$userId,
			$cardId,
			$data
		);
		$this->load->view('json',$data);
	}

}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
