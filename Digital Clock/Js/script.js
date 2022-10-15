function clock() {
    let time = new Date();
    var hrs = time.getHours();
    var mins = time.getMinutes();
    var sec = time.getSeconds();

    let hour = document.getElementById('hours');
    let minutes = document.getElementById('minutes');
    let seconds = document.getElementById('seconds');

    // hour.innerHTML = hrs;
    // minutes.innerHTML = mins;
    // seconds.innerHTML = sec;

    // Hours

    if (hrs > 12) {
        hour.innerHTML = "0" + (hrs - 12);
    }
    else {
        hour.innerHTML = hrs;
    }

    // Minutes

    if (mins < 10) {
        minutes.innerHTML = '0' + mins;
    }
    else {
        minutes.innerHTML = mins;
    }

    // Seconds

    if (sec < 10) {
        seconds.innerHTML = '0' + sec;
    }
    else {
        seconds.innerHTML = sec;
    }

}

var dt = setInterval(clock, 1000);