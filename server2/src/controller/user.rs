use crate::util::data::WebData;
use crate::util::error::Error;
use crate::util::response::JsonResponse;
use crate::model::userAo;
use actix_web::web;
use futures::future::Future;

pub fn router(cfg:&mut web::ServiceConfig){
	cfg.route("/search",web::get().to(search))
		.route("/get",web::get().to_async(get));
}

fn search(data:web::Data<WebData>)->Result<String,Error>{
	return Err(Error::new(1,"sadf"));
}

fn get(data:web::Data<WebData>)->impl Future<Item=JsonResponse<userAo::User>,Error=Error>{
	return userAo::get(&data.pool,10001)
		.map(|data|{
			JsonResponse::new(data)
		});
}