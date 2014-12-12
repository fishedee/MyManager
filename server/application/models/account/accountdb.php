<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class AccountDb extends CI_Model {
	var $tableName = "t_account";
	var $TYPE_IN = 1;
	var $TYPE_OUT = 2;
	var $TYPE_TRANSFER_IN = 3;
	var $TYPE_TRANSFER_OUT = 4;
	public function __construct(){
		parent::__construct();
	}

	public function search($where,$limit){
		foreach( $where as $key=>$value ){
			if( $key == "name" || $key == 'remark' )
				$this->db->like($key,$value);
			else if( $key == "categoryId" || $key == 'cardId' || $key == 'type' || $key == "userId")
				$this->db->where($key,$value);
		}
		
		$count = $this->db->count_all_results($this->tableName);
		
		foreach( $where as $key=>$value ){
			if( $key == "name" || $key == 'remark' )
				$this->db->like($key,$value);
			else if( $key == "categoryId" || $key == 'cardId' || $key == 'type' || $key == "userId")
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

	public function get($accountId){
		$this->db->where("accountId",$accountId);
		$query = $this->db->get($this->tableName)->result_array();
		if( count($query) == 0 )
			return array(
					"code"=>1,
					"msg"=>"不存在此数据",
					"data"=>""
				    );
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>$query[0]
			    );
	}

	public function del( $accountId ){
		$this->db->where("accountId",$accountId);
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

	public function mod( $accountId , $data ){
		$this->db->where("accountId",$accountId);
		$query = $this->db->update($this->tableName,$data);
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>""
			    );
	}
	
	public function resetCategory( $categoryId ){
		$data = array('categoryId'=>0);
		$this->db->where("categoryId",$categoryId);
		$query = $this->db->update($this->tableName,$data);
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>""
			    );
	}
	
	public function resetCard( $cardId ){
		$data = array('cardId'=>0);
		$this->db->where("cardId",$cardId);
		$query = $this->db->update($this->tableName,$data);
		return array(
				"code"=>0,
				"msg"=>"",
				"data"=>""
			    );
	}
	
	public function getMonthTypeStatisticByUser( $userId ){
		$sql = "select DATE_FORMAT(createTime,'%Y') as year,DATE_FORMAT(createTime,'%m') as month,type,SUM(money) as money ".
			'from '.$this->tableName.' '.
			'where userId = ? '.
			'group by year,month,type '.
			'order by year desc,month desc';
		$argv = array($userId);
		$query = $this->db->query($sql,$argv)->result_array();
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$query
		);
	}
	
	public function getDetailTypeStatisticByUser( $userId ,$month,$year,$type){
		$sql = "select categoryId , sum(money) as money ".
			'from '.$this->tableName.' '.
			"where DATE_FORMAT(createTime,'%Y') = ? and DATE_FORMAT(createTime,'%m') = ? and userId = ? and type = ? ".
			'group by categoryId';
		$argv = array($year,$month,$userId,$type);
		$query = $this->db->query($sql,$argv)->result_array();
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$query
		);
	}
	
	public function getMonthCardStatisticByUser( $userId ){
		$sql = "select DATE_FORMAT(createTime,'%Y') as year,DATE_FORMAT(createTime,'%m') as month,cardId,type,SUM(money) as money ".
			'from '.$this->tableName.' '.
			'where userId = ? '.
			'group by year,month,cardId,type '.
			'order by year asc,month asc';
		$argv = array($userId);
		$query = $this->db->query($sql,$argv)->result_array();
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$query
		);
	}
	
	public function getDetailCardStatisticByUser( $userId ,$month,$year,$cardId){
		$sql = "select type , sum(money) as money ".
			'from '.$this->tableName.' '.
			"where DATE_FORMAT(createTime,'%Y') = ? and DATE_FORMAT(createTime,'%m') = ? and userId = ? and cardId = ? ".
			'group by type';
		$argv = array($year,$month,$userId,$cardId);
		$query = $this->db->query($sql,$argv)->result_array();
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$query
		);
	}

}
