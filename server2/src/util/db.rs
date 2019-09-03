extern crate mysql_async as mysql;

pub type Pool = mysql::Pool;

pub fn get()->mysql::Pool{
	let database_url = "mysql://root:Yinghao23367847@localhost:3306/FishMoney";
	let pool = mysql::Pool::new(database_url);
	return pool;
}