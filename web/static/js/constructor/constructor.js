const usersList = document.querySelector(".users__list")
// const usersPosList = document.querySelector(".users__list--const")

let filters = []

const getAndDrawUsers = async () => {
  await fetch("http://91.142.94.150:8192/api/filters/get-all")
    .then(res => res.json())
    .then(data => {
      if (data.data !== null) {
        filters = data.data.filter(d => !d.disable)
        drawUsers(filters)
      } else {
        usersList.insertAdjacentHTML("beforeend", `
                    <li class="users__item users-item">
                        <div class="users-item__column">Пока нет записей</div>
                    </li>
                `)
      }
    })
}

const removeUsers = () => {
  usersList.querySelectorAll(".users__item").forEach(elem => elem.remove())
}

const submitFilters = async (data) => {
  await fetch("http://91.142.94.150:8192/api/filters/edit-position", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  })

}

const drawUsers = (data) => {
  data.forEach((filter, i) => {
    usersList.insertAdjacentHTML("beforeend", `
                <li id="filter-${filter.id}" class="users__item users-item">
                    <div class="users-item__column users-item__name">
                        <div style="cursor: default" class="users-item__link">${filter.name}</div>
                    </div>
                    <div class="users-item__column users-item__up">
                        <div style="cursor: pointer" class="users-item__link">Выше</div>
                    </div>
                    <div class="users-item__column users-item__down">
                        <div style="cursor: pointer" class="users-item__link">Ниже</div>
                    </div>
                </li>
            `)
    filter.position = i

    const currElem = document.querySelector(`#filter-${filter.id}`)
    const upBtn = currElem.querySelector('.users-item__up')
    const downBtn = currElem.querySelector('.users-item__down')

    upBtn.addEventListener('click', e => {
      const ind = filters.findIndex(filt => filt.id === filter.id)

      if (ind === 0) {
        const temp = filters[filters.length - 1].position
        filters[filters.length - 1].position = filters[ind].position
        filters[ind].position = temp

        ;[filters[filters.length - 1], filters[ind]] = [filters[ind], filters[filters.length - 1]]
        submitFilters(filters)
          .then(() => {
            removeUsers()
            getAndDrawUsers()
          })
        console.log(filters)
      } else {
        const temp = filters[ind - 1].position
        filters[ind - 1].position = filters[ind].position
        filters[ind].position = temp

        ;[filters[ind - 1], filters[ind]] = [filters[ind], filters[ind - 1]]
        submitFilters(filters)
          .then(() => {
            removeUsers()
            getAndDrawUsers()
          })
      }
    })

    downBtn.addEventListener('click', e => {
      const ind = filters.findIndex(filt => filt.id === filter.id)
      console.log(filters.length, ind)

      if (ind + 1 === filters.length) {
        const temp = filters[0].position
        filters[0].position = filters[ind].position
        filters[ind].position = temp

        ;[filters[0], filters[ind]] = [filters[ind], filters[0]]
        submitFilters(filters).then(() => {
          removeUsers()
          getAndDrawUsers()
        })
      } else {
        const temp = filters[ind + 1].position
        filters[ind + 1].position = filters[ind].position
        filters[ind].position = temp

        ;[filters[ind + 1], filters[ind]] = [filters[ind], filters[ind + 1]]
        submitFilters(filters).then(() => {
          removeUsers()
          getAndDrawUsers()
        })
      }
    })
  })
}

getAndDrawUsers().then(() => {
})