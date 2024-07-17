import React from 'react';
import { Link } from 'react-router-dom';

function QuestionList({ questions, subjectId }) {
  return (
    <div className="question-list">
      {questions.map(question => (
        <div key={question.id} className="question-item">
          <div className="question-header">
            <div className="question-icon">Q</div>
            <h3>{question.question_text}</h3>
          </div>
          <div className="answer-section">
            <div className="answer-icon">A</div>
            <div className="answer-content">
              {question.answers && question.answers.length > 0 ? (
                <div className="answers">
                  {question.answers.map((answer, index) => (
                    <p key={index} className="answer">{answer}</p>
                  ))}
                </div>
              ) : (
                <p>回答がありません。</p>
              )}
            </div>
          </div>
          <div className="question-actions">
            <Link to={`/subject/${subjectId}/question/${question.id}/answer`} className="button answer-button">
              回答する
            </Link>
          </div>
        </div>
      ))}
    </div>
  );
}

export default QuestionList;