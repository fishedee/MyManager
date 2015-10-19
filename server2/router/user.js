var express = require('express');
var session = require('../config/session');
var router = express.Router();

router.get('/', async function(req, res) {
	try{
		var sessionData = await session.getSessionData(req);
		if( sessionData.user ){
			res.send("You are login as "+sessionData.user);
		}else{
			res.send("You are not login");
		}
	}catch(e){
		console.log(e);
	}
});

router.get('/login',async function(req,res){
	try{
		var data = {
			user:new Date().valueOf()
		}
		await session.setSessionData(req,data);
		res.send("You have login in as "+data.user);
	}catch(e){
		console.log(e);
	}
});

router.get('/logout',async function(req,res){
	try{
		await session.destorySession(req);
		res.send("You have logout");
	}catch(e){
		console.log(e);
	}
});

module.exports = router;
