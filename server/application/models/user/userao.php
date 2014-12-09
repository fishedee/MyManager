<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class UserAo extends CI_Model {
	
	var $TYPE_ADMIN = 0;
	var $TYPE_USER = 1;
	public function __construct(){
		parent::__construct();
		$this->load->model('user/userDb','userDb');
	}
	
	public function search($dataWhere,$dataLimit){
		return $this->userDb->search($dataWhere,$dataLimit);
	}
	
	public function get($userId){
		return $this->userDb->get($userId);
	}
	
	public function del($userId){
		return $this->userDb->del($userId);
	}
	
	public function add($data){
		//检查是否有重名
		$result = $this->userDb->getByName($data['name']);
		if( $result['code'] != 0 )
			return $result;
		$user = $result['data'];
		if( count($user) != 0 )
			return array(
				'code'=>1,
				'msg'=>'存在重复的用户名',
				'data'=>''
			);
		
		//添加用户
		$data['password'] = sha1($data['password']);
		return $this->userDb->add($data);
	}
	
	public function modType($userId,$type){
		$data = array();
		$data['type'] = $type;
		return $this->userDb->mod($userId,$data);
	}
	
	public function modPassword($userId,$password){
		$data = array();
		$data['password'] = sha1($password);
		return $this->userDb->mod($userId,$data);
	}
	
	public function modPasswordByOld($userId,$oldPassword,$newPassword){
		//检查是否有重名
		$result = $this->userDb->getByIdAndPass($userId,sha1($oldPassword));
		if( $result['code'] != 0 )
			return $result;
		$users = $result['data'];
		if( count($users) == 0 )
			return array(
				'code'=>1,
				'msg'=>'原密码错误',
				'data'=>''
			);
		
		//修改密码
		$data = array();
		$data['password'] = sha1($newPassword);
		return $this->userDb->mod($userId,$data);
	}
}
