const usersList = document.querySelector(".users__list")

const getAndDrawUsers = () => {
    fetch("http://172.20.10.7:8192/api/users/get-all")
        .then(res => res.json())
        .then(data => drawUsers(data.data))
}

const ucFirst = (str) => {
    return str[0].toUpperCase() + str.slice(1)
}

const drawUsers = (data) => {
    data.forEach(user => {
        usersList.insertAdjacentHTML("beforeend", `
            <li class="users__item users-item">
                <div class="users-item__column users-item__name">
                    <a class="users-item__link" href="/users/edit/${user.id}">${ucFirst(user.name)}</a>
                </div>
                <div class="users-item__column users-item__group">${ucFirst(user.group)}</div>
                <div class="users-item__column users-item__plot">${ucFirst(user.plot)}</div>
            </li>
        `)
    })
}

getAndDrawUsers()
