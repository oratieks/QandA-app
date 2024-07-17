import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import './AnswerForm.css'; // 新しいCSSファイルをインポート

function AnswerForm() {
  const { subjectId, questionId } = useParams();
  const navigate = useNavigate();
  const [question, setQuestion] = useState(null);
  const [answer, setAnswer] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    setLoading(true);
    axios.get(`http://localhost:8080/api/questions/${questionId}`)
      .then(response => {
        setQuestion(response.data);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching question:', error);
        setError('質問の取得に失敗しました。');
        setLoading(false);
      });
  }, [questionId]);

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    axios.post(`http://localhost:8080/api/questions/${questionId}/answers`, { answer_text: answer })
      .then(() => {
        navigate(`/subject/${subjectId}`);
      })
      .catch(error => {
        console.error('Error submitting answer:', error);
        setError('回答の送信に失敗しました。');
        setLoading(false);
      });
  };

  if (loading) return <div className="loading">Loading...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="answer-form-container">
      <div className="question-card">
        <div className="question-icon">Q</div>
        <h2 className="question-text">{question.question_text}</h2>
      </div>
      <form onSubmit={handleSubmit} className="answer-form">
        <div className="form-group">
          <label htmlFor="answer">あなたの回答：</label>
          <textarea
            id="answer"
            value={answer}
            onChange={(e) => setAnswer(e.target.value)}
            placeholder="ここに回答を入力してください..."
            required
          />
        </div>
        <div className="form-actions">
          <button type="submit" className="submit-button">回答を送信</button>
          <button type="button" className="cancel-button" onClick={() => navigate(`/subject/${subjectId}`)}>キャンセル</button>
        </div>
      </form>
    </div>
  );
}

export default AnswerForm;