import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import QuestionList from '../components/QuestionList';

function SubjectQuestions() {
  const { id } = useParams();
  const [questions, setQuestions] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get(`http://localhost:8080/api/subjects/${id}/questions`)
      .then(response => {
        console.log('Raw response:', response.data);
        // 質問をグループ化
        const groupedQuestions = response.data.reduce((acc, item) => {
          if (!acc[item.id]) {
            acc[item.id] = { ...item, answers: [] };
          }
          if (item.answer_text) {
            acc[item.id].answers.push(item.answer_text);
          }
          return acc;
        }, {});
        setQuestions(Object.values(groupedQuestions));
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching questions:', error);
        setError('質問一覧を取得できませんでした。');
        setLoading(false);
      });
  }, [id]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (questions.length === 0) {
    return <div className="message">この科目には質問がありません。</div>;
  }

  return (
    <div className="subject-questions">
      <h2>質問一覧</h2>
      <QuestionList questions={questions} subjectId={id} />
    </div>
  );
}

export default SubjectQuestions;