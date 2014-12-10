<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class CardAo extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('card/cardDb','cardDb');
		$this->load->model('account/accountWhen','accountWhen');
	}

	public function search($userId,$where,$limit){
		$where['userId'] = $userId;
		return $this->cardDb->search($where,$limit);
	}

	public function get($userId,$cardId){
		//��ȡ����
		$result = $this->cardDb->get($cardId);
		if( $result['code'] != 0 )
			return $result;
		
		//���Ȩ��
		if( $result['data']['userId'] != $userId )
			return array(
				'code'=>1,
				'msg'=>'Ȩ�޲���',
				'data'=>''
			);
		
		return $result;
	}
	
	public function del($userId,$cardId){
		//���Ȩ��
		$result = $this->get($userId,$cardId);
		if( $result['code'] != 0 )
			return $result;
			
		//ɾ��card�����������
		$result = $this->cardDb->del($cardId);
		if( $result['code'] != 0 )
			return $result;
			
		//֪ͨaccountWhen��ָ����cardId��ɾ����
		$this->accountWhen->whenCardDelete($cardId);
		return $result;
	}
	
	public function add($userId,$data){
		$data['userId'] = $userId;
		return $this->cardDb->add($data);
	}
	
	public function mod($userId,$cardId,$data){
		//���Ȩ��
		$result = $this->get($userId,$cardId);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->cardDb->mod($cardId,$data);
	}

}
