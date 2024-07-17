import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

function SubjectList() {
  const [subjects, setSubjects] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    axios.get('http://localhost:8080/api/subjects')
      .then(response => {
        console.log('Raw response:', response.data);
        setSubjects(response.data);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching subjects:', error);
        setError('科目一覧が取得できませんでした。');
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (!subjects || subjects.length === 0) {
    return <div className="message">科目が存在しません。</div>;
  }

  return (
    <div className="subject-list">
      {subjects.map(subject => (
        <Link key={subject.id} to={`/subject/${subject.id}`} className="subject-item">
          {subject.name}
        </Link>
      ))}
    </div>
  );
}

export default SubjectList;