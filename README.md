# mathmatic-test-server

Go 로 만든 서버 프로젝트입니다.

## 실행

```bash
make run
# 다른 터미널에서
curl localhost:8080/healthz
```

포트는 `ADDR` 환경 변수로 바꿀 수 있습니다. (예: `ADDR=:9000 make run`)

## 명령어

| 명령어 | 설명 |
| --- | --- |
| `make run` | 서버 실행 |
| `make build` | `bin/server` 빌드 |
| `make test` | 테스트 실행 |
| `make fmt` | 코드 포맷팅 |
| `make vet` | 정적 분석 |

## 구조 (헥사고날 아키텍처)

```
cmd/server/                 진입점
configs/
├── routes/                 (비어 있음)
└── app_config.{development,docker,production,stage,template}.json
internal/
├── adapters/
│   ├── driving/httpapi/    앱을 호출하는 쪽 (HTTP 핸들러)
│   └── driven/memory/      앱이 호출하는 쪽 (메모리 저장소)
├── core/
│   ├── domain/             엔티티, 도메인 규칙 (외부 의존 없음)
│   ├── application/        유스케이스 구현
│   └── routing/            (비어 있음)
├── di/                     의존성 조립
├── golden/                 (비어 있음)
└── ports/
    ├── inbound/            인바운드 포트 (유스케이스 인터페이스)
    └── outbound/           아웃바운드 포트 (저장소 등 인터페이스)
```

의존성 방향: `adapters -> ports -> core/domain`, 조립은 `di` 에서만 합니다.

## API

| 메서드 | 경로 | 설명 |
| --- | --- | --- |
| GET | `/healthz` | 상태 확인 |
| GET | `/problems/{id}` | 문제 조회 |
| POST | `/problems/{id}/answer` | 답 제출 (`{"answer": 2}`) |
