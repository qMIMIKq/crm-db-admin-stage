const addForm = document.querySelector(".edit-user__form")
const plotsSection = document.querySelector("#plot")

const exceptions = ["все"]

const drawData = (url, block) => {
    fetch(`http://192.168.1.231:8192/api/${url}`)
        .then(res => res.json())
        .then(data => {
            data.data.forEach(group => {
                if (!checkExc(exceptions, group.name)) {
                    block.insertAdjacentHTML("beforeend", `
                    <option value="${group.id}">${group.name}</option>
                `)
                }
            })
        })
}
drawData("plots/get-all", plotsSection)

const checkExc = (data, val) => {
    let flag = false
    data.forEach(d => {
        if (d === val) {
            flag = true
        }
    })

    return flag
}

let ok = false
let err = false

addForm.addEventListener("submit", e => {
    e.preventDefault()

    const formData = new FormData(addForm)
    const obj = {}

    const startTime = formData.get("start_time")
    const endTime = formData.get("end_time")


    if (startTime.length || endTime.length) {
        if (!validateTime(startTime, endTime)) {
            err = true

            addForm.insertAdjacentHTML("beforeend", `
            <div class="user-form__block user-form__error">
                <h3>Некорректный формат времени</h3>
            </div>
        `)

            return
        }
    }

    formData.forEach(((value, key) => {
        switch (key) {
            case 'disable':
                obj[key] = value === "on"
                break
            default:
                obj[key] = value.toLowerCase()
        }
    }))
    fetch("http://192.168.1.231:8192/api/filters/add", {
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
                    <h3>Фильтр успешно добавлен</h3>
                </div>
            `)
            }
        }
    })
})

const validateTime = (startTime, endTime) => {
    const arrStartTime = startTime.split(":")
    const arrEndTime = endTime.split(":")

    if (arrStartTime.length !== 2 || arrEndTime.length !== 2) {
        return false
    }

    if (Number(arrStartTime[0]) > Number(arrEndTime[0])) {
        return false
    }

    for (let i = 0; i < arrStartTime.length; i++) {
        if (arrStartTime[i].length !== 2 || arrEndTime[i].length !== 2) {
            return false
        }
    }

    return true
}