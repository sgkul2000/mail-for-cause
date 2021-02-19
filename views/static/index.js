new Vue({
    el: "#ContentMain",
    delimiters:["[[", "]]"],
    data:{
        name: "shreesh"
    },
    methods:{
        redirect() {
            window.open("mailto:sgkul2000@gmail.com?subject=maimadharchodhoon&body=yebodyhai", "_blank")
        }
    }
})