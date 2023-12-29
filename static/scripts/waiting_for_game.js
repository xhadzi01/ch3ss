// URLs for ready check and also for redirect
let isReadyURL = 'http://localhost:32000/is-ready-to-proceed';
let redirectURL = 'http://localhost:32000/proceed-to-game';
// variable to store our intervalID
let nIntervId;


function RedirectToGame(){
    window.location.assign(redirectURL);
}

function checkIsReadyCallback() {
    // to wait for results use 'then'
    fetch(isReadyURL).then(r=> r.json().then(isReadyVal=> {
        if (isReadyVal) {
            console.log('Oponent is connected.')
            UninitializePeriodicTask();
            RedirectToGame()
        } else {
            console.log('Oponent is NOT yet connected.')
        }
    }));
}

function InitializePeriodicTask() {
    // check if an interval has already been set up
    if (!nIntervId) {
      nIntervId = setInterval(checkIsReadyCallback, 1000);
    }
}

function UninitializePeriodicTask() {
    if (nIntervId) {
        clearInterval(nIntervId);
    }
}

// enable periodic start that will be checking state
InitializePeriodicTask();