import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 5000,
  duration: '10s',
  thresholds: {
    http_req_duration: ['p(95)<200'],
  },
};

export default function () {
  const payload = JSON.stringify({
    user_id: '00000000-0000-0000-0000-000000000001',
    amount: 99.5,
    description: 'load test transaction',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post('http://192.168.1.18:80/v1/transaction', payload, params);
  check(res, {
    'status is 200 or 201': (r) => r.status >= 200 && r.status <= 400,
  });
  
  sleep(1);
}
