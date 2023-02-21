let app = new Vue({
    el: '#FinalProject',
    data() {
        return {
            // tab page
            activeName: 'CompaniePage',
            contentListHeight: 2000,
            item:{
                label:"",
                value:""
              },
              list: [
              {
                 label: "Companies",
                 value: "CompaniePage"
              },
              {
                 label: "Layoff Charts",
                 value: "ChartsPage"
              },
              {
                 label: "List of Employees Laid Off",
                 value: "ListLaidOffPage"
              },
             ],
        
        }
    },
    components: {
        "HeadCon": {
            template: "#header",
            data(){
                return{
                    options: [{
                        value: 'In 2022',
                        label: 'In 2022',
                      }, {
                        value: 'In 2023',
                        label: 'In 2023'
                      }], 
                    selectValue:'In 2023'
                }
            }
        },
    },
    methods: {
        handleClick:function(tab) {
            let result = document.getElementById('mainContent').src = './'+ tab.name +'.html';
            console.log('result',result)
         },
        getContentListHeight: function () {
            var parentHeight = document.getElementById('mainContent').parentNode.offsetHeight;
            return parentHeight - 100;
        }
    },
    mounted() {
        this.$nextTick(function () {
            app.contentListHeight = app.getContentListHeight();
        })
    }
});