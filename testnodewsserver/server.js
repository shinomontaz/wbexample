// const WebSocket = require('ws');
//
// const wss = new WebSocket.Server({ port: 81 });
//
// wss.on("connect_error", (err) => {
//   console.log(`connect_error due to ${err.message}`);
// });
//
// wss.on('connection', function connection(ws) {
//   console.log('Соединение установлено');
//
//   ws.emit('hello', 'hello from the server');
//
//
//   ws.on('message', function incoming(message) {
//     console.log(`Получено сообщение: ${message}`);
//   });
//
//   ws.on('close', function close() {
//     console.log('Соединение закрыто');
//   });
// });

const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

io.on('connection', (socket) => {
  console.log('a user connected');
});

server.listen(3000, () => {
  console.log('listening on *:3000');
});

//
// const express = require('express')
// const app = express()
// const server = require('http').server(app)
// const io = require('socket.io')(server)
// io.on('connection', function (socket) {
//   // #2 - this will run for the new connection 'socket' and set up its callbacks
//   // #3 - send to the new client a 'hello' message
//   socket.emit('hello', 'hello from the server')
//   socket.on('clientdata', (data) => {
//     // #6 - handle this clients 'clientdata'
//     console.log(data)
//   })
// })
// server.listen(81, () => {
//   console.log("server started: http://localhost:${port}")
// })
