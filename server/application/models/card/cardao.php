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
		//获取数据
		$result = $this->cardDb->get($cardId);
		if( $result['code'] != 0 )
			return $result;
		
		//检查权限
		if( $result['data']['userId'] != $userId )
			return array(
				'code'=>1,
				'msg'=>'权限不足',
				'data'=>''
			);
		
		return $result;
	}
	
	public function del($userId,$cardId){
		//检查权限
		$result = $this->get($userId,$cardId);
		if( $result['code'] != 0 )
			return $result;
			
		//删除card表里面的数据
		$result = $this->cardDb->del($cardId);
		if( $result['code'] != 0 )
			return $result;
			
		//通知accountWhen，指定的cardId被删除了
		$this->accountWhen->whenCardDelete($cardId);
		return $result;
	}
	
	public function add($userId,$data){
		$data['userId'] = $userId;
		return $this->cardDb->add($data);
	}
	
	public function mod($userId,$cardId,$data){
		//检查权限
		$result = $this->get($userId,$cardId);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->cardDb->mod($cardId,$data);
	}

}
