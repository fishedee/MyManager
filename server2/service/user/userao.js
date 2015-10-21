var userdb = require('./userdb');
var hash = require('../../config/hash');

module.exports = {
	search:userdb.search,
	get:userdb.get,
	del:userdb.del,
	async add(data){
		var result = await userdb.getByName(data.name);
		if( result.length != 0 )
			throw new Error('存在重复的用户名');

		data.password = hash.sha1(data.password);
		return await userdb.add(data);
	},
	async modType(userId,type){
		await userdb.mod(userId,{type:type});
	},
	async modPassword(userId,password){
		await userdb.mod(userId,{password:hash.sha1(password)});
	},
	async modPasswordByOld(userId,oldPassword,newPassword){
		var data = await userdb.getByIdAndPass(userId,hash.sha1(oldPassword));
		if( data.length == 0 )
			throw new Error('原密码错误');
		await userdb.mod(userId,{password:hash.sha1(newPassword)});
	}
};
