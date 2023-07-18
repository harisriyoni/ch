
import { postWithToken, get } from "https://jscroot.github.io/api/croot.js";
import { Urlpost } from "../config/post";
import { data } from "../template/data";

function pushbutton(){

    function AmbilResponse(result) {
        console.log(result);
        alert('Data berhasil disimpan!');
    }
    postWithToken(Urlpost, "Login", "tokenjasmine", data, AmbilResponse)

};  