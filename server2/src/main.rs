use actix_web::{web, App, HttpRequest, HttpServer, Responder};
use actix_web::middleware::Logger;
use listenfd::ListenFd;
use env_logger;

fn main() {
	std::env::set_var("RUST_LOG", "actix_web=info");
	env_logger::init();

	let mut listenfd = ListenFd::from_env();
    let mut server = HttpServer::new(|| {
        App::new()
        	.wrap(Logger::default())
            .route("/",web::get().to(||{
            	"Hello World"
            }))
    });

    server = if let Some(l) = listenfd.take_tcp_listener(0).unwrap() {
        server.listen(l).unwrap()
    } else {
        server.bind("127.0.0.1:8088").unwrap()
    };

    server.run().unwrap();
}