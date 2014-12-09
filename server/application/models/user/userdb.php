<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class UserDb extends CI_Model 
{
	var $tableName = "t_user";
	var $TYPE_ADMIN = 0;
	var $TYPE_USER = 1;

	public function __construct(){
		parent::__construct();
	}

	public function search($where,$limit){
		foreach( $where as $key=>$value ){
			if( $key == "name" )
				$this->db->like($key,$value);
			else if( $key == "type" )
				$this->db->where($key,$value);
		}
		
		$count = $this->db->count_all_results($this->tableName);
		
		foreach( $where as $key=>$value ){
			if( $key == "name" )
				$this->db->like($key,$value);
			else if( $key == "type" )
				$this->db->where($key,$value);
		}
			
		$this->db->order_by('createTime','desc');
		
		if( isset($limit["pageIndex"]) && isset($limit["pageSize"]))
			$this->db->limit($limit["pageSize"],$limit["pageIndex"]);

		$query = $this->db->get($this->tableName);
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>array(
					"count"=>$count,
					"data"=>$query->result_array()
				)
		);
	}

	public function get($userId){
		$this->db->where("userId",$userId);
		$query = $this->db->get($this->tableName)->result_array();
		if( count($query) == 0 )
			return array(
					"code"=>1,
					"msg"=>"不存在此用户",
					"data"=>""
				    );
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>$query[0]
			    );
	}

	public function del( $userId ){
		$this->db->where("userId",$userId);
		$query = $this->db->delete($this->tableName);
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>""
			);
	}
	
	public function add( $data ){
		$query = $this->db->insert($this->tableName,$data);
		return array(
			"code"=>0,
			"msg"=>"",
			"data"=>""
			);
	}

	public function mod( $userId , $data )
	{
		$this->db->where("userId",$userId);
		$query = $this->db->update($this->tableName,$data);
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>""
			    );
	}

	public function modPassword( $userId , $oldPassword , $newPassword ){
		$data = array();
		$data['password'] = $newPassword;
		$this->db->where("userId",$userId);
		$this->db->where("password",$oldPassword);

		$query = $this->db->update($this->tableName,$data);
		$rows = $this->db->affected_rows();
		if( $rows == 0 ){
			return array(
					"code"=>1,
					"msg"=>"密码错误",
					"data"=>""
				    );
		}

		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>""
			    );
	}

	public function getByIdAndPass($userId,$password){
		$this->db->where("userId",$userId);
		$this->db->where("password",$password);
		$query = $this->db->get($this->tableName)->result_array();
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>$query
			    );
	}
	
	public function getByNameAndPass($name,$password){
		$this->db->where("name",$name);
		$this->db->where("password",$password);
		$query = $this->db->get($this->tableName)->result_array();
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>$query
			    );
	}

	public function getByName($name){
		$this->db->where("name",$name);
		$query = $this->db->get($this->tableName)->result_array();
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>$query
			    );
	}

}
