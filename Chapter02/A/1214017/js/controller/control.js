import { postWithToken} from "https://jscroot.github.io/api/croot.js";
import { Urlpost } from "../config/url";
import { formData } from "../template/temp"; 
import { setInner } from "https://jscroot.github.io/element/croot.js";
setInner("namadivisi","Dari croot.js");

form.addEventListener('submit', (e) => {
    e.preventDefault();
    postWithToken(Urlpost, "Token", "tokenrofi", formData, AmbilResponse)
    Postdata(formData);
    resetform();
  })