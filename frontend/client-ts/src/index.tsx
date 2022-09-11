import React from 'react';
import ReactDOM from 'react-dom/client';
import './style/index.scss';
import App from './App';
import { BrowserRouter } from 'react-router-dom';
import "./fonts/fontawesome-free-6.1.2-web/css/all.css"


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
    <BrowserRouter>
          <App />
  </BrowserRouter>
);


