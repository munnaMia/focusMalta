window.onload = function () {
    // All button
    const startBtn = document.getElementById("startBtn");
    const focusBtn = document.getElementById("focusBtn");
    const breakBtn = document.getElementById("breakBtn");
    const restBtn = document.getElementById("restBtn");
    const refreashBtn = document.getElementById("refreashBtn");

    // timer display
    const timerDisplay = document.getElementById("timer");

    // audio source
    let audio_beep = new Audio("/static/audio/beep-warning.mp3");

    // default timer will be set to zero and timer will store as seconds
    let timer = 0;
    let cashtimer = 0;
    let timerInMin = 0;
    let timerInSec = 0;
    let timerID = null; // timer ID to track a particular timer
    let startStopState = false; // state to maintainn start and end btn false mean timer is stop and vise varsa

    if (timer > 0) {
        displayTime(timerInMin, timerInSec);
    }

    // focus means a non stop 25 min of work.
    focusBtn.addEventListener("click", function () {
        if (!allowClick(timer)) {
            return;
        }
        startBtn.textContent = "Start";
        startBtn.style.backgroundColor = "#00ff007a";
        timer = 25 * 60; // time in seconds
        cashtimer = timer;
        calculateTimer(timer);
        displayTime(timerInMin, timerInSec);
    });

    // break means 5 min of break
    breakBtn.addEventListener("click", function () {
        if (!allowClick(timer)) {
            return;
        }
        startBtn.textContent = "Start";
        startBtn.style.backgroundColor = "#00ff007a";
        timer = 5*60; // time in seconds
        cashtimer = timer;
        calculateTimer(timer);
        displayTime(timerInMin, timerInSec);
    });

    // rest means 15 min of rest
    restBtn.addEventListener("click", function () {
        if (!allowClick(timer)) {
            return;
        }
        startBtn.textContent = "Start";
        startBtn.style.backgroundColor = "#00ff007a";
        timer = 15 * 60; // time in seconds
        cashtimer = timer;
        calculateTimer(timer);
        displayTime(timerInMin, timerInSec);
    });

    startBtn.addEventListener("click", function (e) {
        if (e.target.textContent == "Restart") {
            timer = cashtimer; // set to last state of timer
            calculateTimer(timer);
            displayTime(timerInMin, timerInSec);
        }
        if (timer == 0) {
            return;
        }

        // if false means timer is stop and true means timer is running
        if (!startStopState) {
            timerID = setInterval(countDown, 1000); // start the timer
            startStopState = true;
            startBtn.textContent = "Stop";
            startBtn.style.backgroundColor = "#ff000095";
        } else {
            clearInterval(timerID);
            timerID = null;
            startStopState = false;
            startBtn.textContent = "Start";
            startBtn.style.backgroundColor = "#00ff007a";
        }
    });

    // when the timer is stop we can reset it to 0
    refreashBtn.addEventListener("click", function () {
        if (!startStopState) {
            clearInterval(timerID);
            timerID = null;
            timer = 0;
            calculateTimer(timer);
            displayTime(timerInMin, timerInSec);
        }
    });

    function allowClick(timer) {
        switch (timer) {
            case 0:
            case 25 * 60:
            case 5 * 60:
            case 15 * 60:
                return true;
            default:
                return false;
        }
    }

    function countDown() {
        // clear timer id if its get 0 or less then
        if (timer == 0) {
            startBtn.textContent = "Restart";
            startBtn.style.backgroundColor = "#ff990095";
        }
        if (timer <= 0) {
            clearInterval(timerID);
            timerID = null;
            alertBeepSoundPlay();
        } else {
            timer--;
            calculateTimer(timer);
            displayTime(timerInMin, timerInSec);
        }
    }

    function alertBeepSoundPlay() {
        audio_beep.play();
    }

    function displayTime(min, sec) {
        timerDisplay.textContent = String(min).padStart(2, "0") + String(":") + String(sec).padStart(2, "0");
    }

    // calculate timer for each case
    function calculateTimer(timer) {
        timerInMin = Math.floor(timer / 60);
        timerInSec = Math.floor(timer % 60);
    }
};
