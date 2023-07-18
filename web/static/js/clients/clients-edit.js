const editForm = document.querySelector(".edit-user__form")

let err = false
let ok = false

editForm.addEventListener("submit", e => {
    e.preventDefault()

    const formData = new FormData(editForm)
    const obj = {}
    formData.forEach(((value, key) => {
        switch (key) {
            case "id":
                obj[key] = Number(value)
                break
            default:
                obj[key] = value.trim()
        }
    }))

    fetch("http://192.168.1.231:8192/api/clients/edit", {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(obj)
    }).then(res => {
        if (res.ok) {
            if (!ok) {
                ok = true
                editForm.insertAdjacentHTML('beforeend', `
                <div class="user-form__block user-form__succ">
                    <h3>Клиент успешно изменен</h3>
                </div>
            `)
            }
        }

        return res.json()
    }).then(data => {
        console.log(data)
    })
})
