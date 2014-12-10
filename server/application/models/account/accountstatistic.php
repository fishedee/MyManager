<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class AccountStatistic extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('account/accountDb','accountDb');
		$this->load->model('category/categoryAo','categoryAo');
		$this->load->model('card/cardAo','cardAo');
	}
	
	public function getMonthTypeStatistic($userId){
		//获取分析数据
		$result = $this->accountDb->getMonthTypeStatisticByUser($userId);
		if( $result['code'] != 0 )
			return $result;
		
		//整理数据	
		foreach( $result['data'] as $key=>&$value ){
			$typeMap = array(
				$this->accountDb->TYPE_IN =>'收入',
				$this->accountDb->TYPE_OUT =>'支出',
				$this->accountDb->TYPE_TRANSFER_IN =>'转账收入',
				$this->accountDb->TYPE_TRANSFER_OUT =>'转账支出',
			);
			$value['typeName'] = $typeMap[$value['type']];
			$value['name'] = $value['year'].'年'.$value['month'].'月';
		}
		
		return $result;
	}
	private function arrangeStatistic( $userId,$statistic ){
		//获取类目数据
		$result = $this->categoryAo->search($userId,array(),array());
		if( $result['code'] != 0 )
			return $result;
		$category = array();
		foreach( $result['data']['data'] as $single ){
			$category[$single['categoryId']] = $single['name'];
		}
		
		//整理数据
		$totalMoney = 0;
		foreach( $statistic as $key=>$value ){
			$totalMoney += $value['money'];
		}
		foreach( $statistic as $key=>&$value ){
			$value['precent'] = ceil($value['money']/$totalMoney*100).'%';
			if( array_key_exists($value['categoryId'],$category))
				$value['categoryName'] = $category[$value['categoryId']];
			else
				$value['categoryName'] = '未分类';
		}
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$statistic
		);
	}
	
	public function getDetailTypeStatistic($userId,$year,$month,$type){
		//获取分析数据
		$result = $this->accountDb->getDetailTypeStatisticByUser($userId ,$month,$year,$type);
		if( $result['code'] != 0 )
			return $result;
		
		return $this->arrangeStatistic($userId,$result['data']);
	}
	
	public function getMonthCardStatistic($userId){
		//获取分析数据
		$result = $this->accountDb->getMonthCardStatisticByUser($userId);
		if( $result['code'] != 0 )
			return $result;
		$statistic = $result['data'];
		
		//获取银卡数据
		$result = $this->cardAo->search($userId,array(),array());
		if( $result['code'] != 0 )
			return $result;
		$card = array();
		foreach( $result['data']['data'] as $single ){
			$card[$single['cardId']] = $single;
		}
		
		//整理每个月的费用统计	
		$statistic2 = array();
		foreach( $statistic as $key=>$value ){
			$year = $value['year'];
			$month = $value['month'];
			$cardId = $value['cardId'];
			$money = 0;
			if( $value['type'] == $this->accountDb->TYPE_IN ||
				$value['type'] == $this->accountDb->TYPE_TRANSFER_IN )
				$money = $value['money'];
			else
				$money = -$value['money'];
			if( !isset($statistic2[$cardId][$year][$month]))
				$statistic2[$cardId][$year][$month] = $money;
			else
				$statistic2[$cardId][$year][$month] += $money;
		}
		
		//根据每个月费用统计出每个月的剩余费用
		$statistic3 = array();
		foreach( $statistic2 as $cardId=>$value ){
			$cardName = '';
			if( array_key_exists($cardId,$card)){
				$last = $card[$cardId]['money'];
				$cardName = $card[$cardId]['name'];
			}else{
				$last = 0;
				$cardName = '未知银行卡';
			}
			foreach( $value as $year=>$value2 ){
				foreach( $value2 as $month=>$money ){
					$last = $last + $money;
					$statistic3[] = array(
						'year'=>$year,
						'month'=>$month,
						'name'=>$year.'年'.$month.'月',
						'cardId'=>$cardId,
						'cardName'=>$cardName,
						'money'=>$last,
					);
				}
			}
		}
		
		//排序
		usort($statistic3,function($left,$right){
			if( $left['year'] != $right['year']){
				return $left['year'] > $right['year'] ? -1 :1;
			}else if( $left['month'] != $right['month']){
				return $left['month'] > $right['month'] ? -1 :1;
			}else if( $left['cardId'] != $right['cardId']){
				return $left['cardId'] < $right['cardId'] ? -1 :1;
			}else{
				return 0;
			}
		});
		
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$statistic3
		);
	}

}
