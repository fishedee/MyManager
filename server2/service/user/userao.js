var userdb = require('./userdb');

module.exports = {
	search:userdb.search,
	get:userdb.get,
	del:userdb.del,
	async add(data){
		var result = await userdb.getByName(data.name);
		if( result.length != 0 )
			throw new Error('存在重复的用户名');

		return await this.userdb.add(data);
	},
	async modType(userId,type){
		await userdb.mod(userId,{type:type});
	},
	async modPassword(userId,password){
		await userdb.mod(userId,{password:password});
	},
	async modPasswordByOld(userId,oldPassword,newPassword){
		var data = await this.userdb.getByIdAndPass(userId,oldPassword);
		if( data.length == 0 )
			throw new Error('原密码错误');
		await userdb.mod(userId,{password:newPassword});
	}
};
