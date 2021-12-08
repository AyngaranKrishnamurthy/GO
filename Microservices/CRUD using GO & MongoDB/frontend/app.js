import Navigo from "navigo";
import{ EmpServiceClient,AuthUserRequest, CreateEmpRequest, ReadEmpRequest, UpdateEmpRequest, DeleteEmpRequest} from './proto/services_grpc_web_pb'


const router = new Navigo
const EmpClient = new EmpServiceClient('http://localhost:9001')

router
    .on("/crud",function(){
        document.body.innerHTML = ""
        const homeDiv = document.createElement("div")
        homeDiv.classList.add("home-div")
        const user = localStorage.getItem("user")
        const buttonContainer = document.createElement("div")
        buttonContainer.classList.add("button-container")

        const createButton = document.createElement("button")
        createButton.innerText = "Create"
        createButton.addEventListener('click', () => {
            router.navigate("/C")
        })
        buttonContainer.appendChild(createButton)

        const readButton = document.createElement("button")
        readButton.innerText = "Retrieve"
        readButton.addEventListener('click', () => {
            router.navigate("/R")
        })
        buttonContainer.appendChild(readButton)

        const updateButton = document.createElement("button")
        updateButton.innerText = "Update"
        updateButton.addEventListener('click', () => {
            router.navigate("/U")
        })
        buttonContainer.appendChild(updateButton)

        const deleteButton = document.createElement("button")
        deleteButton.innerText = "Delete"
        deleteButton.addEventListener('click', () => {
            router.navigate("/D")
        })
        buttonContainer.appendChild(deleteButton)

        homeDiv.appendChild(buttonContainer)
        document.body.appendChild(homeDiv)
    })
    .on("/C",function(){
        document.body.innerHTML = ""
        const createDiv = document.createElement('div')
        createDiv.classList.add("crud-div")

        const createLabel = document.createElement('h1')
        createLabel.innerText = "Create New Collection"
        createDiv.appendChild(createLabel)

        const createForm = document.createElement("form")

        const idInputLabel = document.createElement("label")
        idInputLabel.classList.add("input-label")
        idInputLabel.innerText = "Employee Id"
        idInputLabel.setAttribute("for","id-input")
        createForm.appendChild(idInputLabel)

        const idInput = document.createElement("input")
        idInput.id = "id-input"
        idInput.setAttribute('type','text')
        idInput.setAttribute('placeholder','12345678')
        createForm.appendChild(idInput)

        const nameInputLabel = document.createElement("label")
        nameInputLabel.classList.add("input-label")
        nameInputLabel.innerText = "Employee Name"
        nameInputLabel.setAttribute("for","name-input")
        createForm.appendChild(nameInputLabel)

        const nameInput = document.createElement("input")
        nameInput.id = "name-input"
        nameInput.setAttribute('type','text')
        nameInput.setAttribute('placeholder','John Doe')
        createForm.appendChild(nameInput)

        const levelInputLabel = document.createElement("label")
        levelInputLabel.classList.add("input-label")
        levelInputLabel.innerText = "Career Level"
        levelInputLabel.setAttribute("for","level-input")
        createForm.appendChild(levelInputLabel)

        const levelInput = document.createElement("input")
        levelInput.id = "level-input"
        levelInput.setAttribute('type','text')
        levelInput.setAttribute('placeholder','00')
        createForm.appendChild(levelInput)

        const streamInputLabel = document.createElement("label")
        streamInputLabel.classList.add("input-label")
        streamInputLabel.innerText = "Assigned Stream"
        streamInputLabel.setAttribute("for","stream-input")
        createForm.appendChild(streamInputLabel)

        const streamInput = document.createElement("input")
        streamInput.id = "stream-input"
        streamInput.setAttribute('type','text')
        streamInput.setAttribute('placeholder','Mainframe')
        createForm.appendChild(streamInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Create!"
        createForm.appendChild(submitBtn)

        createForm.addEventListener('submit',event =>{
            event.preventDefault()
            let req = new CreateEmpRequest()
            req.setId(idInput.value)
            req.setName(nameInput.value)
            req.setLevel(levelInput.value)
            req.setStream(streamInput.value)
            EmpClient.createEmp(req, {},(err,res) => {
                if (err) return alert(err.message)
                //console.log(res.getResult())
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                EmpClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {eid: res.getEid(), empname: res.getEmpname(), Level: res.getLevel(), Stream: res.getStream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    alert(`Created Employee ${user.empname}`)
                    router.navigate("/crud")
                })
            })
        })

        createDiv.appendChild(createForm)

        document.body.appendChild(createDiv)
    })

    .on("/R",function(){
        document.body.innerHTML = ""
        const readDiv = document.createElement('div')
        readDiv.classList.add("crud-div")
        
        const readLabel = document.createElement('h1')
        readLabel.innerText = "Read Employee Details"
        readDiv.appendChild(readLabel)

        const readForm = document.createElement("form")

        const idInputLabel = document.createElement("label")
        idInputLabel.classList.add("input-label")
        idInputLabel.innerText = "Employee Id"
        idInputLabel.setAttribute("for","id-input")
        readForm.appendChild(idInputLabel)

        const idInput = document.createElement("input")
        idInput.id = "id-input"
        idInput.setAttribute('type','text')
        idInput.setAttribute('placeholder','12345678')
        readForm.appendChild(idInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Get Info"
        readForm.appendChild(submitBtn)

        readForm.addEventListener('submit',event =>{
            event.preventDefault()
            let req = new ReadEmpRequest()
            req.setId(idInput.value)
            EmpClient.readEmp(req, {},(err,res) => {
                if (err) return alert(err.messssage)
                //console.log(res.getResult())
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                EmpClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {eid: res.getEid(), empname: res.getEmpname(), Level: res.getLevel(), Stream: res.getStream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    alert(`Employee Details => ID: ${user.eid} | Name: ${user.empname} | Level: ${user.Level} | Stream: ${user.Stream}`)
                    router.navigate("/crud")
                })
            })
        })

        readDiv.appendChild(readForm)

        document.body.appendChild(readDiv)
        
        
    })

    .on("/U",function(){
        document.body.innerHTML = ""
        const UpdDiv = document.createElement('div')
        UpdDiv.classList.add("crud-div")
        
        const UpdLabel = document.createElement('h1')
        UpdLabel.innerText = "Update Employee Details"
        UpdDiv.appendChild(UpdLabel)

        const UpdForm = document.createElement("form")

        const idInputLabel = document.createElement("label")
        idInputLabel.classList.add("input-label")
        idInputLabel.innerText = "Employee Id"
        idInputLabel.setAttribute("for","id-input")
        UpdForm.appendChild(idInputLabel)

        const idInput = document.createElement("input")
        idInput.id = "id-input"
        idInput.setAttribute('type','text')
        idInput.setAttribute('placeholder','12345678')
        UpdForm.appendChild(idInput)

        const newNameLabel = document.createElement("label")
        newNameLabel.classList.add("input-label")
        newNameLabel.innerText = "New Name"
        newNameLabel.setAttribute("for","nname-input")
        UpdForm.appendChild(newNameLabel)

        const newNameInput = document.createElement("input")
        newNameInput.id = "nname-input"
        newNameInput.setAttribute('type','text')
        newNameInput.setAttribute('placeholder','12345678')
        UpdForm.appendChild(newNameInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Update Info"
        UpdForm.appendChild(submitBtn)

        UpdForm.addEventListener('submit',event =>{
            event.preventDefault() 
            let req = new UpdateEmpRequest()
            req.setId(idInput.value)
            req.setNname(newNameInput.value)
            EmpClient.updateEmp(req, {},(err,res) => {
                if (err) return alert(err.messssage)
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                EmpClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)

                    const user = {eid: res.getEid(), empname: res.getEmpname(), Level: res.getLevel(), Stream: res.getStream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    if(user.eid==""){
                        alert(`Employee ID: ${idInput.value} Updated Successfully`)                        
                    }else{ 
                        alert(`Employee ID doesn't exist`)
                                              
                    }
                    router.navigate("/crud")
                })
            })
        })

        UpdDiv.appendChild(UpdForm)

        document.body.appendChild(UpdDiv)
        
        
    })

    .on("/D",function(){
        document.body.innerHTML = ""
        const delDiv = document.createElement('div')
        delDiv.classList.add("crud-div")
        
        const delLabel = document.createElement('h1')
        delLabel.innerText = "Delete Employee Details"
        delDiv.appendChild(delLabel)

        const delForm = document.createElement("form")

        const idInputLabel = document.createElement("label")
        idInputLabel.classList.add("input-label")
        idInputLabel.innerText = "Employee Id"
        idInputLabel.setAttribute("for","id-input")
        delForm.appendChild(idInputLabel)

        const idInput = document.createElement("input")
        idInput.id = "id-input"
        idInput.setAttribute('type','text')
        idInput.setAttribute('placeholder','12345678')
        delForm.appendChild(idInput)

        const submitBtn = document.createElement("button")
        submitBtn.innerText = "Delete Info"
        delForm.appendChild(submitBtn)

        delForm.addEventListener('submit',event =>{
            event.preventDefault()
            let req = new DeleteEmpRequest()
            req.setId(idInput.value)
            EmpClient.deleteEmp(req, {},(err,res) => {
                localStorage.setItem('token', res.getToken())
                req = new AuthUserRequest()
                req.setToken(res.getToken())
                EmpClient.authUser(req, {} , (err,res)=>{
                    if (err) return alert (err.message)
                    const user = {eid: res.getEid(), empname: res.getEmpname(), Level: res.getLevel(), Stream: res.getStream()}
                    localStorage.setItem('user', JSON.stringify(user))
                    alert(`Deleted Employee ID: ${user.eid}`)
                    router.navigate("/crud")
                })
            })
        })

        delDiv.appendChild(delForm)

        document.body.appendChild(delDiv)
        
        
    })

    .resolve()