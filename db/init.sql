CREATE TABLE subjects (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE questions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  subject_id INT,
  question_text TEXT NOT NULL,
  FOREIGN KEY (subject_id) REFERENCES subjects(id)
);

CREATE TABLE answers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  question_id INT,
  answer_text TEXT NOT NULL,
  FOREIGN KEY (question_id) REFERENCES questions(id)
);

-- デフォルトの科目を挿入
INSERT INTO subjects (name) VALUES 
('応用数学'),
('情報理論'),
('ソフトウェア工学'),
('情報工学実験Ⅲ'),
('応用統計'),
('画像処理'),
('ネットワークシステム');

-- デフォルトの質問と回答を挿入（例）
INSERT INTO questions (subject_id, question_text) VALUES 
(1, '応用数学における微分方程式の重要性について教えてください。');

INSERT INTO answers (question_id, answer_text) VALUES 
(1, '微分方程式は応用数学において非常に重要です。物理現象や工学的問題を数学的にモデル化する際に頻繁に使用されます。');