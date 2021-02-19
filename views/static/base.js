new Vue({
    el: "#main",
    delimiters: ["[[", "]]"],
    data:{
        isLogged: false,
        user: null
    },
    mounted() {
        let isLogged = JSON.parse(localStorage.getItem("auth"))
    }
})