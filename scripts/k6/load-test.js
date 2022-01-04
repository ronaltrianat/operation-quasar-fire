import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 100 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    { duration: '2m', target: 100 }, // stay at 100 users for 10 minutes
    { duration: '1m', target: 0 },
  ],
  thresholds: {
    'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
};

const BASE_URL = 'http://operationquasarfire-env.eba-umarwuy2.us-east-1.elasticbeanstalk.com';


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
            "message": [
              "este",
              "",
              "",
              "mensaje",
              ""
            ],
            "position": {
              "x": 100,
              "y": -100
            }
          },
          {
            "name": "sato",
            "distance": 728.0110,
            "message": [
              "",
              "es",
              "",
              "",
                      "secreto"
            ],
            "position": {
              "x": 500,
              "y": 100
            }
          },
          {
            "name": "kenobi",
            "distance": 583.0952,
            "message": [
              "este",
              "",
              "un",
              "",
              ""
            ],
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


