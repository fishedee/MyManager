var mysql = require("mysql");
var pool = mysql.createPool({
	host:'localhost',
	user:'root',
	password:'1',
	database:'FishMoney',
	port:3306
});
function getConnection(){
	return new Promise(function(resolve,reject){
		pool.getConnection(function(error,client){
			if( !!error )
				reject(error);
			else
				resolve(client);
		});
	});
}
function releaseConnection(connection){
	pool.releaseConnection(connection);
}
function query(connection,sql){
	return new Promise(function(resolve,reject){
		connection.query(sql,function(error,rows){
			if( !!error )
				reject(error);
			else
				resolve(rows);
		});
	});
}
function buildWhereSql(whereSql = null){
	var sql = '';
	if( whereSql ){
		if( typeof whereSql == 'string'){
			sql += 'where ';
			sql += whereSql;
		}else{
			sql += 'where ';
			var temp = whereSql.map(function(value,index){
				return index + ' = "' + value + '"'; 
			});
			sql += temp.join(' and ');
		}
	}
	return sql+' ';
}
function buildValueSql(valueSql=null){
	var sql = '';
	if( valueSql ){
		if( typeof valueSql == 'string'){
			sql += valueSql;
		}else{
			sql += '('+valueSql[0].keys().join(',')+')';
			var tempSql = valueSql.map(function(value,index){
				return '('+value.values().join(',')+')';
			});
			sql += 'values' + tempSql.join(',');
		}
	}
	return sql+' ';
}
function buildSetSql(setSql = null){
	var sql = '';
	if( setSql ){
		if( typeof setSql == 'string'){
			sql += setSql;
		}else{
			sql += 'set ';
			sql += setSql.map(function(value,index){
				return index + ' = ' + value;
			}).join(',');
		}
	}
	return sql+' ';
}
function buildSelectSql(selectSql=null){
	var sql = '';
	if( selectSql ){
		if( typeof selectSql == 'string'){
			sql += selectSql;
		}else{
			sql += selectSql.join(',');
		}
	}else{
		sql += '*';
	}
	return sql+' ';
}
function buildOrderSql(orderSql=null){
	var sql = '';
	if( orderSql ){
		if( typeof orderSql == 'string'){
			sql += 'order by ';
			sql += orderSql;
		}else{
			sql += 'order by ';
			var tempSql = orderSql.map(function(value,index){
				return index+' '+value;
			});
			sql += tempSql.join(',');
		}
	}
	return sql + ' ';
}
function buildLimitSql(limitSql=null){
	var sql = '';
	if( limitSql && limitSql.pageIndex && limitSql.pageSize )
		sql += ' limit '+limitSql.pageSize+','+limitSql.pageIndex;
	return sql+' ';
}
module.exports = {
	async select(sql){
		try{
			var connection = await getConnection();
			var result = await query(connection,sql);
			releaseConnection(connection);
			return result;
		}catch(err){
			if( connection )
				releaseConnection(connection);
			throw err;
		}
	},
	async insert(sql){
		try{
			var connection = await getConnection();
			var result = await query(connection,sql);
			releaseConnection(connection);
			return result.insertId;
		}catch(err){
			if( connection )
				releaseConnection(connection);
			throw err;
		}
	},
	async delete(sql){
		try{
			var connection = await getConnection();
			var result = await query(connection,sql);
			releaseConnection(connection);
			return result.affectedRows;
		}catch(err){
			if( connection )
				releaseConnection(connection);
			throw err;
		}
	},
	async update(sql){
		try{
			var connection = await getConnection();
			var result = await query(connection,sql);
			releaseConnection(connection);
			return result.affectedRows;
		}catch(err){
			if( connection )
				releaseConnection(connection);
			throw err;
		}
	},
	buildSelect(tableName,whereSql=null,selectSql=null,orderSql=null,limitSql=null){
		var sql = 'select ';
		sql += buildSelectSql(selectSql);
		sql += ' from '+tableName+' ';
		sql += buildWhereSql(selectSql);
		sql += buildOrderSql(orderSql);
		sql += buildLimitSql(limitSql);
		return sql;
	},
	buildInsert(tableName,valueSql=[]){
		var sql = 'insert into '+tableName+' ';
		sql += buildValueSql(valueSql);
		return sql;
	},
	buildDelete(tableName,whereSql=null){
		var sql = 'delete from '+tableName+' ';
		sql += buildWhereSql(whereSql);
		return sql;
	},
	buildUpdate(tableName,setSql=null,whereSql=null){
		var sql = 'update '+tableName+' ';
		sql += buildSetSql(setSql);
		sql += buildWhereSql(whereSql);
		return sql;
	}

}