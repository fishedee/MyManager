var express = require('express');
var json = require('../view/json');
var loginao = require('../service/user/loginao');
var router = express.Router();

router.get('/islogin',json(async function(req, res,next) {
	return await loginao.islogin(req);
}));

router.get('/checkout',json(async function(req, res,next) {
	return await loginao.logout(req);
}));

router.post('/checkin',json(async function(req, res,next) {
	return await loginao.login(req,req.body.name,req.body.password);
}));

module.exports = router;
