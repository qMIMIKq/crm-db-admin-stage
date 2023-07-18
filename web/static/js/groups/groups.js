const groupsList = document.querySelector(".groups__list")

const getAndDrawGroups = () => {
    fetch("http://192.168.1.231:8192/api/groups/get-all")
        .then(res => res.json())
        .then(data => drawGroups(data.data))
}

const ucFirst = (str) => {
    return str[0].toUpperCase() + str.slice(1)
}

const drawGroups = (data) => {
    data.forEach(group => {
        groupsList.insertAdjacentHTML("beforeend", `
            <li class="groups__item groups-item">
                <div class="groups__column users-item__name">
                     <a class="groups-item__link" href="/groups/edit/${group.id}">${ucFirst(group.name)}</a>
                </div>
                
                <div class="groups__column users-item__name">
                    ${ucFirst(group.description)}
                </div>
            </li>
        `)
    })
}

getAndDrawGroups()
