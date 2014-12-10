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
		//��ȡ����
		$result = $this->categoryDb->get($categoryId);
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
	
	public function del($userId,$categoryId){
		//���Ȩ��
		$result = $this->get($userId,$categoryId);
		if( $result['code'] != 0 )
			return $result;
			
		//ɾ��category�����������
		$result = $this->categoryDb->del($categoryId);
		if( $result['code'] != 0 )
			return $result;
			
		//֪ͨaccountWhen��ָ����categoryId��ɾ����
		$this->accountWhen->whenCategoryDelete($categoryId);
		return $result;
	}
	
	public function add($userId,$data){
		$data['userId'] = $userId;
		return $this->categoryDb->add($data);
	}
	
	public function mod($userId,$categoryId,$data){
		//���Ȩ��
		$result = $this->get($userId,$categoryId);
		if( $result['code'] != 0 )
			return $result;
		
		$data['userId'] = $userId;
		return $this->categoryDb->mod($categoryId,$data);
	}

}
