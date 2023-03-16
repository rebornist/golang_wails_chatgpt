import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {ChatGPT, GoogleTranslate, Greet} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const [lang, setLang] = useState('en');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    function gpt() {
        ChatGPT(name).then(updateResultText);
    }

    function translate() {
        GoogleTranslate(lang, name).then(updateResultText);
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={translate}>Greet</button>
            </div>
        </div>
    )
}

export default App
