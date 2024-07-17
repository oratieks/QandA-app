import React from 'react';
import { Link } from 'react-router-dom';

function Header() {
  return (
    <header className="header">
      <h1>応用情報工学科 Q&Aアプリ</h1>
      <nav>
        <Link to="/" className="button">ホーム</Link>
        <Link to="/ask" className="button">質問する</Link>
      </nav>
    </header>
  );
}

export default Header;
第15回です