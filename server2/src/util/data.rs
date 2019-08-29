use actix_web::web;
use crate::util::db::Pool;

pub struct WebData{
	pool:Pool
}

pub fn get(pool:Pool)->web::Data<WebData>{
	return web::Data::new(WebData{
		pool:pool
	});
}