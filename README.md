# monitorTheProcess
monitor the process

## 개요
이 프로젝트는 특정 프로세스의 CPU와 메모리 사용량을 모니터링합니다.

## 사용 방법
1. 리포지토리를 클론합니다:
    ```sh
    git clone https://github.com/lmk/monitorTheProcess.git
    ```
2. 디렉토리로 이동합니다:
    ```sh
    cd monitorTheProcess
    ```
3. 프로그램을 빌드합니다:
    ```sh
    go build
    ```
4. 프로그램을 실행합니다:
    ```sh
    ./monitorTheProcess -pid <프로세스 ID> -duration <기간(초)>
    ```

## 예제
```sh
./monitorTheProcess -pid 1234 -duration 60
