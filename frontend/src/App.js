import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './pages/Home';
import SubjectQuestions from './pages/SubjectQuestions';
import QuestionForm from './pages/QuestionForm';
import AnswerForm from './pages/AnswerForm';  // 新しく追加

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/subject/:id" element={<SubjectQuestions />} />
            <Route path="/ask" element={<QuestionForm />} />
            <Route path="/subject/:subjectId/question/:questionId/answer" element={<AnswerForm />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;