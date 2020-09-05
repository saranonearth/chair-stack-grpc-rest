import {
    AddServiceClient,
    Item,
    ItemList,
} from "./proto/addGarment_grpc_web_pb";

const addclient = new AddServiceClient("http://localhost:9003");

let req = new Item();

req.setGarment("TShirt");

addclient.addGarment(req, {}, (err, res) => {
    console.log(err);
    console.log(res.getListList());
});