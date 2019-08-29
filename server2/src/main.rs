use listenfd::ListenFd;
use actix_web::{web, App, Error,HttpRequest, HttpServer, Responder,HttpResponse,error};

mod util;
mod controller;

use crate::util::db;
use crate::util::logger;
use crate::util::session;
use crate::util::data;
use crate::controller::user;

fn main() {
    logger::init();
	let db = db::get();
	let data = data::get(db);

	let mut listenfd = ListenFd::from_env();
    let mut server = HttpServer::new(move|| {
        return App::new()
        	.wrap(logger::get())
        	.wrap(session::get())
        	.register_data(data.clone())
            .service(web::scope("/user").configure(user::router));
    });

    server = if let Some(l) = listenfd.take_tcp_listener(0).unwrap() {
        server.listen(l).unwrap()
    } else {
        server.bind("127.0.0.1:8088").unwrap()
    };

    server.run().unwrap();
}