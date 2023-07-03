import { setInner } from "https://jscroot.github.io/element/croot.js";


export let Urlpost = "http://gocroot-baru.herokuapp.com/insertjum"

export function AmbilResponse(result) {
    console.log(result);
    setInner("inisemua", result.message);
}

export function resetform(){

  }