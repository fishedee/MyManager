use crate::util::db::Pool;
use crate::util::error::Error;
use mysql_async::prelude::*;
use futures::future::{ok,err,Future};

pub struct User{
	userId:i32,
	name:i32,
	password:String,
	r#type:i32,
	createTime:String,
	modifyTime:String,
}

pub struct Users{
	Count:u32,
	Data:Vec<User>,
}

pub struct UserSearch{

}

pub fn search(db:&Pool,search:&UserSearch)->Result<Users,Error>{
	return Ok(Users{
		Count:0,
		Data:Vec::new()
	});
}

pub fn get(db:&Pool,userId:i32)->impl Future<Item=User,Error=Error>{
	let conn = db.get_conn();
	return conn.and_then(move|conn|{
		let sql = format!("select userId,name,password,type,createTime,modifyTime from t_user where userId = {}",userId);
		return conn.query(sql).map_err(|e|{
			Error::new(500,format!("{:?}",e));
		}).and_then(|data|{
			return data.collect::<(String,String)>();
		}).and_then(|_,data|{
			if data.len() == 0{
				return err(Error::new(1,"不存在该用户"));
			}else{
				return ok(data[0]);
			}
		});
	});
}