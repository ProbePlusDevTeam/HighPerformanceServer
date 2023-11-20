var http = require('http');
var createStatsd = require('uber-statsd-client');
var sdc = createStatsd({host: '192.168.2.1',port: 8126});
qs = require('querystring');
var argv = process.argv;var port = argv[2];
function randomIntInc (low, high){    return Math.floor(Math.random() * (high - low + 1) + low);}
function sendResponse(res,times, old_sleep){    res.write('pong');    if(times==0)    {        res.end();    }    else    {         sleep = randomIntInc(0, old_sleep+1);        setTimeout(sendResponse, sleep, res,times-1, old_sleep);    }}
var server = http.createServer(function(req, res){   headers = req.headers;   old_sleep = parseInt(headers["sleep"]);   times = headers["times"] || 0;   sleep = randomIntInc(0, old_sleep+1);   sdc.increment("ssl.server.http"); res.writeHead(200);   setTimeout(sendResponse, sleep, res, times, old_sleep)
});
server.timeout = 3600000;server.listen(port);