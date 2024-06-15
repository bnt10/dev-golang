# Golang 1.21.10 이미지를 기반으로 합니다.
FROM golang:1.21.10

# 작업 디렉토리 설정
WORKDIR /workspace

# 필요한 도구 설치
RUN apt-get update && apt-get install -y \
    git \
    curl \
    bash \
    && rm -rf /var/lib/apt/lists/*

# golangci-lint 설치
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/go/bin v1.50.0

# 환경 변수 설정
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# vscode 사용자를 생성
RUN groupadd -g 1000 vscode \
    && useradd -u 1000 -g vscode -m vscode \
    && echo "vscode ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# 작업 디렉토리 소유권 변경
RUN chown -R vscode:vscode /workspace

# git safe directory 설정
RUN git config --global --add safe.directory /workspace

# 사용자의 권한 설정
USER vscode

# 기본 명령 설정
CMD [ "sleep", "infinity" ]