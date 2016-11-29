/**
 * Created by gensh on 2016/11/3.
 */
const Config = {
    base: "/",
    appName: "Virtual Judge",
    logo: "/static/img/logo.png",
    auth: {
        github: {
            name: "Github",
            logo: "/static/img/github.png",
            enable: true,
            client_id: "94b2759733e24ee2994d",
            redirect_uri: "http://localhost:8080/auth/callback/github",
            auth_url: "https://github.com/login/oauth/authorize/?redirect_uri=R&client_id=C",
        }
    },
    //homePage do not end with "/"
    OJs: [{name: "Local", homePage: ""}, {name: "POJ", homePage: "http://poj.org"}]
};
var userInfo = {name: "", avatar: "", id: 0, is_login: false};

//init config
// Config.authUrl += "&redirect_uri=" + Config.authCallback;
Vue.filter('formatTime', formatTime);

var Home = Vue.extend({
    template: '#template-home',
    created: function () {
        var toolbar = $('#vj_app_toolbar');
        var docToolbarHeight = toolbar.outerHeight();
        $(window).on('scroll', function () {
            if ($(this).scrollTop() > docToolbarHeight) {
                toolbar.addClass('waterfall');
            } else {
                toolbar.removeClass('waterfall');
            }
        });
    },
    mounted: function () {
        $('#vj_app_toolbar').removeClass('waterfall');
    },
    destroyed: function () {
        $(window).off('scroll');
        $('#vj_app_toolbar').addClass('waterfall');
    }
});

var Problems = Vue.extend({
    template: '#template-problems',
    data: function () {
        return {problems: [], ojs: Config.OJs};
    },
    methods: {
        OJName: function (type) {
            return getOJNameByType(type);
        }, formatUpdatedAt: function (time) {
            return formatTime(time);
        }, sourceUrl: function (type, url) {
            return formatSourceUrl(type, url)
        },
        //
        filterProblemsByOj: function (index) { //remember set button text
            console.log(index);
        }
    },
    created: function () {
        $.ajax({
            url: Config.base + "problems",
            context: this,
            success: function (data) {
                if (data == null) {
                    data = [];
                }
                try {
                    var self = this;
                    data.forEach(function (e) {
                        self.problems.push(e)
                    });
                } catch (e) {
                    new Snackbar("error happened while loading problema data.", {timeout: 3500});
                }
            }, error: function (err) {
                new Snackbar("error happened while loading problems data.", {timeout: 3500});
            }
        });
    }
});

var ProblemDetail = Vue.extend({
    template: "#template-problem-detail",
    methods: {},
    computed: {
        ojName: function () {
            return getOJNameByType(this.detail.oj);
        }, formatUpdatedAt: function () {
            return formatTime(this.detail.updated_at);
        }, sourceUrl: function () {
            return formatSourceUrl(this.detail.oj, this.detail.source_url);
        }
    },
    data: function () {
        return {detail: {}, loading_status: -1, ojs: Config.OJs}; //-1 is not loaded,0 loading,1:loaded, 2:error loading,3:not found
    },
    created: function () {
        var id = this.$route.params.id;
        this.loading_status = 0;
        $.ajax({
            url: Config.base + "problem/detail/" + id,
            context: this,
            success: function (data) {
                if (data != null) {
                    try {
                        this.detail = data;
                        /*"id","problem_id","title","describe","hint","input","input_sample", "output","output_sample",
                         "ac_count","submitted_count","mem_limit","time_limit","oj","origin_id","origin_url",
                         "source","source_url","created_at","updated_at":*/
                        this.loading_status = 1;
                    } catch (e) {
                        new Snackbar("error happened while loading problema data.", {timeout: 3500});
                        this.loading_status = 2;
                    }
                } else {
                    this.loading_status = 3; //not found
                }
            }, error: function (err) {
                this.loading_status = 2;
                new Snackbar("error happened while loading problems data.", {timeout: 3500});
            }
        });
    }
});

var Submit = Vue.extend({
    template:"#template-problem-submit",
    data:function(){
        return {summary: {},language:0,code:"",submit_status:1,user:userInfo};
    },
    methods: {
        onLoginSuccess:function(){
            console.log("Success");
        },submit:function () {
            if(!this.code){
                new Snackbar("source code can not be blank.", {timeout: 3500});
                return;
            }
            this.submit_status = 0; //submitting
            var self = this;
            Util.postData.init(Config.base+"submit",{
                code:Base64.encode(this.code),language:this.language,problem_id:this.$route.params.id
            },null,function(data){
                new Snackbar("source code submitted,waiting for judge.", {timeout: 3500});
                self.code = "";
            },null,null,null,function(){//submitted but error
                self.submit_status = 1;
            });

        }
    }, mounted:function () {
    },created: function () {
        var id = this.$route.params.id;
        this.summary.id = id;
        $.ajax({
            url: Config.base + "problem/summary/" + id,
            context: this,
            success: function (data) {
                if (data != null) {
                    try {
                        this.summary = data;
                    } catch (e) {
                        new Snackbar("error happened while loading problema data.", {timeout: 3500});
                    }
                }
            }, error: function (err) {
                new Snackbar("error happened while loading problems data.", {timeout: 3500});
            }
        });
    }
});

const router = new VueRouter({
    base: Config.base,
    routes: [
        {path: '/', name: 'home', component: Home},
        {path: '/problems', name: 'problems', component: Problems},
        {path: '/problem/:id', name: 'detail', component: ProblemDetail},
        {path: '/submit/:id', name: 'submit', component: Submit}
    ]
});

var app = new Vue({
    router: router,
    data: {config: Config},
    methods: {
        goSignIn: function (key) { //eg:go to github.com
            var instance = this.config.auth[key];
            var url = instance.auth_url.replace("R", instance.redirect_uri).replace("C", instance.client_id);
            // console.log(url);
            window.open(url, "", "location=no,status=no");
            $("#sign-in-modal").modal("hide");
        }
    },
    mounted: function () {
        $.ajax({
            url: Config.base + "auth/user_status",
            success: function (data) {
                try {
                    if (data.is_login) {
                        userInfo.name = data.name;
                        userInfo.id = data.id;
                        userInfo.avatar = data.avatar;
                        userInfo.is_login = true;
                    }
                } catch (e) {
                    new Snackbar("error loading website data.", {timeout: 3500});
                }
            }, error: function (err) {
                new Snackbar("error loading website data.", {timeout: 3500});
            }
        });
    }
}).$mount('#app');
//
// setTimeout(function () {
//     console.log(app.$refs.app.onLoginSuccess());
// }, 2000);

window.addEventListener('message', function (e) {
    if (e.origin == location.origin) {
        var data = e.data;
        if (data.status == 1) {
            userInfo.name = data.name;
            userInfo.id = data.id;
            userInfo.avatar = data.avatar;
            userInfo.is_login = true;
            new Snackbar("登录认证成功", {timeout: 3500});
            if(app.$refs.app.onLoginSuccess){ //callback
                app.$refs.app.onLoginSuccess()
            }
        }
    }
});

function getOJNameByType(type) {
    if (type < Config.OJs.length && type >= 0) {
        return Config.OJs[type].name;
    } else {
        return "Unknown";
    }
}

function formatSourceUrl(type, url) {
    if (url == undefined) {
        return "";
    }
    if (url == "" || url.indexOf("http") == 0) {
        return url;
    }
    if (type < Config.OJs.length && type >= 0) {
        if (url.indexOf("/") != 0) {
            url = "/" + url;
        }
        return Config.OJs[type].homePage + url;
    } else {
        return "#";
    }
}