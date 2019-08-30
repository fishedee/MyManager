use crate::util::db::Pool;
use crate::util::error::Error;
use crate::util::string::implode;
use mysql_async::prelude::*;
use futures::future::{ok,err,Future};
use serde::{Serialize,Deserialize};

#[derive(Serialize)]
pub struct User{
	userId:u32,
	name:String,
	password:String,
	r#type:u32,
	createTime:String,
	modifyTime:String,
}

#[derive(Serialize)]
pub struct Users{
	count:u32,
	data:Vec<User>,
}

#[derive(Deserialize)]
pub struct UserSearch{
	userId:Option<u32>,
	r#type:Option<u32>,
	pageIndex:u32,
	pageSize:u32,
}

fn getWhere(search:&UserSearch)->String{
	let mut sql_vec:Vec<String> = Vec::new();
	if let Some(userId) = search.userId{
		sql_vec.push(format!("userId = {}",userId));
	}
	if let Some(r#type) = search.r#type{
		sql_vec.push(format!("type = {}",r#type));
	}
	let mut sql_str = implode(&sql_vec," and ");
	if sql_str.len() != 0{
		sql_str = "where ".to_string()+&sql_str;
	}
	return sql_str;
}


pub fn search(db:&Pool,search:&UserSearch)->impl Future<Item=Users,Error=Error>{

	let whereSql = getWhere(&search);
	let dataSql = format!("select userId,name,password,type,createTime,modifyTime from t_user {} limit {},{}",whereSql,search.pageIndex,search.pageSize);
	let countSql = format!("select count(*) from t_user {}",whereSql);

	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		return conn.query(dataSql).and_then(|data|{
			data.collect_and_drop::<(u32,String,String,u32,String,String)>()
		}).map(move|(conn, data)|{
			let rows = data.into_iter().map(|single|{
				return User{
					userId:single.0,
					name:single.1,
					password:single.2,
					r#type:single.3,
					createTime:single.4,
					modifyTime:single.5,
				};
			}).collect::<Vec<User>>();
			return (conn,rows);
		});
	}).and_then(move|(conn,rows)|{
			return conn.query(countSql)
		.and_then(|data|{
			data.collect::<u32>()
		}).and_then(move|(_,mut data)|{
			let single = data.pop().unwrap();
			return ok(Users{
				data:rows,
				count:single,
			});
		})
	}).map_err(|e|{
		return Error::new(500,format!("{:?}",e));
	});
}

pub fn get(db:&Pool,userId:i32)->impl Future<Item=User,Error=Error>{
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		let sql = format!("select userId,name,password,type,createTime,modifyTime from t_user where userId = {}",userId);
		return conn.query(sql)
	}).map_err(|e|{
		return Error::new(500,format!("{:?}",e));
	}).and_then(|data|{
		return data.collect::<(u32,String,String,u32,String,String)>()
			.map_err(|e|{
				return Error::new(500,format!("{:?}",e));
			})
			.and_then(|(_,mut data)|{
				if data.len() == 0{
					return err(Error::new(1,"user dos not exist!"));
				}else{
					let single = data.pop().unwrap();
					return ok(User{
						userId:single.0,
						name:single.1,
						password:single.2,
						r#type:single.3,
						createTime:single.4,
						modifyTime:single.5,
					});
				}
			});
	});
}