    var randomNumber1 = Math.floor(Math.random()*6)+1;
    var randomNumber2 = Math.floor(Math.random()*6)+1;

    document.getElementById('image1').src = `images/dice${randomNumber1}.png`;
    document.getElementById('image2').src = `images/dice${randomNumber2}.png`;
    if(randomNumber1>randomNumber2){
        document.getElementById('heading').innerHTML = `&#128681; Player1 Wins!`;
    }
    else if(randomNumber2>randomNumber1){
        document.getElementById('heading').innerHTML = ` Player2 Wins! &#128681;`;
    }
    else  document.getElementById('heading').innerHTML = ` Draw!`;

