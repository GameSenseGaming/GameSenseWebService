import React from 'react';
import ReactDOM from 'react-dom';
import './style/index.css';
import App from './App';
import {I18nextProvider} from "react-i18next";
import i18next from "i18next";

i18next.init({
  interpolation: { escapeValue: false },  // React already does escaping
});

ReactDOM.render(
    <App/>
    ,
  document.getElementById('root')
);
