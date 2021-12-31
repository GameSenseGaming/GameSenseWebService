import React from 'react';
import './style/App.css';
import Footer from'./components/Footer';
import Navbar from './components/navbar/Navbar';




function App() {
  return (
    <div className="page-container">
      <div className='content-wrap'>
      <Navbar />
      
      </div>

      <div>
      <Footer />
      </div>
      
    </div>
  );
}

export default App;
