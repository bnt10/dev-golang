-- 사용자 테이블 생성
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    auth_token VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE
);

-- 초기 데이터 삽입
INSERT INTO users (username, email, password, last_login, auth_token, is_active) VALUES
('testuser1', 'testuser1@example.com', 'password1', NOW(), 'token1', TRUE),
('testuser2', 'testuser2@example.com', 'password2', NOW(), 'token2', TRUE);
