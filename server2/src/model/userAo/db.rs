use crate::util::db::Pool;
use crate::util::error::Error;
use crate::util::string::implode;
use mysql_async::prelude::*;
use futures::future::{ok,err,Future};
use super::data;

/*
pub fn add<'a>(db:&'a Pool,userAdd:&'a data::UserAdd)->impl Future<Item=u64,Error=Error>+'a{
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		conn.prep_exec("insert into t_user value(?,?,?)",(userAdd.name,userAdd.password,userAdd.r#type))
	}).map(|result|{
		result.last_insert_id().unwrap()
	}).map_err(|e|{
		Error::new(500,format!("{:?}",e))
	});
}*/

pub fn r#mod<'a>(prev:impl Future<Item=(&'a Pool,&'a data::UserMod)>)->impl Future<Item=(),Error=Error>+'a{
	prev.and_then(|(db,data)|{
		let conn = db.get_conn().map_err(|e|{
			Error::new(500,format!("{:?}",e))
		});
		let data = ok(data);
		return conn.join(data);
	}).and_then(|(conn,userMod)|{
		let sql = "update t_user set name = ? and password = ? and type = ? where userId = ?";
		let args = (userMod.name,userMod.password,userMod.r#type,userMod.userId);
		return conn.prep_exec(sql,args).map_err(|e|{
			Error::new(500,format!("{:?}",e))
		});
	});
}
/*
pub fn r#mod(db:& Pool,userMod:& data::UserMod)->impl Future<Item=(),Error=Error>{
	let sql = "update t_user set name = ? and password = ? and type = ? where userId = ?";
	let argv = (userMod.name.clone(),userMod.password.clone(),userMod.r#type,userMod.userId);
	let conn = db.get_conn();
	
}
pub fn r#del(db:&Pool,userId:u64)->impl Future<Item=(),Error=Error>{
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		conn.prep_exec("delete t_user where userId = ?",(userId,))
	}).map(|result|{
		()
	}).map_err(|e|{
		Error::new(500,format!("{:?}",e))
	});
}


fn getWhere(search:&data::UserSearch)->String{
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


pub fn search(db:&Pool,search:&data::UserSearch)->impl Future<Item=data::Users,Error=Error>{

	let whereSql = getWhere(&search);
	let dataSql = format!("select userId,name,password,type,createTime,modifyTime from t_user {} limit {},{}",whereSql,search.pageIndex,search.pageSize);
	let countSql = format!("select count(*) from t_user {}",whereSql);

	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		return conn.query(dataSql).and_then(|data|{
			data.collect_and_drop::<(u64,String,String,u64,String,String)>()
		}).map(move|(conn, data)|{
			let rows = data.into_iter().map(|single|{
				return data::User{
					userId:single.0,
					name:single.1,
					password:single.2,
					r#type:single.3,
					createTime:single.4,
					modifyTime:single.5,
				};
			}).collect::<Vec<data::User>>();
			return (conn,rows);
		});
	}).and_then(move|(conn,rows)|{
			return conn.query(countSql)
		.and_then(|data|{
			data.collect::<u64>()
		}).and_then(move|(_,mut data)|{
			let single = data.pop().unwrap();
			return ok(data::Users{
				data:rows,
				count:single,
			});
		})
	}).map_err(|e|{
		return Error::new(500,format!("{:?}",e));
	});
}*/

pub fn get<'a>(prev:impl Future<Item=(&'a Pool,u64),Error=Error>)->impl Future<Item=data::User,Error=Error>+'a{
	return prev.and_then(|(db,userId)|{
		return db.get_conn()
			.and_then(move|conn|{
				return conn.prep_exec("select userId,name,password,type,createTime,modifyTime from t_user where userId = ?",(userId,))
			}).map_err(|e|{
				return Error::new(500,format!("{:?}",e));
			}).and_then(|data|{
				return data.collect::<(u64,String,String,u64,String,String)>()
					.map_err(|e|{
						return Error::new(500,format!("{:?}",e));
					})
					.and_then(|(_,mut data)|{
						if data.len() == 0{
							return err(Error::new(1,"user dos not exist!"));
						}else{
							let single = data.pop().unwrap();
							return ok(data::User{
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
	});
}


/*
pub fn get(db:&Pool,userId:u64)->impl Future<Item=data::User,Error=Error>{
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		return conn.prep_exec("select userId,name,password,type,createTime,modifyTime from t_user where userId = ?",(userId,))
	})
}


pub fn getByName<'a>(db:&'a Pool,name:&'a str)->impl Future<Item=Vec<data::User>,Error=Error>+'a {
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		return conn.prep_exec("select userId,name,password,type,createTime,modifyTime from t_user where name = ?",(name,))
	}).map_err(|e|{
		return Error::new(500,format!("{:?}",e));
	}).and_then(|data|{
		return data.collect::<(u64,String,String,u64,String,String)>()
			.map_err(|e|{
				return Error::new(500,format!("{:?}",e));
			})
			.map(|(_,mut data)|{
				return data.into_iter().map(|single|{
					return data::User{
						userId:single.0,
						name:single.1,
						password:single.2,
						r#type:single.3,
						createTime:single.4,
						modifyTime:single.5,
					}; 
				}).collect::<Vec<data::User>>();
			});
	});
}*/