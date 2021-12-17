import Navigo from 'navigo'
import{crudServiceClient,CreateRequest,AuthUserRequest,RetrieveRequest, UpdateRequest, DeleteRequest} from './proto/service_grpc_web_pb'

const router = new Navigo()
const crudClient = new crudServiceClient("http://localhost:9090")

router

    .on("/",function(){
        document.body.innerHTML = ""
        const homeDiv = document.createElement("div")
        homeDiv.classList.add("home-div")
        const user = localStorage.getItem("user")
        
        const buttonContainer = document.createElement("div")
        buttonContainer.classList.add("button-container")

        const createButton = document.createElement("button")
        createButton.innerText = "Create"
        createButton.addEventListener('click', () => {
            router.navigate("/NewEmployee")
        })
        buttonContainer.appendChild(createButton)

        const readButton = document.createElement("button")
        readButton.innerText = "Retrieve"
        readButton.addEventListener('click', () => {
            router.navigate("/RetrieveEmployee")
        })
        buttonContainer.appendChild(readButton)

        homeDiv.appendChild(buttonContainer)

        document.body.appendChild(homeDiv)
    })

    .on("/NewEmployee", function(){
        document.body.innerHTML=""
        const createDiv = document.createElement("div")
        createDiv.classList.add("crud-div")
        
        const createLabel = document.createElement('h1')
        createLabel.innerText = "New Employee Registration"
        createDiv.appendChild(createLabel)
        
        const createForm = document.createElement("form")

        const empidLabel = document.createElement("label")
        empidLabel.classList.add("input-label")
        empidLabel.setAttribute("for","empid-input")
        empidLabel.innerText = "Employee Id "
        createForm.appendChild(empidLabel)

        const empidInput = document.createElement("input")
        empidInput.setAttribute("type","text")
        empidInput.id = "empid-input"
        empidInput.setAttribute("placeholder","Eg: 12345678")
        createForm.appendChild(empidInput)

        empidInput.addEventListener('input',() =>{
            empidError.innerText = ""
            const empid = empidInput.value
            if (empid.length !=8) {
                empidError.innerText = "Warning: Invalid Employee Id"
                return
            }
        })

        const empidError = document.createElement("div")
        empidError.id = "empid-error"
        empidError.classList.add("error")
        createForm.appendChild(empidError)

        const nameLabel = document.createElement("label")
        nameLabel.classList.add("input-label")
        nameLabel.setAttribute("for","name-input")
        nameLabel.innerText = "Employee Name "
        createForm.appendChild(nameLabel)

        const nameInput = document.createElement("input")
        nameInput.setAttribute("type","text")
        nameInput.id = "name-input"
        nameInput.setAttribute("placeholder","Eg: Jane Doe")
        createForm.appendChild(nameInput)

        nameInput.addEventListener('input',() =>{
            nameError.innerText = ""
            const name = nameInput.value
            if (name.length <8) {
                nameError.innerText = "Employee name must be a full name. i.e:{firstname lastname}"
                return
            }
        })

        
        const nameError = document.createElement("div")
        nameError.id = "name-error"
        nameError.classList.add("error")
        createForm.appendChild(nameError)

        const levelLabel = document.createElement("label")
        levelLabel.classList.add("input-label")
        levelLabel.setAttribute("for","level-input")
        levelLabel.innerText = "Level"
        createForm.appendChild(levelLabel)

        const levelInput = document.createElement("input")
        levelInput.setAttribute("type","int")
        levelInput.id = "level-input"
        levelInput.setAttribute("placeholder","09")
        createForm.appendChild(levelInput)

        levelInput.addEventListener('input',() =>{
            levelError.innerText = ""
            const level = levelInput.value
            if (level>13) {
                levelError.innerText = "Career Level should be between 1 & 13"
                return
            }
        })

        const levelError = document.createElement("div")
        levelError.id = "level-error"
        levelError.classList.add("error")
        createForm.appendChild(levelError)

        const streamLabel = document.createElement("label")
        streamLabel.classList.add("input-label")
        streamLabel.setAttribute("for","stream-input")
        streamLabel.innerText = "Stream"
        createForm.appendChild(streamLabel)

        const streamInput = document.createElement("input")
        streamInput.setAttribute("type","int")
        streamInput.id = "stream-input"
        streamInput.setAttribute("placeholder","Java")
        createForm.appendChild(streamInput)

        const createBtn = document.createElement("button")
        createBtn.innerText = "Register"
        createForm.appendChild(createBtn)

        createForm.addEventListener('submit', event => {
            event.preventDefault()
            if (empidInput.value == "" || empidError.innerText != "" || nameInput.value == "" || nameError.innerText != "" || levelInput.value == "" || levelError.innerText != "" || streamInput.value == "") return
            let request = new CreateRequest()
            request.setEid(empidInput.value)
            request.setEmpname(nameInput.value)
            request.setElevel(levelInput.value)
            request.setEstream(streamInput.value)
            crudClient.create(request, {}, (err, res) => {
                if(err) return alert(err)
                localStorage.setItem('token', res.getToken())
                request = new AuthUserRequest()
                request.setToken(res.getToken())
                crudClient.authUser(request, {}, (err, res) => {
                    if(err) return alert(err)
                    const user = { eid: res.getEid(), empname: res.getEmpname(), elevel: res.getElevel(), estream: res.getEstream() }
                    localStorage.setItem("user", JSON.stringify(user))
                    alert (`Employee ${user.empname} was registered successfully`)
                    router.navigate("/")
                })
            })
        })

        createDiv.appendChild(createForm)

        document.body.appendChild(createDiv)
    })

    .on("/RetrieveEmployee", function(){
        document.body.innerHTML = ""
        const retrieveDiv = document.createElement('div')
        retrieveDiv.classList.add("crud-div")

        const retrieveLabel = document.createElement('h1')
        retrieveLabel.innerText = "Retrieve Employee Details"
        retrieveDiv.appendChild(retrieveLabel)

        const retrieveForm = document.createElement("form")

        const empid1Label = document.createElement("label")
        empid1Label.classList.add("input-label")
        empid1Label.innerText = "Employee Id"
        empid1Label.setAttribute("for","retrieve-input")
        retrieveForm.appendChild(empid1Label)

        const empid1Input = document.createElement("input")
        empid1Input.id = "retrieve-input"
        empid1Input.setAttribute('type','text')
        empid1Input.setAttribute('placeholder','Eg: 12345678')
        retrieveForm.appendChild(empid1Input)

        const retrieveBtn = document.createElement("button")
        retrieveBtn.innerText = "Retrieve" 
        retrieveForm.appendChild(retrieveBtn)

        retrieveForm.addEventListener("submit", event => {
            event.preventDefault()
            let req = new RetrieveRequest()
            req.setEid(empid1Input.value)
            crudClient.retrieve(req, {} , (err,res) =>{
                if (err) return alert (err.message)
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                crudClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {eid: res.getEid(), empname: res.getEmpname(), elevel: res.getElevel(), estream: res.getEstream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    alert(`Data retrieved successfully`)
                    router.navigate("/EmployeeDetails")
                })
            })
        })

        retrieveDiv.appendChild(retrieveForm)

        document.body.appendChild(retrieveDiv)
    })

    .on("/EmployeeDetails",function(){
            document.body.innerHTML = ""
            const crudDiv = document.createElement("div")
            crudDiv.classList.add("crud-div")
            let user = JSON.parse(localStorage.getItem("user"))
            
            const buttonContainer = document.createElement("div")
            buttonContainer.classList.add("button-container")
    
            const crudLabel = document.createElement('h6')
            crudLabel.innerText = `Employee Details: ${user.eid} | ${user.empname} | ${user.elevel} | ${user.estream}`
            buttonContainer.appendChild(crudLabel)
    
            const updBtn = document.createElement("button")
                updBtn.innerText = "Update"
    
                updBtn.addEventListener('click', () => {
                    router.navigate("/UpdateDetails")
                })
                buttonContainer.appendChild(updBtn)
    
            const crudBtn = document.createElement("button")
                crudBtn.innerText = "Delete"
    
                crudBtn.addEventListener('click', () => {
                    //event.preventDefault()
                    let req = new DeleteRequest()
                    req.setEid(user.eid)
                    crudClient.delete(req, {},(err,res) => {
                    if (err) return alert(err.messssage)
                    alert(`Employee details deleted Successfully`)
                    localStorage.setItem("user", "null")
                    localStorage.setItem("token","null")
                    router.navigate("/")
                })
                //newly added ends
            
                })
                buttonContainer.appendChild(crudBtn)
    
            const createButton = document.createElement("button")
            createButton.innerText = "CRUD"
            createButton.addEventListener('click', () => {
                router.navigate("/")
            })
            buttonContainer.appendChild(createButton)
    
            crudDiv.appendChild(buttonContainer)
            
        
            document.body.appendChild(crudDiv)
            
        
    })

    .on("/UpdateDetails",function(){

        document.body.innerHTML = ""
        const UpdDiv = document.createElement('div')
        UpdDiv.classList.add("crud-div")
        let user = JSON.parse(localStorage.getItem("user"))

        const UpdLabel = document.createElement('h1')
        UpdLabel.innerText = "Update Employee Details"
        UpdDiv.appendChild(UpdLabel)

        const UpdForm = document.createElement("form")

        const newNameLabel = document.createElement("label")
        newNameLabel.classList.add("input-label")
        newNameLabel.innerText = "New Name"
        newNameLabel.setAttribute("for","nname-input")
        UpdForm.appendChild(newNameLabel)

        const newNameInput = document.createElement("input")
        newNameInput.id = "nname-input"
        newNameInput.setAttribute('type','text')
        newNameInput.setAttribute('placeholder','John Doe')
        UpdForm.appendChild(newNameInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Update Name"
        UpdForm.appendChild(submitBtn)

        UpdForm.addEventListener('submit',event =>{
            event.preventDefault()
            let req = new UpdateRequest()
            req.setEid(user.eid)
            req.setEmpname1(newNameInput.value)
            crudClient.update(req, {},(err,res) => {
                if (err) return alert(err.message)
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                crudClient.authUser(req, {} , (err,res)=>{
                if (err) return alert (err.message)
                
                //retrieves again
                let req = new RetrieveRequest()
                req.setEid(user.eid)
                crudClient.retrieve(req, {} , (err,res) =>{
                    if (err) return alert (err.message)
                    localStorage.setItem('token', res.getToken())
                    req = new AuthUserRequest()
                    req.setToken(res.getToken())
                    crudClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {eid: res.getEid(), empname: res.getEmpname(), elevel: res.getElevel(), estream: res.getEstream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    alert('Name Updated Successfully')
                    router.navigate("/EmployeeDetails")
                    })
                })
                //

                })
            })
        })
        const crudBtn = document.createElement("button")
            crudBtn.innerText = "Cancel"

            crudBtn.addEventListener('click', () => {
                router.navigate("/EmployeeDetails")
            })
            UpdForm.appendChild(crudBtn)

        UpdDiv.appendChild(UpdForm)

        document.body.appendChild(UpdDiv)
    
    })

.resolve()