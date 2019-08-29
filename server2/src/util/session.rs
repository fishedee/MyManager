use actix_session::{CookieSession};

pub fn get()->CookieSession{
	return CookieSession::signed(&[0; 32]).secure(false);
}