use crate::util::db::Pool;
use crate::util::error::Error;
use crypto::sha1::Sha1;
use crypto::digest::Digest;
use futures::future::{ok,err,Future};
use super::db;
use super::data;

pub fn checkMustValidPassword(password:&str,passwordHash:&str)->Result<(),Error>{
	if getPasswordHash(password) != passwordHash{
		return Err(Error::new(1,"密码不正确"));
	}else{
		return Ok(());
	}
}

fn getPasswordHash(password:&str)->String{
	let mut hasher = Sha1::new();

	hasher.input_str(password);

	let hex = hasher.result_str();

	return hex;
}

pub fn search(db:&Pool,userSearch:&data::UserSearch)->impl Future<Item=data::Users,Error=Error>{
	return db::search(db,userSearch);
}

pub fn get(db:&Pool,userId:u64)->impl Future<Item=data::User,Error=Error>{
	return db::get(db,userId);
}

pub fn getByName(db:&Pool,name:&str)->impl Future<Item=Vec<data::User>,Error=Error>{
	return db::getByName(db,name);
}

pub fn del(db:&Pool,userId:u64)->impl Future<Item=(),Error=Error>{
	return db::del(db,userId);
}

pub fn add(db:& Pool,userAdd:& data::UserAdd)->impl Future<Item=u64,Error=Error>{
	let newUserAdd = data::UserAdd{
		password:getPasswordHash(&userAdd.password),
		name:userAdd.name.clone(),
		r#type:userAdd.r#type.clone(),
	};
	let newDb = db.clone();
	return db::getByName(db,&userAdd.name)
		.and_then(|users|{
			if users.len() != 0{
				return err(Error::new(1,"存在重复的用户名"));
			}else{
				return ok(());
			}
		})
		.and_then(move|_|{
			return db::add(&newDb,&newUserAdd);
		});
}

pub fn modType(db:&Pool,userModType:& data::UserModType)->impl Future<Item=(),Error=Error>{
	let newDb = db.clone();
	let userType = userModType.r#type;
	return db::get(db,userModType.userId)
		.and_then(move|user|{
			return db::r#mod(&newDb,&data::UserMod{
				userId:user.userId,
				r#type:userType,
				password:user.password,
				name:user.name,
			});
		});
}

pub fn modPassword(db:&Pool,userModPassword:&data::UserModPassword)->impl Future<Item=(),Error=Error>{
	let newPassword = userModPassword.password.clone();
	let newDb = db.clone();
	return db::get(db,userModPassword.userId)
		.and_then(move|user|{
			return db::r#mod(&newDb,&data::UserMod{
				userId:user.userId,
				r#type:user.r#type,
				password:getPasswordHash(&newPassword),
				name:user.name,
			});
		});
}

pub fn modPasswordByOld(db:& Pool,userId:u64,userModMyPassword:&data::UserModMyPassword)->impl Future<Item=(),Error=Error>{
	let newDb = db.clone();
	let oldPassword = userModMyPassword.oldPassword.clone();
	let newPassword = userModMyPassword.newPassword.clone();
	return db::get(db,userId)
		.and_then(move|user|{
			let result = checkMustValidPassword(&oldPassword,&user.password);
			if let Err(e) = result{
				return err(e);
			}else{
				return ok(user);
			}
		})
		.and_then(move|user|{
			return db::r#mod(&newDb,&data::UserMod{
				userId:user.userId,
				r#type:user.r#type,
				password:getPasswordHash(&newPassword),
				name:user.name,
			});
		});
}