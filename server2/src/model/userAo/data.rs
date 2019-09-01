use serde::{Serialize,Deserialize};

#[derive(Serialize)]
pub struct User{
	pub userId:u64,
	pub name:String,
	pub password:String,
	pub r#type:u64,
	pub createTime:String,
	pub modifyTime:String,
}

#[derive(Serialize)]
pub struct Users{
	pub count:u64,
	pub data:Vec<User>,
}

#[derive(Deserialize)]
pub struct UserSearch{
	pub userId:Option<u64>,
	pub r#type:Option<u64>,
	pub pageIndex:u64,
	pub pageSize:u64,
}

#[derive(Deserialize)]
pub struct UserAdd{
	pub name:String,
	pub password:String,
	pub r#type:u64,
}

#[derive(Deserialize)]
pub struct UserMod{
	pub userId:u64,
	pub name:String,
	pub password:String,
	pub r#type:u64,
}

#[derive(Deserialize)]
pub struct UserModType{
	pub userId:u64,
	pub r#type:u64,
}