use env_logger;
use actix_web::middleware::Logger;

pub fn init(){
	std::env::set_var("RUST_LOG", "actix_web=info");
	env_logger::init();
}

pub fn get()->actix_web::middleware::Logger{
	return Logger::default();
}