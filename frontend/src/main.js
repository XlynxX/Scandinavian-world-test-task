import './style.css';
import './app.css';

import {Generate} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <div class="result" id="result">Type number of symbols</div>
        <div class="input-box" id="input">
            <input class="input" id="number" type="number" autocomplete="off" />
        </div>
        
        <div style="margin-top: 1rem">
            <input type="checkbox" id="numbers" name="numbers" checked />
            <label for="numbers">Numbers (0-9)</label>
        </div>

        <div>
            <input type="checkbox" id="lowercaseLetters" name="lowercaseLetters" checked />
            <label for="lowercaseLetters">Lowercase letters (a-z)</label>
        </div>

        <div>
            <input type="checkbox" id="uppercaseLetters" name="uppercaseLetters" checked />
            <label for="uppercaseLetters">Uppercase letters (A-Z)</label>
        </div>

        <button class="btn" style="width:200px;height: 35px;margin-top:1rem" onclick="generate()">Generate</button>
        
        <h1 id="password">PASSWORD</h1>
    </div>
`;

let numberElement = document.getElementById("number");
let passwordElement = document.getElementById("password");

let numEl = document.getElementById("numbers");
let lowEl = document.getElementById("lowercaseLetters");
let upEl = document.getElementById("uppercaseLetters");

// Setup the greet function
window.generate = function () {
    // Get name
    let number = numberElement.value;
    let useNumbers = numEl.checked;
    let useLowerCaseLetters = lowEl.checked;
    let useUpperCaseLetters = upEl.checked;

    console.log(number, useNumbers, useLowerCaseLetters, useUpperCaseLetters);
    
    // Check if the input is empty
    if (number === "") return;

    try {
        Generate(parseInt(number), useNumbers, useLowerCaseLetters, useUpperCaseLetters)
            .then((result) => {
                // Update result with data back from App.Greet()
                passwordElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
