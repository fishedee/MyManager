var express = require('express');
var mysql = require('../config/database.js');
var router = express.Router();

router.get('/',async function(req, res,next) {
	try{
		var connection = await mysql.getConnection();
		var data = await mysql.query(connection,'select * from t_user');
		res.send('this is index'+JSON.stringify(data));
	}catch(e){
		console.log(e);
		next(e);
	}
});

module.exports = router;
