# API 문서
API 문서는 [링크](https://documenter.getpostman.com/view/9858969/2sA2rCUh2K)에서 확인할 수 있습니다.

# 환경 세팅
### mysql5.7 설치
프로젝트에 있는 `mysql-docker-compose.yaml` 파일을 통해 실행할 수 있습니다.
```yaml
version: "3"
services:
  mysql:
    image: mysql:5.7.38
    container_name: mysql5.7
    environment:
      MYSQL_DATABASE: productmgm
      MYSQL_USER: admin
      MYSQL_ROOT_PASSWORD: passwd
      MYSQL_PASSWORD: passwd
      TZ: Asia/Seoul
      LANG: C.UTF-8
    volumes:
      - /tmp/mysql_data:/var/lib/mysql
    command: '--character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --skip-character-set-client-handshake'
    ports:
      - "3306:3306"
```
### Dockerfile 빌드 방법
1. Dockerfile의 환경변수 설정
    ```Dockerfile
    ## 환경변수 설정 (환경에 따라 값 설정 필요)
    ENV DB_HOST=localhost\
        DB_PORT=3306\
        DB_USERNAME=admin\
        DB_PASSWORD=passwd\
        DB_DATABASE=productmgm\
        JWT_KEY=example\
        JWT_TIME_DURATION=50000\
        IS_PRODUCTION=true
    ```
2. product-management-app 프로젝트의 루트 경로 이동
3. Dockerfile 기반 빌드 명령어
`docker build -t IMAGE_NAME .`  
3. 실행 명령어
`docker run -p 8080:8080 IMAGE_NAME`


# 프로젝트 개발 컨벤션
### 린트 체크
`make lint` 명령을 통해 린트를 체크해야합니다.
