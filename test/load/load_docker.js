import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [

        { duration: '10s', target: 5 },
        { duration: '10s', target: 25 },
    { duration: '20s', target: 100 },
        { duration: '30s', target: 200 },
        { duration: '10s', target: 25 },
    { duration: '10s', target: 0 },
    ],
};

export default function () {

    let getUrls = [
      'http://172.18.0.4:8080/api/v1/wallets/103960a0-9a79-43ed-bff0-052a19eaa98e',
      'http://172.18.0.4:8080/api/v1/wallets/f2bfc720-c67a-4f79-9bcb-47c146deb8e3',
      'http://172.18.0.4:8080/api/v1/wallets/c6a51a4e-7c5e-4354-8576-d22c7649ad65',
    ];

    let getResponses = getUrls.map(url => http.get(url));

    check(getResponses, {
    'all requests succeeded': (rs) => rs.every((r) => r.status === 200),
    });

    let postUrl = 'http://172.18.0.4:8080/api/v1/wallet';

    let postWithdrawPayload = {
    id:'8e4c1a17-4dbd-4910-a0d1-7e6eced20a5e',
    amount: 1,
    operation: 'withdraw',
  };

  let postDepositPayload = {
    id:'8e4c1a17-4dbd-4910-a0d1-7e6eced20a5e',
    amount: 1,
    operation: 'deposit',
  };

  let postHeaders = { 'Content-Type': 'application/json' };
  let withdrawRes = http.post(postUrl, JSON.stringify(postWithdrawPayload),{headers: postHeaders});
  let depositRes = http.post(postUrl, JSON.stringify(postDepositPayload), {headers: postHeaders});

  check(withdrawRes, {
    'status was 200': (r) => r.status === 200,
  });

  check(depositRes, {
   'status was 200': (r) => r.status === 200,
  });


    sleep(0.001);
}
