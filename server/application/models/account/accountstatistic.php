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
		$statistic = array();
		foreach( $result['data'] as $key=>$value ){
			$year = $value['year'];
			$month = $value['month'];
			$type = $value['type'];
			$money = $value['money'];
			$statistic[$year][$month][$type] = $money;
		}
		
		//计算时间段
		$minTime = null;
		$maxTime = null;
		foreach( $result['data'] as $key=>$value ){
			if( $minTime == null || ($minTime['year']*100+$minTime['month'] > $value['year']*100+$value['month'])){
				$minTime = array();
				$minTime['year'] = $value['year'];
				$minTime['month'] = $value['month'];
			}
			if( $maxTime == null || ($maxTime['year']*100+$maxTime['month'] < $value['year']*100+$value['month'])){
				$maxTime = array();
				$maxTime['year'] = $value['year'];
				$maxTime['month'] = $value['month'];
			}
		}
		
		//计算结果
		$data = array();
		for( $year = $maxTime['year'] ; $year >= $minTime['year'] ; $year -- ){
			$minMonth = 1;
			$maxMonth = 12;
			if( $year == $minTime['year'])
				$minMonth = $minTime['month'];
			if( $year == $maxTime['year'])
				$maxMonth = $maxTime['month'];
			for( $month = $maxMonth ; $month >= $minMonth ; $month -- ){
				for( $type = 1 ; $type <= 4 ; $type++){
					$money = 0;
					if( isset($statistic[$year][$month][$type]))
						$money = $statistic[$year][$month][$type];
					$typeMap = array(
						$this->accountDb->TYPE_IN =>'收入',
						$this->accountDb->TYPE_OUT =>'支出',
						$this->accountDb->TYPE_TRANSFER_IN =>'转账收入',
						$this->accountDb->TYPE_TRANSFER_OUT =>'转账支出',
					);
					$data[] = array(
						'name'=>$year.'年'.$month.'月',
						'year'=>$year,
						'month'=>$month,
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
	
	public function getDetailTypeStatistic($userId,$year,$month,$type){
		//获取分析数据
		$result = $this->accountDb->getDetailTypeStatisticByUser($userId ,$month,$year,$type);
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
	
	public function getMonthCardStatistic($userId){
		//获取分析数据
		$result = $this->accountDb->getMonthCardStatisticByUser($userId);
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
			$month = $value['month'];
			$cardId = $value['cardId'];
			$money = 0;
			if( $value['type'] == $this->accountDb->TYPE_IN ||
				$value['type'] == $this->accountDb->TYPE_TRANSFER_IN )
				$money = $value['money'];
			else
				$money = -$value['money'];
			if( !isset($statistic[$year][$month][$cardId]))
				$statistic[$year][$month][$cardId] = $money;
			else
				$statistic[$year][$month][$cardId] += $money;
		}
		
		//计算时间段
		$minTime = null;
		$maxTime = null;
		foreach( $accountInfo as $key=>$value ){
			if( $minTime == null || ($minTime['year']*100+$minTime['month'] > $value['year']*100+$value['month'])){
				$minTime = array();
				$minTime['year'] = $value['year'];
				$minTime['month'] = $value['month'];
			}
			if( $maxTime == null || ($maxTime['year']*100+$maxTime['month'] < $value['year']*100+$value['month'])){
				$maxTime = array();
				$maxTime['year'] = $value['year'];
				$maxTime['month'] = $value['month'];
			}
		}

		//计算结果
		$data = array();
		$statistic2 = array();
		for( $year = $minTime['year'] ; $year <= $maxTime['year'] ; $year ++ ){
			$minMonth = 1;
			$maxMonth = 12;
			if( $year == $minTime['year'])
				$minMonth = $minTime['month'];
			if( $year == $maxTime['year'])
				$maxMonth = $maxTime['month'];
			for( $month = $minMonth ; $month <= $maxMonth ; $month ++ ){
				foreach( $card as $cardId=>$cardData ){
					$money = 0;
					if( isset($statistic[$year][$month][$cardId]))
						$money = $statistic[$year][$month][$cardId];
					if( $year == $minTime['year'] && $month == $minTime['month'] ){
						//初始时间
						$money = $cardData['money'] + intval($money);
					}else{
						//非初始时间
						$lastYear = $year;
						$lastMonth = $month - 1 ;
						if( $lastMonth == 0 ){
							$lastYear = $lastYear - 1;
							$lastMonth = 12;
						}
						$money = $statistic2[$lastYear][$lastMonth][$cardId] + $money;
					}
					$data[] = array(
						'name'=>$year.'年'.$month.'月',
						'year'=>$year,
						'month'=>$month,
						'cardId'=>$cardId,
						'cardName'=>$cardData['name'],
						'money'=>$money
					);
					$statistic2[$year][$month][$cardId] = $money;
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
	
	public function getDetailCardStatistic($userId,$year,$month,$cardId){
		//获取分析数据
		$result = $this->accountDb->getDetailCardStatisticByUser($userId ,$month,$year,$cardId);
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
