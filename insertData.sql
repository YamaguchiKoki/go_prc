-- articlesテーブルにデータを挿入
INSERT INTO articles (article_id, title, contents, user_name, nice, created_at)
VALUES 
(1, 'first blog', 'This is the my first blog.', 'saki', 1, NOW()),
(2, 'second article', 'This is the test article.', 'saki', 2, NOW());

-- commentsテーブルにデータを挿入
INSERT INTO comments (comment_id, article_id, message, created_at)
VALUES 
(1, 1, 'test comment1', NOW()),
(2, 1, 'second comment', NOW());
