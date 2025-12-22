# monitorTheProcess

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

실시간으로 특정 프로세스의 CPU와 메모리 사용량을 모니터링하는 경량 Go 애플리케이션입니다.

## 특징

- 📊 **실시간 모니터링**: 지정된 프로세스의 CPU 및 메모리 사용량을 초 단위로 추적
- ⏱️ **사용자 정의 가능한 간격**: 모니터링 기간을 유연하게 설정
- 📈 **평균 계산**: 설정된 기간 동안의 평균 리소스 사용량 제공
- 🔍 **경량**: 최소한의 시스템 리소스로 효율적인 모니터링
- 🖥️ **크로스 플랫폼**: Linux, macOS, Windows 지원

## 필수 요구사항

- Go 1.16 이상
- 모니터링 대상 프로세스의 PID(Process ID)

## 설치

### 소스에서 빌드

```bash
# 리포지토리 클론
git clone https://github.com/lmk/monitorTheProcess.git

# 프로젝트 디렉토리로 이동
cd monitorTheProcess

# 의존성 다운로드
go mod download

# 빌드
go build -o monitorTheProcess
```

### Go Install 사용

```bash
go install github.com/lmk/monitorTheProcess@latest
```

## 사용 방법

### 기본 사용법

```bash
./monitorTheProcess -pid <프로세스_ID> -duration <기간_초>
```

### 매개변수

| 매개변수 | 필수 | 기본값 | 설명 |
|---------|------|--------|------|
| `-pid` | ✓ | - | 모니터링할 프로세스의 PID |
| `-duration` | | 60 | 평균 계산 기간(초) |

### 예제

**1. 기본 모니터링 (60초 간격)**
```bash
./monitorTheProcess -pid 1234
```

**2. 사용자 정의 간격 (30초)**
```bash
./monitorTheProcess -pid 1234 -duration 30
```

**3. 장기 모니터링 (5분 간격)**
```bash
./monitorTheProcess -pid 5678 -duration 300
```

## 출력 형식

프로그램은 다음과 같은 형식으로 출력합니다:

```
INFO: 2025/12/22 10:30:45 main.go:23: Start App
PID: 1234
Duration Second: 60
INFO: 2025/12/22 10:30:45 main.go:29: ps -o pcpu,rss -p 1234
INFO: 2025/12/22 10:31:45 main.go:47: pcpu,rss: 15.32 524288
INFO: 2025/12/22 10:32:45 main.go:47: pcpu,rss: 12.45 516096
```

- **pcpu**: 기간 동안의 평균 CPU 사용률 (%)
- **rss**: 기간 동안의 평균 메모리 사용량 (바이트)

## 작동 원리

1. 지정된 PID의 프로세스를 1초마다 샘플링
2. 설정된 기간 동안 CPU 및 메모리 사용량 누적
3. 기간 종료 시 평균 값 계산 및 출력
4. 다음 기간을 위해 카운터 초기화 및 반복

이 방식은 `ps -o pcpu,rss -p <PID>` 명령과 유사한 정보를 제공하지만 시간 경과에 따른 평균 값을 계산합니다.

## 의존성

- [struCoder/pidusage](https://github.com/struCoder/pidusage) v0.2.1 - 프로세스 사용량 통계 수집

## 프로젝트 구조

```
monitorTheProcess/
├── main.go        # 메인 애플리케이션 로직 및 모니터링 루프
├── Config.go      # 설정 관리 및 입력 검증
├── logger.go      # 로깅 유틸리티
├── go.mod         # Go 모듈 정의
├── go.sum         # 의존성 체크섬
└── README.md      # 프로젝트 문서
```

## 문제 해결

### "Invalid PID!" 오류
- PID가 0이 아닌 유효한 값인지 확인
- `-pid` 플래그가 올바르게 지정되었는지 확인

### 프로세스를 찾을 수 없음
- PID가 실행 중인 프로세스인지 확인
- 프로세스에 접근할 수 있는 권한이 있는지 확인 (필요시 sudo 사용)

### 높은 CPU 사용량
- `-duration` 값을 늘려 샘플링 빈도 감소
- 모니터링이 대상 프로세스 성능에 영향을 주지 않음

## 기여

기여는 언제나 환영합니다! 다음과 같이 기여해 주세요:

1. 이 저장소를 포크합니다
2. 기능 브랜치를 생성합니다 (`git checkout -b feature/AmazingFeature`)
3. 변경사항을 커밋합니다 (`git commit -m 'Add some AmazingFeature'`)
4. 브랜치에 푸시합니다 (`git push origin feature/AmazingFeature`)
5. Pull Request를 생성합니다

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다.

## 작성자

**lmk** - [GitHub](https://github.com/lmk)

## 감사의 글

- [struCoder/pidusage](https://github.com/struCoder/pidusage) - 크로스 플랫폼 프로세스 통계 제공

---

**참고**: 이 도구는 개발 및 디버깅 목적으로 설계되었습니다. 프로덕션 환경에서의 모니터링은 전문 APM 솔루션을 고려하세요
