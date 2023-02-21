let app = new Vue({
    el: '#Companie',
    data() {
        return {
            tableList: [
                {
                    Company:'Bloomberg',
                    Location:'New York',
                    LaidOff:'1200',
                    Date:'9/2/2023',
                    Industry:'',
                    Source:'',
                    Country:'',

                },
                {
                    Company:'Bloomberg',
                    Location:'New York',
                    LaidOff:'1200',
                    Date:'9/2/2023',
                    Industry:'',
                    Source:'',
                    Country:'',

                },
            ], 
        }
    },
    components: {
       
    },
    methods: {
        deleteClick(val) {
            let id = val;
            console.log('the id is:',id) 
            this.$emit("remove-asset",id);// the value of sending to father compoent
        },
    },
    mounted() {
       
    }
});