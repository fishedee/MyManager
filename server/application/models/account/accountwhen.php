<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class AccountWhen extends CI_Model {

	public function __construct(){
		parent::__construct();
		$this->load->model('account/accountDb','accountDb');
	}
	
	public function whenCategoryDelete( $categoryId ){
		return $this->accountDb->resetCategory($categoryId);
	}
	
	public function whenCardDelete( $cardId ){
		return $this->accountDb->resetCard($cardId);
	}	

}
