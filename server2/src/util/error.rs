use actix_web::{error,HttpResponse,http};
use serde::Serialize;
use std::fmt;

#[derive(Serialize,Debug)]
pub struct Error{
	code:i32,
	msg:String
}

impl Error{
	pub fn new<T:fmt::Display>(code:i32,msg:T)->Error{
		return Error{
			code:code,
			msg:format!("{}",msg)
		}
	}
}

impl error::ResponseError for Error{
	fn error_response(&self) -> HttpResponse {
		if self.code == 500{
			return HttpResponse::new(http::StatusCode::INTERNAL_SERVER_ERROR);
		}else{
			return HttpResponse::new(http::StatusCode::OK);
		}
	}
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    	if self.code == 500{
    		return write!(f,"Internal Server Error");
    	}else{
    		let body = serde_json::to_string(&self).unwrap();
    		return write!(f, "{}", body);
    	}
    }
}
