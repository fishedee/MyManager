var express = require('express');
var json = require('../view/json');
var loginao = require('../services/user/loginao');
var usertypeenum = require('../services/user/usertypeenum');
var userao = require('../services/user/userao');
var router = express.Router();

async function checkAdmin(){
	var data = await loginao.islogin();
	if( data.type != usertypeenum.ADMIN )
		throw new Error('你没有权限执行此操作');
}

router.get('/search', json(async function(req, res,next) {
	await checkAdmin();

	var where = {};
	if( req.params.name )
		where.name = req.params.name;
	if( req.params.type )
		where.type = req.params.type;

	var limit = {};
	if( req.params.pageIndex )
		limit.pageIndex = req.params.pageIndex;
	if( req.params.pageSize )
		limit.pageSize = req.params.pageSize;

	return await userao.search(where,limit);
)});

router.get('/get', json(async function(req, res,next) {
	await checkAdmin();

	return await userao.get(req.params.userId);
)});

router.post('/add', json(async function(req, res,next) {
	await checkAdmin();

	return await userao.add({
		name:req.params.name,
		type:req.params.type,
		password:req.params.password
	});
)});

router.post('/del', json(async function(req, res,next) {
	await checkAdmin();

	return await userao.del(req.params.userId);
)});

router.post('/modType', json(async function(req, res,next) {
	await checkAdmin();

	return await userao.modType(req.params.userId,req.params.type);
)});

router.post('/modPassword', json(async function(req, res,next) {
	await checkAdmin();

	return await userao.modPassword(req.params.userId,req.params.password);
)});

router.post('/modMyPassword', json(async function(req, res,next) {
	await userId = loginao.islogin().userId;

	return await userao.modPassword(userId,req.params.oldPassword,req.params.newPassword);
)});

module.exports = router;

/* End of file welcome.php */
/* Location: ./application/controllers/welcome.php */
