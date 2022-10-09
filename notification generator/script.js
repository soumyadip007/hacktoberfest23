const ask=document.querySelector(".askbtn"), trybtn = document.querySelector(".trybtn"), generate = document.querySelector(".genbtn");
i=0;
ask.addEventListener("click", () =>{
    Notification.requestPermission().then(perm => {
        if (perm === "granted") {
            document.querySelector(".askaccess").classList.remove("active");
            document.querySelector(".notif").classList.add("active");
        }
        else{
            document.querySelector(".askaccess").classList.remove("active");
            document.querySelector(".disabled").classList.add("active");
        }
    })
} )
generate.addEventListener("click", ()=>{
   if (document.querySelector(".user").value != "" && document.querySelector(".body").value!="" && document.querySelector(".subject").value!="") {
	const usermsg="sent from: "+ (document.querySelector(".user").value);
	   const bodymsg= ", The message says: "+ (document.querySelector(".body").value);
	   const bod = usermsg+bodymsg;
	   const subject= "Notification subject: "+ (document.querySelector(".subject").value);
	    new Notification(subject, {
	        body: bod,
            icon: "https://cdn-icons-png.flaticon.com/512/1592/1592461.png",
	    })
}
else{
    alert("All the fields are required, please fill out!");
}
})
trybtn.addEventListener("click", ()=>{
    i++;
    location.reload();

})
if (i>=2) {
    alert("Seems like it isn't working. I suggest to check if you have blocked notifications in the site settings.");
}
