
function checkPassword(){
    const password=document.getElementById("password").value.trim();
    const confirmPassword=document.getElementById("confirmPassword").value.trim();
    if (password!==confirmPassword){
        error.textContent="The Password you've entered does not match!";
    }
    else{
        error.textContent="";
}

const form=document.querySelector("form");
form.addEventListener("submit",function(e){
    const password=document.getElementById("password").value;
    const confirmPassword=document.getElementById("confirmPassword").value;
    const name1=document.getElementById("name").value;
    const email=document.getElementById("email").value;
    if(!name1){
        e.preventDefault();
        error.textContent="Please enter a name";
    }
    else if(!email){
        e.preventDefault();
        error.textContent="Please enter an email";
    }
    else if(!password){
        error.textContent="You'll have to enter a password"
    }
    else if(password!==confirmPassword){
        e.preventDefault();
        error.textContent="Please make sure the passwords match before you click submit!";
    }
    else{
        error.textContent="";
        
    }
})
}
async function backendEmailCheck(){
    const email=document.getElementById("email").value;
     const response = await fetch("/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ email, password })
    });

    const data = await response.json();

    if (!response.ok) {
        document.getElementById("error").innerText = data.error;
        return;
    }
    window.location.href = "/login";
}