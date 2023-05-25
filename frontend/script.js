const button = document.querySelector("#btn");
const input = document.querySelector('#input');
const order_area = document.querySelector('#order_area')
const delivery_area = document.querySelector('#delivery_area')
const payment_area = document.querySelector('#payment_area')
const items_area = document.querySelector('#items_area')
const e_label = document.querySelector('#e_label')


button.addEventListener('click', (e) => {
    console.log("click")
    console.log(input.value)
    getOrderById()
});


async function getOrderById() {
    let items_res = '';
    e_label.textContent = '';

    const response = await fetch('http://localhost:8008/get-order?id=' + input.value,
        {
            method: 'GET',
        })

    if (response.status == 200) {
        console.log(response)
        let res = await response.json();
        let i = 0;
        while (i < res.items.length) {
            item_res = 'chrt id: ' + res.items[i].chrt_id + '<br>' +
                'track number: ' + res.items[i].track_number + '<br>' +
                'price: ' + res.items[i].price + '<br>' +
                'rid: ' + res.items[i].rid + '<br>' +
                'name: ' + res.items[i].name + '<br>' +
                'sale: ' + res.items[i].sale + '<br>' +
                'size: ' + res.items[i].size + '<br>' +
                'total price: ' + res.items[i].total_price + '<br>' +
                'nm id: ' + res.items[i].nm_id + '<br>' +
                'brand: ' + res.items[i].brand + '<br>' +
                'status: ' + res.items[i].status +'<hr>'

            items_res += item_res
            console.log(res.items[i]);
            i++;
        }

        order_res = 'order uid: ' + res.order_uid + '<br>' +
            'track number: ' + res.track_number + '<br>' +
            'entry: ' + res.entry + '<br>' +
            'locale: ' + res.locale + '<br>' +
            'internal signature: ' + res.internal_signature + '<br>' +
            'customer id: ' + res.customer_id + '<br>' +
            'delivery service: ' + res.delivery_service + '<br>' +
            'shardkey: ' + res.shardkey + '<br>' +
            'sm id: ' + res.sm_id + '<br>' +
            'date created: ' + res.date_created + '<br>' +
            'oof_shard: ' + res.oof_shard

        delivery_res = 'name: ' + res.delivery.name + '<br>' +
            'phone: ' + res.delivery.phone + '<br>' +
            'zip: ' + res.delivery.zip + '<br>' +
            'city: ' + res.delivery.city + '<br>' +
            'address: ' + res.delivery.address + '<br>' +
            'region: ' + res.delivery.region + '<br>' +
            'email: ' + res.delivery.email

        payment_res = 'transaction: ' + res.payment.transaction + '<br>' +
            'request id: ' + res.payment.request_id + '<br>' +
            'currency: ' + res.payment.currency + '<br>' +
            'provider: ' + res.payment.provider + '<br>' +
            'amount: ' + res.payment.amount + '<br>' +
            'payment dt: ' + res.payment.payment_dt + '<br>' +
            'bank: ' + res.payment.bank + '<br>' +
            'delivery cost: ' + res.payment.delivery_cost + '<br>' +
            'goods total: ' + res.payment.goods_total + '<br>' +
            'custom fee: ' + res.payment.custom_fee

        order_area.innerHTML = order_res
        delivery_area.innerHTML = delivery_res
        payment_area.innerHTML = payment_res
        items_area.innerHTML = items_res
        console.log("ответ:", res.track_number)
    }
    if (response.status == 204) {
        e_label.textContent = "заказ не найден"
    }

}
