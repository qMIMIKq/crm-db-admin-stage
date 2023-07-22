const usersList = document.querySelector(".users__list")

const getAndDrawUsers = () => {
  fetch("http://192.168.1.231:8192/api/filters/get-all-hidden")
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
  data.forEach(filter => {
    usersList.insertAdjacentHTML("beforeend", `
            <li class="users__item users-item">
                <div class="users-item__column users-item__name">
                    <a class="users-item__link" href="/filters/edit/${filter.id}">${filter.name}</a>
                </div>
                <div class="users-item__column users-item__group">${filter.plot}</div>
                <div class="users-item__column users-item__disabled">
                    ${filter.disable ? 'Да' : 'Нет'}
                </div>
            </li>
        `)
  })
}

getAndDrawUsers()
