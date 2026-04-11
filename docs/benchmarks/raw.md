jules ~/Dev/master/benchlab main ≡*?1 ~3  1.134s 
❯ make k6-test s=unit-read-k6-test.js
k6 run benchmarks/scripts/unit-read-k6-test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 


     execution: local
        script: benchmarks/scripts/unit-read-k6-test.js
        output: -

     scenarios: (100.00%) 2 scenarios, 20 max VUs, 10m30s max duration (incl. graceful stop):
              * grpc_scenario: 1000 iterations shared among 10 VUs (maxDuration: 10m0s, exec: test_grpc, gracefulStop: 30s)
              * rest_scenario: 1000 iterations shared among 10 VUs (maxDuration: 10m0s, exec: test_rest, gracefulStop: 30s)



  █ TOTAL RESULTS 

    checks_total.......: 2000    1260.420396/s
    checks_succeeded...: 100.00% 2000 out of 2000
    checks_failed......: 0.00%   0 out of 2000

    ✓ gRPC status is OK
    ✓ REST status is 200

    HTTP
    http_req_duration..............: avg=10.11ms min=534.41µs med=3.52ms max=64.88ms p(90)=35.29ms p(95)=42.62ms
      { expected_response:true }...: avg=10.11ms min=534.41µs med=3.52ms max=64.88ms p(90)=35.29ms p(95)=42.62ms
    http_req_failed................: 0.00%  0 out of 1000
    http_reqs......................: 1000   630.210198/s

    EXECUTION
    iteration_duration.............: avg=12.93ms min=607.97µs med=5.81ms max=73.57ms p(90)=39.52ms p(95)=46.38ms
    iterations.....................: 2000   1260.420396/s
    vus............................: 20     min=20        max=20
    vus_max........................: 20     min=20        max=20

    NETWORK
    data_received..................: 570 kB 359 kB/s
    data_sent......................: 347 kB 219 kB/s

    GRPC
    grpc_req_duration..............: avg=12.35ms min=542.94µs med=4.86ms max=71.85ms p(90)=39.42ms p(95)=46.59ms




running (00m01.6s), 00/20 VUs, 2000 complete and 0 interrupted iterations
grpc_scenario ✓ [======================================] 10 VUs  00m01.6s/10m0s  1000/1000 shared iters
rest_scenario ✓ [======================================] 10 VUs  00m01.1s/10m0s  1000/1000 shared iters

jules ~/Dev/master/benchlab main ≡*?1 ~3 
❯ make k6-test s=write-k6-test.js
k6 run benchmarks/scripts/write-k6-test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 


     execution: local
        script: benchmarks/scripts/write-k6-test.js
        output: -

     scenarios: (100.00%) 2 scenarios, 10 max VUs, 10m30s max duration (incl. graceful stop):
              * grpc_scenario: 500 iterations shared among 5 VUs (maxDuration: 10m0s, exec: test_grpc, gracefulStop: 30s)
              * rest_scenario: 500 iterations shared among 5 VUs (maxDuration: 10m0s, exec: test_rest, gracefulStop: 30s)



  █ TOTAL RESULTS 

    checks_total.......: 1000    1352.297344/s
    checks_succeeded...: 100.00% 1000 out of 1000
    checks_failed......: 0.00%   0 out of 1000

    ✓ REST status is 200
    ✓ gRPC status is OK

    HTTP
    http_req_duration..............: avg=5.34ms min=920.37µs med=2.92ms max=39.36ms p(90)=15.03ms p(95)=18.77ms
      { expected_response:true }...: avg=5.34ms min=920.37µs med=2.92ms max=39.36ms p(90)=15.03ms p(95)=18.77ms
    http_req_failed................: 0.00%  0 out of 500
    http_reqs......................: 500    676.148672/s

    EXECUTION
    iteration_duration.............: avg=6.44ms min=1.01ms   med=3.76ms max=46.95ms p(90)=16.29ms p(95)=20.23ms
    iterations.....................: 1000   1352.297344/s

    NETWORK
    data_received..................: 179 kB 242 kB/s
    data_sent......................: 317 kB 429 kB/s

    GRPC
    grpc_req_duration..............: avg=5.49ms min=791.35µs med=2.58ms max=44.48ms p(90)=14.74ms p(95)=19.24ms




running (00m00.7s), 00/10 VUs, 1000 complete and 0 interrupted iterations
grpc_scenario ✓ [======================================] 5 VUs  00m00.7s/10m0s  500/500 shared iters
rest_scenario ✓ [======================================] 5 VUs  00m00.6s/10m0s  500/500 shared iters

jules ~/Dev/master/benchlab main ≡*?1 ~3 
❯ make k6-test s=ramp-up-K6-test.js
k6 run benchmarks/scripts/ramp-up-K6-test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 


     execution: local
        script: benchmarks/scripts/ramp-up-K6-test.js
        output: -

     scenarios: (100.00%) 2 scenarios, 200 max VUs, 2m30s max duration (incl. graceful stop):
              * grpc_scenario: Up to 100 looping VUs for 2m0s over 3 stages (gracefulRampDown: 30s, exec: test_grpc, gracefulStop: 30s)
              * rest_scenario: Up to 100 looping VUs for 2m0s over 3 stages (gracefulRampDown: 30s, exec: test_rest, gracefulStop: 30s)



  █ TOTAL RESULTS 

    checks_total.......: 132843 1106.973795/s
    checks_succeeded...: 96.15% 127733 out of 132843
    checks_failed......: 3.84%  5110 out of 132843

    ✗ gRPC status is OK
      ↳  95% — ✓ 49401 / ✗ 2215
    ✗ REST status is 200
      ↳  96% — ✓ 78332 / ✗ 2895

    HTTP
    http_req_duration..............: avg=112.57ms min=313µs    med=10.58ms max=1.82s p(90)=445.31ms p(95)=536.42ms
      { expected_response:true }...: avg=110.45ms min=313µs    med=9.73ms  max=1.82s p(90)=451.89ms p(95)=539.99ms
    http_req_failed................: 3.56%  2895 out of 81227
    http_reqs......................: 81227  676.860358/s

    EXECUTION
    iteration_duration.............: avg=137.9ms  min=348.21µs med=20.36ms max=2.02s p(90)=483.41ms p(95)=571.76ms
    iterations.....................: 132843 1106.973795/s
    vus............................: 2      min=2             max=200
    vus_max........................: 200    min=200           max=200

    NETWORK
    data_received..................: 40 MB  330 kB/s
    data_sent......................: 20 MB  168 kB/s

    GRPC
    grpc_req_duration..............: avg=148.27ms min=570.14µs med=18.47ms max=1.92s p(90)=496.14ms p(95)=570.75ms




running (2m00.0s), 000/200 VUs, 132843 complete and 0 interrupted iterations
grpc_scenario ✓ [======================================] 000/100 VUs  2m0s
rest_scenario ✓ [======================================] 000/100 VUs  2m0s