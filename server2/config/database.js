var mysql = require("mysql");
var pool = mysql.createPool({
	host:'localhost',
	user:'root',
	password:'1',
	database:'FishMoney',
	port:3306
});
module.exports = {
	getConnection(){
		return new Promise(function(resolve,reject){
			pool.getConnection(function(error,client){
				if( !!error )
					reject(error);
				else
					resolve(client);
			});
		});
	},
	releaseConnection(client){
		pool.releaseConnection(client);
	},
	query(connection,sql){
		return new Promise(function(resolve,reject){
			connection.query(sql,function(error,rows){
				if( !!error )
					reject(error);
				else
					resolve(rows);
			});
		});
	},
}