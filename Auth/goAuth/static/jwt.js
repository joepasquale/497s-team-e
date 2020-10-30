async function login(){
    const loginURL = "http://localhost:5555/login"
    let username = document.getElementById("username").value
    let password = document.getElementById("password").value
    let data = {'username':username, 'password':password}
    let r = await fetch(loginURL, {
        method: 'POST',
        mode:"cors",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        body: JSON.stringify(data),
    })
    let j;
    try {
        j = await r.json();
    } catch (e) {
        document.getElementById("output").innerHTML = e
    }

    document.getElementById("output").innerHTML = JSON.stringify(j)
}