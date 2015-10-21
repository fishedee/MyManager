var express = require('express');
var json = require('../view/json');
var loginao = require('../service/user/loginao');
var usertypeenum = require('../service/user/usertypeenum');
var userao = require('../service/user/userao');
var router = express.Router();

async function checkAdmin(req){
	var data = await loginao.islogin(req);
	if( data.type != usertypeenum.ADMIN )
		throw new Error('你没有权限执行此操作');
}

router.get('/search', json(async function(req, res,next) {
	await checkAdmin(req);

	var where = {};
	if( req.query.name )
		where.name = req.query.name;
	if( req.query.type )
		where.type = req.query.type;

	var limit = {};
	if( req.query.pageIndex )
		limit.pageIndex = req.query.pageIndex;
	if( req.query.pageSize )
		limit.pageSize = req.query.pageSize;

	return await userao.search(where,limit);
}));

router.get('/get', json(async function(req, res,next) {
	await checkAdmin(req);

	return await userao.get(req.query.userId);
}));

router.post('/add', json(async function(req, res,next) {
	await checkAdmin(req);

	return await userao.add({
		name:req.body.name,
		type:req.body.type,
		password:req.body.password
	});
}));

router.post('/del', json(async function(req, res,next) {
	await checkAdmin(req);

	return await userao.del(req.body.userId);
}));

router.post('/modType', json(async function(req, res,next) {
	await checkAdmin(req);

	return await userao.modType(req.body.userId,req.body.type);
}));

router.post('/modPassword', json(async function(req, res,next) {
	await checkAdmin(req);

	return await userao.modPassword(req.body.userId,req.body.password);
}));

router.post('/modMyPassword', json(async function(req, res,next) {
	var userId = (await loginao.islogin(req)).userId;

	return await userao.modPasswordByOld(userId,req.body.oldPassword,req.body.newPassword);
}));

module.exports = router;

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
