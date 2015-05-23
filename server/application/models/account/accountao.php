<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class AccountAo extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('account/accountDb','accountDb');
		$this->load->model('category/categoryAo','categoryAo');
		$this->load->model('card/cardAo','cardAo');
	}
	private function checkAccountData( $userId,$data ){
		//校验分类ID
		$result = $this->categoryAo->get( $userId,$data['categoryId']);
		if( $result['code'] != 0 ){
			return array(
				'code'=>1,
				'msg'=>'分类ID不合法',
				'data'=>''
			);
		}
			
		//校验银卡ID
		$result = $this->cardAo->get( $userId,$data['cardId']);
		if( $result['code'] != 0 ){
			return array(
				'code'=>1,
				'msg'=>'银卡ID不合法',
				'data'=>''
			);
		}
		
		//校验类型
		if( $data['type'] != $this->accountDb->TYPE_IN
			&& $data['type'] != $this->accountDb->TYPE_OUT
			&& $data['type'] != $this->accountDb->TYPE_TRANSFER_IN
			&& $data['type'] != $this->accountDb->TYPE_TRANSFER_OUT	
			&& $data['type'] != $this->accountDb->TYPE_BORROW_IN 
			&& $data['type'] != $this->accountDb->TYPE_BORROW_OUT ){
			return array(
				'code'=>1,
				'msg'=>'类型ID不合法',
				'data'=>''
			);
		}
		
		//校验金额
		if( $data['money'] < 0 ){
			return array(
				'code'=>1,
				'msg'=>'金额必须大于或等于0',
				'data'=>''
			);
		}
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>''
		);
	}
	
	public function search($userId,$where,$limit){
		$where['userId'] = $userId;
		return $this->accountDb->search($where,$limit);
	}

	public function get($userId,$accountId){
		//获取数据
		$result = $this->accountDb->get($accountId);
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
	
	public function del($userId,$accountId){
		//检查权限
		$result = $this->get($userId,$accountId);
		if( $result['code'] != 0 )
			return $result;
			
		return $this->accountDb->del($accountId);
	}
	
	public function add($userId,$data){
		//检查数据
		$result = $this->checkAccountData($userId,$data);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->accountDb->add($data);
	}
	
	public function mod($userId,$accountId,$data){
		//检查数据
		$result = $this->checkAccountData($userId,$data);
		if( $result['code'] != 0 )
			return $result;
			
		//检查权限
		$result = $this->get($userId,$accountId);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->accountDb->mod($accountId,$data);
	}

}
