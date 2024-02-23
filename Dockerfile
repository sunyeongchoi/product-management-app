## 첫 번째 스테이지: 빌드
FROM golang:alpine AS builder

## 라벨 설정
LABEL maintainer="suoung0716@gmail.com" \
      version="1.0.0" \
      description="product management application dockerfile"

## 환경변수 설정
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## 작업 디렉토리 설정
WORKDIR /app

## 소스 코드 복사
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

## 빌드 수행
RUN go build -o main .


## 두 번째 스테이지: 최소 크기 이미지 생성
FROM scratch

## 이전 스테이지에서 빌드된 실행 파일 복사
COPY --from=builder /app/main .

## 포트번호
EXPOSE 8080

## 실행
ENTRYPOINT ["./main"]