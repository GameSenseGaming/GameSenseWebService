import React from 'react';
import './App.css';
import Footer from'./Components/Footer';
import Navbar from './Components/Navbar/Navbar';




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
