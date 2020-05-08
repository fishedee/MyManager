mod ao;
mod data;
mod db;

pub use data::*;
pub use ao::*;
#[derive(Serialize)]
pub struct User{
	userId:u32,
	name:String,
	password:String,
	r#type:u32,
	createTime:String,
	modifyTime:String,
}

pub struct Users{
	Count:u32,
	Data:Vec<User>,
}

pub struct UserSearch{
	userId:Option<u32>,
	r#type:Option<u32>,
	pageIndex:u32,
	pageSize:u32,
}

fn getSearchWhere(where:&userSearch)->String{
	
}
pub fn search(_db:&Pool,_search:&UserSearch)->Result<Users,Error>{
	return Ok(Users{
		Count:0,
		Data:Vec::new()
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
