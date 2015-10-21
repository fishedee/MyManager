var express = require('express');
var util = require('./config/util');
var session = require('./config/session');
var bodyParser = require('body-parser');
var loginRouter = require('./router/login');
var userRouter = require('./router/user');

var app = express();
app.use(session.getMiddleware());
app.use(bodyParser.urlencoded({ extended: false }));
app.use('/login',loginRouter);
app.use('/user',userRouter);
app.use(function(err, req, res, next) {
  console.error(err.stack);
  res.status(500).send('Something broke!');
});

var server = app.listen(3000, function () {
	 var host = server.address().address;
	 var port = server.address().port;
	 console.log('Example app listening at http://%s:%s', host, port);
});