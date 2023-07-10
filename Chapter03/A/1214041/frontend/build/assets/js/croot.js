import { setInner } from "https://jscroot.github.io/element/croot.js";
import {get} from "https://jscroot.github.io/api/croot.js";
import {rowtabel} from "../js/view/table"
setInner("namadivisi","Dari croot.js");
get("",);

function tableContent(result){
    console.log(result); //object json
}

function userTable(jsonParse){
    jsonParse.forEach((elemet,index) =>
    {
        rowtabel.replace("", element.).replace("",element);
        console.log(element._id);
        addEventListener("",row);
    });
}