use crate::util::data::WebData;
use crate::util::db::Pool;
use crate::util::error::Error;
use crate::util::response::JsonResponse;
use crate::model::userAo;
use crate::model::loginAo;
use actix_session::{Session};
use actix_web::web;
use futures::future::{ok,Future};
use std::sync::{Mutex, Arc,RwLock};

pub fn router(cfg:&mut web::ServiceConfig){
	cfg.route("/islogin",web::get().to_async(isLogin))
		.route("/checkin",web::post().to_async(checkIn))
		.route("/checkout",web::get().to_async(checkOut));
}

fn isLogin(data:web::Data<WebData>,session:Session)->impl Future<Item=JsonResponse<userAo::User>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::isLogin(&data.pool,&session)
		.map(|data|{
			JsonResponse::new(data)
		});
}

fn checkIn(data:web::Data<WebData>,form:web::Form<loginAo::LoginCheckIn>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::login(&data.pool,&session,&form)
		.map(|data|{
			JsonResponse::new(data)
		});
}

fn checkOut(data:web::Data<WebData>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::logout(&data.pool,&session)
		.map(|data|{
			JsonResponse::new(data)
		});
}