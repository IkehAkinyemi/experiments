const crypto = require('crypto');
const axios = require('axios');

// Set the API endpoint and headers
const apiUrl = 'http://localhost:8080';
const apiKey = 'MY_API_KEY';
const apiSecret = 'MY_API_SECRET';
const expiry = new Date(Date.now() + 300000).toISOString(); // 5 minute expiry

// Define the data to be sent with the request
const requestData = {param1: 'value1', param2: 'value2'};

// Generate the signature using the HMAC-SHA256 algorithm
const requestHash = crypto.createHmac('sha256', apiSecret)
  .update(`${apiKey}${expiry}${JSON.stringify(requestData)}`)
  .digest('hex');

console.log(requestHash)

// Set the request headers, including the signature and expiry
const headers = {
  'X-API-KEY': apiKey,
  'X-EXPIRY': expiry,
  'X-REQUEST-SIGNATURE': requestHash
};

// Send the request using Axios
axios.post(apiUrl, requestData, { headers })
  .then(response => console.log(response.data))
  .catch(error => console.log(error.response.data));
