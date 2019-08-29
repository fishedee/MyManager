use crate::util::data::WebData;
use crate::util::error::Error;
use crate::model::userAo;
use actix_web::web;
use futures::future::{ok,err,Future};

pub fn router(cfg:&mut web::ServiceConfig){
	cfg.route("/search",web::get().to(search))
		.route("/get",web::get().to_async(get));
}

fn search(data:web::Data<WebData>)->Result<String,Error>{
	return Err(Error::new(1,"sadf"));
}

fn get(data:web::Data<WebData>)->impl Future<Item=userAo::User,Error=Error>{
	return userAo::get(&data.pool,10001);
}