import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './QuestionForm.css'; // 新しいCSSファイルをインポート

function QuestionForm() {
  const navigate = useNavigate();
  const [subjects, setSubjects] = useState([]);
  const [formData, setFormData] = useState({
    subject: '',
    question: ''
  });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    setLoading(true);
    axios.get('http://localhost:8080/api/subjects')
      .then(response => {
        setSubjects(response.data);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching subjects:', error);
        setError('科目の取得に失敗しました。');
        setLoading(false);
      });
  }, []);

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    const payload = {
      ...formData,
      subject: parseInt(formData.subject, 10)
    };
    axios.post('http://localhost:8080/api/questions', payload)
      .then(response => {
        navigate(`/subject/${payload.subject}`);
      })
      .catch(error => {
        console.error('Error creating question:', error);
        setError('質問の作成に失敗しました。');
        setLoading(false);
      });
  };

  if (loading) return <div className="loading">Loading...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="question-form-container">
      <div className="form-header">
        <div className="form-icon">Q</div>
        <h2>新しい質問を作成</h2>
      </div>
      <form onSubmit={handleSubmit} className="question-form">
        <div className="form-group">
          <label htmlFor="subject">科目：</label>
          <select
            id="subject"
            name="subject"
            value={formData.subject}
            onChange={handleChange}
            required
          >
            <option value="">科目を選択してください</option>
            {subjects.map(subject => (
              <option key={subject.id} value={subject.id}>{subject.name}</option>
            ))}
          </select>
        </div>
        <div className="form-group">
          <label htmlFor="question">質問内容：</label>
          <textarea
            id="question"
            name="question"
            value={formData.question}
            onChange={handleChange}
            placeholder="ここに質問を入力してください..."
            required
          />
        </div>
        <div className="form-actions">
          <button type="submit" className="submit-button">質問を投稿</button>
          <button type="button" className="cancel-button" onClick={() => navigate('/')}>キャンセル</button>
        </div>
      </form>
    </div>
  );
}

export default QuestionForm;