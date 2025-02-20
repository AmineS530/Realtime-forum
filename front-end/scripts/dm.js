function dms_ToggleShowSidebar(event) {
    console.log("dms_ShowSidebar",document.getElementById("backdrop").classList, event);
    document.getElementById("backdrop").classList.toggle("show");
    console.log("dms_ShowSidebar after",document.getElementById("backdrop").classList);
}