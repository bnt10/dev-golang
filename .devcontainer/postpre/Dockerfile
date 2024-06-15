# PostgreSQL 이미지를 기반으로 합니다.
FROM postgres:14

# 데이터베이스 초기화 스크립트 복사
COPY init.sql /docker-entrypoint-initdb.d/
