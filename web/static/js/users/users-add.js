const editForm = document.querySelector(".edit-user__form")
const editGroup = document.querySelector(".user-form__group")
const editPlot = document.querySelector(".user-form__plot")

let err = false
let ok = false

const drawData = (url, block) => {
  fetch(`http://172.20.10.7:8192/api/${url}`)
    .then(res => res.json())
    .then(data => {
      data.data.forEach(group => {
        block.insertAdjacentHTML("beforeend", `
                    <option value="${group.id}">${group.name}</option>
                `)
      })
    })
}
drawData("groups/get-all", editGroup)
drawData("plots/get-all", editPlot)

editForm.addEventListener("submit", e => {
  e.preventDefault()

  const formData = new FormData(editForm)

  if (formData.get("password") !== formData.get("password_repeat")) {
    if (!err) {
      ok = false
      err = true
      editForm.insertAdjacentHTML("beforeend", `
            <div class="user-form__block user-form__error">
                <h3>Пароли не совпадают</h3>
            </div>
        `)
    }

    return
  }

  const obj = {}
  formData.forEach(((value, key) => {
    switch (key) {
      case "id":
        obj[key] = Number(value)
        break
      case "password_repeat":
        break
      default:
        obj[key] = value
    }
  }))

  fetch("http://172.20.10.7:8192/api/users/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(obj)
  }).then(res => {
    if (res.ok) {
      if (!ok) {
        ok = true
        err = false
        editForm.insertAdjacentHTML('beforeend', `
                <div class="user-form__block user-form__succ">
                    <h3>Пользователь успешно добавлен</h3>
                </div>
            `)
      }
    }
  })
})