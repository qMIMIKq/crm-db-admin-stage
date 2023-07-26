const usersList = document.querySelector(".users__list")

const getAndDrawUsers = () => {
  fetch("http://91.142.94.150:8192/api/users/get-all")
    .then(res => res.json())
    .then(data => {
      console.log(data.data)
      drawUsers(data.data)
    })
}

const ucFirst = (str) => {
  return str[0].toUpperCase() + str.slice(1)
}

const drawUsers = (data) => {
  data.forEach(user => {
    usersList.insertAdjacentHTML("beforeend", `
            <li class="users__item users-item">
                <div class="users-item__column users-item__name">
                    <a class="users-item__link" href="/users/edit/${user.id}">${user.name}</a>
                </div>
                <div class="users-item__column users-item__name">${user.nickname}</div>
                <div class="users-item__column users-item__group">${ucFirst(user.group)}</div>
                <div class="users-item__column users-item__plot">${user.plot}</div>
                <div class="users-item__column users-item__disabled">
                    ${user.disable ? 'Да' : 'Нет'}
                </div>
            </li>
        `)
  })
}

getAndDrawUsers()
