import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10m', target: 100 },
    { duration: '10m', target: 200 },
    { duration: '10m', target: 0 },
  ],
  thresholds: {
    'http_req_duration': ['p(99)<1500'],
    'logged in successfully': ['p(99)<1500'],
  },
};

const BASE_URL = 'https://f0g0tobrb6.execute-api.us-east-1.amazonaws.com/operation-quasar-fire/api';


export default () => {

    const params = {
        headers: {
          'Content-Type': 'application/json',
        },
      };
    
  
    const payload = JSON.stringify({
        "satellites": [
          {
            "name": "skywalker",
            "distance": 500.0,
            "message": ["este", "", "", "mensaje", ""],
            "position": {
              "x": 100,
              "y": -100
            }
          },
          {
            "name": "sato",
            "distance": 728.0110,
            "message": ["", "es", "", "", "secreto"],
            "position": {
              "x": 500,
              "y": 100
            }
          },
          {
            "name": "kenobi",
            "distance": 583.0952,
            "message": ["este", "", "un", "", ""],
            "position": {
              "x": -500,
              "y": -200
            }
          }
        ]
      });

      const response = http.post(`${BASE_URL}/topsecret`, payload, params);
      
      const checkRes = check(response, {
        'status is 200': (r) => r.status === 200,
      });

  sleep(1);
};


