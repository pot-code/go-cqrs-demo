import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 100 }, // below normal load
    { duration: '20s', target: 100 },
    { duration: '1m', target: 200 }, // normal load
    // { duration: '5m', target: 200 },
    // { duration: '2m', target: 300 }, // around the breaking point
    // { duration: '5m', target: 300 },
    // { duration: '2m', target: 400 }, // beyond the breaking point
    // { duration: '5m', target: 400 },
    // { duration: '10m', target: 0 }, // scale down. Recovery stage.
  ],
};

export default function () {
  const BASE_URL = 'http://localhost:8081/api/order'; // make sure this is not production

  let responses = http.batch([
    [
      'POST',
      `${BASE_URL}`,
      JSON.stringify({
        note: 'balaasdsa',
        customer_id: 'CUSTOMER_123456712318',
        seller_id: 'SELLER_12345678',
      }),
      { tags: { name: 'CreateOrder' }, headers: { 'Content-Type': 'application/json' } },
    ], //
  ]);

  responses.forEach(function (response) {
    check(response, {
      200: (r) => r.status === 200,
    });
  });
  sleep(1);
}
