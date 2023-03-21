import { postWithToken, get } from "https://jscroot.github.io/api/croot.js";
import { Urlpost } from "../config/url";
import { data } from "../template/temp";

function pushbutton(){

    function AmbilResponse(result) {
        console.log(result);
        alert('Data berhasil disimpan!');
    }

    postWithToken(Urlpost, "Token", "tokenrofi", data, AmbilResponse)

};  