use actix_session::{CookieSession};
use std::sync::{Mutex, Arc,RwLock};

pub type Session = Arc<actix_session::Session>;

pub fn get()->CookieSession{
	return CookieSession::signed(&[0; 32]).secure(false);
}