var express = require('express');
var json = require('../view/json');
var loginao = require('../services/user/loginao');
var router = express.Router();

router.get('/islogin', async json(function(req, res,next) {
	return await loginao->islogin();
)});

router.get('/checkout', async json(function(req, res,next) {
	return await loginao->checkout();
)});

router.get('/checkin', async json(function(req, res,next) {
	return await loginao->checkin(req.params.name,req.params.password);
)});

module.exports = router;
