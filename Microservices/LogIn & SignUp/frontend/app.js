import Navigo from 'navigo'
import{AuthServiceClient, LoginRequest,SignupRequest, AuthUserRequest} from './proto/service_grpc_web_pb'

const router = new Navigo()

const authClient = new AuthServiceClient('http://localhost:9001')

router
    .on("/", function(){
        document.body.innerHTML = ""
        const homeDiv = document.createElement("div")
        homeDiv.classList.add("home-div")
        const user = JSON.parse(localStorage.getItem("user"))
        if (user==null){
            const buttonContainer = document.createElement("div")
            buttonContainer.classList.add("button-container")

            const loginButton = document.createElement("button")
            loginButton.innerText = "Login"
            loginButton.addEventListener('click', () =>{
                router.navigate("/login")
            })
            buttonContainer.appendChild(loginButton)

            
            const signupButton = document.createElement("button")
            signupButton.innerText = "Sign Up"
            signupButton.addEventListener('click', () =>{
                router.navigate("/signup")
            })
            buttonContainer.appendChild(signupButton)

            homeDiv.appendChild(buttonContainer)

        }else{
            const authText = document.createElement("div")
            authText.classList.add("auth-text")
            authText.innerText = `Logged in as ${user.username}`

            const logoutBtn = document.createElement("button")
            logoutBtn.innerText = "Logout"

            logoutBtn.addEventListener('click', () => {
                localStorage.setItem("user", "null")
                localStorage.setItem("token","null")
                window.location.reload()
            })

            authText.appendChild(logoutBtn)
            homeDiv.appendChild(authText)
        }
        document.body.appendChild(homeDiv)
    })
    /* .on("/about", function(){
        document.body.innerHTML = "About"
    }) */
    .on("/login", function(){
        document.body.innerHTML = ""
        const loginDiv = document.createElement('div')
        loginDiv.classList.add("authentication-div")

        const loginLabel = document.createElement('h1')
        loginLabel.innerText = "Login"
        loginDiv.appendChild(loginLabel)

        const loginForm = document.createElement("form")

        const loginInputLabel = document.createElement("label")
        loginInputLabel.classList.add("input-label")
        loginInputLabel.innerText = "Username / E-mail"
        loginInputLabel.setAttribute("for","login-input")
        loginForm.appendChild(loginInputLabel)

        const loginInput = document.createElement("input")
        loginInput.id = "login-input"
        loginInput.setAttribute('type','text')
        loginInput.setAttribute('placeholder','Eg: John Doe')
        loginForm.appendChild(loginInput)

        const passwordInputLabel = document.createElement("label")
        passwordInputLabel.classList.add("input-label")
        passwordInputLabel.innerText = "Password"
        passwordInputLabel.setAttribute("for","password-input")
        loginForm.appendChild(passwordInputLabel)

        const passwordInput = document.createElement("input")
        passwordInput.setAttribute("type","password")
        passwordInput.id = "password-input"
        passwordInput.setAttribute("placeholder","Password")
        loginForm.appendChild(passwordInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Login" 
        loginForm.appendChild(submitBtn)

        loginForm.addEventListener("submit", event => {
            event.preventDefault()
            let req = new LoginRequest()
            req.setLogin(loginInput.value)
            req.setPassword(passwordInput.value)
            authClient.login(req, {} , (err,res) =>{
                if (err) return alert (err.message)
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                authClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {id: res.getId(), username: res.getUsername(), email: res.getEmail()}
                    localStorage.setItem('user', JSON.stringify(user))
                    router.navigate("/")
                })
            })
        })

        loginDiv.appendChild(loginForm)

        document.body.appendChild(loginDiv)
    })
    .on("/signup", function(){
        document.body.innerHTML=""
        const signupDiv = document.createElement("div")
        signupDiv.classList.add("authentication-div")
        
        const signupLabel = document.createElement('h1')
        signupLabel.innerText = "Signup"
        signupDiv.appendChild(signupLabel)
        
        const signupForm = document.createElement("form")

        const usernameLabel = document.createElement("label")
        usernameLabel.classList.add("input-label")
        usernameLabel.setAttribute("for","username-input")
        usernameLabel.innerText = "Username "
        signupForm.appendChild(usernameLabel)

        const usernameInput = document.createElement("input")
        usernameInput.setAttribute("type","text")
        usernameInput.id = "username-input"
        usernameInput.setAttribute("placeholder","Eg: John Doe")
        signupForm.appendChild(usernameInput)

        usernameInput.addEventListener('input',() =>{
            usernameError.innerText = ""
            const username = usernameInput.value
            if (username.length <4 || username.length >20) {
                usernameError.innerText = "Warning: Username must be minimum 4 characters and maximum 20 characters"
                return
            }
        })

        const usernameError = document.createElement("div")
        usernameError.id = "username-error"
        usernameError.classList.add("error")
        signupForm.appendChild(usernameError)

        const emailLabel = document.createElement("label")
        emailLabel.classList.add("input-label")
        emailLabel.setAttribute("for","email-input")
        emailLabel.innerText = "Email Id "
        signupForm.appendChild(emailLabel)

        const emailInput = document.createElement("input")
        emailInput.setAttribute("type","email")
        emailInput.id = "email-input"
        emailInput.setAttribute("placeholder","Eg: john_doe@gmail.com")
        signupForm.appendChild(emailInput)

        emailInput.addEventListener('input',() =>{
            emailError.innerText = ""
            const email = emailInput.value
            if (email.length <12 || email.length >40) {
                emailError.innerText = "Invalid Email Address"
                return
            }
            if (!(new RegExp("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").test(email))){
                emailError.innerText = "Invalid E-mail Id"
            }
        })

        
        const emailError = document.createElement("div")
        emailError.id = "email-error"
        emailError.classList.add("error")
        signupForm.appendChild(emailError)

        const passwordLabel = document.createElement("label")
        passwordLabel.classList.add("input-label")
        passwordLabel.setAttribute("for","password-input")
        passwordLabel.innerText = "Password"
        signupForm.appendChild(passwordLabel)

        const passwordInput = document.createElement("input")
        passwordInput.setAttribute("type","password")
        passwordInput.id = "password-input"
        passwordInput.setAttribute("placeholder","######")
        signupForm.appendChild(passwordInput)

        passwordInput.addEventListener('input',() =>{
            passwordError.innerText = ""
            const password = passwordInput.value
            if (password.length <8 || password.length >20) {
                passwordError.innerText = "Password must be minimum of 8 characters and maximum of 20 characters"
                return
            }
        })

        const passwordError = document.createElement("div")
        passwordError.id = "password-error"
        passwordError.classList.add("error")
        signupForm.appendChild(passwordError)

        const signupBtn = document.createElement("button")
        signupBtn.innerText = "SignUp"
        signupForm.appendChild(signupBtn)

        signupForm.addEventListener('submit', event => {
            event.preventDefault()
            if (usernameInput.value == "" || usernameError.innerText != "" || emailInput.value == "" || emailError.innerText != "" || passwordInput.value == "" || passwordError.innerText != "") return
            let request = new SignupRequest()
            request.setUsername(usernameInput.value)
            request.setEmail(emailInput.value)
            request.setPassword(passwordInput.value)
            authClient.signup(request, {}, (err, res) => {
                if(err) return alert(err)
                localStorage.setItem('token', res.getToken())
                request = new AuthUserRequest()
                request.setToken(res.getToken())
                authClient.authUser(request, {}, (err, res) => {
                    if(err) return alert(err)
                    const user = { id: res.getId(), username: res.getUsername(), email: res.getEmail() }
                    localStorage.setItem("user", JSON.stringify(user))
                    router.navigate("/")
                })
            })
        })

        signupDiv.appendChild(signupForm)

        document.body.appendChild(signupDiv)
    })
    .resolve()