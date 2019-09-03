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
	cfg.route("/search",web::get().to_async(search))
		.route("/get",web::get().to_async(get))
		.route("/add",web::post().to_async(add))
		.route("/del",web::post().to_async(del))
		.route("/modType",web::post().to_async(modType))
		.route("/modPassword",web::post().to_async(modPassword))
		.route("/modMyPassword",web::post().to_async(modMyPassword));
}

fn search(data:web::Data<WebData>,query:web::Query<userAo::UserSearch>,session:Session)->impl Future<Item=JsonResponse<userAo::Users>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			return userAo::search(&data.pool,&query);
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn get(data:web::Data<WebData>,query:web::Query<u64>,session:Session)->impl Future<Item=JsonResponse<userAo::User>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			return userAo::get(&data.pool,*query);
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn add(data:web::Data<WebData>,form:web::Form<userAo::UserAdd>,session:Session)->impl Future<Item=JsonResponse<u64>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			return userAo::add(&data.pool,&form);
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn del(data:web::Data<WebData>,form:web::Form<u64>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);
	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			return userAo::del(&data.pool,*form)
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn r#modType(data: web::Data<WebData>,form:web::Form<userAo::UserModType>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);

	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			userAo::r#modType(&data.pool,&form)
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn modPassword(data: web::Data<WebData>,form:web::Form<userAo::UserModPassword>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);

	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|_|{
			userAo::modPassword(&data.pool,&form)
		}).map(|data|{
			JsonResponse::new(data)
		});
}

fn modMyPassword(data: web::Data<WebData>,form:web::Form<userAo::UserModMyPassword>,session:Session)->impl Future<Item=JsonResponse<()>,Error=Error>{
	let session = Arc::new(session);

	return loginAo::checkMustAdmin(&data.pool,&session)
		.and_then(move|user|{
			userAo::modPasswordByOld(&data.pool,user.userId,&form)
		}).map(|data|{
			JsonResponse::new(data)
		});
}
