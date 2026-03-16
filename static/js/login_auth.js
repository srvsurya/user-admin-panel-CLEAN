
console.log("JS loaded")
const form = document.querySelector("form");
const error = document.getElementById("error");

form.addEventListener("submit",function(e){
const email = document.getElementById("email").value.trim();
const password = document.getElementById("password").value.trim();
    if (!email && !password){
        e.preventDefault();
        error.textContent="No email or password entered";
    }
    else if(!email){
        e.preventDefault();
        error.textContent="No email entered";
    }
    else if(!password){
        e.preventDefault();
        error.textContent="No password entered";
    }
    else{
        error.textContent="";
    }
})