import http from 'k6/http';

export const options = {
  stages: [
    { duration: '15s', target: 50 },
    { duration: '30s', target: 100 },
    { duration: '1m30s', target: 200 },
    { duration: '20s', target: 100 },
    { duration: '10s', target: 50 },
  ],
};

function makeid(length) {
    var result           = '';
    var characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < length; i++ ) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
   }
   return result;
}

const baseUrl = 'http://localhost:53441';
const email = makeid(25) + '@email.com';
const password = 'password';

const register = () => {
  const url = baseUrl + '/api/user';
  const payload = JSON.stringify({
    email,
    password,
    username: 'john',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
};

const getToken = () => {
  const url = baseUrl + '/api/auth/login';
  const payload = JSON.stringify({
    email,
    password,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);
  return res.json().token;
};

const getUser = (token) => {
  const url = baseUrl + '/api/user';
  const params = {
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
  };

  http.get(url, params);
};

export default function () {
  register();
  const token = getToken();
  getUser(token);
}
