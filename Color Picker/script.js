// get elements
const container=document.getElementById("container");
const infoContainer=document.getElementById("Info")
const proBtn=document.getElementById("proceedBtn");
let rgbVal,start=false;

// add click listener on proceed button
proBtn.addEventListener("click",(e)=>{
    infoContainer.style.display="none";
    let timer=setTimeout(()=>{
        start=true;
        clearTimeout(timer);
    },1000)
})

    // add click event on document and change background color of container
    document.addEventListener("mousemove",(e)=>{
        let x=e.clientX;
        let y=e.clientY;
        rgbVal=`rgb(${x%255},${y%255},${Math.abs(x+y)%255})`;
        start?container.style.backgroundColor=rgbVal:null;
    });

    // add mouseclick event on document to copy it to clipboard
    document.addEventListener("click",(e)=>{
        start?navigator.clipboard.writeText(rgbVal):null;
        start?alert("Copied the color: "+rgbVal+" to clipboard"):null;
    })

