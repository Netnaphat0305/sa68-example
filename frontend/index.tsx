import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './src/App';
import 'antd/dist/antd.css'; 
import './index.css'; // ถ้ามี

const root = ReactDOM.createRoot(document.getElementById('root')!);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
