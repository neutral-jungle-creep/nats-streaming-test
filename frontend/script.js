const button = document.querySelector("#btn");
const input = document.querySelector('#input');
const textarea = document.querySelector('#textarea')

console.log(button)

button.addEventListener('click', (e) => {
    console.log("click")
    console.log(input.value)
    getOrderById()
});


async function getOrderById() {
    const response = await fetch('http://localhost:8008/get-order?id=' + input.value,
        {
            method: 'GET',
        })

    try {
        console.log(response)
        let res = await response.json();
        textarea.innerHTML = res.order_uid  + ' ' + res.track_number + res.entry
        console.log("ответ:", res.track_number)
    } catch (err) {
        console.log(err)
    }
}
