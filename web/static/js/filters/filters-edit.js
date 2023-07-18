const editForm = document.querySelector(".edit-user__form")
const editPlot = document.querySelector("#plot")
const filterPlot = document.querySelector("#plot option").textContent

let err = false
let ok = false

const exceptions = ["все"]

const drawData = (url, block, userData) => {
    fetch(`http://172.20.10.7:8192/api/${url}`)
        .then(res => res.json())
        .then(data => {
            data.data.forEach(group => {
                if (group.name !== userData && !checkExc(exceptions, group.name)) {
                    block.insertAdjacentHTML("beforeend", `
                    <option value="${group.id}">${group.name}</option>
                `)
                }
            })
        })
}
drawData("plots/get-all", editPlot, filterPlot)

const checkExc = (data, val) => {
    let flag = false
    data.forEach(d => {
        if (d === val) {
            flag = true
        }
    })

    return flag
}

editForm.addEventListener("submit", e => {
    e.preventDefault()

    const formData = new FormData(editForm)

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

    const obj = {}
    formData.forEach(((value, key) => {
        switch (key) {
            case "id":
                obj[key] = Number(value)
                break
            case 'disable':
                obj[key] = value === "on"
                console.log(value === "on")
                break
            default:
                obj[key] = value.trim().toLowerCase()
        }
    }))


    fetch("http://172.20.10.7:8192/api/filters/edit", {
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
                    <h3>Фильтр успешно изменен</h3>
                </div>
            `)
            }
        }

        return res.json()
    }).then(data => {
        console.log(data)
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