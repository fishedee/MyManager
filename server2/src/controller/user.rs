use crate::util::data::WebData;
use crate::util::error::Error;
use crate::util::response::JsonResponse;
use crate::model::userAo;
use actix_web::web;
use futures::future::{ok,Future};

pub fn router(cfg:&mut web::ServiceConfig){
	cfg//.route("/search",web::get().to_async(search))
		//.route("/get",web::get().to_async(get))
		//.route("/add",web::post().to_async(add))
		//.route("/mod",web::post().to_async(r#mod))
		.route("/modType",web::post().to_async(modType));
}

/*
fn search(data:web::Data<WebData>,query:web::Query<userAo::UserSearch>)->impl Future<Item=JsonResponse<userAo::Users>,Error=Error>{
	return userAo::search(&data.pool,&query)
		.map(|data|{
			JsonResponse::new(data)
		});
}

fn get(data:web::Data<WebData>,query:web::Query<u64>)->impl Future<Item=JsonResponse<userAo::User>,Error=Error>{
	return userAo::get(&data.pool,*query)
		.map(|data|{
			JsonResponse::new(data)
		});
}

fn add(data:web::Data<WebData>,form:web::Form<userAo::UserAdd>)->impl Future<Item=JsonResponse<u64>,Error=Error>{
	return userAo::add(&data.pool,&form)
		.map(|data|{
			JsonResponse::new(data)
		});
}*/

fn r#modType(data:web::Data<WebData>,form:web::Form<userAo::UserModType>)->impl Future<Item=JsonResponse<()>,Error=Error>{
	return ok::<(),Error>(()).and_then(move|_|{
		let pool = &data.pool;
		let form:&userAo::UserModType = &form;
		return ok::<_,Error>((pool,form))
			.and_then(|(pool,form)|{
				return userAo::r#modType(ok::<_,Error>((pool,form)));
			}).map(|_|{
				return JsonResponse::new(());
			});
	});
}

/*
fn del(data:web::Data<WebData>,form:web::Form<u64>)->impl Future<Item=JsonResponse<()>,Error=Error>{
	return userAo::del(&data.pool,*form)
		.map(|data|{
			JsonResponse::new(data)
		});
}*/