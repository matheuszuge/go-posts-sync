CREATE TABLE IF NOT EXISTS posts (
  id      INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  title   VARCHAR(255) NOT NULL,
  body    TEXT NOT NULL,
  CONSTRAINT fk_posts_users FOREIGN KEY (user_id) REFERENCES users(id)
);
