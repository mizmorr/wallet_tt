import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [

        { duration: '5s', target: 25 },
        { duration: '10s', target: 50 },
    { duration: '10s', target: 100 }, // Стабилизация на 30 виртуальных пользователей
        { duration: '30s', target: 200 }, // Стабилизация на 30 виртуальных пользователей
        { duration: '10s', target: 0 },  // Спад нагрузки до 0
    ],
};

export default function () {

    let res = http.get('http://localhost:8080/api/v1/wallets/839fd732-e489-4f30-8cb9-7f7cb03651e8');
    check(res, {
        'status was 200': (r) => r.status === 200,
    });
    sleep(0.001);
}
