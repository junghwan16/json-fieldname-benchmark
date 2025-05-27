# JSON Fieldname Benchmark

## Background & Motivation

이 프로젝트는 우연히 “API 성능을 위해 JSON 컬럼명(필드명)을 줄이면 체감할 만큼 속도가 빨라질까?”라는 얘기를 듣고  
“실제로 어떤 상황에서, 어느 정도 규모에서 그 차이가 눈에 띌까?”를 검증해 보기 위해 시작했습니다.

또한, gRPC(Protocol Buffers)의 바이너리 직렬화＋압축 방식이 JSON 대비 얼마나 성능을 개선하는지도 함께 살펴봅니다.

## What We’re Testing

1. **JSON – Short Keys**

   - 필드명을 `a`, `b`, `c` 같은 짧은 이름으로 사용
   - 경량화된 텍스트 포맷

2. **JSON – Long Keys**

   - `customer_full_name`, `total_transaction_amount_usd` 등 긴 이름 사용
   - 가독성 vs. 전송량 비교

3. **gRPC (Protobuf + Compression)**
   - 동일한 데이터 구조를 Protobuf 스키마로 정의
   - 내장 압축 옵션(gzip) on/off

## When It Becomes Noticeable

- **초당 수백~수천 건** 수준의 고빈도 호출
- **Payload 크기가 작고**, 오버헤드(헤더 + 필드명)가 전체의 큰 비중을 차지할 때
- **모바일/원격지 네트워크(높은 레이턴시)** 환경에서
- **대량 로그 전송**이나 **IoT 디바이스 통신** 처럼 작은 메시지를 자주 주고받을 때

이런 시나리오에서 JSON 필드명 절감과 gRPC 압축의 효과가 실질적으로 체감될 것으로 예상됩니다.

## Project Structure

```

json-fieldname-benchmark/
├── go.mod
├── main.go # Echo 서버 진입점
├── handler.go # 짧은/긴 JSON 핸들러
├── proto/ # Protobuf 정의 & gRPC 서버 코드
├── benchmark/ # 벤치마크 스크립트 (wrk, ghz 등)
└── README.md

```

## Usage

1. **Go + Echo 서버 실행**

   ```bash
   go run main.go handler.go
   ```

2. **gRPC 서버 실행**

   ```bash
   # proto 폴더에서
   protoc --go_out=. --go-grpc_out=. example.proto
   go run grpc_server.go
   ```

3. **벤치마크 스크립트 실행**

   ```bash
   # wrk를 이용한 JSON 테스트
   wrk -t4 -c100 -d30s http://localhost:8080/short
   wrk -t4 -c100 -d30s http://localhost:8080/long

   # ghz를 이용한 gRPC 테스트
   ghz --insecure --proto=proto/example.proto \
       --call=Example.Service/Method -n 100000 \
       -d '{"customerFullName":"Alice",...}' localhost:50051
   ```

## Results

- JSON Short vs Long: Payload 크기, 평균 레이턴시 비교

  - wrk 잘만들었네..

  ```
  === HTTP JSON short keys ===
  Running 30s test @ http://localhost:8080/short
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
     Latency     2.02ms    2.07ms  31.49ms   87.01%
     Req/Sec    14.68k     1.92k   22.80k    70.42%
  1756774 requests in 30.09s, 236.23MB read
  Requests/sec:  58385.60
  Transfer/sec:      7.85MB

  === HTTP JSON long keys ===
  Running 30s test @ http://localhost:8080/long
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
     Latency     1.92ms    2.00ms  26.82ms   87.28%
     Req/Sec    15.44k     2.01k   22.37k    68.25%
  1848642 requests in 30.09s, 393.15MB read
  Requests/sec:  61433.06
  Transfer/sec:     13.06MB
  ```

- JSON vs gRPC: 바이너리 직렬화＋압축 효과
