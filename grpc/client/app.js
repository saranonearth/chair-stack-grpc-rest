import { GarmentServiceClient, GetAllRequest, Item, ClearRequest } from './proto/garmentActions_grpc_web_pb'


/* 
Connecting to the GarmentService by creating a instance
of GarmentServiceClient
*/

const garmentClient = new GarmentServiceClient("http://localhost:9003")



console.log("gRPC Client Running")
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

    let req = new GetAllRequest();


    garmentClient.allGarments(req, {}, (err, res) => {
        if (err) return alert("ERROR while fetching garments!")

        const data = giveMeListofObjects(res.array[0]);

        console.log(data)

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
    let req = new Item();
    req.setGarment(item);

    garmentClient.addGarment(req, {}, (err, res) => {
        if (err) return alert("ERROR while adding garment!")

        const data = giveMeListofObjects(res.array[0]);

        console.log(data)

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
    let req = new ClearRequest()


    garmentClient.clearList(req, {}, (err, res) => {
        console.log(err)
        console.log(res)
    })

    const stack = document.getElementById('chairItems')
    stack.innerHTML = "";

    const yay = document.getElementById('yay')

    yay.style.display = 'block';

    setTimeout(() => {
        yay.style.display = 'none'
    }, 2000)
}


/*
 Util functions
*/

const giveMeListofObjects = (data) => {
    let result = [];

    data.map(item => {
        result = [{ id: item[0], garment: item[1] }, ...result]
    })

    return result
}