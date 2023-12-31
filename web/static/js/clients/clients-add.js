const addForm = document.querySelector(".edit-user__form")

let ok = false

addForm.addEventListener("submit", e => {
    e.preventDefault()

    const formData = new FormData(addForm)
    const obj = {}

    formData.forEach(((value, key) => {
        switch (key) {
            case "id":
                obj[key] = Number(value)
                break
            default:
                obj[key] = value.toLowerCase().trim()
        }
    }))

    fetch("http://192.168.1.230:8192/api/clients/add", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(obj)
    }).then(res => {
        if (res.ok) {
            if (!ok) {
                ok = true
                addForm.insertAdjacentHTML('beforeend', `
                <div class="user-form__block user-form__succ">
                    <h3>Клиент успешно добавлен</h3>
                </div>
            `)
            }
        }
    })
})