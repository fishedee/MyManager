var userdb = require('./userdb');
var session = require('../../config/session');
var hash = require('../../config/hash');
module.exports = {
	async islogin(req){
		var sessionData = await session.getSessionData(req);
		if( sessionData.userId && sessionData.userId >= 1000 ){
			return await userdb.get(sessionData.userId);
		}else{
			throw new Error('账号未登陆');
		}
	},
	async logout(req){
		session.destorySession(req);
	},
	async login(req,name,password){
		var data = await userdb.getByNameAndPass(name,hash.sha1(password));
		if( data.length == 0 )
			throw new Error('账号或密码错误');
		await session.setSessionData(req,{userId:data[0].userId});
	}
};