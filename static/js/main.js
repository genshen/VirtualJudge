/**
 * Created by gensh on 2016/11/3.
 */
const Config = {
    base: "/",
    appName: "Virtual Judge",
    logo: "https://v4-alpha.getbootstrap.com/assets/brand/bootstrap-solid.svg",
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
    data: function () {
        return {detail: {}, ojs: Config.OJs};
    },
    created: function () {
        var id = this.$route.params.id;
        $.ajax({
            url: Config.base + "problem/"+id,
            context: this,
            success: function (data) {
                //todo may be null
                try {
                    var self = this;
                    data.forEach(function (e) {
                        self.problems.push(e);
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

const router = new VueRouter({
    base: Config.base,
    routes: [
        {path: '/', name: 'home', component: Home},
        {path: '/problems', name: 'problems', component: Problems},
        {path: '/problem/:id', name: 'detail', component: ProblemDetail}
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
                    new Snackbar("加载数据出错啦,请刷新重试.", {timeout: 3500});
                }
            }, error: function (err) {
                new Snackbar("加载数据出错啦,请刷新重试.", {timeout: 3500});
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