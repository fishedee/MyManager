use actix_web::{Error, HttpRequest, HttpResponse, Responder};
use serde::Serialize;

#[derive(Serialize)]
pub struct JsonResponse<T> {
    code:i32,
    msg:String,
    data:T,
}

impl<T> JsonResponse<T>{
    pub fn new(data:T)->JsonResponse<T>{
        return JsonResponse{
            code:0,
            msg:"".to_string(),
            data:data,
        }
    }
}

impl<T:Serialize> Responder for JsonResponse<T> {
    type Error = Error;
    type Future = Result<HttpResponse, Error>;

    fn respond_to(self, _req: &HttpRequest) -> Self::Future {
        let body = serde_json::to_string(&self)?;

        Ok(HttpResponse::Ok()
            .content_type("application/json")
            .body(body))
    }
}