const plotsList = document.querySelector(".plots__list")

const getAndDrawPlots = () => {
  fetch("http://192.168.1.231:8192/api/plots/get-all")
    .then(res => res.json())
    .then(data => drawPlots(data.data))
}

const ucFirst = (str) => {
  return str[0].toUpperCase() + str.slice(1)
}

const drawPlots = (data) => {
  data.forEach(plot => {
    plotsList.insertAdjacentHTML("beforeend", `
            <li class="plots__item plots-item">
                <div class="plots__column users-item__name">
                    <a class="plots-item__link" href="/plots/edit/${plot.id}">${plot.name}</a>
                </div>
                <div class="plots__column users-item__name">
                    <div class="">${plot.short_name}</вшм>
                </div>
            </li>
        `)
  })
}

getAndDrawPlots()
