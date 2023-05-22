// Import required modules
const express = require('express');
const app = express();
const http = require('http').createServer(app);
const io = require('socket.io')(http);

// Serve static files from the 'public' directory
app.use(express.static('public'));

// Define a route
app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

// Socket.IO event handling
io.on('connection', (socket) => {
  console.log('A user connected');

  // Handle 'chat message' event
  socket.on('chat message', (msg) => {
    console.log('Message: ' + msg);
    io.emit('chat message', msg); // Broadcast the message to all connected clients
  });

  // Handle 'disconnect' event
  socket.on('disconnect', () => {
    console.log('A user disconnected');
  });
});

// Start the server
const port = 3000;
http.listen(port, () => {
  console.log('Server listening on port ' + port);
});
