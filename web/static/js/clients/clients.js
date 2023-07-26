const usersList = document.querySelector(".users__list")

//91.142.94.150

const getAndDrawUsers = () => {
  fetch("http://91.142.94.150:8192/api/clients/get-all")
    .then(res => res.json())
    .then(data => {
      if (data.data !== null) {
        drawUsers(data.data)
      } else {
        usersList.insertAdjacentHTML("beforeend", `
                    <li class="users__item users-item">
                        <div class="users-item__column">Пока нет записей</div>
                    </li>
                `)
      }
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
                    <a class="users-item__link" href="/clients/edit/${user.id}">${user.name}</a>
                </div>
            </li>
        `)
  })
}

getAndDrawUsers()
