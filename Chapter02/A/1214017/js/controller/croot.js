import { postWithToken } from "https://jscroot.github.io/api/croot.js";
import { onClick, getValue } from "https://jscroot.github.io/element/croot.js";
import { Urlpost, AmbilResponse} from "../config/url.js";


function forOnClick(){
  let data = {
    bulan : getValue("bulan"),
    tahun : getValue("tahun"),
    jumlah : getValue("jumlah")
  }
  postWithToken(Urlpost, "Token", "tokenrofi", data, AmbilResponse);

}

onClick("button", forOnClick);

