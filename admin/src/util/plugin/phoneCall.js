
export default function (number) {
    let a = document.createElement("a")
    a.href = "tel:" + number
    a.click()
    a.remove()
}