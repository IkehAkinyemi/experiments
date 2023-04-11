import { GoogleLogin } from 'react-google-login';

const clientId = 'YOUR_GOOGLE_CLIENT_ID';

const handleGoogleLoginSuccess = async (response) => {
  // Send the authentication token to the backend for verification
  const { data } = await axios.post('/api/login/google', {
    token: response.tokenId
  });

  // Store the JWT in local storage
  localStorage.setItem('token', data.token);

  // Redirect to the home page
  history.push('/');
};

const handleGoogleLoginFailure = (response) => {
  console.log(response);
};

const Login = () => {
  return (
    <GoogleLogin
      clientId={clientId}
      onSuccess={handleGoogleLoginSuccess}
      onFailure={handleGoogleLoginFailure}
      buttonText="Sign in with Google"
      cookiePolicy={'single_host_origin'}
    />
  );
};
