<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class CategoryAo extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('category/categoryDb','categoryDb');
		$this->load->model('account/accountWhen','accountWhen');
	}

	public function search($userId,$where,$limit){
		$where['userId'] = $userId;
		return $this->categoryDb->search($where,$limit);
	}

	public function get($userId,$categoryId){
		//获取数据
		$result = $this->categoryDb->get($categoryId);
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
	
	public function del($userId,$categoryId){
		//检查权限
		$result = $this->get($userId,$categoryId);
		if( $result['code'] != 0 )
			return $result;
			
		//删除category表里面的数据
		$result = $this->categoryDb->del($categoryId);
		if( $result['code'] != 0 )
			return $result;
			
		//通知accountWhen，指定的categoryId被删除了
		$this->accountWhen->whenCategoryDelete($categoryId);
		return $result;
	}
	
	public function add($userId,$data){
		$data['userId'] = $userId;
		return $this->categoryDb->add($data);
	}
	
	public function mod($userId,$categoryId,$data){
		//检查权限
		$result = $this->get($userId,$categoryId);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->categoryDb->mod($categoryId,$data);
	}

}
