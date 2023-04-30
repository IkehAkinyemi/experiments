const ws = new WebSocket('ws://localhost:8080/ws');

ws.addEventListener('open', () => {
  console.log('Connected to WebSocket server!');
  ws.send('Hello, server!');
});

ws.addEventListener('message', (event) => {
  console.log('Received message from server:', event.data);
});

ws.addEventListener('close', () => {
  console.log('Disconnected from WebSocket server!');
});
