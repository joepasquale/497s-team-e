const url = "localhost:5555/register"; 
const loginURL = "localhost:5555/login.html"; 

// const url = "http://localhost:8080/register"
// const loginURL = "http://localhost:8080/login"
// NEW: helper method for posting data
async function postData(url, data) {
    const resp = await fetch(url,
		{
			method: 'POST',
			mode: 'cors',
			cache: 'no-cache',
			credentials: 'same-origin',
			headers: {
				'Content-Type': 'application/json'
			},
			redirect: 'follow',
			body: JSON.stringify(data)
		});
    return resp;
}


function checker() {
	let username = document.getElementById("username").value;
	let email = document.getElementById("email").value;
	let password = document.getElementById("password").value;
	let confirmPassword = document.getElementById("password-confirm").value;
	let flag = true
	if (password == confirmPassword) {
		if (password.match(/[a-z]/g) && password.match(/[A-Z]/g) && password.match(/[0-9]/g)&& password.length>=8){
			document.getElementById("password-prompt").setAttribute("hidden", "true");
			return true
		}
		else {
			let message = "Must contain at least one number and one uppercase and lowercase letter, and at least 8 or more characters";
			document.getElementById("password-prompt").removeAttribute('hidden');
			document.getElementById("password-prompt").innerHTML = message;
			flag = false
		}
	} else {
		let message = "Two password not the same";
		document.getElementById("password-prompt").removeAttribute('hidden');
		document.getElementById("password-prompt").innerHTML = message;
		flag = false;
	}
	if(username === null || password === null) {
		let message = "You must provide a username and email!";
		document.getElementById("email-prompt").removeAttribute('hidden');
		document.getElementById("email-prompt").innerHTML = message;
		document.getElementById("username-prompt").removeAttribute('hidden');
		document.getElementById("username-prompt").innerHTML = message;
		flag = false
	}
	return flag
};


function register(){
	(async () => {
		console.log('running register')
		let username = document.getElementById("username").value;
		let password = document.getElementById("password").value;
		// let confirmPassword = document.getElementById("password-confirm").value;
		console.log('running register\n' + username + " " + email)
		if(checker()===true){
			const data = { 'username' : username, 'password' : password }; // -- (1)
			const resp = await postData(url, data); 
            const j = await resp.json();
            document.getElementById("output").innerHTML =  JSON.stringify(j);
		}
	})()
}