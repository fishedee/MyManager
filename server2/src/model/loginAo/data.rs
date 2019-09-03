use serde::{Serialize,Deserialize};

#[derive(Deserialize)]
pub struct LoginCheckIn{
	pub name:String,
	pub password:String,
}