var session = require('express-session');
module.exports = {
	getMiddleware(){
		return session({
			resave: false, // don't save session if unmodified
	  		saveUninitialized: false, // don't create session until something stored
	  		secret: 'shhhh, very secret'
		});
	},
	getSessionData(req){
		return req.session; 
	},
	setSessionData(req,data){
		return new Promise(function(resolve,reject){
			req.session.regenerate(function(){
				for( var i in data ){
					req.session[i] = data[i];
				}
				resolve(null);
			});
		});
	},
	destorySession(req){
		return new Promise(function(resolve,reject){
			req.session.destroy(resolve);
		});
	}
};