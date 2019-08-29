use crate::util::data::WebData;
use actix_web::web;
use actix_web::get;

pub fn router(cfg:&mut web::ServiceConfig){
	cfg.service(search)
		.service(get);
}

#[get("/search")]
fn search(data:web::Data<WebData>)->String{
	"123".to_string()
}

#[get("/get")]
fn get(data:web::Data<WebData>)->String{
	"456".to_string()
}