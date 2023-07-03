let validation = document.getElementById("validation")
let btn = document.getElementById("insert-data")
let close_insert = document.getElementById("close_validation")
function validations(val, bool) {
    if (bool) {
      validation.classList.remove("hidden");
      document.getElementById("warn_validation").innerText = val;
    } else {
      validation.classList.add("hidden");
    }
  }

btn.addEventListener("click",()=>validations("anu enak",true))
close_insert.addEventListener("click",()=>validations("anu enak",false))