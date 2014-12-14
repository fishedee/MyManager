<?php if ( ! defined('BASEPATH')) exit('No direct script access allowed');

class Welcome extends CI_Controller {

	public function __construct()
    {
        parent::__construct();
		$this->load->model('user/loginAo','loginAo');
    }
	
	public function index()
	{
		$result = $this->loginAo->islogin();
		if( $result["code"] == 0 ){
			require_once(dirname(__FILE__).'/../../../static/build/backstage/index.html');
		}else{
			require_once(dirname(__FILE__).'/../../../static/build/backstage/login.html');
		}
	}
	
}

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
