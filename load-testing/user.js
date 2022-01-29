import http from 'k6/http';
import { sleep } from 'k6';

// Smoke test
// export const options = {
//   stages: [
//     { duration: '1m', target: 1 }, // ramp-down to 0 users
//   ]
// };

// // Load test
// export const options = {
//   stages: [
//     { duration: '5m', target: 100 }, // ramp-down to 0 users
//     { duration: '10m', target: 100 }, // ramp-down to 0 users
//     { duration: '5m', target: 0 }, // ramp-down to 0 users
//   ]
// };

// // Stress test
// export const options = {
//   noConnectionReuse: true,
//   noVuConnectionReuse: true,
//   stages: [
//     { duration: '2m', target: 100 }, // below normal load
//     { duration: '2m', target: 100 },
//     { duration: '2m', target: 200 }, // normal load
//     { duration: '2m', target: 200 },
//     { duration: '2m', target: 300 }, // around the breaking point
//     { duration: '2m', target: 300 },
//     { duration: '2m', target: 400 }, // beyond the breaking point
//     { duration: '2m', target: 400 },
//     { duration: '10m', target: 0 }, // scale down. Recovery stage.
//   ],
// };

// // Spike test
// export const options = {
//   stages: [
//     { duration: '10s', target: 100 }, // below normal load
//     { duration: '1m', target: 100 },
//     { duration: '10s', target: 1400 }, // spike to 1400 users
//     { duration: '3m', target: 1400 }, // stay at 1400 for 3 minutes
//     { duration: '10s', target: 100 }, // scale down. Recovery stage.
//     { duration: '3m', target: 100 },
//     { duration: '10s', target: 0 },
//   ],
// };

// // Soak testing
// export const options = {
//   stages: [
//     { duration: '2m', target: 400 },
//     { duration: '3h56m', target: 400 },
//     { duration: '2m', target: 0 },
//   ],
// };

const makeid = (length) => {
    var result           = '';
    var characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < length; i++ ) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
   }
   return result;
}

const register = (baseUrl, email, password) => {
  const url = baseUrl + '/user';
  const payload = {
    email,
    password
  };

  return http.post(url, payload);
};

const getToken = (baseUrl, email, password) => {
  const url = baseUrl + '/auth/login';
  const payload = {
    email,
    password,
  };

  const res = http.post(url, payload);
  return res.json().token;
};

const getUser = (baseUrl, token) => {
  const url = baseUrl + '/user';
  const params = {
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
  };

  return http.get(url, params);
};

export default function() {
  const port = 8000 
  const baseUrl = `http://localhost:${port}/api`;
  const email = makeid(25) + '@email.com';
  const password = 'password';

  register(baseUrl, email, password);
  sleep(1);
  const token = getToken(baseUrl, email, password);
  sleep(1);
  getUser(baseUrl, token);
  sleep(1);
}
