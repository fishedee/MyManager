<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class AccountStatistic extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('account/accountDb','accountDb');
		$this->load->model('category/categoryAo','categoryAo');
		$this->load->model('card/cardAo','cardAo');
	}
	public function getWeekTypeStatistic($userId){
		//获取分析数据
		$result = $this->accountDb->getWeekTypeStatisticByUser($userId);
		if( $result['code'] != 0 )
			return $result;
		
		//整理数据	
		$statistic = array();
		foreach( $result['data'] as $key=>$value ){
			$year = $value['year'];
			$week = $value['week'];
			$type = $value['type'];
			$money = $value['money'];
			$statistic[$year][$week][$type] = $money;
		}
		
		//计算时间段
		$minTime = null;
		$maxTime = null;
		foreach( $result['data'] as $key=>$value ){
			if( $minTime == null || ($minTime['year']*100+$minTime['week'] > $value['year']*100+$value['week'])){
				$minTime = array();
				$minTime['year'] = $value['year'];
				$minTime['week'] = $value['week'];
			}
			if( $maxTime == null || ($maxTime['year']*100+$maxTime['week'] < $value['year']*100+$value['week'])){
				$maxTime = array();
				$maxTime['year'] = $value['year'];
				$maxTime['week'] = $value['week'];
			}
		}
		
		//计算结果
		$data = array();
		for( $year = $maxTime['year'] ; $year >= $minTime['year'] ; $year -- ){
			$minWeek = 1;
			$maxWeek = 52;
			if( $year == $minTime['year'])
				$minWeek = $minTime['week'];
			if( $year == $maxTime['year'])
				$maxWeek = $maxTime['week'];
			for( $week = $maxWeek ; $week >= $minWeek ; $week -- ){
				for( $type = 1 ; $type <= 4 ; $type++){
					$money = 0;
					if( isset($statistic[$year][$week][$type]))
						$money = $statistic[$year][$week][$type];
					$typeMap = array(
						$this->accountDb->TYPE_IN =>'收入',
						$this->accountDb->TYPE_OUT =>'支出',
						$this->accountDb->TYPE_TRANSFER_IN =>'转账收入',
						$this->accountDb->TYPE_TRANSFER_OUT =>'转账支出',
					);
					$data[] = array(
						'name'=>$year.'年'.$week.'周',
						'year'=>$year,
						'week'=>$week,
						'type'=>$type,
						'typeName'=>$typeMap[$type],
						'money'=>$money
					);
				}
			}
		}
		
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$data
		);
	}
	
	public function getWeekDetailTypeStatistic($userId,$year,$week,$type){
		//获取分析数据
		$result = $this->accountDb->getWeekDetailTypeStatisticByUser($userId ,$year,$week,$type);
		if( $result['code'] != 0 )
			return $result;
		$statistic = $result['data'];
		
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
	
	public function getWeekCardStatistic($userId){
		//获取分析数据
		$result = $this->accountDb->getWeekCardStatisticByUser($userId);
		if( $result['code'] != 0 )
			return $result;
		$accountInfo = $result['data'];
		
		//获取银卡数据
		$result = $this->cardAo->search($userId,array(),array());
		if( $result['code'] != 0 )
			return $result;
		$card = array();
		foreach( $result['data']['data'] as $single ){
			$card[$single['cardId']] = $single;
		}
		foreach( $accountInfo as $key=>$value ){
			if( $value['cardId'] == 0 ){
				$card[0] = array(
					'cardId'=>0,
					'name'=>'无银行卡',
					'money'=>0,
				);
				break;
			}
		}
		
		//整理每个月的费用统计	
		$statistic = array();
		foreach( $accountInfo as $key=>$value ){
			$year = $value['year'];
			$week = $value['week'];
			$cardId = $value['cardId'];
			$money = 0;
			if( $value['type'] == $this->accountDb->TYPE_IN ||
				$value['type'] == $this->accountDb->TYPE_TRANSFER_IN )
				$money = $value['money'];
			else
				$money = -$value['money'];
			if( !isset($statistic[$year][$week][$cardId]))
				$statistic[$year][$week][$cardId] = $money;
			else
				$statistic[$year][$week][$cardId] += $money;
		}
		
		//计算时间段
		$minTime = null;
		$maxTime = null;
		foreach( $accountInfo as $key=>$value ){
			if( $minTime == null || ($minTime['year']*100+$minTime['week'] > $value['year']*100+$value['week'])){
				$minTime = array();
				$minTime['year'] = $value['year'];
				$minTime['week'] = $value['week'];
			}
			if( $maxTime == null || ($maxTime['year']*100+$maxTime['week'] < $value['year']*100+$value['week'])){
				$maxTime = array();
				$maxTime['year'] = $value['year'];
				$maxTime['week'] = $value['week'];
			}
		}

		//计算结果
		$data = array();
		$statistic2 = array();
		for( $year = $minTime['year'] ; $year <= $maxTime['year'] ; $year ++ ){
			$minWeek = 1;
			$maxWeek = 52;
			if( $year == $minTime['year'])
				$minWeek = $minTime['week'];
			if( $year == $maxTime['year'])
				$maxWeek = $maxTime['week'];
			for( $week = $minWeek ; $week <= $maxWeek ; $week ++ ){
				foreach( $card as $cardId=>$cardData ){
					$money = 0;
					if( isset($statistic[$year][$week][$cardId]))
						$money = $statistic[$year][$week][$cardId];
					if( $year == $minTime['year'] && $week == $minTime['week'] ){
						//初始时间
						$money = $cardData['money'] + intval($money);
					}else{
						//非初始时间
						$lastYear = $year;
						$lastWeek = $week - 1 ;
						if( $lastWeek == 0 ){
							$lastYear = $lastYear - 1;
							$lastWeek = 52;
						}
						$money = $statistic2[$lastYear][$lastWeek][$cardId] + $money;
					}
					$data[] = array(
						'name'=>$year.'年'.$week.'周',
						'year'=>$year,
						'week'=>$week,
						'cardId'=>$cardId,
						'cardName'=>$cardData['name'],
						'money'=>$money
					);
					$statistic2[$year][$week][$cardId] = $money;
				}
			}
		}
		$data = array_reverse($data);
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$data
		);
	}
	
	public function getWeekDetailCardStatistic($userId,$year,$week,$cardId){
		//获取分析数据
		$result = $this->accountDb->getWeekDetailCardStatisticByUser($userId ,$year,$week,$cardId);
		if( $result['code'] != 0 )
			return $result;
		$statistic = $result['data'];
		
		//整理数据
		$totalMoney = 0;
		foreach( $statistic as $key=>$value ){
			$totalMoney += $value['money'];
		}
		foreach( $statistic as $key=>&$value ){
			$typeMap = array(
				$this->accountDb->TYPE_IN =>'收入',
				$this->accountDb->TYPE_OUT =>'支出',
				$this->accountDb->TYPE_TRANSFER_IN =>'转账收入',
				$this->accountDb->TYPE_TRANSFER_OUT =>'转账支出',
			);
			$value['typeName'] = $typeMap[$value['type']];
			$value['precent'] = ceil($value['money']/$totalMoney*100).'%';
		}
		return array(
			'code'=>0,
			'msg'=>'',
			'data'=>$statistic
		);
	}

}
