use crate::util::db::Pool;
use crate::util::session::Session;
use crate::util::error::Error;
use crate::model::userAo;
use futures::future::{ok,err,Future};
use super::data;
use chrono::naive::NaiveDateTime;

pub fn isLogin(db:&Pool,session:&Session)->impl Future<Item=userAo::User,Error=Error>{
	let hasUserId = session.get::<u64>("userId").unwrap();
	let newDb = db.clone();
	return ok(())
		.and_then(move|_|{
			let result:Box<dyn Future<Item=userAo::User,Error=Error>>;
			if let Some(userId) = hasUserId{
				result = Box::new(userAo::get(&newDb,userId));
			}else{
				result = Box::new(ok(userAo::User{
					userId:0,
					name:"".to_string(),
					password:"".to_string(),
					r#type:0,
					createTime:NaiveDateTime::from_timestamp(0,0),
					modifyTime:NaiveDateTime::from_timestamp(0,0),
				})); 
			}
			return result;
		})
}

pub fn checkMustLogin(db:&Pool,session:&Session)->impl Future<Item=userAo::User,Error=Error>{
	return isLogin(db,session)
		.and_then(|user|{
			if user.userId == 0{
				return Box::new(err(Error::new(1,"用户未登录")));
			}else{
				return Box::new(ok(user));
			}
		})
}

pub fn checkMustAdmin(db:&Pool,session:&Session)->impl Future<Item=userAo::User,Error=Error>{
	return checkMustLogin(db,session)
		.and_then(|user|{
			if user.r#type != 1{
				return err(Error::new(1,"非管理员没有权限执行此操作"));
			}else{
				return ok(user);
			}
		})
}

pub fn logout(db:&Pool,session:&Session)->impl Future<Item=(),Error=Error>{
	let session = session.clone();
	return ok(())
		.map(move|_|{
			session.set("userId",0 as u64).unwrap();
			return ();
		})
}

pub fn login(db:&Pool,session:&Session,loginCheckIn:&data::LoginCheckIn)->impl Future<Item=(),Error=Error>{
	let name = loginCheckIn.name.clone();
	let password = loginCheckIn.password.clone();
	let session = session.clone();
	return userAo::getByName(db,&name)
		.and_then(move|users|{
			let result:Box<dyn Future<Item=(),Error=Error>>;
			if users.len() == 0{
				result = Box::new(err(Error::new(1,"不存在此账户")));
			}else{
				let re = userAo::checkMustValidPassword(&password,&users[0].password);
				match re{
					Ok(_)=>{
						result = Box::new(ok(()));
						session.set("userId",users[0].userId).unwrap();
					},
					Err(e)=>{
						result = Box::new(err(e));
					}
				}
			}
			return result;
		});
}