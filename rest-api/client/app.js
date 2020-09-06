import axios from 'axios'

/*Base endpoint of the REST api*/
const BASE_ENDPOINT = "http://localhost:1366"


console.log("REST Client Running")
const URI_TSHIRT = "https://i.ibb.co/qYz06y9/t-shirts.png";
const URI_SWEATSHIRT = "https://i.ibb.co/bgKQYNh/tracksuit.png"
const URI_PANT = "https://i.ibb.co/cbFffwV/pants.png"


/*
    Fetching all the garments
*/

const getall = document.getElementById('getall')

getall.addEventListener('click', () => {
    checkChair()
})



const checkChair = () => {

    axios.get(`${BASE_ENDPOINT}/all-garments`).then(RESPONSE => {
        console.log("FETCH ALL GARMENTS RESPONSE", RESPONSE)

        let data = RESPONSE.data.reverse();

        const stack = document.getElementById('chairItems')
        stack.innerHTML = "";

        data.map(item => {
            console.log(item.garment)

            switch (item.garment) {
                case "tshirt":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_TSHIRT}" alt=""></div>
                        </div>`;
                    break;
                case "sweatshirt":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_SWEATSHIRT}" alt=""></div>
                        </div>`;
                    break;
                case "pant":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_PANT}" alt=""></div>
                        </div>`;
            }
        })

    }).catch(err => {
        alert("Error while getting all garments", err)
    })


}


/*
    Adding garments
*/

const tshirt = document.getElementById("tshirt")
const pant = document.getElementById("pant")
const sweatshirt = document.getElementById("sweatshirt")


tshirt.addEventListener('click', () => {
    addGarment('tshirt')
})
sweatshirt.addEventListener('click', () => {
    addGarment('sweatshirt')
})

pant.addEventListener('click', () => {
    addGarment('pant')
})



const addGarment = (item) => {

    //Fetching from /add-garment

    axios.get(`${BASE_ENDPOINT}/add-garment?garment=${item}`).then(RESPONSE => {
        console.log("ADD GARMENT RESPONSE ", RESPONSE)
        let data = RESPONSE.data.reverse()

        const stack = document.getElementById('chairItems')
        stack.innerHTML = "";

        data.map(item => {
            switch (item.garment) {
                case "tshirt":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_TSHIRT}" alt=""></div>
                        </div>`;
                    break;
                case "sweatshirt":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_SWEATSHIRT}" alt=""></div>
                        </div>`;
                    break;
                case "pant":
                    stack.innerHTML = stack.innerHTML + `<div class="garment">
                            <div><img class="icons shirt" src="${URI_PANT}" alt=""></div>
                        </div>`;
            }
        })
    }).catch(err => {
        alert("Error while adding garment", err)
    })


}

/*
    Doing laundry to clear the chair
*/

const wash = document.getElementById("wash")

wash.addEventListener('click', () => {
    console.log("Clicked wash")
    clearChair()

})

const clearChair = () => {

    axios.get(`${BASE_ENDPOINT}/clear`).then(RESPONSE => {
        console.log(RESPONSE)

        const stack = document.getElementById('chairItems')
        stack.innerHTML = "";

        const yay = document.getElementById('yay')

        yay.style.display = 'block';

        setTimeout(() => {
            yay.style.display = 'none'
        }, 2000)
    }).catch(err => {
        alert("Error while clearing", err)
    })
}