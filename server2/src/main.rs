use listenfd::ListenFd;
use actix_web::{web, App, HttpServer};

mod util;
mod controller;
mod model;

use crate::util::db;
use crate::util::logger;
use crate::util::session;
use crate::util::data;
use crate::controller::user;
use crate::controller::login;

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
            .service(web::scope("/user").configure(user::router))
            .service(web::scope("/login").configure(login::router));
    })
    .keep_alive(75)
    .workers(8);

    server = if let Some(l) = listenfd.take_tcp_listener(0).unwrap() {
        server.listen(l).unwrap()
    } else {
        server.bind("0.0.0.0:8180").unwrap()
    };

    server.run().unwrap();
}