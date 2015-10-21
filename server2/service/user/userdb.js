var express = require('express');
var database = require('../../config/database');
var router = express.Router();

var tableName = 't_user';
module.exports = {
	async search(where,limit){
		var whereSql = '';
		for( var i in where ){
			if( whereSql != '')
				whereSql += ' and ';
			if( i == 'name')
				whereSql += 'name like "'+where[i]+'"';
			else if( i == 'type' )
				whereSql += 'type = "'+where[i]+'"';
		}

		var sql = database.buildSelect(tableName,whereSql,'count(*) as count');
		var count = (await database.select(sql))[0]['count'];

		var sql = database.buildSelect(tableName,whereSql,'*','createTime desc',limit);
		var data = await database.select(sql);

		return {
			count:count,
			data:data
		}
	},
	async get(userId){
		var sql = database.buildSelect(tableName,{userId:userId});
		var data = await database.select(sql);
		if( data.length == 0 )
			throw new Error('找不到用户数据'+userId);
		return data[0];
	},
	async add(data){
		var sql = database.buildInsert(tableName,[data]);
		return await database.insert(sql);
	},
	async mod(userId,data){
		var sql = database.buildUpdate(tableName,data,{userId:userId});
		return await database.update(sql);
	},
	async del(userId){
		var sql = database.buildDelete(tableName,{userId:userId});
		return await database.delete(sql);
	},
	async modPassword(userId , oldPassword , newPassword ){
		var sql = database.buildUpdate(tableName,{password:newPassword},{userId:userId,password:oldPassword});
		var affectedRows = await database.update(sql);
		if( affectedRows <= 0 )
			throw new Error('密码错误');
	},
	async getByIdAndPass(userId,password){
		var sql = database.buildSelect(tableName,{userId:userId,password:password});
		return await database.select(sql);
	},
	async getByNameAndPass(name,password){
		var sql = database.buildSelect(tableName,{name:name,password:password});
		return await database.select(sql);
	},
	async getByName(name){
		var sql = database.buildSelect(tableName,{name:name});
		return await database.select(sql);
	}
};
